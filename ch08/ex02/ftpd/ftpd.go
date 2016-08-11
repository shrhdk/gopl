package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type session struct {
	workdir    string
	user       string
	granted    bool
	dataType   string
	ctrlConn   *net.TCPConn
	clientAddr *net.TCPAddr
}

func (s *session) reply(code int, format string, a ...interface{}) error {
	message := fmt.Sprintf(format, a...)
	_, err := fmt.Fprintf(s.ctrlConn, "%d %s\r\n", code, message)
	return err
}

func (s *session) handleCommand(line string) error {
	tokens := strings.Fields(line)
	opc := strings.ToUpper(tokens[0])
	opr := tokens[1:]

	switch opc {

	// ACCESS CONTROL COMMANDS
	case "USER":
		return s.handleUserCommand(opc, opr)
	case "PASS":
		return s.handlePassCommand(opc, opr)
	case "ACCT":
		return s.reply(502, "%s command not implemented", opc)
	case "CWD":
		return s.handleCwdCommand(opc, opr)
	case "CDUP":
		return s.handleCdupCommand(opc, opr)
	case "REIN":
		return s.reply(502, "%s command not implemented", opc)
	case "QUIT":
		return s.handleQuitCommand(opc, opr)

	// TRANSFER PARAMETER COMMANDS
	case "PORT":
		return s.handlePortCommand(opc, opr)
	case "PASV":
		return s.reply(502, "%s command not implemented", opc)
	case "TYPE":
		return s.handleTypeCommand(opc, opr)
	case "STRU":
		return s.handleStruCommand(opc, opr)
	case "MODE":
		return s.handleModeCommand(opc, opr)

	// FTP SERVICE COMMANDS
	case "RETR":
		return s.handleRetrCommand(opc, opr)
	case "LIST":
		return s.handleListCommand(opc, opr)
	case "NLST":
		return s.handleListCommand(opc, opr)
	case "PWD":
		return s.handlePwdCommand(opc, opr)

	default:
		return s.reply(500, "%s not understood", opc)
	}
}

// ACCESS CONTROL COMMANDS

func (s *session) handleUserCommand(opc string, opr []string) error {
	if len(opr) < 1 {
		return s.reply(500, "%s: command requires a parameter", opc)
	}

	if s.granted {
		return s.reply(503, "You are already logged in")
	}

	s.user = opr[0]

	if s.user == "anonymous" {
		return s.reply(331, "Anonymous login ok, send your complete email address as your password")
	}

	return s.reply(331, "Password required for %s", s.user)
}

func (s *session) handlePassCommand(opc string, opr []string) error {
	if s.user == "" {
		return s.reply(503, "Login with USER first")
	}

	if s.granted {
		return s.reply(230, "Already logged in.")
	}

	if s.user == "anonymous" {
		s.granted = true
		return s.reply(230, "Anonymous access granted, restrictions apply")
	}

	return s.reply(530, "Login incorrect.")
}

func (s *session) handleCwdCommand(opc string, opr []string) error {
	if !s.granted {
		return s.reply(530, "Please login with USER and PASS")
	}

	workdir := filepath.Join(s.workdir, opr[0])

	if f, err := os.Stat(workdir); err != nil || !f.IsDir() {
		return s.reply(550, "Failed to change directory.")
	}

	s.workdir = workdir

	return s.reply(250, "Directory successfully changed.")
}

func (s *session) handleCdupCommand(opc string, opr []string) error {
	return s.handleCwdCommand(opc, []string{".."})
}

func (s *session) handleQuitCommand(opc string, opr []string) error {
	s.reply(221, "Goodbye.")
	return s.ctrlConn.Close()
}

// TRANSFER PARAMETER COMMANDS

func (s *session) handlePortCommand(opc string, opr []string) error {
	if len(opr) < 1 {
		return s.reply(500, "%s: command requires a parameter", opc)
	}

	clientAddr, err := parseFTPAddr(opr[0])
	if err != nil {
		return s.reply(500, "Illegal PORT command.")
	}

	s.clientAddr = clientAddr

	return s.reply(200, "PORT OK")
}

func (s *session) handleTypeCommand(opc string, opr []string) error {
	if len(opr) < 1 {
		return s.reply(500, "%s: command requires a parameter", opc)
	}

	if opr[0] == "A" || opr[0] == "I" {
		s.dataType = opr[0]
		return s.reply(200, "%s set to %s", opc, opr[0])
	}

	return s.reply(500, "'%s %s' unsupported value", opc, opr[0])
}

func (s *session) handleStruCommand(opc string, opr []string) error {
	if len(opr) < 1 {
		return s.reply(500, "%s: command requires a parameter", opc)
	}

	if opr[0] == "F" {
		return s.reply(200, "%s set to %s", opc, opr[0])
	}

	return s.reply(500, "'%s %s' unsupported value", opc, opr[0])
}

func (s *session) handleModeCommand(opc string, opr []string) error {
	if len(opr) < 1 {
		return s.reply(500, "%s: command requires a parameter", opc)
	}

	if opr[0] == "S" {
		return s.reply(200, "%s set to %s", opc, opr[0])
	}

	return s.reply(500, "'%s %s' unsupported value", opc, opr[0])
}

// FTP SERVICE COMMANDS

func (s *session) handleRetrCommand(opc string, opr []string) error {
	if s.clientAddr == nil {
		return s.reply(425, "Use PORT first.")
	}

	if len(opr) < 1 {
		return s.reply(500, "%s: command requires a parameter", opc)
	}

	p := filepath.Join(s.workdir, opr[0])

	f, err := os.Open(p)
	if err != nil {
		return s.reply(550, "Failed to open file.")
	}
	defer f.Close()

	s.reply(150, "File status okay; about to open data connection.")

	conn, err := net.DialTCP("tcp4", nil, s.clientAddr)
	if err != nil {
		return s.reply(425, "Can't open data connection.")
	}
	defer conn.Close()

	io.Copy(conn, f)

	return s.reply(226, "Transfer complete.")
}

func (s *session) handleListCommand(opc string, opr []string) error {
	if s.clientAddr == nil {
		return s.reply(425, "Use PORT first.")
	}

	finfos, err := ioutil.ReadDir(s.workdir)
	if err != nil {
		s.reply(550, "Failed to open directory.")
	}

	s.reply(150, "Here comes the directory listing.")

	conn, err := net.DialTCP("tcp4", nil, s.clientAddr)
	if err != nil {
		return s.reply(425, "Can't open data connection.")
	}
	defer conn.Close()

	for _, finfo := range finfos {
		fmt.Fprintf(conn, "%s\r\n", finfo.Name())
	}

	return s.reply(226, "Directory send OK.")
}

func (s *session) handlePwdCommand(opc string, opr []string) error {
	message := fmt.Sprintf("\"%s\" is the current directory", s.workdir)
	return s.reply(257, message)
}

// helper

func parseFTPAddr(addr string) (*net.TCPAddr, error) {
	ss := strings.Split(addr, ",")
	is := make([]byte, 0)
	for _, s := range ss {
		i, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		if i > 255 {
			return nil, fmt.Errorf("invalid addr value: " + s)
		}
		is = append(is, byte(i))
	}

	if len(is) != 6 {
		return nil, fmt.Errorf("invalid addr formant: " + addr)
	}

	ip := net.IPv4(is[0], is[1], is[2], is[3])
	port := int(is[4])*256 + int(is[5])

	tcpAddr, err := net.ResolveTCPAddr("tcp4", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		return nil, err
	}

	return tcpAddr, nil
}

// main

func startSession(conn *net.TCPConn) {
	var s session
	s.ctrlConn = conn
	if workdir, err := os.Getwd(); err == nil {
		s.workdir = workdir
	} else {
		log.Fatal(err)
	}

	s.reply(220, "FTP Server Ready")

	go func() {
		defer func() {
			conn.Close()
			log.Println("close controll connection")
		}()

		input := bufio.NewScanner(conn)
		for input.Scan() {
			s.handleCommand(input.Text())
		}
	}()
}

func main() {
	port := flag.Int("port", 21, "port to bind")
	flag.Parse()

	addr, err := net.ResolveTCPAddr("tcp4", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatal(err)
	}

	ctrlListener, err := net.ListenTCP("tcp4", addr)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := ctrlListener.AcceptTCP()
		if err != nil {
			log.Print(err)
			continue
		}
		startSession(conn)
	}
}
