// internal/art/art.go
package art

var AsciiArt = map[string]string{
	"train": `
      o O___ _________
    _][__|o| |O O O O|
   <_______|-|_______|
    /O-O-O     o   o
`,
	"cat": `
    /\_/\
   ( o.o )
    > ^ <
`,
	"dog": `
     / \__
    (    @\___
    /         O
   /   (_____/
  /_____/   U
`,
}

func GetArt(key string) (string, bool) {
	art, exists := AsciiArt[key]
	return art, exists
}

func GetAvailableArt() []string {
	keys := make([]string, 0, len(AsciiArt))
	for key := range AsciiArt {
		keys = append(keys, key)
	}
	return keys
}
