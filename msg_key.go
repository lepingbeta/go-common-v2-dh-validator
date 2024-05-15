package dhvalidator

func GetMsgKey(validName string) string {
    ErrorList := map[string]string{
        "firstName":  "invalid_firstname",
        "chineseMobile":  "invalid_chinese_mobile",
        "inviteCodeNum":  "invalid_invite_code",
        "lastName":  "invalid_lastname",
        "nickname":  "invalid_nickname",
        "gender":  "invalid_gender",
        "mongoId":  "invalid_mongoid",
        "image":  "invalid_image",
        "account":  "invalid_account",
        "password":  "invalid_password",
    }
    return ErrorList[validName]
}
