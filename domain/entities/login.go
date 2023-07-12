package entities

type SpeakerModel struct {
	SpeakerID           string   `json:"speaker_id" bson:"speaker_id,omitempty"`
	SpeakerName         string   `json:"speaker_name" bson:"speaker_name,omitempty"`
	EngName             string   `json:"eng_name" bson:"eng_name,omitempty"`
	ThaiName            string   `json:"thai_name" bson:"thai_name,omitempty"`
	Image               string   `json:"image" bson:"image,omitempty"`
	FaceImage           string   `json:"face_image" bson:"face_image,omitempty"`
	HorizintalFaceImage string   `json:"horizontal_face_image" bson:"horizontal_face_image,omitempty"`
	SquareImage         string   `json:"square_image" bson:"square_image,omitempty"`
	Audio               string   `json:"audio" bson:"audio,omitempty"`
	VoiceStyle          []string `json:"voice_style" bson:"voice_style,omitempty"`
	AgeStyle            string   `json:"age_style" bson:"age_style,omitempty"`
	SpeechStyle         []string `json:"speech_style" bson:"speech_style,omitempty"`
	Speed               string   `json:"speed" bson:"speed,omitempty"`
	Popularity          string   `json:"popularity" bson:"popularity,omitempty"`
	Type                int      `json:"type" bson:"type,omitempty"`
	Language            string   `json:"language" bson:"language,omitempty"`
	Status              bool     `json:"status" bson:"status,omitempty"`
	Gender              string   `json:"gender" bson:"gender,omitempty"`
	Private             bool     `json:"private" bson:"private,omitempty"`
	AllowUID            []string `json:"allow_uid" bson:"allow_uid,omitempty"`
}

type LineLoginData struct {
	Iss     string   `json:"iss bson:"iss,omitempty"`
	Sub     string   `json:"sub bson:"sub"`
	Aud     string   `json:"aud bson:"aud",omitempty"`
	Exp     int      `json:"exp bson:"exp",omitempty"`
	Iat     int      `json:"iat bson:"iat",omitempty"`
	Amr     []string `json:"amr bson:"amr",omitempty"`
	Name    string   `json:"name" bson:"name"`
	Email   string   `json:"email bson:"email",omitempty"`
	Picture string   `json:"picture bson:"picture",omitempty"`
}
