package dhvalidator

func GetMsgKey(validName string) string {
    ErrorList := map[string]string{
        "firstName":  "invalid_firstname",
        "chineseMobile":  "invalid_chinese_mobile",
        "lastName":  "invalid_lastname",
        "nickname":  "invalid_nickname",
        "gender":  "invalid_gender",
        "image":  "invalid_image",
        "account":  "invalid_account",
        "password":  "invalid_password",
    }
    return ErrorList[validName]
}
