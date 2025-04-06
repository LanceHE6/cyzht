package encrypt

import (
	"crypto/sha256"
	"encoding/hex"
)

// CalculateFileHash
//
//	@Description: 计算文件sha256
//	@param content 文件二进制
//	@return string sha256值
func CalculateFileHash(content []byte) string {
	hash := sha256.New()
	hash.Write(content)
	return hex.EncodeToString(hash.Sum(nil))
}
