package enums

type EnumTag string

const (
	TagNSFW           EnumTag = "nsfw"     // Not Safe For Work
	TagSpoiler        EnumTag = "spoiler"  // May ruin a surprise
	TagMatureContent  EnumTag = "mature"   // Contains mature or adult content
	TagBrandAffiliate EnumTag = "brand"    // Brand affiliate
	TagBusiness       EnumTag = "business" // Made for a brand or business
)

func GetEnumTag(tag *string) EnumTag {
	if tag == nil {
		return ""
	}
	switch *tag {
	case "nsfw":
		return TagNSFW
	case "spoiler":
		return TagSpoiler
	case "mature":
		return TagMatureContent
	case "brand":
		return TagBrandAffiliate
	case "business":
		return TagBusiness
	default:
		return ""
	}
}
