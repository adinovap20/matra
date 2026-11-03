package prosody

import (
	"errors"
	"fmt"
)

type VruttaKey uint64

// marathiVruttaRegistry maps a compressed bitmask to its Aksharaganavrutti name.
// Bitmask pattern: 0 = Laghu (लघु), 1 = Guru (गुरू)
var marathiVruttaRegistry = map[VruttaKey]string{
	// 11 Syllables
	makeKey(11, 0b11011001011): "इंद्रवज्रा",   // त त ज ग ग (Ta Ta Ja Ga Ga)
	makeKey(11, 0b01011001011): "उपेंद्रवज्रा", // ज त ज ग ग (Ja Ta Ja Ga Ga)
	makeKey(11, 0b01101101101): "भुजंगी",       // य य य ल ग (Bhujangaprayat minus last Guru)

	// 12 Syllables
	makeKey(12, 0b011011011011): "भुजंगप्रयात",  // य य य य (Ya Ya Ya Ya)
	makeKey(12, 0b010110010101): "वंशस्थ",       // ज त ज र (Ja Ta Ja Ra)
	makeKey(12, 0b000100100101): "द्रुतविलंबित", // न भ भ र (Na Bha Bha Ra)
	makeKey(12, 0b001001001001): "तोटक",         // स स स स (Sa Sa Sa Sa)

	// 14 Syllables
	makeKey(14, 0b11010001001011): "वसंततिलका", // त भ ज ज ग ग (Ta Bha Ja Ja Ga Ga)

	// 15 Syllables
	makeKey(15, 0b000000111011011): "मालिनी", // न न म य य (Na Na Ma Ya Ya)

	// 16 Syllables
	makeKey(16, 0b0101010101010101): "पंचचामर", // ज र ज र ज ग (Ja Ra Ja Ra Ja Ga)

	// 17 Syllables
	makeKey(17, 0b11110000011011011): "मंदाक्रांता", // म भ न त त ग ग (Ma Bha Na Ta Ta Ga Ga)
	makeKey(17, 0b01000101000101101): "पृथ्वी",      // ज स ज स य ल ग (Ja Sa Ja Sa Ya La Ga)
	makeKey(17, 0b01111100000110001): "शिखरिणी",     // य म न स भ ल ग (Ya Ma Na Sa Bha La Ga)

	// 19 Syllables
	makeKey(19, 0b1110010100011101101): "शार्दूलविक्रीडित", // म स ज स त त ग (Ma Sa Ja Sa Ta Ta Ga)

	// 21 Syllables
	makeKey(21, 0b111101100000011011011): "स्रग्धरा", // म र भ न य य य (Ma Ra Bha Na Ya Ya Ya)

	// 22 Syllables
	makeKey(22, 0b1101101101101101101101): "मंदारमाला", // त त त त त त त ग (Ta x7 + Ga)

	// 23 Syllables
	makeKey(23, 0b01101101101101101101101): "सुमंदारमाला", // य य य य य य य ल ग (Ya x7 + La Ga)
}

// sanskritVruttaRegistry maps a compressed bitmask to its Classical Sanskrit meter name.
// Bitmask pattern: 0 = Laghu (लघु / L), 1 = Guru (गुरू / G)
// Gaṇa reference: M=111, Y=011, R=101, S=001, T=110, J=010, Bh=100, N=000
var sanskritVruttaRegistry = map[VruttaKey]string{
	// --- 8 Syllables (Anuṣṭubh Class - Pure Samavrittas) ---
	makeKey(8, 0b11111111): "विद्युन्माला", // Vidyunmālā (M M G G)
	makeKey(8, 0b01010101): "प्रमाणी",      // Pramāṇī (J R L G)
	makeKey(8, 0b10101010): "समानी",        // Samānī (R J G L)
	makeKey(8, 0b10010011): "चित्रपदा",     // Citrapadā (Bh Bh G G)
	makeKey(8, 0b11100011): "हंसरुतम्",     // Haṃsarutam (M N G G)

	// --- 11 Syllables (Triṣṭubh Class) ---
	makeKey(11, 0b11011001011): "इन्द्रवज्रा",   // Indravajrā (T T J G G)
	makeKey(11, 0b01011001011): "उपेन्द्रवज्रा", // Upendravajrā (J T J G G)
	makeKey(11, 0b11111011011): "शालिनी",        // Śālinī (M T T G G)
	makeKey(11, 0b11110011011): "वातोर्मी",      // Vātormī (M Bh T G G)
	makeKey(11, 0b10100010101): "रथोद्धता",      // Rathoddhatā (R N R L G)
	makeKey(11, 0b10100010011): "स्वागता",       // Svāgatā (R N Bh G G)
	makeKey(11, 0b10010010011): "दोधकम्",        // Dodhakam (Bh Bh Bh G G)
	makeKey(11, 0b11100000001): "भ्रमरविलसिता",  // Bhramaravilasitā (M N N L G)
	makeKey(11, 0b11011001001): "मोटनक",         // Moṭanaka (T T J L G)

	// --- 12 Syllables (Jagatī Class) ---
	makeKey(12, 0b010110010101): "वंशस्थ",        // Vaṃśastha (J T J R)
	makeKey(12, 0b110110010101): "इन्द्रवंशा",    // Indravaṃśā (T T J R)
	makeKey(12, 0b011011011011): "भुजङ्गप्रयात",  // Bhujaṅgaprayāta (Y Y Y Y)
	makeKey(12, 0b000100100101): "द्रुतविलम्बित", // Drutavilambita (N Bh Bh R)
	makeKey(12, 0b001001001001): "तोटक",          // Toṭaka (S S S S)
	makeKey(12, 0b101101101101): "स्रग्विणी",     // Sragviṇī (R R R R)
	makeKey(12, 0b111111011011): "वैश्वदेवी",     // Vaiśvadevī (M M Y Y)
	makeKey(12, 0b010001010001): "जलोद्धतगति",    // Jaloddhatagati (J S J S)
	makeKey(12, 0b001010001001): "प्रमिताक्षरा",  // Pramitākṣarā (S J S S)

	// --- 13 Syllables (Atijagatī Class) ---
	makeKey(13, 0b1110000101011): "प्रहर्षिणी",  // Praharṣiṇī (M N J R G)
	makeKey(13, 0b0101000010101): "रुचिरा",      // Rucirā (J Bh S J G)
	makeKey(13, 0b1111100110011): "मत्तमयूर",    // Mattamayūra (M T Y S G)
	makeKey(13, 0b0010100010101): "मञ्जुभाषिणी", // Mañjubhāṣiṇī (S J S J G)

	// --- 14 Syllables (Śakkarī Class) ---
	makeKey(14, 0b11010001001011): "वसन्ततिलका", // Vasantatilakā (T Bh J J G G)
	makeKey(14, 0b00000010100101): "अपराजिता",   // Aparājitā (N N R S L G)

	// --- 15 Syllables (Atiśakkarī Class) ---
	makeKey(15, 0b000000111011011): "मालिनी", // Mālinī (N N M Y Y)

	// --- 16 Syllables (Aṣṭi Class) ---
	makeKey(16, 0b0101010101010101): "पञ्चचामर", // Pañcacāmara (J R J R J G)

	// --- 17 Syllables (Atyaṣṭi Class) ---
	makeKey(17, 0b11110000011011011): "मन्दाक्रान्ता", // Mandākrāntā (M Bh N T T G G)
	makeKey(17, 0b01111100000110001): "शिखरिणी",       // Śikhariṇī (Y M N S Bh L G)
	makeKey(17, 0b01000101000101101): "पृथ्वी",        // Pṛthvī (J S J S Y L G)
	makeKey(17, 0b00000111110100101): "हरिणी",         // Hariṇī (N S M R S L G)

	// --- 18 Syllables (Dhṛti Class) ---
	makeKey(18, 0b111110000011011011): "कुसुमितलतावेल्लिता", // Kusumitalatāvellitā (M T N Y Y Y)

	// --- 19 Syllables (Atidhṛti Class) ---
	makeKey(19, 0b1110010100011101101): "शार्दूलविक्रीडित", // Śārdūlavikrīḍita (M S J S T T G)

	// --- 21 Syllables (Prakṛti Class) ---
	makeKey(21, 0b111101100000011011011): "स्रग्धरा", // Sragdharā (M R Bh N Y Y Y)

	// --- 23 Syllables (Vikṛti Class) ---
	makeKey(23, 0b00001010001010001010001): "अश्वललित", // Aśvalalita (N J Bh J Bh J Bh L G)
}

func makeKey(length int, bits uint64) VruttaKey {
	return VruttaKey((uint64(length) << 58) | bits)
}

func Encode(syllables []int) (VruttaKey, error) {
	n := len(syllables)
	if n > 58 {
		return 0, errors.New("syllable count exceeds bitmask capacity")
	}

	var bits uint64
	for _, s := range syllables {
		bits <<= 1
		if s == 2 {
			bits |= 1
		} else if s != 1 {
			return 0, fmt.Errorf("invalid syllable marker: %d (expected 1 or 2)", s)
		}
	}

	return makeKey(n, bits), nil
}

func IdentifyMarathiVrutta(syllables []int) (string, error) {
	key, err := Encode(syllables)
	if err != nil {
		return "", err
	}

	name, found := marathiVruttaRegistry[key]
	if !found {
		return "", fmt.Errorf("unrecognized meter for %d syllables", len(syllables))
	}

	return name, nil
}

func IdentifySanskritVrutta(syllables []int) (string, error) {
	key, err := Encode(syllables)
	if err != nil {
		return "", err
	}

	name, found := sanskritVruttaRegistry[key]
	if !found {
		return "", fmt.Errorf("unrecognized meter for %d syllables", len(syllables))
	}

	return name, nil
}
