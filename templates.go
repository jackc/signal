package signal

import "github.com/clipperhouse/typewriter"

var templates = typewriter.TemplateSlice{
	signal,
}

var signal = &typewriter.Template{
	Name: "Signal",
	Text: `
// The primary type that represents a signal
type {{.Name}}Signal struct {
   listeners [](chan <%= type_name %>)
   mutex     sync.Mutex
}

// Add channel c to the signal to receive messages from this Signal
func (s *{{.Name}}Signal) Add(c chan {{.Name}}) {
   s.mutex.Lock()
   defer s.mutex.Unlock()
   s.listeners = append(s.listeners, c)
}

// Remove channel c from the signal
func (s *{{.Name}}Signal) Remove(c chan {{.Name}}) {
   s.mutex.Lock()
   defer s.mutex.Unlock()

   for i, l := range s.listeners {
      if c == l {
         s.listeners[i] = s.listeners[len(s.listeners)-1]
         s.listeners = s.listeners[:len(s.listeners)-1]
         return
      }
   }
}

// Dispatch synchronously sends msg to all channels that have been added to this signal.
func (s *{{.Name}}Signal) Dispatch(msg {{.Name}}) {
   s.mutex.Lock()
   defer s.mutex.Unlock()

   for _, l := range s.listeners {
      l <- msg
   }
}
`,
}
