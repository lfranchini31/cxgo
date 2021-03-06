package libs

import (
	"strings"
	"testing"

	"github.com/gotranspile/cxgo/types"
	"github.com/stretchr/testify/require"
)

func TestStdInt(t *testing.T) {
	c := types.NewEnv(types.Config32())
	s := incStdInt(c)
	s = strings.TrimSpace(s)
	s = strings.TrimSpace(s)
	require.Equal(t, strings.TrimSpace(`
#include <cxgo_builtin.h>
#define int8_t _cxgo_sint8
#define int16_t _cxgo_sint16
#define int32_t _cxgo_sint32
#define int64_t _cxgo_sint64

#define uint8_t _cxgo_uint8
#define uint16_t _cxgo_uint16
#define uint32_t _cxgo_uint32
#define uint64_t _cxgo_uint64

#define int_least8_t _cxgo_sint8
#define int_least16_t _cxgo_sint16
#define int_least32_t _cxgo_sint32
#define int_least64_t _cxgo_sint64

#define uint_least8_t _cxgo_uint8
#define uint_least16_t _cxgo_uint16
#define uint_least32_t _cxgo_uint32
#define uint_least64_t _cxgo_uint64

#define int_fast8_t _cxgo_sint8
#define int_fast16_t _cxgo_sint16
#define int_fast32_t _cxgo_sint32
#define int_fast64_t _cxgo_sint64

#define uint_fast8_t _cxgo_uint8
#define uint_fast16_t _cxgo_uint16
#define uint_fast32_t _cxgo_uint32
#define uint_fast64_t _cxgo_uint64

typedef _cxgo_sint32 intptr_t;
typedef _cxgo_uint32 uintptr_t;

typedef _cxgo_sint64 intmax_t;
typedef _cxgo_uint64 uintmax_t;

#define INT8_MIN -0x80
#define INT8_MAX 0x7f
#define UINT8_MAX 0xff
#define INT16_MIN -0x8000
#define INT16_MAX 0x7fff
#define UINT16_MAX 0xffff
#define INT32_MIN -0x80000000
#define INT32_MAX 0x7fffffff
#define UINT32_MAX 0xffffffff
#define INT64_MIN -0x8000000000000000
#define INT64_MAX 0x7fffffffffffffff
#define UINT64_MAX 0xffffffffffffffff

#define INT_LEAST8_MIN -0x80
#define INT_LEAST8_MAX 0x7f
#define UINT_LEAST8_MAX 0xff
#define INT_LEAST16_MIN -0x8000
#define INT_LEAST16_MAX 0x7fff
#define UINT_LEAST16_MAX 0xffff
#define INT_LEAST32_MIN -0x80000000
#define INT_LEAST32_MAX 0x7fffffff
#define UINT_LEAST32_MAX 0xffffffff
#define INT_LEAST64_MIN -0x8000000000000000
#define INT_LEAST64_MAX 0x7fffffffffffffff
#define UINT_LEAST64_MAX 0xffffffffffffffff

#define INT_FAST8_MIN -0x80
#define INT_FAST8_MAX 0x7f
#define UINT_FAST8_MAX 0xff
#define INT_FAST16_MIN -0x8000
#define INT_FAST16_MAX 0x7fff
#define UINT_FAST16_MAX 0xffff
#define INT_FAST32_MIN -0x80000000
#define INT_FAST32_MAX 0x7fffffff
#define UINT_FAST32_MAX 0xffffffff
#define INT_FAST64_MIN -0x8000000000000000
#define INT_FAST64_MAX 0x7fffffffffffffff
#define UINT_FAST64_MAX 0xffffffffffffffff

#define INTPTR_MIN -0x80000000
#define INTPTR_MAX 0x7fffffff
#define UINTPTR_MAX 0xffffffff

#define INTMAX_MIN -0x8000000000000000
#define INTMAX_MAX 0x7fffffffffffffff
#define UINTMAX_MAX 0xffffffffffffffff

#define PTRDIFF_MIN -0x80000000
#define PTRDIFF_MAX 0x7fffffff

#define SIZE_MAX 0xffffffff

#define WCHAR_MIN 0x0
#define WCHAR_MAX 0xffff

#define WINT_MIN 0x0
#define WINT_MAX 0xffffffff
`), strings.TrimSpace(s))
}
