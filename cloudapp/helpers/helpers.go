package helpers

import (
	"github.com/cloudfoundry-community/go-cfclient"
	"go4.org/sort"
	"math/rand"
	"encoding/json"
	"fmt"
)

// NameSorter sorts planets by name.
type NameSorter []cfclient.Buildpack

func (a NameSorter) Len() int           { return len(a) }
func (a NameSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a NameSorter) Less(i, j int) bool { return a[i].Name < a[j].Name }

func GetConvertToJson(buildpacks []cfclient.Buildpack) string {
	test, err := json.Marshal(buildpacks)
	if err != nil {
		fmt.Printf("Helpers Error: %s", err)
		panic(err)
	}
	return string(test)
}

func Shuffle(buildpacks []cfclient.Buildpack) {
	for i := 1; i < len(buildpacks); i++ {
		r := rand.Intn(i + 1)
		if i != r {
			buildpacks[r], buildpacks[i] = buildpacks[i], buildpacks[r]
		}
	}
}

func GetSortListByName(buildpacks []cfclient.Buildpack) {
	sort.Sort(NameSorter(buildpacks))
}
