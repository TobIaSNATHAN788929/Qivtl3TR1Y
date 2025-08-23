// 代码生成时间: 2025-08-24 01:24:07
Features:
- Supports SHA1, SHA256, and MD5 hash algorithms.

Usage:
- POST /hash with a JSON payload containing the string to be hashed and the desired algorithm.

Example JSON payload:
{
	"string": "Hello, World!",
	"algorithm": "SHA256"
}
*/

package main

import (
	"crypto/sha1"
	