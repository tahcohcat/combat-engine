package entities

type RegionType int

const NumRegions= 2

const (
	UnknownRegion RegionType = 0
	WoodsRegion RegionType = 1
	RiverRegion RegionType = 2
)

type Region struct {
	Region RegionType
	CharacterCapacity uint8

	// Left to Right
	BoardLocationIndex uint8

	// Characters played within region
	Characters [] Character

	GameResource GameResource
}

func NewRegion(characterCapacity uint8) Region {
	return Region{
		Region:             getRandomRegionType(),
		CharacterCapacity:  characterCapacity,
		BoardLocationIndex: 0,
		Characters:         make([] Character, int(characterCapacity)),
		GameResource: GenerateResource(ResRegion, "resource-region"),
	}
}