// Code generated for linux/arm64 by 'ccgo --goptr=off --package-name=standard -o internal/standard/md4c_generated.go third_party/md4c/md4c.c third_party/md4c/md4c-html.c third_party/md4c/entity.c third_party/md4c-bridge/bridge.c', DO NOT EDIT.

//go:build linux && arm64

package standard

import (
	"math"
	"reflect"
	"unsafe"

	"modernc.org/libc"
)

var _ = math.Pi
var _ reflect.Type
var _ unsafe.Pointer

const ARG_MAX = 131072
const BC_BASE_MAX = 99
const BC_DIM_MAX = 2048
const BC_SCALE_MAX = 99
const BC_STRING_MAX = 1000
const BUFSIZ = 1024
const CHAR = "MD_CHAR"
const CHARCLASS_NAME_MAX = 14
const CHAR_BIT = 8
const CHAR_MAX = 255
const CHAR_MIN = 0
const CODESPAN_MARK_MAXLEN = 32
const COLL_WEIGHTS_MAX = 2
const DELAYTIMER_MAX = 0x7fffffff
const EXIT_FAILURE = 1
const EXIT_SUCCESS = 0
const EXPR_NEST_MAX = 32
const FALSE = 0
const FILENAME_MAX = 4096
const FILESIZEBITS = 64
const FOPEN_MAX = 1000
const HOST_NAME_MAX = 255
const INT_MAX = 0x7fffffff
const IOV_MAX = 1024
const LINE_MAX = 4096
const LLONG_MAX = 0x7fffffffffffffff
const LOGIN_NAME_MAX = 256
const LONG_BIT = 64
const LONG_MAX = "__LONG_MAX"
const L_ctermid = 20
const L_cuserid = 20
const L_tmpnam = 20
const MB_LEN_MAX = 4
const MD_ANALYZE_NOSKIP_EMPH = 1
const MD_BLOCK_CONTAINER_CLOSER = 2
const MD_BLOCK_CONTAINER_OPENER = 1
const MD_BLOCK_LOOSE_LIST = 4
const MD_BLOCK_SETEXT_HEADER = 8
const MD_BUILD_ATTR_NO_ESCAPES = 1
const MD_DIALECT_COMMONMARK = 0
const MD_FLAG_COLLAPSEWHITESPACE = 1
const MD_FLAG_HARD_SOFT_BREAKS = 32768
const MD_FLAG_LATEXMATHSPANS = 4096
const MD_FLAG_NOHTMLBLOCKS = 32
const MD_FLAG_NOHTMLSPANS = 64
const MD_FLAG_NOINDENTEDCODEBLOCKS = 16
const MD_FLAG_PERMISSIVEATXHEADERS = 2
const MD_FLAG_PERMISSIVEEMAILAUTOLINKS = 8
const MD_FLAG_PERMISSIVEURLAUTOLINKS = 4
const MD_FLAG_PERMISSIVEWWWAUTOLINKS = 1024
const MD_FLAG_STRIKETHROUGH = 512
const MD_FLAG_TABLES = 256
const MD_FLAG_TASKLISTS = 2048
const MD_FLAG_UNDERLINE = 16384
const MD_FLAG_WIKILINKS = 8192
const MD_FNV1A_BASE = 2166136261
const MD_FNV1A_PRIME = 16777619
const MD_MARK_AUTOLINK = 32
const MD_MARK_AUTOLINK_MISSING_MAILTO = 64
const MD_MARK_CLOSER = 8
const MD_MARK_EMPH_MOD3_0 = 64
const MD_MARK_EMPH_MOD3_1 = 128
const MD_MARK_EMPH_OC = 32
const MD_MARK_HASNESTEDBRACKETS = 32
const MD_MARK_OPENER = 4
const MD_MARK_POTENTIAL_CLOSER = 2
const MD_MARK_POTENTIAL_OPENER = 1
const MD_MARK_RESOLVED = 16
const MD_MARK_VALIDPERMISSIVEAUTOLINK = 32
const MD_ROLLBACK_ALL = 1
const MD_ROLLBACK_CROSSING = 0
const MQ_PRIO_MAX = 32768
const NAME_MAX = 255
const NGROUPS_MAX = 32
const NL_ARGMAX = 9
const NL_LANGMAX = 32
const NL_MSGMAX = 32767
const NL_NMAX = 16
const NL_SETMAX = 255
const NL_TEXTMAX = 2048
const NZERO = 20
const OFF = "MD_OFFSET"
const PATH_MAX = 4096
const PIPE_BUF = 4096
const PTHREAD_DESTRUCTOR_ITERATIONS = 4
const PTHREAD_KEYS_MAX = 128
const PTHREAD_STACK_MIN = 2048
const P_tmpdir = "/tmp"
const RAND_MAX = 0x7fffffff
const RE_DUP_MAX = 255
const SCHAR_MAX = 127
const SEM_NSEMS_MAX = 256
const SEM_VALUE_MAX = 0x7fffffff
const SHRT_MAX = 0x7fff
const SSIZE_MAX = "LONG_MAX"
const SYMLOOP_MAX = 40
const SZ = "MD_SIZE"
const TABLE_MAXCOLCOUNT = 128
const TMP_MAX = 10000
const TRUE = 1
const TTY_NAME_MAX = 32
const TZNAME_MAX = 6
const UCHAR_MAX = 255
const UINT_MAX = 4294967295
const USHRT_MAX = 0xffff
const WNOHANG = 1
const WORD_BIT = 32
const WUNTRACED = 2
const _FORTIFY_SOURCE = 3
const _GNU_SOURCE = 1
const _IOFBF = 0
const _IOLBF = 1
const _IONBF = 2
const _LIBCPP_HARDENING_MODE = "_LIBCPP_HARDENING_MODE_FAST"
const _LP64 = 1
const _POSIX2_BC_BASE_MAX = 99
const _POSIX2_BC_DIM_MAX = 2048
const _POSIX2_BC_SCALE_MAX = 99
const _POSIX2_BC_STRING_MAX = 1000
const _POSIX2_CHARCLASS_NAME_MAX = 14
const _POSIX2_COLL_WEIGHTS_MAX = 2
const _POSIX2_EXPR_NEST_MAX = 32
const _POSIX2_LINE_MAX = 2048
const _POSIX2_RE_DUP_MAX = 255
const _POSIX_AIO_LISTIO_MAX = 2
const _POSIX_AIO_MAX = 1
const _POSIX_ARG_MAX = 4096
const _POSIX_CHILD_MAX = 25
const _POSIX_CLOCKRES_MIN = 20000000
const _POSIX_DELAYTIMER_MAX = 32
const _POSIX_HOST_NAME_MAX = 255
const _POSIX_LINK_MAX = 8
const _POSIX_LOGIN_NAME_MAX = 9
const _POSIX_MAX_CANON = 255
const _POSIX_MAX_INPUT = 255
const _POSIX_MQ_OPEN_MAX = 8
const _POSIX_MQ_PRIO_MAX = 32
const _POSIX_NAME_MAX = 14
const _POSIX_NGROUPS_MAX = 8
const _POSIX_OPEN_MAX = 20
const _POSIX_PATH_MAX = 256
const _POSIX_PIPE_BUF = 512
const _POSIX_RE_DUP_MAX = 255
const _POSIX_RTSIG_MAX = 8
const _POSIX_SEM_NSEMS_MAX = 256
const _POSIX_SEM_VALUE_MAX = 32767
const _POSIX_SIGQUEUE_MAX = 32
const _POSIX_SSIZE_MAX = 32767
const _POSIX_SS_REPL_MAX = 4
const _POSIX_STREAM_MAX = 8
const _POSIX_SYMLINK_MAX = 255
const _POSIX_SYMLOOP_MAX = 8
const _POSIX_THREAD_DESTRUCTOR_ITERATIONS = 4
const _POSIX_THREAD_KEYS_MAX = 128
const _POSIX_THREAD_THREADS_MAX = 64
const _POSIX_TIMER_MAX = 32
const _POSIX_TRACE_EVENT_NAME_MAX = 30
const _POSIX_TRACE_NAME_MAX = 8
const _POSIX_TRACE_SYS_MAX = 8
const _POSIX_TRACE_USER_EVENT_MAX = 32
const _POSIX_TTY_NAME_MAX = 9
const _POSIX_TZNAME_MAX = 6
const _STDC_PREDEF_H = 1
const _XOPEN_IOV_MAX = 16
const _XOPEN_NAME_MAX = 255
const _XOPEN_PATH_MAX = 1024
const __AARCH64EL__ = 1
const __AARCH64_CMODEL_SMALL__ = 1
const __ARM_64BIT_STATE = 1
const __ARM_ALIGN_MAX_PWR = 28
const __ARM_ALIGN_MAX_STACK_PWR = 16
const __ARM_ARCH = 8
const __ARM_ARCH_8A = 1
const __ARM_ARCH_ISA_A64 = 1
const __ARM_ARCH_PROFILE = 65
const __ARM_BUILDATTR64_FV = 65
const __ARM_FEATURE_CLZ = 1
const __ARM_FEATURE_FMA = 1
const __ARM_FEATURE_IDIV = 1
const __ARM_FEATURE_NUMERIC_MAXMIN = 1
const __ARM_FEATURE_UNALIGNED = 1
const __ARM_FP = 14
const __ARM_FP16_ARGS = 1
const __ARM_FP16_FORMAT_IEEE = 1
const __ARM_NEON = 1
const __ARM_NEON_SVE_BRIDGE = 1
const __ARM_PCS_AAPCS64 = 1
const __ARM_SIZEOF_MINIMAL_ENUM = 4
const __ARM_SIZEOF_WCHAR_T = 4
const __ARM_STATE_ZA = 1
const __ARM_STATE_ZT0 = 1
const __ATOMIC_ACQUIRE = 2
const __ATOMIC_ACQ_REL = 4
const __ATOMIC_CONSUME = 1
const __ATOMIC_RELAXED = 0
const __ATOMIC_RELEASE = 3
const __ATOMIC_SEQ_CST = 5
const __BFLT16_DECIMAL_DIG__ = 4
const __BFLT16_DENORM_MIN__ = "9.18354961579912115600575419704879436e-41B"
const __BFLT16_DIG__ = 2
const __BFLT16_EPSILON__ = "7.81250000000000000000000000000000000e-3B"
const __BFLT16_HAS_DENORM__ = 1
const __BFLT16_HAS_INFINITY__ = 1
const __BFLT16_HAS_QUIET_NAN__ = 1
const __BFLT16_IS_IEC_60559__ = 0
const __BFLT16_MANT_DIG__ = 8
const __BFLT16_MAX_10_EXP__ = 38
const __BFLT16_MAX_EXP__ = 128
const __BFLT16_MAX__ = "3.38953138925153547590470800371487867e+38B"
const __BFLT16_MIN__ = "1.17549435082228750796873653722224568e-38B"
const __BFLT16_NORM_MAX__ = "3.38953138925153547590470800371487867e+38B"
const __BIGGEST_ALIGNMENT__ = 16
const __BIG_ENDIAN = 4321
const __BITINT_MAXWIDTH__ = 65535
const __BYTE_ORDER = 1234
const __BYTE_ORDER__ = "__ORDER_LITTLE_ENDIAN__"
const __CCGO__ = 1
const __CHAR_BIT__ = 8
const __CHAR_UNSIGNED__ = 1
const __DBL_DECIMAL_DIG__ = 17
const __DBL_DIG__ = 15
const __DBL_HAS_DENORM__ = 1
const __DBL_HAS_INFINITY__ = 1
const __DBL_HAS_QUIET_NAN__ = 1
const __DBL_IS_IEC_60559__ = 1
const __DBL_MANT_DIG__ = 53
const __DBL_MAX_10_EXP__ = 308
const __DBL_MAX_EXP__ = 1024
const __DEC128_EPSILON__ = 1e-33
const __DEC128_MANT_DIG__ = 34
const __DEC128_MAX_EXP__ = 6145
const __DEC128_MAX__ = "9.999999999999999999999999999999999E6144"
const __DEC128_MIN__ = 1e-6143
const __DEC128_SUBNORMAL_MIN__ = 0.000000000000000000000000000000001e-6143
const __DEC32_EPSILON__ = 1e-6
const __DEC32_MANT_DIG__ = 7
const __DEC32_MAX_EXP__ = 97
const __DEC32_MAX__ = 9.999999e96
const __DEC32_MIN__ = 1e-95
const __DEC32_SUBNORMAL_MIN__ = 0.000001e-95
const __DEC64X_EPSILON__ = "1E-33D64x"
const __DEC64X_MANT_DIG__ = 34
const __DEC64X_MAX_EXP__ = 6145
const __DEC64X_MAX__ = "9.999999999999999999999999999999999E6144D64x"
const __DEC64X_MIN__ = "1E-6143D64x"
const __DEC64X_SUBNORMAL_MIN__ = "0.000000000000000000000000000000001E-6143D64x"
const __DEC64_EPSILON__ = 1e-15
const __DEC64_MANT_DIG__ = 16
const __DEC64_MAX_EXP__ = 385
const __DEC64_MAX__ = "9.999999999999999E384"
const __DEC64_MIN__ = 1e-383
const __DEC64_SUBNORMAL_MIN__ = 0.000000000000001e-383
const __DECIMAL_BID_FORMAT__ = 1
const __DECIMAL_DIG__ = 36
const __DEC_EVAL_METHOD__ = 2
const __ELF__ = 1
const __FINITE_MATH_ONLY__ = 0
const __FLOAT_WORD_ORDER__ = "__ORDER_LITTLE_ENDIAN__"
const __FLT128_DECIMAL_DIG__ = 36
const __FLT128_DENORM_MIN__ = 6.47517511943802511092443895822764655e-4966
const __FLT128_DIG__ = 33
const __FLT128_EPSILON__ = 1.92592994438723585305597794258492732e-34
const __FLT128_HAS_DENORM__ = 1
const __FLT128_HAS_INFINITY__ = 1
const __FLT128_HAS_QUIET_NAN__ = 1
const __FLT128_IS_IEC_60559__ = 1
const __FLT128_MANT_DIG__ = 113
const __FLT128_MAX_10_EXP__ = 4932
const __FLT128_MAX_EXP__ = 16384
const __FLT128_MAX__ = "1.18973149535723176508575932662800702e+4932"
const __FLT128_MIN__ = 3.36210314311209350626267781732175260e-4932
const __FLT128_NORM_MAX__ = "1.18973149535723176508575932662800702e+4932"
const __FLT16_DECIMAL_DIG__ = 5
const __FLT16_DENORM_MIN__ = 5.96046447753906250000000000000000000e-8
const __FLT16_DIG__ = 3
const __FLT16_EPSILON__ = 9.76562500000000000000000000000000000e-4
const __FLT16_HAS_DENORM__ = 1
const __FLT16_HAS_INFINITY__ = 1
const __FLT16_HAS_QUIET_NAN__ = 1
const __FLT16_IS_IEC_60559__ = 1
const __FLT16_MANT_DIG__ = 11
const __FLT16_MAX_10_EXP__ = 4
const __FLT16_MAX_EXP__ = 16
const __FLT16_MAX__ = 6.55040000000000000000000000000000000e+4
const __FLT16_MIN__ = 6.10351562500000000000000000000000000e-5
const __FLT16_NORM_MAX__ = 6.55040000000000000000000000000000000e+4
const __FLT32X_DECIMAL_DIG__ = 17
const __FLT32X_DENORM_MIN__ = 4.94065645841246544176568792868221372e-324
const __FLT32X_DIG__ = 15
const __FLT32X_EPSILON__ = 2.22044604925031308084726333618164062e-16
const __FLT32X_HAS_DENORM__ = 1
const __FLT32X_HAS_INFINITY__ = 1
const __FLT32X_HAS_QUIET_NAN__ = 1
const __FLT32X_IS_IEC_60559__ = 1
const __FLT32X_MANT_DIG__ = 53
const __FLT32X_MAX_10_EXP__ = 308
const __FLT32X_MAX_EXP__ = 1024
const __FLT32X_MAX__ = 1.79769313486231570814527423731704357e+308
const __FLT32X_MIN__ = 2.22507385850720138309023271733240406e-308
const __FLT32X_NORM_MAX__ = 1.79769313486231570814527423731704357e+308
const __FLT32_DECIMAL_DIG__ = 9
const __FLT32_DENORM_MIN__ = 1.40129846432481707092372958328991613e-45
const __FLT32_DIG__ = 6
const __FLT32_EPSILON__ = 1.19209289550781250000000000000000000e-7
const __FLT32_HAS_DENORM__ = 1
const __FLT32_HAS_INFINITY__ = 1
const __FLT32_HAS_QUIET_NAN__ = 1
const __FLT32_IS_IEC_60559__ = 1
const __FLT32_MANT_DIG__ = 24
const __FLT32_MAX_10_EXP__ = 38
const __FLT32_MAX_EXP__ = 128
const __FLT32_MAX__ = 3.40282346638528859811704183484516925e+38
const __FLT32_MIN__ = 1.17549435082228750796873653722224568e-38
const __FLT32_NORM_MAX__ = 3.40282346638528859811704183484516925e+38
const __FLT64X_DECIMAL_DIG__ = 36
const __FLT64X_DENORM_MIN__ = 6.47517511943802511092443895822764655e-4966
const __FLT64X_DIG__ = 33
const __FLT64X_EPSILON__ = 1.92592994438723585305597794258492732e-34
const __FLT64X_HAS_DENORM__ = 1
const __FLT64X_HAS_INFINITY__ = 1
const __FLT64X_HAS_QUIET_NAN__ = 1
const __FLT64X_IS_IEC_60559__ = 1
const __FLT64X_MANT_DIG__ = 113
const __FLT64X_MAX_10_EXP__ = 4932
const __FLT64X_MAX_EXP__ = 16384
const __FLT64X_MAX__ = "1.18973149535723176508575932662800702e+4932"
const __FLT64X_MIN__ = 3.36210314311209350626267781732175260e-4932
const __FLT64X_NORM_MAX__ = "1.18973149535723176508575932662800702e+4932"
const __FLT64_DECIMAL_DIG__ = 17
const __FLT64_DENORM_MIN__ = 4.94065645841246544176568792868221372e-324
const __FLT64_DIG__ = 15
const __FLT64_EPSILON__ = 2.22044604925031308084726333618164062e-16
const __FLT64_HAS_DENORM__ = 1
const __FLT64_HAS_INFINITY__ = 1
const __FLT64_HAS_QUIET_NAN__ = 1
const __FLT64_IS_IEC_60559__ = 1
const __FLT64_MANT_DIG__ = 53
const __FLT64_MAX_10_EXP__ = 308
const __FLT64_MAX_EXP__ = 1024
const __FLT64_MAX__ = 1.79769313486231570814527423731704357e+308
const __FLT64_MIN__ = 2.22507385850720138309023271733240406e-308
const __FLT64_NORM_MAX__ = 1.79769313486231570814527423731704357e+308
const __FLT_DECIMAL_DIG__ = 9
const __FLT_DENORM_MIN__ = 1.40129846432481707092372958328991613e-45
const __FLT_DIG__ = 6
const __FLT_EPSILON__ = 1.19209289550781250000000000000000000e-7
const __FLT_EVAL_METHOD_C99__ = 0
const __FLT_EVAL_METHOD_TS_18661_3__ = 0
const __FLT_EVAL_METHOD__ = 0
const __FLT_HAS_DENORM__ = 1
const __FLT_HAS_INFINITY__ = 1
const __FLT_HAS_QUIET_NAN__ = 1
const __FLT_IS_IEC_60559__ = 1
const __FLT_MANT_DIG__ = 24
const __FLT_MAX_10_EXP__ = 38
const __FLT_MAX_EXP__ = 128
const __FLT_MAX__ = 3.40282346638528859811704183484516925e+38
const __FLT_MIN__ = 1.17549435082228750796873653722224568e-38
const __FLT_NORM_MAX__ = 3.40282346638528859811704183484516925e+38
const __FLT_RADIX__ = 2
const __FP_FAST_FMA = 1
const __FP_FAST_FMAF = 1
const __FP_FAST_FMAF32 = 1
const __FP_FAST_FMAF32x = 1
const __FP_FAST_FMAF64 = 1
const __FUNCTION__ = "__func__"
const __GCC_ASM_FLAG_OUTPUTS__ = 1
const __GCC_ATOMIC_BOOL_LOCK_FREE = 2
const __GCC_ATOMIC_CHAR16_T_LOCK_FREE = 2
const __GCC_ATOMIC_CHAR32_T_LOCK_FREE = 2
const __GCC_ATOMIC_CHAR8_T_LOCK_FREE = 2
const __GCC_ATOMIC_CHAR_LOCK_FREE = 2
const __GCC_ATOMIC_INT_LOCK_FREE = 2
const __GCC_ATOMIC_LLONG_LOCK_FREE = 2
const __GCC_ATOMIC_LONG_LOCK_FREE = 2
const __GCC_ATOMIC_POINTER_LOCK_FREE = 2
const __GCC_ATOMIC_SHORT_LOCK_FREE = 2
const __GCC_ATOMIC_TEST_AND_SET_TRUEVAL = 1
const __GCC_ATOMIC_WCHAR_T_LOCK_FREE = 2
const __GCC_CONSTRUCTIVE_SIZE = 64
const __GCC_DESTRUCTIVE_SIZE = 256
const __GCC_HAVE_DWARF2_CFI_ASM = 1
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_1 = 1
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_16 = 1
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_2 = 1
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_4 = 1
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_8 = 1
const __GCC_IEC_559 = 2
const __GCC_IEC_559_COMPLEX = 2
const __GNUC_EXECUTION_CHARSET_NAME = "UTF-8"
const __GNUC_MINOR__ = 2
const __GNUC_PATCHLEVEL__ = 0
const __GNUC_STDC_INLINE__ = 1
const __GNUC_WIDE_EXECUTION_CHARSET_NAME = "UTF-32LE"
const __GNUC__ = 16
const __GXX_ABI_VERSION = 1021
const __HAVE_FUNCTION_MULTI_VERSIONING = 1
const __HAVE_SPECULATION_SAFE_VALUE = 1
const __INT16_MAX__ = 0x7fff
const __INT32_MAX__ = 0x7fffffff
const __INT32_TYPE__ = "int"
const __INT64_MAX__ = 0x7fffffffffffffff
const __INT8_MAX__ = 0x7f
const __INTMAX_MAX__ = 0x7fffffffffffffff
const __INTMAX_WIDTH__ = 64
const __INTPTR_MAX__ = 0x7fffffffffffffff
const __INTPTR_WIDTH__ = 64
const __INT_FAST16_MAX__ = 0x7fffffffffffffff
const __INT_FAST16_WIDTH__ = 64
const __INT_FAST32_MAX__ = 0x7fffffffffffffff
const __INT_FAST32_WIDTH__ = 64
const __INT_FAST64_MAX__ = 0x7fffffffffffffff
const __INT_FAST64_WIDTH__ = 64
const __INT_FAST8_MAX__ = 0x7f
const __INT_FAST8_WIDTH__ = 8
const __INT_LEAST16_MAX__ = 0x7fff
const __INT_LEAST16_WIDTH__ = 16
const __INT_LEAST32_MAX__ = 0x7fffffff
const __INT_LEAST32_TYPE__ = "int"
const __INT_LEAST32_WIDTH__ = 32
const __INT_LEAST64_MAX__ = 0x7fffffffffffffff
const __INT_LEAST64_WIDTH__ = 64
const __INT_LEAST8_MAX__ = 0x7f
const __INT_LEAST8_WIDTH__ = 8
const __INT_MAX__ = 0x7fffffff
const __INT_WIDTH__ = 32
const __LDBL_DECIMAL_DIG__ = 36
const __LDBL_DENORM_MIN__ = 6.47517511943802511092443895822764655e-4966
const __LDBL_DIG__ = 33
const __LDBL_EPSILON__ = 1.92592994438723585305597794258492732e-34
const __LDBL_HAS_DENORM__ = 1
const __LDBL_HAS_INFINITY__ = 1
const __LDBL_HAS_QUIET_NAN__ = 1
const __LDBL_IS_IEC_60559__ = 1
const __LDBL_MANT_DIG__ = 113
const __LDBL_MAX_10_EXP__ = 4932
const __LDBL_MAX_EXP__ = 16384
const __LDBL_MAX__ = "1.18973149535723176508575932662800702e+4932"
const __LDBL_MIN__ = 3.36210314311209350626267781732175260e-4932
const __LDBL_NORM_MAX__ = "1.18973149535723176508575932662800702e+4932"
const __LITTLE_ENDIAN = 1234
const __LONG_LONG_MAX__ = 0x7fffffffffffffff
const __LONG_LONG_WIDTH__ = 64
const __LONG_MAX = 0x7fffffffffffffff
const __LONG_MAX__ = 0x7fffffffffffffff
const __LONG_WIDTH__ = 64
const __LP64__ = 1
const __OPTIMIZE__ = 1
const __ORDER_BIG_ENDIAN__ = 4321
const __ORDER_LITTLE_ENDIAN__ = 1234
const __ORDER_PDP_ENDIAN__ = 3412
const __PIC__ = 2
const __PRAGMA_REDEFINE_EXTNAME = 1
const __PRETTY_FUNCTION__ = "__func__"
const __PTRDIFF_MAX__ = 0x7fffffffffffffff
const __PTRDIFF_WIDTH__ = 64
const __SCHAR_MAX__ = 0x7f
const __SCHAR_WIDTH__ = 8
const __SHRT_MAX__ = 0x7fff
const __SHRT_WIDTH__ = 16
const __SIG_ATOMIC_MAX__ = 0x7fffffff
const __SIG_ATOMIC_TYPE__ = "int"
const __SIG_ATOMIC_WIDTH__ = 32
const __SIZEOF_DOUBLE__ = 8
const __SIZEOF_FLOAT__ = 4
const __SIZEOF_INT128__ = 16
const __SIZEOF_INT__ = 4
const __SIZEOF_LONG_DOUBLE__ = 8
const __SIZEOF_LONG_LONG__ = 8
const __SIZEOF_LONG__ = 8
const __SIZEOF_POINTER__ = 8
const __SIZEOF_PTRDIFF_T__ = 8
const __SIZEOF_SHORT__ = 2
const __SIZEOF_SIZE_T__ = 8
const __SIZEOF_WCHAR_T__ = 4
const __SIZEOF_WINT_T__ = 4
const __SIZE_MAX__ = 0xffffffffffffffff
const __SIZE_WIDTH__ = 64
const __SSP_STRONG__ = 3
const __STDC_EMBED_EMPTY__ = 2
const __STDC_EMBED_FOUND__ = 1
const __STDC_EMBED_NOT_FOUND__ = 0
const __STDC_HOSTED__ = 1
const __STDC_IEC_559_COMPLEX__ = 1
const __STDC_IEC_559__ = 1
const __STDC_IEC_60559_BFP__ = 201404
const __STDC_IEC_60559_COMPLEX__ = 201404
const __STDC_ISO_10646__ = 201706
const __STDC_UTF_16__ = 1
const __STDC_UTF_32__ = 1
const __STDC_VERSION__ = 202311
const __STDC__ = 1
const __UINT16_MAX__ = 0xffff
const __UINT32_MAX__ = 0xffffffff
const __UINT64_MAX__ = 0xffffffffffffffff
const __UINT8_MAX__ = 0xff
const __UINTMAX_MAX__ = 0xffffffffffffffff
const __UINTPTR_MAX__ = 0xffffffffffffffff
const __UINT_FAST16_MAX__ = 0xffffffffffffffff
const __UINT_FAST32_MAX__ = 0xffffffffffffffff
const __UINT_FAST64_MAX__ = 0xffffffffffffffff
const __UINT_FAST8_MAX__ = 0xff
const __UINT_LEAST16_MAX__ = 0xffff
const __UINT_LEAST32_MAX__ = 0xffffffff
const __UINT_LEAST64_MAX__ = 0xffffffffffffffff
const __UINT_LEAST8_MAX__ = 0xff
const __USE_TIME_BITS64 = 1
const __VERSION__ = "16.2.0"
const __WCHAR_MAX__ = 0xffffffff
const __WCHAR_MIN__ = 0
const __WCHAR_WIDTH__ = 32
const __WINT_MAX__ = 0xffffffff
const __WINT_MIN__ = 0
const __WINT_WIDTH__ = 32
const __aarch64__ = 1
const __gnu_linux__ = 1
const __inline = "inline"
const __linux = 1
const __linux__ = 1
const __pic__ = 2
const __restrict = "restrict"
const __restrict_arr = "restrict"
const __unix = 1
const __unix__ = 1
const alloca1 = "__builtin_alloca"
const false1 = 0
const linux = 1
const md_strchr = "strchr"
const true1 = 1
const unix = 1

/* C23 keywords, no longer defined as macros by <stdbool.h> and <stddef.h>. */

type __builtin_va_list = uintptr

type __predefined_size_t = uint64

type __predefined_wchar_t = uint32

type __predefined_ptrdiff_t = int64

type MD_CHAR = uint8

type MD_SIZE = uint32

type MD_OFFSET = uint32

type MD_BLOCKTYPE1 = int32

type MD_BLOCKTYPE = int32

const MD_BLOCK_DOC = 0
const MD_BLOCK_QUOTE = 1
const MD_BLOCK_UL = 2
const MD_BLOCK_OL = 3
const MD_BLOCK_LI = 4
const MD_BLOCK_HR = 5
const MD_BLOCK_H = 6
const MD_BLOCK_CODE = 7
const MD_BLOCK_HTML = 8
const MD_BLOCK_P = 9
const MD_BLOCK_TABLE = 10
const MD_BLOCK_THEAD = 11
const MD_BLOCK_TBODY = 12
const MD_BLOCK_TR = 13
const MD_BLOCK_TH = 14
const MD_BLOCK_TD = 15

type MD_SPANTYPE1 = int32

type MD_SPANTYPE = int32

const MD_SPAN_EM = 0
const MD_SPAN_STRONG = 1
const MD_SPAN_A = 2
const MD_SPAN_IMG = 3
const MD_SPAN_CODE = 4
const MD_SPAN_DEL = 5
const MD_SPAN_LATEXMATH = 6
const MD_SPAN_LATEXMATH_DISPLAY = 7
const MD_SPAN_WIKILINK1 = 8
const MD_SPAN_U = 9

type MD_TEXTTYPE1 = int32

type MD_TEXTTYPE = int32

const MD_TEXT_NORMAL = 0
const MD_TEXT_NULLCHAR = 1
const MD_TEXT_BR = 2
const MD_TEXT_SOFTBR = 3
const MD_TEXT_ENTITY = 4
const MD_TEXT_CODE = 5
const MD_TEXT_HTML = 6
const MD_TEXT_LATEXMATH = 7

type MD_ALIGN1 = int32

type MD_ALIGN = int32

const MD_ALIGN_DEFAULT = 0
const MD_ALIGN_LEFT = 1
const MD_ALIGN_CENTER = 2
const MD_ALIGN_RIGHT = 3

type MD_ATTRIBUTE = struct {
	Ftext           uintptr
	Fsize           MD_SIZE
	Fsubstr_types   uintptr
	Fsubstr_offsets uintptr
}

type MD_BLOCK_UL_DETAIL = struct {
	Fis_tight int32
	Fmark     MD_CHAR
}

type MD_BLOCK_OL_DETAIL = struct {
	Fstart          uint32
	Fis_tight       int32
	Fmark_delimiter MD_CHAR
}

type MD_BLOCK_LI_DETAIL = struct {
	Fis_task          int32
	Ftask_mark        MD_CHAR
	Ftask_mark_offset MD_OFFSET
}

type MD_BLOCK_H_DETAIL = struct {
	Flevel uint32
}

type MD_BLOCK_CODE_DETAIL = struct {
	Finfo       MD_ATTRIBUTE
	Flang       MD_ATTRIBUTE
	Ffence_char MD_CHAR
}

type MD_BLOCK_TABLE_DETAIL = struct {
	Fcol_count      uint32
	Fhead_row_count uint32
	Fbody_row_count uint32
}

type MD_BLOCK_TD_DETAIL = struct {
	Falign MD_ALIGN1
}

type MD_SPAN_A_DETAIL = struct {
	Fhref        MD_ATTRIBUTE
	Ftitle       MD_ATTRIBUTE
	Fis_autolink int32
}

type MD_SPAN_IMG_DETAIL = struct {
	Fsrc   MD_ATTRIBUTE
	Ftitle MD_ATTRIBUTE
}

type MD_SPAN_WIKILINK_DETAIL = struct {
	Ftarget MD_ATTRIBUTE
}

type MD_SPAN_WIKILINK = MD_SPAN_WIKILINK_DETAIL

type MD_PARSER = struct {
	Fabi_version uint32
	Fflags       uint32
	Fenter_block uintptr
	Fleave_block uintptr
	Fenter_span  uintptr
	Fleave_span  uintptr
	Ftext        uintptr
	Fdebug_log   uintptr
	Fsyntax      uintptr
}

type MD_RENDERER = struct {
	Fabi_version uint32
	Fflags       uint32
	Fenter_block uintptr
	Fleave_block uintptr
	Fenter_span  uintptr
	Fleave_span  uintptr
	Ftext        uintptr
	Fdebug_log   uintptr
	Fsyntax      uintptr
}

type size_t = uint64

type ssize_t = int64

type off_t = int64

type va_list = uintptr

type __isoc_va_list = uintptr

type fpos_t = struct {
	F__lldata [0]int64
	F__align  [0]float64
	F__opaque [16]uint8
}

type _G_fpos64_t = fpos_t

type cookie_io_functions_t = struct {
	Fread   uintptr
	Fwrite  uintptr
	Fseek   uintptr
	Fclose1 uintptr
}

type _IO_cookie_io_functions_t = cookie_io_functions_t

type wchar_t = uint32

type div_t = struct {
	Fquot int32
	Frem  int32
}

type ldiv_t = struct {
	Fquot int64
	Frem  int64
}

type lldiv_t = struct {
	Fquot int64
	Frem  int64
}

type locale_t = uintptr

/*****************************
 ***  Miscellaneous Stuff  ***
 *****************************/

/* Make the UTF-8 support the default. */

/* Magic for making wide literals with MD4C_USE_UTF16. */

/* Misc. macros. */

/* For falling through case labels in switch statements. */

/* Suppress "unused parameter" warnings. */

/******************************
 ***  Some internal limits  ***
 ******************************/

/* We limit code span marks to lower than 32 backticks. This solves the
 * pathologic case of too many openers, each of different length: Their
 * resolving would be then O(n^2). */

/* We limit column count of tables to prevent quadratic explosion of output
 * from pathological input of a table thousands of columns and thousands
 * of rows where rows are requested with as little as single character
 * per-line, relying on us to "helpfully" fill all the missing "<td></td>". */

/************************
 ***  Internal Types  ***
 ************************/

/* These are omnipresent so lets save some typing. */

type MD_MARK = struct {
	Fbeg   MD_OFFSET
	Fend   MD_OFFSET
	Fprev  int32
	Fnext  int32
	Fch    MD_CHAR
	Fflags uint8
}

/*****************************
 ***  Miscellaneous Stuff  ***
 *****************************/

/* Make the UTF-8 support the default. */

/* Magic for making wide literals with MD4C_USE_UTF16. */

/* Misc. macros. */

/* For falling through case labels in switch statements. */

/* Suppress "unused parameter" warnings. */

/******************************
 ***  Some internal limits  ***
 ******************************/

/* We limit code span marks to lower than 32 backticks. This solves the
 * pathologic case of too many openers, each of different length: Their
 * resolving would be then O(n^2). */

/* We limit column count of tables to prevent quadratic explosion of output
 * from pathological input of a table thousands of columns and thousands
 * of rows where rows are requested with as little as single character
 * per-line, relying on us to "helpfully" fill all the missing "<td></td>". */

/************************
 ***  Internal Types  ***
 ************************/

/* These are omnipresent so lets save some typing. */

type MD_MARK_tag = MD_MARK

type MD_BLOCK = struct {
	F__ccgo0 uint32
	Fn_lines MD_SIZE
}

type MD_BLOCK_tag = MD_BLOCK

type MD_CONTAINER = struct {
	Fch              MD_CHAR
	F__ccgo_align1   [2]byte
	F__ccgo4         uint16
	Fstart           uint32
	Fmark_indent     uint32
	Fcontents_indent uint32
	Fblock_byte_off  MD_OFFSET
	Ftask_mark_off   MD_OFFSET
}

type MD_CONTAINER_tag = MD_CONTAINER

type MD_REF_DEF = struct {
	Flabel      uintptr
	Ftitle      uintptr
	Fhash       uint32
	Flabel_size MD_SIZE
	Ftitle_size MD_SIZE
	Fdest_beg   MD_OFFSET
	Fdest_end   MD_OFFSET
	F__ccgo36   uint8
}

type MD_REF_DEF_tag = MD_REF_DEF

// C documentation
//
//	/* During analyzes of inline marks, we need to manage stacks of unresolved
//	 * openers of the given type.
//	 * The stack connects the marks via MD_MARK::next;
//	 */
type MD_MARKSTACK = struct {
	Ftop int32
}

// C documentation
//
//	/* During analyzes of inline marks, we need to manage stacks of unresolved
//	 * openers of the given type.
//	 * The stack connects the marks via MD_MARK::next;
//	 */
type MD_MARKSTACK_tag = MD_MARKSTACK

// C documentation
//
//	/* Context propagated through all the parsing. */
type MD_CTX = struct {
	Ftext                                       uintptr
	Fsize                                       MD_SIZE
	Fparser                                     MD_PARSER
	Fuserdata                                   uintptr
	Fdoc_ends_with_newline                      int32
	Fbuffer                                     uintptr
	Falloc_buffer                               uint32
	Fref_defs                                   uintptr
	Fn_ref_defs                                 int32
	Falloc_ref_defs                             int32
	Fref_def_hashtable                          uintptr
	Fref_def_hashtable_size                     int32
	Fmarks                                      uintptr
	Fn_marks                                    int32
	Falloc_marks                                int32
	Fmark_char_map                              [256]uint8
	Fopener_stacks                              [16]MD_MARKSTACK
	Fptr_stack                                  MD_MARKSTACK
	Fn_table_cell_boundaries                    int32
	Ftable_cell_boundaries_head                 int32
	Ftable_cell_boundaries_tail                 int32
	Funresolved_link_head                       int32
	Funresolved_link_tail                       int32
	Fhtml_comment_horizon                       MD_OFFSET
	Fhtml_proc_instr_horizon                    MD_OFFSET
	Fhtml_decl_horizon                          MD_OFFSET
	Fhtml_cdata_horizon                         MD_OFFSET
	Fblock_bytes                                uintptr
	Fcurrent_block                              uintptr
	Fn_block_bytes                              int32
	Falloc_block_bytes                          int32
	Fcontainers                                 uintptr
	Fn_containers                               int32
	Falloc_containers                           int32
	Fcode_indent_offset                         uint32
	Fcode_fence_length                          MD_SIZE
	Fhtml_block_type                            int32
	Flast_line_has_list_loosening_effect        int32
	Flast_list_item_starts_with_two_blank_lines int32
}

// C documentation
//
//	/* Context propagated through all the parsing. */
type MD_CTX_tag = MD_CTX

type MD_LINETYPE_tag = int32

const MD_LINE_BLANK = 0
const MD_LINE_HR = 1
const MD_LINE_ATXHEADER = 2
const MD_LINE_SETEXTHEADER = 3
const MD_LINE_SETEXTUNDERLINE = 4
const MD_LINE_INDENTEDCODE = 5
const MD_LINE_FENCEDCODE = 6
const MD_LINE_HTML = 7
const MD_LINE_TEXT = 8
const MD_LINE_TABLE = 9
const MD_LINE_TABLEUNDERLINE = 10

type MD_LINETYPE = int32

type MD_LINE_ANALYSIS = struct {
	Ftype1             MD_LINETYPE
	Fdata              uint32
	Fenforce_new_block int32
	Fbeg               MD_OFFSET
	Fend               MD_OFFSET
	Findent            uint32
}

type MD_LINE_ANALYSIS_tag = MD_LINE_ANALYSIS

type MD_LINE = struct {
	Fbeg MD_OFFSET
	Fend MD_OFFSET
}

type MD_LINE_tag = MD_LINE

type MD_VERBATIMLINE = struct {
	Fbeg    MD_OFFSET
	Fend    MD_OFFSET
	Findent MD_OFFSET
}

type MD_VERBATIMLINE_tag = MD_VERBATIMLINE

/*****************
 ***  Helpers  ***
 *****************/

/* Character accessors. */

/* Character classification.
 * Note we assume ASCII compatibility of code points < 128 here. */

// C documentation
//
//	/* Case insensitive check of string equality. */
func md_ascii_case_eq(tls *libc.TLS, s1 uintptr, s2 uintptr, n MD_SIZE) (r int32) {
	var ch1, ch2 MD_CHAR
	var i MD_OFFSET
	_, _, _ = ch1, ch2, i
	i = uint32(0)
	for {
		if !(i < n) {
			break
		}
		ch1 = *(*MD_CHAR)(__ccgo_ptr(s1 + uintptr(i)))
		ch2 = *(*MD_CHAR)(__ccgo_ptr(s2 + uintptr(i)))
		if libc.Uint32FromInt32(libc.Int32FromUint8('a')) <= uint32(ch1) && uint32(ch1) <= libc.Uint32FromInt32(libc.Int32FromUint8('z')) {
			ch1 = libc.Uint8FromInt32(int32(ch1) + (libc.Int32FromUint8('A') - libc.Int32FromUint8('a')))
		}
		if libc.Uint32FromInt32(libc.Int32FromUint8('a')) <= uint32(ch2) && uint32(ch2) <= libc.Uint32FromInt32(libc.Int32FromUint8('z')) {
			ch2 = libc.Uint8FromInt32(int32(ch2) + (libc.Int32FromUint8('A') - libc.Int32FromUint8('a')))
		}
		if libc.Int32FromUint8(ch1) != libc.Int32FromUint8(ch2) {
			return FALSE
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	return int32(TRUE)
}

func md_ascii_eq(tls *libc.TLS, s1 uintptr, s2 uintptr, n MD_SIZE) (r int32) {
	return libc.BoolInt32(libc.Xmemcmp(tls, s1, s2, uint64(n)*uint64(1)) == 0)
}

func md_text_with_null_replacement(tls *libc.TLS, ctx uintptr, type1 MD_TEXTTYPE1, str uintptr, size MD_SIZE) (r int32) {
	var off MD_OFFSET
	var ret int32
	_, _ = off, ret
	off = uint32(0)
	ret = 0
	for int32(1) != 0 {
		for off < size && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr(str + uintptr(off)))) != int32('\000') {
			off = off + 1
		}
		if off > uint32(0) {
			ret = (*(*func(*libc.TLS, MD_TEXTTYPE1, uintptr, MD_SIZE, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Ftext})))(tls, type1, str, off, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
			if ret != 0 {
				return ret
			}
			str = str + uintptr(off)
			size = size - off
			off = uint32(0)
		}
		if off >= size {
			return 0
		}
		ret = (*(*func(*libc.TLS, MD_TEXTTYPE1, uintptr, MD_SIZE, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Ftext})))(tls, int32(MD_TEXT_NULLCHAR), __ccgo_ts, uint32(1), (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
		if ret != 0 {
			return ret
		}
		off = off + 1
	}
	return r
}

// C documentation
//
//	/* If the offset falls into a gap between line, we return the following
//	 * line. */
func md_lookup_line(tls *libc.TLS, off MD_OFFSET, lines uintptr, n_lines MD_SIZE, p_line_index uintptr) (r uintptr) {
	var hi, lo, pivot MD_SIZE
	var line uintptr
	_, _, _, _ = hi, line, lo, pivot
	lo = uint32(0)
	hi = n_lines - uint32(1)
	for lo <= hi {
		pivot = (lo + hi) / uint32(2)
		line = lines + uintptr(pivot)*8
		if off < (*MD_LINE)(unsafe.Pointer(line)).Fbeg {
			if hi == uint32(0) || (*(*MD_LINE)(__ccgo_ptr(lines + uintptr(hi-uint32(1))*8))).Fend < off {
				if p_line_index != libc.UintptrFromInt32(0) {
					*(*MD_SIZE)(__ccgo_ptr(p_line_index)) = pivot
				}
				return line
			}
			hi = pivot - uint32(1)
		} else {
			if off > (*MD_LINE)(unsafe.Pointer(line)).Fend {
				lo = pivot + uint32(1)
			} else {
				if p_line_index != libc.UintptrFromInt32(0) {
					*(*MD_SIZE)(__ccgo_ptr(p_line_index)) = pivot
				}
				return line
			}
		}
	}
	return libc.UintptrFromInt32(0)
}

/*************************
 ***  Unicode Support  ***
 *************************/

type MD_UNICODE_FOLD_INFO = struct {
	Fcodepoints   [3]uint32
	Fn_codepoints uint32
}

/*************************
 ***  Unicode Support  ***
 *************************/

type MD_UNICODE_FOLD_INFO_tag = MD_UNICODE_FOLD_INFO

// C documentation
//
//	/* Binary search over sorted "map" of codepoints. Consecutive sequences
//	     * of codepoints may be encoded in the map by just using the
//	     * (MIN_CODEPOINT | 0x40000000) and (MAX_CODEPOINT | 0x80000000).
//	     *
//	     * Returns index of the found record in the map (in the case of ranges,
//	     * the minimal value is used); or -1 on failure. */
func md_unicode_bsearch__(tls *libc.TLS, codepoint uint32, map1 uintptr, map_size size_t) (r int32) {
	var beg, end, pivot_beg, pivot_end, v1 int32
	_, _, _, _, _ = beg, end, pivot_beg, pivot_end, v1
	beg = 0
	end = libc.Int32FromUint64(map_size) - int32(1)
	for beg <= end {
		/* Pivot may be a range, not just a single value. */
		v1 = (beg + end) / libc.Int32FromInt32(2)
		pivot_end = v1
		pivot_beg = v1
		if *(*uint32)(__ccgo_ptr(map1 + uintptr(pivot_end)*4))&uint32(0x40000000) != 0 {
			pivot_end = pivot_end + 1
		}
		if *(*uint32)(__ccgo_ptr(map1 + uintptr(pivot_beg)*4))&uint32(0x80000000) != 0 {
			pivot_beg = pivot_beg - 1
		}
		if codepoint < *(*uint32)(__ccgo_ptr(map1 + uintptr(pivot_beg)*4))&uint32(0x00ffffff) {
			end = pivot_beg - int32(1)
		} else {
			if codepoint > *(*uint32)(__ccgo_ptr(map1 + uintptr(pivot_end)*4))&uint32(0x00ffffff) {
				beg = pivot_end + int32(1)
			} else {
				return pivot_beg
			}
		}
	}
	return -int32(1)
}

func md_is_unicode_whitespace__(tls *libc.TLS, codepoint uint32) (r int32) {
	/* The ASCII ones are the most frequently used ones, also CommonMark
	 * specification requests few more in this range. */
	if codepoint <= uint32(0x7f) {
		return libc.BoolInt32(codepoint == libc.Uint32FromInt32(libc.Int32FromUint8(' ')) || codepoint == libc.Uint32FromInt32(libc.Int32FromUint8('\t')) || (codepoint == libc.Uint32FromInt32(libc.Int32FromUint8('\v')) || codepoint == libc.Uint32FromInt32(libc.Int32FromUint8('\f'))))
	}
	return libc.BoolInt32(md_unicode_bsearch__(tls, codepoint, uintptr(unsafe.Pointer(&WHITESPACE_MAP)), libc.Uint64FromInt64(32)/libc.Uint64FromInt64(4)) >= 0)
}

/* Unicode "Zs" category.
 * (generated by scripts/build_whitespace_map.py) */
var WHITESPACE_MAP = [8]uint32{
	0: libc.Uint32FromInt32(libc.Int32FromInt32(0x0020)),
	1: libc.Uint32FromInt32(libc.Int32FromInt32(0x00a0)),
	2: libc.Uint32FromInt32(libc.Int32FromInt32(0x1680)),
	3: libc.Uint32FromInt32(libc.Int32FromInt32(0x2000) | libc.Int32FromInt32(0x40000000)),
	4: libc.Uint32FromInt32(libc.Int32FromInt32(0x200a)) | libc.Uint32FromUint32(0x80000000),
	5: libc.Uint32FromInt32(libc.Int32FromInt32(0x202f)),
	6: libc.Uint32FromInt32(libc.Int32FromInt32(0x205f)),
	7: libc.Uint32FromInt32(libc.Int32FromInt32(0x3000)),
}

func md_is_unicode_punct__(tls *libc.TLS, codepoint uint32) (r int32) {
	/* The ASCII ones are the most frequently used ones, also CommonMark
	 * specification requests few more in this range. */
	if codepoint <= uint32(0x7f) {
		return libc.BoolInt32(libc.Uint32FromInt32(libc.Int32FromInt32(33)) <= codepoint && codepoint <= libc.Uint32FromInt32(libc.Int32FromInt32(47)) || libc.Uint32FromInt32(libc.Int32FromInt32(58)) <= codepoint && codepoint <= libc.Uint32FromInt32(libc.Int32FromInt32(64)) || libc.Uint32FromInt32(libc.Int32FromInt32(91)) <= codepoint && codepoint <= libc.Uint32FromInt32(libc.Int32FromInt32(96)) || libc.Uint32FromInt32(libc.Int32FromInt32(123)) <= codepoint && codepoint <= libc.Uint32FromInt32(libc.Int32FromInt32(126)))
	}
	return libc.BoolInt32(md_unicode_bsearch__(tls, codepoint, uintptr(unsafe.Pointer(&PUNCT_MAP)), libc.Uint64FromInt64(2304)/libc.Uint64FromInt64(4)) >= 0)
}

/* Unicode general "P" and "S" categories.
 * (generated by scripts/build_punct_map.py) */
var PUNCT_MAP = [576]uint32{
	0:   libc.Uint32FromInt32(libc.Int32FromInt32(0x0021) | libc.Int32FromInt32(0x40000000)),
	1:   libc.Uint32FromInt32(libc.Int32FromInt32(0x002f)) | libc.Uint32FromUint32(0x80000000),
	2:   libc.Uint32FromInt32(libc.Int32FromInt32(0x003a) | libc.Int32FromInt32(0x40000000)),
	3:   libc.Uint32FromInt32(libc.Int32FromInt32(0x0040)) | libc.Uint32FromUint32(0x80000000),
	4:   libc.Uint32FromInt32(libc.Int32FromInt32(0x005b) | libc.Int32FromInt32(0x40000000)),
	5:   libc.Uint32FromInt32(libc.Int32FromInt32(0x0060)) | libc.Uint32FromUint32(0x80000000),
	6:   libc.Uint32FromInt32(libc.Int32FromInt32(0x007b) | libc.Int32FromInt32(0x40000000)),
	7:   libc.Uint32FromInt32(libc.Int32FromInt32(0x007e)) | libc.Uint32FromUint32(0x80000000),
	8:   libc.Uint32FromInt32(libc.Int32FromInt32(0x00a1) | libc.Int32FromInt32(0x40000000)),
	9:   libc.Uint32FromInt32(libc.Int32FromInt32(0x00a9)) | libc.Uint32FromUint32(0x80000000),
	10:  libc.Uint32FromInt32(libc.Int32FromInt32(0x00ab) | libc.Int32FromInt32(0x40000000)),
	11:  libc.Uint32FromInt32(libc.Int32FromInt32(0x00ac)) | libc.Uint32FromUint32(0x80000000),
	12:  libc.Uint32FromInt32(libc.Int32FromInt32(0x00ae) | libc.Int32FromInt32(0x40000000)),
	13:  libc.Uint32FromInt32(libc.Int32FromInt32(0x00b1)) | libc.Uint32FromUint32(0x80000000),
	14:  libc.Uint32FromInt32(libc.Int32FromInt32(0x00b4)),
	15:  libc.Uint32FromInt32(libc.Int32FromInt32(0x00b6) | libc.Int32FromInt32(0x40000000)),
	16:  libc.Uint32FromInt32(libc.Int32FromInt32(0x00b8)) | libc.Uint32FromUint32(0x80000000),
	17:  libc.Uint32FromInt32(libc.Int32FromInt32(0x00bb)),
	18:  libc.Uint32FromInt32(libc.Int32FromInt32(0x00bf)),
	19:  libc.Uint32FromInt32(libc.Int32FromInt32(0x00d7)),
	20:  libc.Uint32FromInt32(libc.Int32FromInt32(0x00f7)),
	21:  libc.Uint32FromInt32(libc.Int32FromInt32(0x02c2) | libc.Int32FromInt32(0x40000000)),
	22:  libc.Uint32FromInt32(libc.Int32FromInt32(0x02c5)) | libc.Uint32FromUint32(0x80000000),
	23:  libc.Uint32FromInt32(libc.Int32FromInt32(0x02d2) | libc.Int32FromInt32(0x40000000)),
	24:  libc.Uint32FromInt32(libc.Int32FromInt32(0x02df)) | libc.Uint32FromUint32(0x80000000),
	25:  libc.Uint32FromInt32(libc.Int32FromInt32(0x02e5) | libc.Int32FromInt32(0x40000000)),
	26:  libc.Uint32FromInt32(libc.Int32FromInt32(0x02eb)) | libc.Uint32FromUint32(0x80000000),
	27:  libc.Uint32FromInt32(libc.Int32FromInt32(0x02ed)),
	28:  libc.Uint32FromInt32(libc.Int32FromInt32(0x02ef) | libc.Int32FromInt32(0x40000000)),
	29:  libc.Uint32FromInt32(libc.Int32FromInt32(0x02ff)) | libc.Uint32FromUint32(0x80000000),
	30:  libc.Uint32FromInt32(libc.Int32FromInt32(0x0375)),
	31:  libc.Uint32FromInt32(libc.Int32FromInt32(0x037e)),
	32:  libc.Uint32FromInt32(libc.Int32FromInt32(0x0384) | libc.Int32FromInt32(0x40000000)),
	33:  libc.Uint32FromInt32(libc.Int32FromInt32(0x0385)) | libc.Uint32FromUint32(0x80000000),
	34:  libc.Uint32FromInt32(libc.Int32FromInt32(0x0387)),
	35:  libc.Uint32FromInt32(libc.Int32FromInt32(0x03f6)),
	36:  libc.Uint32FromInt32(libc.Int32FromInt32(0x0482)),
	37:  libc.Uint32FromInt32(libc.Int32FromInt32(0x055a) | libc.Int32FromInt32(0x40000000)),
	38:  libc.Uint32FromInt32(libc.Int32FromInt32(0x055f)) | libc.Uint32FromUint32(0x80000000),
	39:  libc.Uint32FromInt32(libc.Int32FromInt32(0x0589) | libc.Int32FromInt32(0x40000000)),
	40:  libc.Uint32FromInt32(libc.Int32FromInt32(0x058a)) | libc.Uint32FromUint32(0x80000000),
	41:  libc.Uint32FromInt32(libc.Int32FromInt32(0x058d) | libc.Int32FromInt32(0x40000000)),
	42:  libc.Uint32FromInt32(libc.Int32FromInt32(0x058f)) | libc.Uint32FromUint32(0x80000000),
	43:  libc.Uint32FromInt32(libc.Int32FromInt32(0x05be)),
	44:  libc.Uint32FromInt32(libc.Int32FromInt32(0x05c0)),
	45:  libc.Uint32FromInt32(libc.Int32FromInt32(0x05c3)),
	46:  libc.Uint32FromInt32(libc.Int32FromInt32(0x05c6)),
	47:  libc.Uint32FromInt32(libc.Int32FromInt32(0x05f3) | libc.Int32FromInt32(0x40000000)),
	48:  libc.Uint32FromInt32(libc.Int32FromInt32(0x05f4)) | libc.Uint32FromUint32(0x80000000),
	49:  libc.Uint32FromInt32(libc.Int32FromInt32(0x0606) | libc.Int32FromInt32(0x40000000)),
	50:  libc.Uint32FromInt32(libc.Int32FromInt32(0x060f)) | libc.Uint32FromUint32(0x80000000),
	51:  libc.Uint32FromInt32(libc.Int32FromInt32(0x061b)),
	52:  libc.Uint32FromInt32(libc.Int32FromInt32(0x061d) | libc.Int32FromInt32(0x40000000)),
	53:  libc.Uint32FromInt32(libc.Int32FromInt32(0x061f)) | libc.Uint32FromUint32(0x80000000),
	54:  libc.Uint32FromInt32(libc.Int32FromInt32(0x066a) | libc.Int32FromInt32(0x40000000)),
	55:  libc.Uint32FromInt32(libc.Int32FromInt32(0x066d)) | libc.Uint32FromUint32(0x80000000),
	56:  libc.Uint32FromInt32(libc.Int32FromInt32(0x06d4)),
	57:  libc.Uint32FromInt32(libc.Int32FromInt32(0x06de)),
	58:  libc.Uint32FromInt32(libc.Int32FromInt32(0x06e9)),
	59:  libc.Uint32FromInt32(libc.Int32FromInt32(0x06fd) | libc.Int32FromInt32(0x40000000)),
	60:  libc.Uint32FromInt32(libc.Int32FromInt32(0x06fe)) | libc.Uint32FromUint32(0x80000000),
	61:  libc.Uint32FromInt32(libc.Int32FromInt32(0x0700) | libc.Int32FromInt32(0x40000000)),
	62:  libc.Uint32FromInt32(libc.Int32FromInt32(0x070d)) | libc.Uint32FromUint32(0x80000000),
	63:  libc.Uint32FromInt32(libc.Int32FromInt32(0x07f6) | libc.Int32FromInt32(0x40000000)),
	64:  libc.Uint32FromInt32(libc.Int32FromInt32(0x07f9)) | libc.Uint32FromUint32(0x80000000),
	65:  libc.Uint32FromInt32(libc.Int32FromInt32(0x07fe) | libc.Int32FromInt32(0x40000000)),
	66:  libc.Uint32FromInt32(libc.Int32FromInt32(0x07ff)) | libc.Uint32FromUint32(0x80000000),
	67:  libc.Uint32FromInt32(libc.Int32FromInt32(0x0830) | libc.Int32FromInt32(0x40000000)),
	68:  libc.Uint32FromInt32(libc.Int32FromInt32(0x083e)) | libc.Uint32FromUint32(0x80000000),
	69:  libc.Uint32FromInt32(libc.Int32FromInt32(0x085e)),
	70:  libc.Uint32FromInt32(libc.Int32FromInt32(0x0888)),
	71:  libc.Uint32FromInt32(libc.Int32FromInt32(0x0964) | libc.Int32FromInt32(0x40000000)),
	72:  libc.Uint32FromInt32(libc.Int32FromInt32(0x0965)) | libc.Uint32FromUint32(0x80000000),
	73:  libc.Uint32FromInt32(libc.Int32FromInt32(0x0970)),
	74:  libc.Uint32FromInt32(libc.Int32FromInt32(0x09f2) | libc.Int32FromInt32(0x40000000)),
	75:  libc.Uint32FromInt32(libc.Int32FromInt32(0x09f3)) | libc.Uint32FromUint32(0x80000000),
	76:  libc.Uint32FromInt32(libc.Int32FromInt32(0x09fa) | libc.Int32FromInt32(0x40000000)),
	77:  libc.Uint32FromInt32(libc.Int32FromInt32(0x09fb)) | libc.Uint32FromUint32(0x80000000),
	78:  libc.Uint32FromInt32(libc.Int32FromInt32(0x09fd)),
	79:  libc.Uint32FromInt32(libc.Int32FromInt32(0x0a76)),
	80:  libc.Uint32FromInt32(libc.Int32FromInt32(0x0af0) | libc.Int32FromInt32(0x40000000)),
	81:  libc.Uint32FromInt32(libc.Int32FromInt32(0x0af1)) | libc.Uint32FromUint32(0x80000000),
	82:  libc.Uint32FromInt32(libc.Int32FromInt32(0x0b70)),
	83:  libc.Uint32FromInt32(libc.Int32FromInt32(0x0bf3) | libc.Int32FromInt32(0x40000000)),
	84:  libc.Uint32FromInt32(libc.Int32FromInt32(0x0bfa)) | libc.Uint32FromUint32(0x80000000),
	85:  libc.Uint32FromInt32(libc.Int32FromInt32(0x0c77)),
	86:  libc.Uint32FromInt32(libc.Int32FromInt32(0x0c7f)),
	87:  libc.Uint32FromInt32(libc.Int32FromInt32(0x0c84)),
	88:  libc.Uint32FromInt32(libc.Int32FromInt32(0x0d4f)),
	89:  libc.Uint32FromInt32(libc.Int32FromInt32(0x0d79)),
	90:  libc.Uint32FromInt32(libc.Int32FromInt32(0x0df4)),
	91:  libc.Uint32FromInt32(libc.Int32FromInt32(0x0e3f)),
	92:  libc.Uint32FromInt32(libc.Int32FromInt32(0x0e4f)),
	93:  libc.Uint32FromInt32(libc.Int32FromInt32(0x0e5a) | libc.Int32FromInt32(0x40000000)),
	94:  libc.Uint32FromInt32(libc.Int32FromInt32(0x0e5b)) | libc.Uint32FromUint32(0x80000000),
	95:  libc.Uint32FromInt32(libc.Int32FromInt32(0x0f01) | libc.Int32FromInt32(0x40000000)),
	96:  libc.Uint32FromInt32(libc.Int32FromInt32(0x0f17)) | libc.Uint32FromUint32(0x80000000),
	97:  libc.Uint32FromInt32(libc.Int32FromInt32(0x0f1a) | libc.Int32FromInt32(0x40000000)),
	98:  libc.Uint32FromInt32(libc.Int32FromInt32(0x0f1f)) | libc.Uint32FromUint32(0x80000000),
	99:  libc.Uint32FromInt32(libc.Int32FromInt32(0x0f34)),
	100: libc.Uint32FromInt32(libc.Int32FromInt32(0x0f36)),
	101: libc.Uint32FromInt32(libc.Int32FromInt32(0x0f38)),
	102: libc.Uint32FromInt32(libc.Int32FromInt32(0x0f3a) | libc.Int32FromInt32(0x40000000)),
	103: libc.Uint32FromInt32(libc.Int32FromInt32(0x0f3d)) | libc.Uint32FromUint32(0x80000000),
	104: libc.Uint32FromInt32(libc.Int32FromInt32(0x0f85)),
	105: libc.Uint32FromInt32(libc.Int32FromInt32(0x0fbe) | libc.Int32FromInt32(0x40000000)),
	106: libc.Uint32FromInt32(libc.Int32FromInt32(0x0fc5)) | libc.Uint32FromUint32(0x80000000),
	107: libc.Uint32FromInt32(libc.Int32FromInt32(0x0fc7) | libc.Int32FromInt32(0x40000000)),
	108: libc.Uint32FromInt32(libc.Int32FromInt32(0x0fcc)) | libc.Uint32FromUint32(0x80000000),
	109: libc.Uint32FromInt32(libc.Int32FromInt32(0x0fce) | libc.Int32FromInt32(0x40000000)),
	110: libc.Uint32FromInt32(libc.Int32FromInt32(0x0fda)) | libc.Uint32FromUint32(0x80000000),
	111: libc.Uint32FromInt32(libc.Int32FromInt32(0x104a) | libc.Int32FromInt32(0x40000000)),
	112: libc.Uint32FromInt32(libc.Int32FromInt32(0x104f)) | libc.Uint32FromUint32(0x80000000),
	113: libc.Uint32FromInt32(libc.Int32FromInt32(0x109e) | libc.Int32FromInt32(0x40000000)),
	114: libc.Uint32FromInt32(libc.Int32FromInt32(0x109f)) | libc.Uint32FromUint32(0x80000000),
	115: libc.Uint32FromInt32(libc.Int32FromInt32(0x10fb)),
	116: libc.Uint32FromInt32(libc.Int32FromInt32(0x1360) | libc.Int32FromInt32(0x40000000)),
	117: libc.Uint32FromInt32(libc.Int32FromInt32(0x1368)) | libc.Uint32FromUint32(0x80000000),
	118: libc.Uint32FromInt32(libc.Int32FromInt32(0x1390) | libc.Int32FromInt32(0x40000000)),
	119: libc.Uint32FromInt32(libc.Int32FromInt32(0x1399)) | libc.Uint32FromUint32(0x80000000),
	120: libc.Uint32FromInt32(libc.Int32FromInt32(0x1400)),
	121: libc.Uint32FromInt32(libc.Int32FromInt32(0x166d) | libc.Int32FromInt32(0x40000000)),
	122: libc.Uint32FromInt32(libc.Int32FromInt32(0x166e)) | libc.Uint32FromUint32(0x80000000),
	123: libc.Uint32FromInt32(libc.Int32FromInt32(0x169b) | libc.Int32FromInt32(0x40000000)),
	124: libc.Uint32FromInt32(libc.Int32FromInt32(0x169c)) | libc.Uint32FromUint32(0x80000000),
	125: libc.Uint32FromInt32(libc.Int32FromInt32(0x16eb) | libc.Int32FromInt32(0x40000000)),
	126: libc.Uint32FromInt32(libc.Int32FromInt32(0x16ed)) | libc.Uint32FromUint32(0x80000000),
	127: libc.Uint32FromInt32(libc.Int32FromInt32(0x1735) | libc.Int32FromInt32(0x40000000)),
	128: libc.Uint32FromInt32(libc.Int32FromInt32(0x1736)) | libc.Uint32FromUint32(0x80000000),
	129: libc.Uint32FromInt32(libc.Int32FromInt32(0x17d4) | libc.Int32FromInt32(0x40000000)),
	130: libc.Uint32FromInt32(libc.Int32FromInt32(0x17d6)) | libc.Uint32FromUint32(0x80000000),
	131: libc.Uint32FromInt32(libc.Int32FromInt32(0x17d8) | libc.Int32FromInt32(0x40000000)),
	132: libc.Uint32FromInt32(libc.Int32FromInt32(0x17db)) | libc.Uint32FromUint32(0x80000000),
	133: libc.Uint32FromInt32(libc.Int32FromInt32(0x1800) | libc.Int32FromInt32(0x40000000)),
	134: libc.Uint32FromInt32(libc.Int32FromInt32(0x180a)) | libc.Uint32FromUint32(0x80000000),
	135: libc.Uint32FromInt32(libc.Int32FromInt32(0x1940)),
	136: libc.Uint32FromInt32(libc.Int32FromInt32(0x1944) | libc.Int32FromInt32(0x40000000)),
	137: libc.Uint32FromInt32(libc.Int32FromInt32(0x1945)) | libc.Uint32FromUint32(0x80000000),
	138: libc.Uint32FromInt32(libc.Int32FromInt32(0x19de) | libc.Int32FromInt32(0x40000000)),
	139: libc.Uint32FromInt32(libc.Int32FromInt32(0x19ff)) | libc.Uint32FromUint32(0x80000000),
	140: libc.Uint32FromInt32(libc.Int32FromInt32(0x1a1e) | libc.Int32FromInt32(0x40000000)),
	141: libc.Uint32FromInt32(libc.Int32FromInt32(0x1a1f)) | libc.Uint32FromUint32(0x80000000),
	142: libc.Uint32FromInt32(libc.Int32FromInt32(0x1aa0) | libc.Int32FromInt32(0x40000000)),
	143: libc.Uint32FromInt32(libc.Int32FromInt32(0x1aa6)) | libc.Uint32FromUint32(0x80000000),
	144: libc.Uint32FromInt32(libc.Int32FromInt32(0x1aa8) | libc.Int32FromInt32(0x40000000)),
	145: libc.Uint32FromInt32(libc.Int32FromInt32(0x1aad)) | libc.Uint32FromUint32(0x80000000),
	146: libc.Uint32FromInt32(libc.Int32FromInt32(0x1b5a) | libc.Int32FromInt32(0x40000000)),
	147: libc.Uint32FromInt32(libc.Int32FromInt32(0x1b6a)) | libc.Uint32FromUint32(0x80000000),
	148: libc.Uint32FromInt32(libc.Int32FromInt32(0x1b74) | libc.Int32FromInt32(0x40000000)),
	149: libc.Uint32FromInt32(libc.Int32FromInt32(0x1b7e)) | libc.Uint32FromUint32(0x80000000),
	150: libc.Uint32FromInt32(libc.Int32FromInt32(0x1bfc) | libc.Int32FromInt32(0x40000000)),
	151: libc.Uint32FromInt32(libc.Int32FromInt32(0x1bff)) | libc.Uint32FromUint32(0x80000000),
	152: libc.Uint32FromInt32(libc.Int32FromInt32(0x1c3b) | libc.Int32FromInt32(0x40000000)),
	153: libc.Uint32FromInt32(libc.Int32FromInt32(0x1c3f)) | libc.Uint32FromUint32(0x80000000),
	154: libc.Uint32FromInt32(libc.Int32FromInt32(0x1c7e) | libc.Int32FromInt32(0x40000000)),
	155: libc.Uint32FromInt32(libc.Int32FromInt32(0x1c7f)) | libc.Uint32FromUint32(0x80000000),
	156: libc.Uint32FromInt32(libc.Int32FromInt32(0x1cc0) | libc.Int32FromInt32(0x40000000)),
	157: libc.Uint32FromInt32(libc.Int32FromInt32(0x1cc7)) | libc.Uint32FromUint32(0x80000000),
	158: libc.Uint32FromInt32(libc.Int32FromInt32(0x1cd3)),
	159: libc.Uint32FromInt32(libc.Int32FromInt32(0x1fbd)),
	160: libc.Uint32FromInt32(libc.Int32FromInt32(0x1fbf) | libc.Int32FromInt32(0x40000000)),
	161: libc.Uint32FromInt32(libc.Int32FromInt32(0x1fc1)) | libc.Uint32FromUint32(0x80000000),
	162: libc.Uint32FromInt32(libc.Int32FromInt32(0x1fcd) | libc.Int32FromInt32(0x40000000)),
	163: libc.Uint32FromInt32(libc.Int32FromInt32(0x1fcf)) | libc.Uint32FromUint32(0x80000000),
	164: libc.Uint32FromInt32(libc.Int32FromInt32(0x1fdd) | libc.Int32FromInt32(0x40000000)),
	165: libc.Uint32FromInt32(libc.Int32FromInt32(0x1fdf)) | libc.Uint32FromUint32(0x80000000),
	166: libc.Uint32FromInt32(libc.Int32FromInt32(0x1fed) | libc.Int32FromInt32(0x40000000)),
	167: libc.Uint32FromInt32(libc.Int32FromInt32(0x1fef)) | libc.Uint32FromUint32(0x80000000),
	168: libc.Uint32FromInt32(libc.Int32FromInt32(0x1ffd) | libc.Int32FromInt32(0x40000000)),
	169: libc.Uint32FromInt32(libc.Int32FromInt32(0x1ffe)) | libc.Uint32FromUint32(0x80000000),
	170: libc.Uint32FromInt32(libc.Int32FromInt32(0x2010) | libc.Int32FromInt32(0x40000000)),
	171: libc.Uint32FromInt32(libc.Int32FromInt32(0x2027)) | libc.Uint32FromUint32(0x80000000),
	172: libc.Uint32FromInt32(libc.Int32FromInt32(0x2030) | libc.Int32FromInt32(0x40000000)),
	173: libc.Uint32FromInt32(libc.Int32FromInt32(0x205e)) | libc.Uint32FromUint32(0x80000000),
	174: libc.Uint32FromInt32(libc.Int32FromInt32(0x207a) | libc.Int32FromInt32(0x40000000)),
	175: libc.Uint32FromInt32(libc.Int32FromInt32(0x207e)) | libc.Uint32FromUint32(0x80000000),
	176: libc.Uint32FromInt32(libc.Int32FromInt32(0x208a) | libc.Int32FromInt32(0x40000000)),
	177: libc.Uint32FromInt32(libc.Int32FromInt32(0x208e)) | libc.Uint32FromUint32(0x80000000),
	178: libc.Uint32FromInt32(libc.Int32FromInt32(0x20a0) | libc.Int32FromInt32(0x40000000)),
	179: libc.Uint32FromInt32(libc.Int32FromInt32(0x20c0)) | libc.Uint32FromUint32(0x80000000),
	180: libc.Uint32FromInt32(libc.Int32FromInt32(0x2100) | libc.Int32FromInt32(0x40000000)),
	181: libc.Uint32FromInt32(libc.Int32FromInt32(0x2101)) | libc.Uint32FromUint32(0x80000000),
	182: libc.Uint32FromInt32(libc.Int32FromInt32(0x2103) | libc.Int32FromInt32(0x40000000)),
	183: libc.Uint32FromInt32(libc.Int32FromInt32(0x2106)) | libc.Uint32FromUint32(0x80000000),
	184: libc.Uint32FromInt32(libc.Int32FromInt32(0x2108) | libc.Int32FromInt32(0x40000000)),
	185: libc.Uint32FromInt32(libc.Int32FromInt32(0x2109)) | libc.Uint32FromUint32(0x80000000),
	186: libc.Uint32FromInt32(libc.Int32FromInt32(0x2114)),
	187: libc.Uint32FromInt32(libc.Int32FromInt32(0x2116) | libc.Int32FromInt32(0x40000000)),
	188: libc.Uint32FromInt32(libc.Int32FromInt32(0x2118)) | libc.Uint32FromUint32(0x80000000),
	189: libc.Uint32FromInt32(libc.Int32FromInt32(0x211e) | libc.Int32FromInt32(0x40000000)),
	190: libc.Uint32FromInt32(libc.Int32FromInt32(0x2123)) | libc.Uint32FromUint32(0x80000000),
	191: libc.Uint32FromInt32(libc.Int32FromInt32(0x2125)),
	192: libc.Uint32FromInt32(libc.Int32FromInt32(0x2127)),
	193: libc.Uint32FromInt32(libc.Int32FromInt32(0x2129)),
	194: libc.Uint32FromInt32(libc.Int32FromInt32(0x212e)),
	195: libc.Uint32FromInt32(libc.Int32FromInt32(0x213a) | libc.Int32FromInt32(0x40000000)),
	196: libc.Uint32FromInt32(libc.Int32FromInt32(0x213b)) | libc.Uint32FromUint32(0x80000000),
	197: libc.Uint32FromInt32(libc.Int32FromInt32(0x2140) | libc.Int32FromInt32(0x40000000)),
	198: libc.Uint32FromInt32(libc.Int32FromInt32(0x2144)) | libc.Uint32FromUint32(0x80000000),
	199: libc.Uint32FromInt32(libc.Int32FromInt32(0x214a) | libc.Int32FromInt32(0x40000000)),
	200: libc.Uint32FromInt32(libc.Int32FromInt32(0x214d)) | libc.Uint32FromUint32(0x80000000),
	201: libc.Uint32FromInt32(libc.Int32FromInt32(0x214f)),
	202: libc.Uint32FromInt32(libc.Int32FromInt32(0x218a) | libc.Int32FromInt32(0x40000000)),
	203: libc.Uint32FromInt32(libc.Int32FromInt32(0x218b)) | libc.Uint32FromUint32(0x80000000),
	204: libc.Uint32FromInt32(libc.Int32FromInt32(0x2190) | libc.Int32FromInt32(0x40000000)),
	205: libc.Uint32FromInt32(libc.Int32FromInt32(0x2426)) | libc.Uint32FromUint32(0x80000000),
	206: libc.Uint32FromInt32(libc.Int32FromInt32(0x2440) | libc.Int32FromInt32(0x40000000)),
	207: libc.Uint32FromInt32(libc.Int32FromInt32(0x244a)) | libc.Uint32FromUint32(0x80000000),
	208: libc.Uint32FromInt32(libc.Int32FromInt32(0x249c) | libc.Int32FromInt32(0x40000000)),
	209: libc.Uint32FromInt32(libc.Int32FromInt32(0x24e9)) | libc.Uint32FromUint32(0x80000000),
	210: libc.Uint32FromInt32(libc.Int32FromInt32(0x2500) | libc.Int32FromInt32(0x40000000)),
	211: libc.Uint32FromInt32(libc.Int32FromInt32(0x2775)) | libc.Uint32FromUint32(0x80000000),
	212: libc.Uint32FromInt32(libc.Int32FromInt32(0x2794) | libc.Int32FromInt32(0x40000000)),
	213: libc.Uint32FromInt32(libc.Int32FromInt32(0x2b73)) | libc.Uint32FromUint32(0x80000000),
	214: libc.Uint32FromInt32(libc.Int32FromInt32(0x2b76) | libc.Int32FromInt32(0x40000000)),
	215: libc.Uint32FromInt32(libc.Int32FromInt32(0x2b95)) | libc.Uint32FromUint32(0x80000000),
	216: libc.Uint32FromInt32(libc.Int32FromInt32(0x2b97) | libc.Int32FromInt32(0x40000000)),
	217: libc.Uint32FromInt32(libc.Int32FromInt32(0x2bff)) | libc.Uint32FromUint32(0x80000000),
	218: libc.Uint32FromInt32(libc.Int32FromInt32(0x2ce5) | libc.Int32FromInt32(0x40000000)),
	219: libc.Uint32FromInt32(libc.Int32FromInt32(0x2cea)) | libc.Uint32FromUint32(0x80000000),
	220: libc.Uint32FromInt32(libc.Int32FromInt32(0x2cf9) | libc.Int32FromInt32(0x40000000)),
	221: libc.Uint32FromInt32(libc.Int32FromInt32(0x2cfc)) | libc.Uint32FromUint32(0x80000000),
	222: libc.Uint32FromInt32(libc.Int32FromInt32(0x2cfe) | libc.Int32FromInt32(0x40000000)),
	223: libc.Uint32FromInt32(libc.Int32FromInt32(0x2cff)) | libc.Uint32FromUint32(0x80000000),
	224: libc.Uint32FromInt32(libc.Int32FromInt32(0x2d70)),
	225: libc.Uint32FromInt32(libc.Int32FromInt32(0x2e00) | libc.Int32FromInt32(0x40000000)),
	226: libc.Uint32FromInt32(libc.Int32FromInt32(0x2e2e)) | libc.Uint32FromUint32(0x80000000),
	227: libc.Uint32FromInt32(libc.Int32FromInt32(0x2e30) | libc.Int32FromInt32(0x40000000)),
	228: libc.Uint32FromInt32(libc.Int32FromInt32(0x2e5d)) | libc.Uint32FromUint32(0x80000000),
	229: libc.Uint32FromInt32(libc.Int32FromInt32(0x2e80) | libc.Int32FromInt32(0x40000000)),
	230: libc.Uint32FromInt32(libc.Int32FromInt32(0x2e99)) | libc.Uint32FromUint32(0x80000000),
	231: libc.Uint32FromInt32(libc.Int32FromInt32(0x2e9b) | libc.Int32FromInt32(0x40000000)),
	232: libc.Uint32FromInt32(libc.Int32FromInt32(0x2ef3)) | libc.Uint32FromUint32(0x80000000),
	233: libc.Uint32FromInt32(libc.Int32FromInt32(0x2f00) | libc.Int32FromInt32(0x40000000)),
	234: libc.Uint32FromInt32(libc.Int32FromInt32(0x2fd5)) | libc.Uint32FromUint32(0x80000000),
	235: libc.Uint32FromInt32(libc.Int32FromInt32(0x2ff0) | libc.Int32FromInt32(0x40000000)),
	236: libc.Uint32FromInt32(libc.Int32FromInt32(0x2fff)) | libc.Uint32FromUint32(0x80000000),
	237: libc.Uint32FromInt32(libc.Int32FromInt32(0x3001) | libc.Int32FromInt32(0x40000000)),
	238: libc.Uint32FromInt32(libc.Int32FromInt32(0x3004)) | libc.Uint32FromUint32(0x80000000),
	239: libc.Uint32FromInt32(libc.Int32FromInt32(0x3008) | libc.Int32FromInt32(0x40000000)),
	240: libc.Uint32FromInt32(libc.Int32FromInt32(0x3020)) | libc.Uint32FromUint32(0x80000000),
	241: libc.Uint32FromInt32(libc.Int32FromInt32(0x3030)),
	242: libc.Uint32FromInt32(libc.Int32FromInt32(0x3036) | libc.Int32FromInt32(0x40000000)),
	243: libc.Uint32FromInt32(libc.Int32FromInt32(0x3037)) | libc.Uint32FromUint32(0x80000000),
	244: libc.Uint32FromInt32(libc.Int32FromInt32(0x303d) | libc.Int32FromInt32(0x40000000)),
	245: libc.Uint32FromInt32(libc.Int32FromInt32(0x303f)) | libc.Uint32FromUint32(0x80000000),
	246: libc.Uint32FromInt32(libc.Int32FromInt32(0x309b) | libc.Int32FromInt32(0x40000000)),
	247: libc.Uint32FromInt32(libc.Int32FromInt32(0x309c)) | libc.Uint32FromUint32(0x80000000),
	248: libc.Uint32FromInt32(libc.Int32FromInt32(0x30a0)),
	249: libc.Uint32FromInt32(libc.Int32FromInt32(0x30fb)),
	250: libc.Uint32FromInt32(libc.Int32FromInt32(0x3190) | libc.Int32FromInt32(0x40000000)),
	251: libc.Uint32FromInt32(libc.Int32FromInt32(0x3191)) | libc.Uint32FromUint32(0x80000000),
	252: libc.Uint32FromInt32(libc.Int32FromInt32(0x3196) | libc.Int32FromInt32(0x40000000)),
	253: libc.Uint32FromInt32(libc.Int32FromInt32(0x319f)) | libc.Uint32FromUint32(0x80000000),
	254: libc.Uint32FromInt32(libc.Int32FromInt32(0x31c0) | libc.Int32FromInt32(0x40000000)),
	255: libc.Uint32FromInt32(libc.Int32FromInt32(0x31e3)) | libc.Uint32FromUint32(0x80000000),
	256: libc.Uint32FromInt32(libc.Int32FromInt32(0x31ef)),
	257: libc.Uint32FromInt32(libc.Int32FromInt32(0x3200) | libc.Int32FromInt32(0x40000000)),
	258: libc.Uint32FromInt32(libc.Int32FromInt32(0x321e)) | libc.Uint32FromUint32(0x80000000),
	259: libc.Uint32FromInt32(libc.Int32FromInt32(0x322a) | libc.Int32FromInt32(0x40000000)),
	260: libc.Uint32FromInt32(libc.Int32FromInt32(0x3247)) | libc.Uint32FromUint32(0x80000000),
	261: libc.Uint32FromInt32(libc.Int32FromInt32(0x3250)),
	262: libc.Uint32FromInt32(libc.Int32FromInt32(0x3260) | libc.Int32FromInt32(0x40000000)),
	263: libc.Uint32FromInt32(libc.Int32FromInt32(0x327f)) | libc.Uint32FromUint32(0x80000000),
	264: libc.Uint32FromInt32(libc.Int32FromInt32(0x328a) | libc.Int32FromInt32(0x40000000)),
	265: libc.Uint32FromInt32(libc.Int32FromInt32(0x32b0)) | libc.Uint32FromUint32(0x80000000),
	266: libc.Uint32FromInt32(libc.Int32FromInt32(0x32c0) | libc.Int32FromInt32(0x40000000)),
	267: libc.Uint32FromInt32(libc.Int32FromInt32(0x33ff)) | libc.Uint32FromUint32(0x80000000),
	268: libc.Uint32FromInt32(libc.Int32FromInt32(0x4dc0) | libc.Int32FromInt32(0x40000000)),
	269: libc.Uint32FromInt32(libc.Int32FromInt32(0x4dff)) | libc.Uint32FromUint32(0x80000000),
	270: libc.Uint32FromInt32(libc.Int32FromInt32(0xa490) | libc.Int32FromInt32(0x40000000)),
	271: libc.Uint32FromInt32(libc.Int32FromInt32(0xa4c6)) | libc.Uint32FromUint32(0x80000000),
	272: libc.Uint32FromInt32(libc.Int32FromInt32(0xa4fe) | libc.Int32FromInt32(0x40000000)),
	273: libc.Uint32FromInt32(libc.Int32FromInt32(0xa4ff)) | libc.Uint32FromUint32(0x80000000),
	274: libc.Uint32FromInt32(libc.Int32FromInt32(0xa60d) | libc.Int32FromInt32(0x40000000)),
	275: libc.Uint32FromInt32(libc.Int32FromInt32(0xa60f)) | libc.Uint32FromUint32(0x80000000),
	276: libc.Uint32FromInt32(libc.Int32FromInt32(0xa673)),
	277: libc.Uint32FromInt32(libc.Int32FromInt32(0xa67e)),
	278: libc.Uint32FromInt32(libc.Int32FromInt32(0xa6f2) | libc.Int32FromInt32(0x40000000)),
	279: libc.Uint32FromInt32(libc.Int32FromInt32(0xa6f7)) | libc.Uint32FromUint32(0x80000000),
	280: libc.Uint32FromInt32(libc.Int32FromInt32(0xa700) | libc.Int32FromInt32(0x40000000)),
	281: libc.Uint32FromInt32(libc.Int32FromInt32(0xa716)) | libc.Uint32FromUint32(0x80000000),
	282: libc.Uint32FromInt32(libc.Int32FromInt32(0xa720) | libc.Int32FromInt32(0x40000000)),
	283: libc.Uint32FromInt32(libc.Int32FromInt32(0xa721)) | libc.Uint32FromUint32(0x80000000),
	284: libc.Uint32FromInt32(libc.Int32FromInt32(0xa789) | libc.Int32FromInt32(0x40000000)),
	285: libc.Uint32FromInt32(libc.Int32FromInt32(0xa78a)) | libc.Uint32FromUint32(0x80000000),
	286: libc.Uint32FromInt32(libc.Int32FromInt32(0xa828) | libc.Int32FromInt32(0x40000000)),
	287: libc.Uint32FromInt32(libc.Int32FromInt32(0xa82b)) | libc.Uint32FromUint32(0x80000000),
	288: libc.Uint32FromInt32(libc.Int32FromInt32(0xa836) | libc.Int32FromInt32(0x40000000)),
	289: libc.Uint32FromInt32(libc.Int32FromInt32(0xa839)) | libc.Uint32FromUint32(0x80000000),
	290: libc.Uint32FromInt32(libc.Int32FromInt32(0xa874) | libc.Int32FromInt32(0x40000000)),
	291: libc.Uint32FromInt32(libc.Int32FromInt32(0xa877)) | libc.Uint32FromUint32(0x80000000),
	292: libc.Uint32FromInt32(libc.Int32FromInt32(0xa8ce) | libc.Int32FromInt32(0x40000000)),
	293: libc.Uint32FromInt32(libc.Int32FromInt32(0xa8cf)) | libc.Uint32FromUint32(0x80000000),
	294: libc.Uint32FromInt32(libc.Int32FromInt32(0xa8f8) | libc.Int32FromInt32(0x40000000)),
	295: libc.Uint32FromInt32(libc.Int32FromInt32(0xa8fa)) | libc.Uint32FromUint32(0x80000000),
	296: libc.Uint32FromInt32(libc.Int32FromInt32(0xa8fc)),
	297: libc.Uint32FromInt32(libc.Int32FromInt32(0xa92e) | libc.Int32FromInt32(0x40000000)),
	298: libc.Uint32FromInt32(libc.Int32FromInt32(0xa92f)) | libc.Uint32FromUint32(0x80000000),
	299: libc.Uint32FromInt32(libc.Int32FromInt32(0xa95f)),
	300: libc.Uint32FromInt32(libc.Int32FromInt32(0xa9c1) | libc.Int32FromInt32(0x40000000)),
	301: libc.Uint32FromInt32(libc.Int32FromInt32(0xa9cd)) | libc.Uint32FromUint32(0x80000000),
	302: libc.Uint32FromInt32(libc.Int32FromInt32(0xa9de) | libc.Int32FromInt32(0x40000000)),
	303: libc.Uint32FromInt32(libc.Int32FromInt32(0xa9df)) | libc.Uint32FromUint32(0x80000000),
	304: libc.Uint32FromInt32(libc.Int32FromInt32(0xaa5c) | libc.Int32FromInt32(0x40000000)),
	305: libc.Uint32FromInt32(libc.Int32FromInt32(0xaa5f)) | libc.Uint32FromUint32(0x80000000),
	306: libc.Uint32FromInt32(libc.Int32FromInt32(0xaa77) | libc.Int32FromInt32(0x40000000)),
	307: libc.Uint32FromInt32(libc.Int32FromInt32(0xaa79)) | libc.Uint32FromUint32(0x80000000),
	308: libc.Uint32FromInt32(libc.Int32FromInt32(0xaade) | libc.Int32FromInt32(0x40000000)),
	309: libc.Uint32FromInt32(libc.Int32FromInt32(0xaadf)) | libc.Uint32FromUint32(0x80000000),
	310: libc.Uint32FromInt32(libc.Int32FromInt32(0xaaf0) | libc.Int32FromInt32(0x40000000)),
	311: libc.Uint32FromInt32(libc.Int32FromInt32(0xaaf1)) | libc.Uint32FromUint32(0x80000000),
	312: libc.Uint32FromInt32(libc.Int32FromInt32(0xab5b)),
	313: libc.Uint32FromInt32(libc.Int32FromInt32(0xab6a) | libc.Int32FromInt32(0x40000000)),
	314: libc.Uint32FromInt32(libc.Int32FromInt32(0xab6b)) | libc.Uint32FromUint32(0x80000000),
	315: libc.Uint32FromInt32(libc.Int32FromInt32(0xabeb)),
	316: libc.Uint32FromInt32(libc.Int32FromInt32(0xfb29)),
	317: libc.Uint32FromInt32(libc.Int32FromInt32(0xfbb2) | libc.Int32FromInt32(0x40000000)),
	318: libc.Uint32FromInt32(libc.Int32FromInt32(0xfbc2)) | libc.Uint32FromUint32(0x80000000),
	319: libc.Uint32FromInt32(libc.Int32FromInt32(0xfd3e) | libc.Int32FromInt32(0x40000000)),
	320: libc.Uint32FromInt32(libc.Int32FromInt32(0xfd4f)) | libc.Uint32FromUint32(0x80000000),
	321: libc.Uint32FromInt32(libc.Int32FromInt32(0xfdcf)),
	322: libc.Uint32FromInt32(libc.Int32FromInt32(0xfdfc) | libc.Int32FromInt32(0x40000000)),
	323: libc.Uint32FromInt32(libc.Int32FromInt32(0xfdff)) | libc.Uint32FromUint32(0x80000000),
	324: libc.Uint32FromInt32(libc.Int32FromInt32(0xfe10) | libc.Int32FromInt32(0x40000000)),
	325: libc.Uint32FromInt32(libc.Int32FromInt32(0xfe19)) | libc.Uint32FromUint32(0x80000000),
	326: libc.Uint32FromInt32(libc.Int32FromInt32(0xfe30) | libc.Int32FromInt32(0x40000000)),
	327: libc.Uint32FromInt32(libc.Int32FromInt32(0xfe52)) | libc.Uint32FromUint32(0x80000000),
	328: libc.Uint32FromInt32(libc.Int32FromInt32(0xfe54) | libc.Int32FromInt32(0x40000000)),
	329: libc.Uint32FromInt32(libc.Int32FromInt32(0xfe66)) | libc.Uint32FromUint32(0x80000000),
	330: libc.Uint32FromInt32(libc.Int32FromInt32(0xfe68) | libc.Int32FromInt32(0x40000000)),
	331: libc.Uint32FromInt32(libc.Int32FromInt32(0xfe6b)) | libc.Uint32FromUint32(0x80000000),
	332: libc.Uint32FromInt32(libc.Int32FromInt32(0xff01) | libc.Int32FromInt32(0x40000000)),
	333: libc.Uint32FromInt32(libc.Int32FromInt32(0xff0f)) | libc.Uint32FromUint32(0x80000000),
	334: libc.Uint32FromInt32(libc.Int32FromInt32(0xff1a) | libc.Int32FromInt32(0x40000000)),
	335: libc.Uint32FromInt32(libc.Int32FromInt32(0xff20)) | libc.Uint32FromUint32(0x80000000),
	336: libc.Uint32FromInt32(libc.Int32FromInt32(0xff3b) | libc.Int32FromInt32(0x40000000)),
	337: libc.Uint32FromInt32(libc.Int32FromInt32(0xff40)) | libc.Uint32FromUint32(0x80000000),
	338: libc.Uint32FromInt32(libc.Int32FromInt32(0xff5b) | libc.Int32FromInt32(0x40000000)),
	339: libc.Uint32FromInt32(libc.Int32FromInt32(0xff65)) | libc.Uint32FromUint32(0x80000000),
	340: libc.Uint32FromInt32(libc.Int32FromInt32(0xffe0) | libc.Int32FromInt32(0x40000000)),
	341: libc.Uint32FromInt32(libc.Int32FromInt32(0xffe6)) | libc.Uint32FromUint32(0x80000000),
	342: libc.Uint32FromInt32(libc.Int32FromInt32(0xffe8) | libc.Int32FromInt32(0x40000000)),
	343: libc.Uint32FromInt32(libc.Int32FromInt32(0xffee)) | libc.Uint32FromUint32(0x80000000),
	344: libc.Uint32FromInt32(libc.Int32FromInt32(0xfffc) | libc.Int32FromInt32(0x40000000)),
	345: libc.Uint32FromInt32(libc.Int32FromInt32(0xfffd)) | libc.Uint32FromUint32(0x80000000),
	346: libc.Uint32FromInt32(libc.Int32FromInt32(0x10100) | libc.Int32FromInt32(0x40000000)),
	347: libc.Uint32FromInt32(libc.Int32FromInt32(0x10102)) | libc.Uint32FromUint32(0x80000000),
	348: libc.Uint32FromInt32(libc.Int32FromInt32(0x10137) | libc.Int32FromInt32(0x40000000)),
	349: libc.Uint32FromInt32(libc.Int32FromInt32(0x1013f)) | libc.Uint32FromUint32(0x80000000),
	350: libc.Uint32FromInt32(libc.Int32FromInt32(0x10179) | libc.Int32FromInt32(0x40000000)),
	351: libc.Uint32FromInt32(libc.Int32FromInt32(0x10189)) | libc.Uint32FromUint32(0x80000000),
	352: libc.Uint32FromInt32(libc.Int32FromInt32(0x1018c) | libc.Int32FromInt32(0x40000000)),
	353: libc.Uint32FromInt32(libc.Int32FromInt32(0x1018e)) | libc.Uint32FromUint32(0x80000000),
	354: libc.Uint32FromInt32(libc.Int32FromInt32(0x10190) | libc.Int32FromInt32(0x40000000)),
	355: libc.Uint32FromInt32(libc.Int32FromInt32(0x1019c)) | libc.Uint32FromUint32(0x80000000),
	356: libc.Uint32FromInt32(libc.Int32FromInt32(0x101a0)),
	357: libc.Uint32FromInt32(libc.Int32FromInt32(0x101d0) | libc.Int32FromInt32(0x40000000)),
	358: libc.Uint32FromInt32(libc.Int32FromInt32(0x101fc)) | libc.Uint32FromUint32(0x80000000),
	359: libc.Uint32FromInt32(libc.Int32FromInt32(0x1039f)),
	360: libc.Uint32FromInt32(libc.Int32FromInt32(0x103d0)),
	361: libc.Uint32FromInt32(libc.Int32FromInt32(0x1056f)),
	362: libc.Uint32FromInt32(libc.Int32FromInt32(0x10857)),
	363: libc.Uint32FromInt32(libc.Int32FromInt32(0x10877) | libc.Int32FromInt32(0x40000000)),
	364: libc.Uint32FromInt32(libc.Int32FromInt32(0x10878)) | libc.Uint32FromUint32(0x80000000),
	365: libc.Uint32FromInt32(libc.Int32FromInt32(0x1091f)),
	366: libc.Uint32FromInt32(libc.Int32FromInt32(0x1093f)),
	367: libc.Uint32FromInt32(libc.Int32FromInt32(0x10a50) | libc.Int32FromInt32(0x40000000)),
	368: libc.Uint32FromInt32(libc.Int32FromInt32(0x10a58)) | libc.Uint32FromUint32(0x80000000),
	369: libc.Uint32FromInt32(libc.Int32FromInt32(0x10a7f)),
	370: libc.Uint32FromInt32(libc.Int32FromInt32(0x10ac8)),
	371: libc.Uint32FromInt32(libc.Int32FromInt32(0x10af0) | libc.Int32FromInt32(0x40000000)),
	372: libc.Uint32FromInt32(libc.Int32FromInt32(0x10af6)) | libc.Uint32FromUint32(0x80000000),
	373: libc.Uint32FromInt32(libc.Int32FromInt32(0x10b39) | libc.Int32FromInt32(0x40000000)),
	374: libc.Uint32FromInt32(libc.Int32FromInt32(0x10b3f)) | libc.Uint32FromUint32(0x80000000),
	375: libc.Uint32FromInt32(libc.Int32FromInt32(0x10b99) | libc.Int32FromInt32(0x40000000)),
	376: libc.Uint32FromInt32(libc.Int32FromInt32(0x10b9c)) | libc.Uint32FromUint32(0x80000000),
	377: libc.Uint32FromInt32(libc.Int32FromInt32(0x10ead)),
	378: libc.Uint32FromInt32(libc.Int32FromInt32(0x10f55) | libc.Int32FromInt32(0x40000000)),
	379: libc.Uint32FromInt32(libc.Int32FromInt32(0x10f59)) | libc.Uint32FromUint32(0x80000000),
	380: libc.Uint32FromInt32(libc.Int32FromInt32(0x10f86) | libc.Int32FromInt32(0x40000000)),
	381: libc.Uint32FromInt32(libc.Int32FromInt32(0x10f89)) | libc.Uint32FromUint32(0x80000000),
	382: libc.Uint32FromInt32(libc.Int32FromInt32(0x11047) | libc.Int32FromInt32(0x40000000)),
	383: libc.Uint32FromInt32(libc.Int32FromInt32(0x1104d)) | libc.Uint32FromUint32(0x80000000),
	384: libc.Uint32FromInt32(libc.Int32FromInt32(0x110bb) | libc.Int32FromInt32(0x40000000)),
	385: libc.Uint32FromInt32(libc.Int32FromInt32(0x110bc)) | libc.Uint32FromUint32(0x80000000),
	386: libc.Uint32FromInt32(libc.Int32FromInt32(0x110be) | libc.Int32FromInt32(0x40000000)),
	387: libc.Uint32FromInt32(libc.Int32FromInt32(0x110c1)) | libc.Uint32FromUint32(0x80000000),
	388: libc.Uint32FromInt32(libc.Int32FromInt32(0x11140) | libc.Int32FromInt32(0x40000000)),
	389: libc.Uint32FromInt32(libc.Int32FromInt32(0x11143)) | libc.Uint32FromUint32(0x80000000),
	390: libc.Uint32FromInt32(libc.Int32FromInt32(0x11174) | libc.Int32FromInt32(0x40000000)),
	391: libc.Uint32FromInt32(libc.Int32FromInt32(0x11175)) | libc.Uint32FromUint32(0x80000000),
	392: libc.Uint32FromInt32(libc.Int32FromInt32(0x111c5) | libc.Int32FromInt32(0x40000000)),
	393: libc.Uint32FromInt32(libc.Int32FromInt32(0x111c8)) | libc.Uint32FromUint32(0x80000000),
	394: libc.Uint32FromInt32(libc.Int32FromInt32(0x111cd)),
	395: libc.Uint32FromInt32(libc.Int32FromInt32(0x111db)),
	396: libc.Uint32FromInt32(libc.Int32FromInt32(0x111dd) | libc.Int32FromInt32(0x40000000)),
	397: libc.Uint32FromInt32(libc.Int32FromInt32(0x111df)) | libc.Uint32FromUint32(0x80000000),
	398: libc.Uint32FromInt32(libc.Int32FromInt32(0x11238) | libc.Int32FromInt32(0x40000000)),
	399: libc.Uint32FromInt32(libc.Int32FromInt32(0x1123d)) | libc.Uint32FromUint32(0x80000000),
	400: libc.Uint32FromInt32(libc.Int32FromInt32(0x112a9)),
	401: libc.Uint32FromInt32(libc.Int32FromInt32(0x1144b) | libc.Int32FromInt32(0x40000000)),
	402: libc.Uint32FromInt32(libc.Int32FromInt32(0x1144f)) | libc.Uint32FromUint32(0x80000000),
	403: libc.Uint32FromInt32(libc.Int32FromInt32(0x1145a) | libc.Int32FromInt32(0x40000000)),
	404: libc.Uint32FromInt32(libc.Int32FromInt32(0x1145b)) | libc.Uint32FromUint32(0x80000000),
	405: libc.Uint32FromInt32(libc.Int32FromInt32(0x1145d)),
	406: libc.Uint32FromInt32(libc.Int32FromInt32(0x114c6)),
	407: libc.Uint32FromInt32(libc.Int32FromInt32(0x115c1) | libc.Int32FromInt32(0x40000000)),
	408: libc.Uint32FromInt32(libc.Int32FromInt32(0x115d7)) | libc.Uint32FromUint32(0x80000000),
	409: libc.Uint32FromInt32(libc.Int32FromInt32(0x11641) | libc.Int32FromInt32(0x40000000)),
	410: libc.Uint32FromInt32(libc.Int32FromInt32(0x11643)) | libc.Uint32FromUint32(0x80000000),
	411: libc.Uint32FromInt32(libc.Int32FromInt32(0x11660) | libc.Int32FromInt32(0x40000000)),
	412: libc.Uint32FromInt32(libc.Int32FromInt32(0x1166c)) | libc.Uint32FromUint32(0x80000000),
	413: libc.Uint32FromInt32(libc.Int32FromInt32(0x116b9)),
	414: libc.Uint32FromInt32(libc.Int32FromInt32(0x1173c) | libc.Int32FromInt32(0x40000000)),
	415: libc.Uint32FromInt32(libc.Int32FromInt32(0x1173f)) | libc.Uint32FromUint32(0x80000000),
	416: libc.Uint32FromInt32(libc.Int32FromInt32(0x1183b)),
	417: libc.Uint32FromInt32(libc.Int32FromInt32(0x11944) | libc.Int32FromInt32(0x40000000)),
	418: libc.Uint32FromInt32(libc.Int32FromInt32(0x11946)) | libc.Uint32FromUint32(0x80000000),
	419: libc.Uint32FromInt32(libc.Int32FromInt32(0x119e2)),
	420: libc.Uint32FromInt32(libc.Int32FromInt32(0x11a3f) | libc.Int32FromInt32(0x40000000)),
	421: libc.Uint32FromInt32(libc.Int32FromInt32(0x11a46)) | libc.Uint32FromUint32(0x80000000),
	422: libc.Uint32FromInt32(libc.Int32FromInt32(0x11a9a) | libc.Int32FromInt32(0x40000000)),
	423: libc.Uint32FromInt32(libc.Int32FromInt32(0x11a9c)) | libc.Uint32FromUint32(0x80000000),
	424: libc.Uint32FromInt32(libc.Int32FromInt32(0x11a9e) | libc.Int32FromInt32(0x40000000)),
	425: libc.Uint32FromInt32(libc.Int32FromInt32(0x11aa2)) | libc.Uint32FromUint32(0x80000000),
	426: libc.Uint32FromInt32(libc.Int32FromInt32(0x11b00) | libc.Int32FromInt32(0x40000000)),
	427: libc.Uint32FromInt32(libc.Int32FromInt32(0x11b09)) | libc.Uint32FromUint32(0x80000000),
	428: libc.Uint32FromInt32(libc.Int32FromInt32(0x11c41) | libc.Int32FromInt32(0x40000000)),
	429: libc.Uint32FromInt32(libc.Int32FromInt32(0x11c45)) | libc.Uint32FromUint32(0x80000000),
	430: libc.Uint32FromInt32(libc.Int32FromInt32(0x11c70) | libc.Int32FromInt32(0x40000000)),
	431: libc.Uint32FromInt32(libc.Int32FromInt32(0x11c71)) | libc.Uint32FromUint32(0x80000000),
	432: libc.Uint32FromInt32(libc.Int32FromInt32(0x11ef7) | libc.Int32FromInt32(0x40000000)),
	433: libc.Uint32FromInt32(libc.Int32FromInt32(0x11ef8)) | libc.Uint32FromUint32(0x80000000),
	434: libc.Uint32FromInt32(libc.Int32FromInt32(0x11f43) | libc.Int32FromInt32(0x40000000)),
	435: libc.Uint32FromInt32(libc.Int32FromInt32(0x11f4f)) | libc.Uint32FromUint32(0x80000000),
	436: libc.Uint32FromInt32(libc.Int32FromInt32(0x11fd5) | libc.Int32FromInt32(0x40000000)),
	437: libc.Uint32FromInt32(libc.Int32FromInt32(0x11ff1)) | libc.Uint32FromUint32(0x80000000),
	438: libc.Uint32FromInt32(libc.Int32FromInt32(0x11fff)),
	439: libc.Uint32FromInt32(libc.Int32FromInt32(0x12470) | libc.Int32FromInt32(0x40000000)),
	440: libc.Uint32FromInt32(libc.Int32FromInt32(0x12474)) | libc.Uint32FromUint32(0x80000000),
	441: libc.Uint32FromInt32(libc.Int32FromInt32(0x12ff1) | libc.Int32FromInt32(0x40000000)),
	442: libc.Uint32FromInt32(libc.Int32FromInt32(0x12ff2)) | libc.Uint32FromUint32(0x80000000),
	443: libc.Uint32FromInt32(libc.Int32FromInt32(0x16a6e) | libc.Int32FromInt32(0x40000000)),
	444: libc.Uint32FromInt32(libc.Int32FromInt32(0x16a6f)) | libc.Uint32FromUint32(0x80000000),
	445: libc.Uint32FromInt32(libc.Int32FromInt32(0x16af5)),
	446: libc.Uint32FromInt32(libc.Int32FromInt32(0x16b37) | libc.Int32FromInt32(0x40000000)),
	447: libc.Uint32FromInt32(libc.Int32FromInt32(0x16b3f)) | libc.Uint32FromUint32(0x80000000),
	448: libc.Uint32FromInt32(libc.Int32FromInt32(0x16b44) | libc.Int32FromInt32(0x40000000)),
	449: libc.Uint32FromInt32(libc.Int32FromInt32(0x16b45)) | libc.Uint32FromUint32(0x80000000),
	450: libc.Uint32FromInt32(libc.Int32FromInt32(0x16e97) | libc.Int32FromInt32(0x40000000)),
	451: libc.Uint32FromInt32(libc.Int32FromInt32(0x16e9a)) | libc.Uint32FromUint32(0x80000000),
	452: libc.Uint32FromInt32(libc.Int32FromInt32(0x16fe2)),
	453: libc.Uint32FromInt32(libc.Int32FromInt32(0x1bc9c)),
	454: libc.Uint32FromInt32(libc.Int32FromInt32(0x1bc9f)),
	455: libc.Uint32FromInt32(libc.Int32FromInt32(0x1cf50) | libc.Int32FromInt32(0x40000000)),
	456: libc.Uint32FromInt32(libc.Int32FromInt32(0x1cfc3)) | libc.Uint32FromUint32(0x80000000),
	457: libc.Uint32FromInt32(libc.Int32FromInt32(0x1d000) | libc.Int32FromInt32(0x40000000)),
	458: libc.Uint32FromInt32(libc.Int32FromInt32(0x1d0f5)) | libc.Uint32FromUint32(0x80000000),
	459: libc.Uint32FromInt32(libc.Int32FromInt32(0x1d100) | libc.Int32FromInt32(0x40000000)),
	460: libc.Uint32FromInt32(libc.Int32FromInt32(0x1d126)) | libc.Uint32FromUint32(0x80000000),
	461: libc.Uint32FromInt32(libc.Int32FromInt32(0x1d129) | libc.Int32FromInt32(0x40000000)),
	462: libc.Uint32FromInt32(libc.Int32FromInt32(0x1d164)) | libc.Uint32FromUint32(0x80000000),
	463: libc.Uint32FromInt32(libc.Int32FromInt32(0x1d16a) | libc.Int32FromInt32(0x40000000)),
	464: libc.Uint32FromInt32(libc.Int32FromInt32(0x1d16c)) | libc.Uint32FromUint32(0x80000000),
	465: libc.Uint32FromInt32(libc.Int32FromInt32(0x1d183) | libc.Int32FromInt32(0x40000000)),
	466: libc.Uint32FromInt32(libc.Int32FromInt32(0x1d184)) | libc.Uint32FromUint32(0x80000000),
	467: libc.Uint32FromInt32(libc.Int32FromInt32(0x1d18c) | libc.Int32FromInt32(0x40000000)),
	468: libc.Uint32FromInt32(libc.Int32FromInt32(0x1d1a9)) | libc.Uint32FromUint32(0x80000000),
	469: libc.Uint32FromInt32(libc.Int32FromInt32(0x1d1ae) | libc.Int32FromInt32(0x40000000)),
	470: libc.Uint32FromInt32(libc.Int32FromInt32(0x1d1ea)) | libc.Uint32FromUint32(0x80000000),
	471: libc.Uint32FromInt32(libc.Int32FromInt32(0x1d200) | libc.Int32FromInt32(0x40000000)),
	472: libc.Uint32FromInt32(libc.Int32FromInt32(0x1d241)) | libc.Uint32FromUint32(0x80000000),
	473: libc.Uint32FromInt32(libc.Int32FromInt32(0x1d245)),
	474: libc.Uint32FromInt32(libc.Int32FromInt32(0x1d300) | libc.Int32FromInt32(0x40000000)),
	475: libc.Uint32FromInt32(libc.Int32FromInt32(0x1d356)) | libc.Uint32FromUint32(0x80000000),
	476: libc.Uint32FromInt32(libc.Int32FromInt32(0x1d6c1)),
	477: libc.Uint32FromInt32(libc.Int32FromInt32(0x1d6db)),
	478: libc.Uint32FromInt32(libc.Int32FromInt32(0x1d6fb)),
	479: libc.Uint32FromInt32(libc.Int32FromInt32(0x1d715)),
	480: libc.Uint32FromInt32(libc.Int32FromInt32(0x1d735)),
	481: libc.Uint32FromInt32(libc.Int32FromInt32(0x1d74f)),
	482: libc.Uint32FromInt32(libc.Int32FromInt32(0x1d76f)),
	483: libc.Uint32FromInt32(libc.Int32FromInt32(0x1d789)),
	484: libc.Uint32FromInt32(libc.Int32FromInt32(0x1d7a9)),
	485: libc.Uint32FromInt32(libc.Int32FromInt32(0x1d7c3)),
	486: libc.Uint32FromInt32(libc.Int32FromInt32(0x1d800) | libc.Int32FromInt32(0x40000000)),
	487: libc.Uint32FromInt32(libc.Int32FromInt32(0x1d9ff)) | libc.Uint32FromUint32(0x80000000),
	488: libc.Uint32FromInt32(libc.Int32FromInt32(0x1da37) | libc.Int32FromInt32(0x40000000)),
	489: libc.Uint32FromInt32(libc.Int32FromInt32(0x1da3a)) | libc.Uint32FromUint32(0x80000000),
	490: libc.Uint32FromInt32(libc.Int32FromInt32(0x1da6d) | libc.Int32FromInt32(0x40000000)),
	491: libc.Uint32FromInt32(libc.Int32FromInt32(0x1da74)) | libc.Uint32FromUint32(0x80000000),
	492: libc.Uint32FromInt32(libc.Int32FromInt32(0x1da76) | libc.Int32FromInt32(0x40000000)),
	493: libc.Uint32FromInt32(libc.Int32FromInt32(0x1da83)) | libc.Uint32FromUint32(0x80000000),
	494: libc.Uint32FromInt32(libc.Int32FromInt32(0x1da85) | libc.Int32FromInt32(0x40000000)),
	495: libc.Uint32FromInt32(libc.Int32FromInt32(0x1da8b)) | libc.Uint32FromUint32(0x80000000),
	496: libc.Uint32FromInt32(libc.Int32FromInt32(0x1e14f)),
	497: libc.Uint32FromInt32(libc.Int32FromInt32(0x1e2ff)),
	498: libc.Uint32FromInt32(libc.Int32FromInt32(0x1e95e) | libc.Int32FromInt32(0x40000000)),
	499: libc.Uint32FromInt32(libc.Int32FromInt32(0x1e95f)) | libc.Uint32FromUint32(0x80000000),
	500: libc.Uint32FromInt32(libc.Int32FromInt32(0x1ecac)),
	501: libc.Uint32FromInt32(libc.Int32FromInt32(0x1ecb0)),
	502: libc.Uint32FromInt32(libc.Int32FromInt32(0x1ed2e)),
	503: libc.Uint32FromInt32(libc.Int32FromInt32(0x1eef0) | libc.Int32FromInt32(0x40000000)),
	504: libc.Uint32FromInt32(libc.Int32FromInt32(0x1eef1)) | libc.Uint32FromUint32(0x80000000),
	505: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f000) | libc.Int32FromInt32(0x40000000)),
	506: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f02b)) | libc.Uint32FromUint32(0x80000000),
	507: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f030) | libc.Int32FromInt32(0x40000000)),
	508: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f093)) | libc.Uint32FromUint32(0x80000000),
	509: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f0a0) | libc.Int32FromInt32(0x40000000)),
	510: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f0ae)) | libc.Uint32FromUint32(0x80000000),
	511: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f0b1) | libc.Int32FromInt32(0x40000000)),
	512: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f0bf)) | libc.Uint32FromUint32(0x80000000),
	513: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f0c1) | libc.Int32FromInt32(0x40000000)),
	514: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f0cf)) | libc.Uint32FromUint32(0x80000000),
	515: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f0d1) | libc.Int32FromInt32(0x40000000)),
	516: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f0f5)) | libc.Uint32FromUint32(0x80000000),
	517: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f10d) | libc.Int32FromInt32(0x40000000)),
	518: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f1ad)) | libc.Uint32FromUint32(0x80000000),
	519: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f1e6) | libc.Int32FromInt32(0x40000000)),
	520: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f202)) | libc.Uint32FromUint32(0x80000000),
	521: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f210) | libc.Int32FromInt32(0x40000000)),
	522: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f23b)) | libc.Uint32FromUint32(0x80000000),
	523: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f240) | libc.Int32FromInt32(0x40000000)),
	524: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f248)) | libc.Uint32FromUint32(0x80000000),
	525: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f250) | libc.Int32FromInt32(0x40000000)),
	526: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f251)) | libc.Uint32FromUint32(0x80000000),
	527: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f260) | libc.Int32FromInt32(0x40000000)),
	528: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f265)) | libc.Uint32FromUint32(0x80000000),
	529: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f300) | libc.Int32FromInt32(0x40000000)),
	530: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f6d7)) | libc.Uint32FromUint32(0x80000000),
	531: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f6dc) | libc.Int32FromInt32(0x40000000)),
	532: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f6ec)) | libc.Uint32FromUint32(0x80000000),
	533: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f6f0) | libc.Int32FromInt32(0x40000000)),
	534: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f6fc)) | libc.Uint32FromUint32(0x80000000),
	535: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f700) | libc.Int32FromInt32(0x40000000)),
	536: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f776)) | libc.Uint32FromUint32(0x80000000),
	537: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f77b) | libc.Int32FromInt32(0x40000000)),
	538: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f7d9)) | libc.Uint32FromUint32(0x80000000),
	539: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f7e0) | libc.Int32FromInt32(0x40000000)),
	540: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f7eb)) | libc.Uint32FromUint32(0x80000000),
	541: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f7f0)),
	542: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f800) | libc.Int32FromInt32(0x40000000)),
	543: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f80b)) | libc.Uint32FromUint32(0x80000000),
	544: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f810) | libc.Int32FromInt32(0x40000000)),
	545: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f847)) | libc.Uint32FromUint32(0x80000000),
	546: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f850) | libc.Int32FromInt32(0x40000000)),
	547: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f859)) | libc.Uint32FromUint32(0x80000000),
	548: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f860) | libc.Int32FromInt32(0x40000000)),
	549: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f887)) | libc.Uint32FromUint32(0x80000000),
	550: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f890) | libc.Int32FromInt32(0x40000000)),
	551: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f8ad)) | libc.Uint32FromUint32(0x80000000),
	552: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f8b0) | libc.Int32FromInt32(0x40000000)),
	553: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f8b1)) | libc.Uint32FromUint32(0x80000000),
	554: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f900) | libc.Int32FromInt32(0x40000000)),
	555: libc.Uint32FromInt32(libc.Int32FromInt32(0x1fa53)) | libc.Uint32FromUint32(0x80000000),
	556: libc.Uint32FromInt32(libc.Int32FromInt32(0x1fa60) | libc.Int32FromInt32(0x40000000)),
	557: libc.Uint32FromInt32(libc.Int32FromInt32(0x1fa6d)) | libc.Uint32FromUint32(0x80000000),
	558: libc.Uint32FromInt32(libc.Int32FromInt32(0x1fa70) | libc.Int32FromInt32(0x40000000)),
	559: libc.Uint32FromInt32(libc.Int32FromInt32(0x1fa7c)) | libc.Uint32FromUint32(0x80000000),
	560: libc.Uint32FromInt32(libc.Int32FromInt32(0x1fa80) | libc.Int32FromInt32(0x40000000)),
	561: libc.Uint32FromInt32(libc.Int32FromInt32(0x1fa88)) | libc.Uint32FromUint32(0x80000000),
	562: libc.Uint32FromInt32(libc.Int32FromInt32(0x1fa90) | libc.Int32FromInt32(0x40000000)),
	563: libc.Uint32FromInt32(libc.Int32FromInt32(0x1fabd)) | libc.Uint32FromUint32(0x80000000),
	564: libc.Uint32FromInt32(libc.Int32FromInt32(0x1fabf) | libc.Int32FromInt32(0x40000000)),
	565: libc.Uint32FromInt32(libc.Int32FromInt32(0x1fac5)) | libc.Uint32FromUint32(0x80000000),
	566: libc.Uint32FromInt32(libc.Int32FromInt32(0x1face) | libc.Int32FromInt32(0x40000000)),
	567: libc.Uint32FromInt32(libc.Int32FromInt32(0x1fadb)) | libc.Uint32FromUint32(0x80000000),
	568: libc.Uint32FromInt32(libc.Int32FromInt32(0x1fae0) | libc.Int32FromInt32(0x40000000)),
	569: libc.Uint32FromInt32(libc.Int32FromInt32(0x1fae8)) | libc.Uint32FromUint32(0x80000000),
	570: libc.Uint32FromInt32(libc.Int32FromInt32(0x1faf0) | libc.Int32FromInt32(0x40000000)),
	571: libc.Uint32FromInt32(libc.Int32FromInt32(0x1faf8)) | libc.Uint32FromUint32(0x80000000),
	572: libc.Uint32FromInt32(libc.Int32FromInt32(0x1fb00) | libc.Int32FromInt32(0x40000000)),
	573: libc.Uint32FromInt32(libc.Int32FromInt32(0x1fb92)) | libc.Uint32FromUint32(0x80000000),
	574: libc.Uint32FromInt32(libc.Int32FromInt32(0x1fb94) | libc.Int32FromInt32(0x40000000)),
	575: libc.Uint32FromInt32(libc.Int32FromInt32(0x1fbca)) | libc.Uint32FromUint32(0x80000000),
}

func md_get_unicode_fold_info(tls *libc.TLS, codepoint uint32, info uintptr) {
	var codepoints, map1 uintptr
	var i, index, v2 int32
	var n_codepoints uint32
	_, _, _, _, _, _ = codepoints, i, index, map1, n_codepoints, v2
	/* Fast path for ASCII characters. */
	if codepoint <= uint32(0x7f) {
		*(*uint32)(__ccgo_ptr(info)) = codepoint
		if libc.Uint32FromInt32(libc.Int32FromUint8('A')) <= codepoint && codepoint <= libc.Uint32FromInt32(libc.Int32FromUint8('Z')) {
			*(*uint32)(__ccgo_ptr(info)) += libc.Uint32FromInt32(libc.Int32FromUint8('a') - libc.Int32FromUint8('A'))
		}
		(*MD_UNICODE_FOLD_INFO)(unsafe.Pointer(info)).Fn_codepoints = uint32(1)
		return
	}
	/* Try to locate the codepoint in any of the maps. */
	i = 0
	for {
		if !(i < libc.Int32FromUint64(libc.Uint64FromInt64(96)/libc.Uint64FromInt64(32))) {
			break
		}
		index = md_unicode_bsearch__(tls, codepoint, FOLD_MAP_LIST[i].Fmap1, FOLD_MAP_LIST[i].Fmap_size)
		if index >= 0 {
			/* Found the mapping. */
			n_codepoints = FOLD_MAP_LIST[i].Fn_codepoints
			map1 = FOLD_MAP_LIST[i].Fmap1
			codepoints = FOLD_MAP_LIST[i].Fdata + uintptr(libc.Uint32FromInt32(index)*n_codepoints)*4
			libc.Xmemcpy(tls, info, codepoints, uint64(4)*uint64(n_codepoints))
			(*MD_UNICODE_FOLD_INFO)(unsafe.Pointer(info)).Fn_codepoints = n_codepoints
			if *(*uint32)(__ccgo_ptr(FOLD_MAP_LIST[i].Fmap1 + uintptr(index)*4)) != codepoint {
				/* The found mapping maps whole range of codepoints,
				 * i.e. we have to offset info->codepoints[0] accordingly. */
				if *(*uint32)(__ccgo_ptr(map1 + uintptr(index)*4))&uint32(0x00ffffff)+uint32(1) == *(*uint32)(__ccgo_ptr(codepoints)) {
					/* Alternating type of the range. */
					if codepoint&uint32(0x1) == *(*uint32)(__ccgo_ptr(map1 + uintptr(index)*4))&uint32(0x1) {
						v2 = int32(1)
					} else {
						v2 = 0
					}
					*(*uint32)(__ccgo_ptr(info)) = codepoint + libc.Uint32FromInt32(v2)
				} else {
					/* Range to range kind of mapping. */
					*(*uint32)(__ccgo_ptr(info)) += codepoint - *(*uint32)(__ccgo_ptr(map1 + uintptr(index)*4))&uint32(0x00ffffff)
				}
			}
			return
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	/* No mapping found. Map the codepoint to itself. */
	*(*uint32)(__ccgo_ptr(info)) = codepoint
	(*MD_UNICODE_FOLD_INFO)(unsafe.Pointer(info)).Fn_codepoints = uint32(1)
}

/* Unicode "Pc", "Pd", "Pe", "Pf", "Pi", "Po", "Ps" categories.
 * (generated by scripts/build_folding_map.py) */
var FOLD_MAP_1 = [283]uint32{
	0:   libc.Uint32FromInt32(libc.Int32FromInt32(0x0041) | libc.Int32FromInt32(0x40000000)),
	1:   libc.Uint32FromInt32(libc.Int32FromInt32(0x005a)) | libc.Uint32FromUint32(0x80000000),
	2:   libc.Uint32FromInt32(libc.Int32FromInt32(0x00b5)),
	3:   libc.Uint32FromInt32(libc.Int32FromInt32(0x00c0) | libc.Int32FromInt32(0x40000000)),
	4:   libc.Uint32FromInt32(libc.Int32FromInt32(0x00d6)) | libc.Uint32FromUint32(0x80000000),
	5:   libc.Uint32FromInt32(libc.Int32FromInt32(0x00d8) | libc.Int32FromInt32(0x40000000)),
	6:   libc.Uint32FromInt32(libc.Int32FromInt32(0x00de)) | libc.Uint32FromUint32(0x80000000),
	7:   libc.Uint32FromInt32(libc.Int32FromInt32(0x0100) | libc.Int32FromInt32(0x40000000)),
	8:   libc.Uint32FromInt32(libc.Int32FromInt32(0x012e)) | libc.Uint32FromUint32(0x80000000),
	9:   libc.Uint32FromInt32(libc.Int32FromInt32(0x0132) | libc.Int32FromInt32(0x40000000)),
	10:  libc.Uint32FromInt32(libc.Int32FromInt32(0x0136)) | libc.Uint32FromUint32(0x80000000),
	11:  libc.Uint32FromInt32(libc.Int32FromInt32(0x0139) | libc.Int32FromInt32(0x40000000)),
	12:  libc.Uint32FromInt32(libc.Int32FromInt32(0x0147)) | libc.Uint32FromUint32(0x80000000),
	13:  libc.Uint32FromInt32(libc.Int32FromInt32(0x014a) | libc.Int32FromInt32(0x40000000)),
	14:  libc.Uint32FromInt32(libc.Int32FromInt32(0x0176)) | libc.Uint32FromUint32(0x80000000),
	15:  libc.Uint32FromInt32(libc.Int32FromInt32(0x0178)),
	16:  libc.Uint32FromInt32(libc.Int32FromInt32(0x0179) | libc.Int32FromInt32(0x40000000)),
	17:  libc.Uint32FromInt32(libc.Int32FromInt32(0x017d)) | libc.Uint32FromUint32(0x80000000),
	18:  libc.Uint32FromInt32(libc.Int32FromInt32(0x017f)),
	19:  libc.Uint32FromInt32(libc.Int32FromInt32(0x0181)),
	20:  libc.Uint32FromInt32(libc.Int32FromInt32(0x0182)),
	21:  libc.Uint32FromInt32(libc.Int32FromInt32(0x0184)),
	22:  libc.Uint32FromInt32(libc.Int32FromInt32(0x0186)),
	23:  libc.Uint32FromInt32(libc.Int32FromInt32(0x0187)),
	24:  libc.Uint32FromInt32(libc.Int32FromInt32(0x0189)),
	25:  libc.Uint32FromInt32(libc.Int32FromInt32(0x018a)),
	26:  libc.Uint32FromInt32(libc.Int32FromInt32(0x018b)),
	27:  libc.Uint32FromInt32(libc.Int32FromInt32(0x018e)),
	28:  libc.Uint32FromInt32(libc.Int32FromInt32(0x018f)),
	29:  libc.Uint32FromInt32(libc.Int32FromInt32(0x0190)),
	30:  libc.Uint32FromInt32(libc.Int32FromInt32(0x0191)),
	31:  libc.Uint32FromInt32(libc.Int32FromInt32(0x0193)),
	32:  libc.Uint32FromInt32(libc.Int32FromInt32(0x0194)),
	33:  libc.Uint32FromInt32(libc.Int32FromInt32(0x0196)),
	34:  libc.Uint32FromInt32(libc.Int32FromInt32(0x0197)),
	35:  libc.Uint32FromInt32(libc.Int32FromInt32(0x0198)),
	36:  libc.Uint32FromInt32(libc.Int32FromInt32(0x019c)),
	37:  libc.Uint32FromInt32(libc.Int32FromInt32(0x019d)),
	38:  libc.Uint32FromInt32(libc.Int32FromInt32(0x019f)),
	39:  libc.Uint32FromInt32(libc.Int32FromInt32(0x01a0) | libc.Int32FromInt32(0x40000000)),
	40:  libc.Uint32FromInt32(libc.Int32FromInt32(0x01a4)) | libc.Uint32FromUint32(0x80000000),
	41:  libc.Uint32FromInt32(libc.Int32FromInt32(0x01a6)),
	42:  libc.Uint32FromInt32(libc.Int32FromInt32(0x01a7)),
	43:  libc.Uint32FromInt32(libc.Int32FromInt32(0x01a9)),
	44:  libc.Uint32FromInt32(libc.Int32FromInt32(0x01ac)),
	45:  libc.Uint32FromInt32(libc.Int32FromInt32(0x01ae)),
	46:  libc.Uint32FromInt32(libc.Int32FromInt32(0x01af)),
	47:  libc.Uint32FromInt32(libc.Int32FromInt32(0x01b1)),
	48:  libc.Uint32FromInt32(libc.Int32FromInt32(0x01b2)),
	49:  libc.Uint32FromInt32(libc.Int32FromInt32(0x01b3)),
	50:  libc.Uint32FromInt32(libc.Int32FromInt32(0x01b5)),
	51:  libc.Uint32FromInt32(libc.Int32FromInt32(0x01b7)),
	52:  libc.Uint32FromInt32(libc.Int32FromInt32(0x01b8)),
	53:  libc.Uint32FromInt32(libc.Int32FromInt32(0x01bc)),
	54:  libc.Uint32FromInt32(libc.Int32FromInt32(0x01c4)),
	55:  libc.Uint32FromInt32(libc.Int32FromInt32(0x01c5)),
	56:  libc.Uint32FromInt32(libc.Int32FromInt32(0x01c7)),
	57:  libc.Uint32FromInt32(libc.Int32FromInt32(0x01c8)),
	58:  libc.Uint32FromInt32(libc.Int32FromInt32(0x01ca)),
	59:  libc.Uint32FromInt32(libc.Int32FromInt32(0x01cb) | libc.Int32FromInt32(0x40000000)),
	60:  libc.Uint32FromInt32(libc.Int32FromInt32(0x01db)) | libc.Uint32FromUint32(0x80000000),
	61:  libc.Uint32FromInt32(libc.Int32FromInt32(0x01de) | libc.Int32FromInt32(0x40000000)),
	62:  libc.Uint32FromInt32(libc.Int32FromInt32(0x01ee)) | libc.Uint32FromUint32(0x80000000),
	63:  libc.Uint32FromInt32(libc.Int32FromInt32(0x01f1)),
	64:  libc.Uint32FromInt32(libc.Int32FromInt32(0x01f2)),
	65:  libc.Uint32FromInt32(libc.Int32FromInt32(0x01f4)),
	66:  libc.Uint32FromInt32(libc.Int32FromInt32(0x01f6)),
	67:  libc.Uint32FromInt32(libc.Int32FromInt32(0x01f7)),
	68:  libc.Uint32FromInt32(libc.Int32FromInt32(0x01f8) | libc.Int32FromInt32(0x40000000)),
	69:  libc.Uint32FromInt32(libc.Int32FromInt32(0x021e)) | libc.Uint32FromUint32(0x80000000),
	70:  libc.Uint32FromInt32(libc.Int32FromInt32(0x0220)),
	71:  libc.Uint32FromInt32(libc.Int32FromInt32(0x0222) | libc.Int32FromInt32(0x40000000)),
	72:  libc.Uint32FromInt32(libc.Int32FromInt32(0x0232)) | libc.Uint32FromUint32(0x80000000),
	73:  libc.Uint32FromInt32(libc.Int32FromInt32(0x023a)),
	74:  libc.Uint32FromInt32(libc.Int32FromInt32(0x023b)),
	75:  libc.Uint32FromInt32(libc.Int32FromInt32(0x023d)),
	76:  libc.Uint32FromInt32(libc.Int32FromInt32(0x023e)),
	77:  libc.Uint32FromInt32(libc.Int32FromInt32(0x0241)),
	78:  libc.Uint32FromInt32(libc.Int32FromInt32(0x0243)),
	79:  libc.Uint32FromInt32(libc.Int32FromInt32(0x0244)),
	80:  libc.Uint32FromInt32(libc.Int32FromInt32(0x0245)),
	81:  libc.Uint32FromInt32(libc.Int32FromInt32(0x0246) | libc.Int32FromInt32(0x40000000)),
	82:  libc.Uint32FromInt32(libc.Int32FromInt32(0x024e)) | libc.Uint32FromUint32(0x80000000),
	83:  libc.Uint32FromInt32(libc.Int32FromInt32(0x0345)),
	84:  libc.Uint32FromInt32(libc.Int32FromInt32(0x0370)),
	85:  libc.Uint32FromInt32(libc.Int32FromInt32(0x0372)),
	86:  libc.Uint32FromInt32(libc.Int32FromInt32(0x0376)),
	87:  libc.Uint32FromInt32(libc.Int32FromInt32(0x037f)),
	88:  libc.Uint32FromInt32(libc.Int32FromInt32(0x0386)),
	89:  libc.Uint32FromInt32(libc.Int32FromInt32(0x0388) | libc.Int32FromInt32(0x40000000)),
	90:  libc.Uint32FromInt32(libc.Int32FromInt32(0x038a)) | libc.Uint32FromUint32(0x80000000),
	91:  libc.Uint32FromInt32(libc.Int32FromInt32(0x038c)),
	92:  libc.Uint32FromInt32(libc.Int32FromInt32(0x038e)),
	93:  libc.Uint32FromInt32(libc.Int32FromInt32(0x038f)),
	94:  libc.Uint32FromInt32(libc.Int32FromInt32(0x0391) | libc.Int32FromInt32(0x40000000)),
	95:  libc.Uint32FromInt32(libc.Int32FromInt32(0x03a1)) | libc.Uint32FromUint32(0x80000000),
	96:  libc.Uint32FromInt32(libc.Int32FromInt32(0x03a3) | libc.Int32FromInt32(0x40000000)),
	97:  libc.Uint32FromInt32(libc.Int32FromInt32(0x03ab)) | libc.Uint32FromUint32(0x80000000),
	98:  libc.Uint32FromInt32(libc.Int32FromInt32(0x03c2)),
	99:  libc.Uint32FromInt32(libc.Int32FromInt32(0x03cf)),
	100: libc.Uint32FromInt32(libc.Int32FromInt32(0x03d0)),
	101: libc.Uint32FromInt32(libc.Int32FromInt32(0x03d1)),
	102: libc.Uint32FromInt32(libc.Int32FromInt32(0x03d5)),
	103: libc.Uint32FromInt32(libc.Int32FromInt32(0x03d6)),
	104: libc.Uint32FromInt32(libc.Int32FromInt32(0x03d8) | libc.Int32FromInt32(0x40000000)),
	105: libc.Uint32FromInt32(libc.Int32FromInt32(0x03ee)) | libc.Uint32FromUint32(0x80000000),
	106: libc.Uint32FromInt32(libc.Int32FromInt32(0x03f0)),
	107: libc.Uint32FromInt32(libc.Int32FromInt32(0x03f1)),
	108: libc.Uint32FromInt32(libc.Int32FromInt32(0x03f4)),
	109: libc.Uint32FromInt32(libc.Int32FromInt32(0x03f5)),
	110: libc.Uint32FromInt32(libc.Int32FromInt32(0x03f7)),
	111: libc.Uint32FromInt32(libc.Int32FromInt32(0x03f9)),
	112: libc.Uint32FromInt32(libc.Int32FromInt32(0x03fa)),
	113: libc.Uint32FromInt32(libc.Int32FromInt32(0x03fd) | libc.Int32FromInt32(0x40000000)),
	114: libc.Uint32FromInt32(libc.Int32FromInt32(0x03ff)) | libc.Uint32FromUint32(0x80000000),
	115: libc.Uint32FromInt32(libc.Int32FromInt32(0x0400) | libc.Int32FromInt32(0x40000000)),
	116: libc.Uint32FromInt32(libc.Int32FromInt32(0x040f)) | libc.Uint32FromUint32(0x80000000),
	117: libc.Uint32FromInt32(libc.Int32FromInt32(0x0410) | libc.Int32FromInt32(0x40000000)),
	118: libc.Uint32FromInt32(libc.Int32FromInt32(0x042f)) | libc.Uint32FromUint32(0x80000000),
	119: libc.Uint32FromInt32(libc.Int32FromInt32(0x0460) | libc.Int32FromInt32(0x40000000)),
	120: libc.Uint32FromInt32(libc.Int32FromInt32(0x0480)) | libc.Uint32FromUint32(0x80000000),
	121: libc.Uint32FromInt32(libc.Int32FromInt32(0x048a) | libc.Int32FromInt32(0x40000000)),
	122: libc.Uint32FromInt32(libc.Int32FromInt32(0x04be)) | libc.Uint32FromUint32(0x80000000),
	123: libc.Uint32FromInt32(libc.Int32FromInt32(0x04c0)),
	124: libc.Uint32FromInt32(libc.Int32FromInt32(0x04c1) | libc.Int32FromInt32(0x40000000)),
	125: libc.Uint32FromInt32(libc.Int32FromInt32(0x04cd)) | libc.Uint32FromUint32(0x80000000),
	126: libc.Uint32FromInt32(libc.Int32FromInt32(0x04d0) | libc.Int32FromInt32(0x40000000)),
	127: libc.Uint32FromInt32(libc.Int32FromInt32(0x052e)) | libc.Uint32FromUint32(0x80000000),
	128: libc.Uint32FromInt32(libc.Int32FromInt32(0x0531) | libc.Int32FromInt32(0x40000000)),
	129: libc.Uint32FromInt32(libc.Int32FromInt32(0x0556)) | libc.Uint32FromUint32(0x80000000),
	130: libc.Uint32FromInt32(libc.Int32FromInt32(0x10a0) | libc.Int32FromInt32(0x40000000)),
	131: libc.Uint32FromInt32(libc.Int32FromInt32(0x10c5)) | libc.Uint32FromUint32(0x80000000),
	132: libc.Uint32FromInt32(libc.Int32FromInt32(0x10c7)),
	133: libc.Uint32FromInt32(libc.Int32FromInt32(0x10cd)),
	134: libc.Uint32FromInt32(libc.Int32FromInt32(0x13f8) | libc.Int32FromInt32(0x40000000)),
	135: libc.Uint32FromInt32(libc.Int32FromInt32(0x13fd)) | libc.Uint32FromUint32(0x80000000),
	136: libc.Uint32FromInt32(libc.Int32FromInt32(0x1c80)),
	137: libc.Uint32FromInt32(libc.Int32FromInt32(0x1c81)),
	138: libc.Uint32FromInt32(libc.Int32FromInt32(0x1c82)),
	139: libc.Uint32FromInt32(libc.Int32FromInt32(0x1c83)),
	140: libc.Uint32FromInt32(libc.Int32FromInt32(0x1c84)),
	141: libc.Uint32FromInt32(libc.Int32FromInt32(0x1c85)),
	142: libc.Uint32FromInt32(libc.Int32FromInt32(0x1c86)),
	143: libc.Uint32FromInt32(libc.Int32FromInt32(0x1c87)),
	144: libc.Uint32FromInt32(libc.Int32FromInt32(0x1c88)),
	145: libc.Uint32FromInt32(libc.Int32FromInt32(0x1c90) | libc.Int32FromInt32(0x40000000)),
	146: libc.Uint32FromInt32(libc.Int32FromInt32(0x1cba)) | libc.Uint32FromUint32(0x80000000),
	147: libc.Uint32FromInt32(libc.Int32FromInt32(0x1cbd) | libc.Int32FromInt32(0x40000000)),
	148: libc.Uint32FromInt32(libc.Int32FromInt32(0x1cbf)) | libc.Uint32FromUint32(0x80000000),
	149: libc.Uint32FromInt32(libc.Int32FromInt32(0x1e00) | libc.Int32FromInt32(0x40000000)),
	150: libc.Uint32FromInt32(libc.Int32FromInt32(0x1e94)) | libc.Uint32FromUint32(0x80000000),
	151: libc.Uint32FromInt32(libc.Int32FromInt32(0x1e9b)),
	152: libc.Uint32FromInt32(libc.Int32FromInt32(0x1ea0) | libc.Int32FromInt32(0x40000000)),
	153: libc.Uint32FromInt32(libc.Int32FromInt32(0x1efe)) | libc.Uint32FromUint32(0x80000000),
	154: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f08) | libc.Int32FromInt32(0x40000000)),
	155: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f0f)) | libc.Uint32FromUint32(0x80000000),
	156: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f18) | libc.Int32FromInt32(0x40000000)),
	157: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f1d)) | libc.Uint32FromUint32(0x80000000),
	158: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f28) | libc.Int32FromInt32(0x40000000)),
	159: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f2f)) | libc.Uint32FromUint32(0x80000000),
	160: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f38) | libc.Int32FromInt32(0x40000000)),
	161: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f3f)) | libc.Uint32FromUint32(0x80000000),
	162: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f48) | libc.Int32FromInt32(0x40000000)),
	163: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f4d)) | libc.Uint32FromUint32(0x80000000),
	164: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f59)),
	165: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f5b)),
	166: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f5d)),
	167: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f5f)),
	168: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f68) | libc.Int32FromInt32(0x40000000)),
	169: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f6f)) | libc.Uint32FromUint32(0x80000000),
	170: libc.Uint32FromInt32(libc.Int32FromInt32(0x1fb8)),
	171: libc.Uint32FromInt32(libc.Int32FromInt32(0x1fb9)),
	172: libc.Uint32FromInt32(libc.Int32FromInt32(0x1fba)),
	173: libc.Uint32FromInt32(libc.Int32FromInt32(0x1fbb)),
	174: libc.Uint32FromInt32(libc.Int32FromInt32(0x1fbe)),
	175: libc.Uint32FromInt32(libc.Int32FromInt32(0x1fc8) | libc.Int32FromInt32(0x40000000)),
	176: libc.Uint32FromInt32(libc.Int32FromInt32(0x1fcb)) | libc.Uint32FromUint32(0x80000000),
	177: libc.Uint32FromInt32(libc.Int32FromInt32(0x1fd8)),
	178: libc.Uint32FromInt32(libc.Int32FromInt32(0x1fd9)),
	179: libc.Uint32FromInt32(libc.Int32FromInt32(0x1fda)),
	180: libc.Uint32FromInt32(libc.Int32FromInt32(0x1fdb)),
	181: libc.Uint32FromInt32(libc.Int32FromInt32(0x1fe8)),
	182: libc.Uint32FromInt32(libc.Int32FromInt32(0x1fe9)),
	183: libc.Uint32FromInt32(libc.Int32FromInt32(0x1fea)),
	184: libc.Uint32FromInt32(libc.Int32FromInt32(0x1feb)),
	185: libc.Uint32FromInt32(libc.Int32FromInt32(0x1fec)),
	186: libc.Uint32FromInt32(libc.Int32FromInt32(0x1ff8)),
	187: libc.Uint32FromInt32(libc.Int32FromInt32(0x1ff9)),
	188: libc.Uint32FromInt32(libc.Int32FromInt32(0x1ffa)),
	189: libc.Uint32FromInt32(libc.Int32FromInt32(0x1ffb)),
	190: libc.Uint32FromInt32(libc.Int32FromInt32(0x2126)),
	191: libc.Uint32FromInt32(libc.Int32FromInt32(0x212a)),
	192: libc.Uint32FromInt32(libc.Int32FromInt32(0x212b)),
	193: libc.Uint32FromInt32(libc.Int32FromInt32(0x2132)),
	194: libc.Uint32FromInt32(libc.Int32FromInt32(0x2160) | libc.Int32FromInt32(0x40000000)),
	195: libc.Uint32FromInt32(libc.Int32FromInt32(0x216f)) | libc.Uint32FromUint32(0x80000000),
	196: libc.Uint32FromInt32(libc.Int32FromInt32(0x2183)),
	197: libc.Uint32FromInt32(libc.Int32FromInt32(0x24b6) | libc.Int32FromInt32(0x40000000)),
	198: libc.Uint32FromInt32(libc.Int32FromInt32(0x24cf)) | libc.Uint32FromUint32(0x80000000),
	199: libc.Uint32FromInt32(libc.Int32FromInt32(0x2c00) | libc.Int32FromInt32(0x40000000)),
	200: libc.Uint32FromInt32(libc.Int32FromInt32(0x2c2f)) | libc.Uint32FromUint32(0x80000000),
	201: libc.Uint32FromInt32(libc.Int32FromInt32(0x2c60)),
	202: libc.Uint32FromInt32(libc.Int32FromInt32(0x2c62)),
	203: libc.Uint32FromInt32(libc.Int32FromInt32(0x2c63)),
	204: libc.Uint32FromInt32(libc.Int32FromInt32(0x2c64)),
	205: libc.Uint32FromInt32(libc.Int32FromInt32(0x2c67) | libc.Int32FromInt32(0x40000000)),
	206: libc.Uint32FromInt32(libc.Int32FromInt32(0x2c6b)) | libc.Uint32FromUint32(0x80000000),
	207: libc.Uint32FromInt32(libc.Int32FromInt32(0x2c6d)),
	208: libc.Uint32FromInt32(libc.Int32FromInt32(0x2c6e)),
	209: libc.Uint32FromInt32(libc.Int32FromInt32(0x2c6f)),
	210: libc.Uint32FromInt32(libc.Int32FromInt32(0x2c70)),
	211: libc.Uint32FromInt32(libc.Int32FromInt32(0x2c72)),
	212: libc.Uint32FromInt32(libc.Int32FromInt32(0x2c75)),
	213: libc.Uint32FromInt32(libc.Int32FromInt32(0x2c7e)),
	214: libc.Uint32FromInt32(libc.Int32FromInt32(0x2c7f)),
	215: libc.Uint32FromInt32(libc.Int32FromInt32(0x2c80) | libc.Int32FromInt32(0x40000000)),
	216: libc.Uint32FromInt32(libc.Int32FromInt32(0x2ce2)) | libc.Uint32FromUint32(0x80000000),
	217: libc.Uint32FromInt32(libc.Int32FromInt32(0x2ceb)),
	218: libc.Uint32FromInt32(libc.Int32FromInt32(0x2ced)),
	219: libc.Uint32FromInt32(libc.Int32FromInt32(0x2cf2)),
	220: libc.Uint32FromInt32(libc.Int32FromInt32(0xa640) | libc.Int32FromInt32(0x40000000)),
	221: libc.Uint32FromInt32(libc.Int32FromInt32(0xa66c)) | libc.Uint32FromUint32(0x80000000),
	222: libc.Uint32FromInt32(libc.Int32FromInt32(0xa680) | libc.Int32FromInt32(0x40000000)),
	223: libc.Uint32FromInt32(libc.Int32FromInt32(0xa69a)) | libc.Uint32FromUint32(0x80000000),
	224: libc.Uint32FromInt32(libc.Int32FromInt32(0xa722) | libc.Int32FromInt32(0x40000000)),
	225: libc.Uint32FromInt32(libc.Int32FromInt32(0xa72e)) | libc.Uint32FromUint32(0x80000000),
	226: libc.Uint32FromInt32(libc.Int32FromInt32(0xa732) | libc.Int32FromInt32(0x40000000)),
	227: libc.Uint32FromInt32(libc.Int32FromInt32(0xa76e)) | libc.Uint32FromUint32(0x80000000),
	228: libc.Uint32FromInt32(libc.Int32FromInt32(0xa779)),
	229: libc.Uint32FromInt32(libc.Int32FromInt32(0xa77b)),
	230: libc.Uint32FromInt32(libc.Int32FromInt32(0xa77d)),
	231: libc.Uint32FromInt32(libc.Int32FromInt32(0xa77e) | libc.Int32FromInt32(0x40000000)),
	232: libc.Uint32FromInt32(libc.Int32FromInt32(0xa786)) | libc.Uint32FromUint32(0x80000000),
	233: libc.Uint32FromInt32(libc.Int32FromInt32(0xa78b)),
	234: libc.Uint32FromInt32(libc.Int32FromInt32(0xa78d)),
	235: libc.Uint32FromInt32(libc.Int32FromInt32(0xa790)),
	236: libc.Uint32FromInt32(libc.Int32FromInt32(0xa792)),
	237: libc.Uint32FromInt32(libc.Int32FromInt32(0xa796) | libc.Int32FromInt32(0x40000000)),
	238: libc.Uint32FromInt32(libc.Int32FromInt32(0xa7a8)) | libc.Uint32FromUint32(0x80000000),
	239: libc.Uint32FromInt32(libc.Int32FromInt32(0xa7aa)),
	240: libc.Uint32FromInt32(libc.Int32FromInt32(0xa7ab)),
	241: libc.Uint32FromInt32(libc.Int32FromInt32(0xa7ac)),
	242: libc.Uint32FromInt32(libc.Int32FromInt32(0xa7ad)),
	243: libc.Uint32FromInt32(libc.Int32FromInt32(0xa7ae)),
	244: libc.Uint32FromInt32(libc.Int32FromInt32(0xa7b0)),
	245: libc.Uint32FromInt32(libc.Int32FromInt32(0xa7b1)),
	246: libc.Uint32FromInt32(libc.Int32FromInt32(0xa7b2)),
	247: libc.Uint32FromInt32(libc.Int32FromInt32(0xa7b3)),
	248: libc.Uint32FromInt32(libc.Int32FromInt32(0xa7b4) | libc.Int32FromInt32(0x40000000)),
	249: libc.Uint32FromInt32(libc.Int32FromInt32(0xa7c2)) | libc.Uint32FromUint32(0x80000000),
	250: libc.Uint32FromInt32(libc.Int32FromInt32(0xa7c4)),
	251: libc.Uint32FromInt32(libc.Int32FromInt32(0xa7c5)),
	252: libc.Uint32FromInt32(libc.Int32FromInt32(0xa7c6)),
	253: libc.Uint32FromInt32(libc.Int32FromInt32(0xa7c7)),
	254: libc.Uint32FromInt32(libc.Int32FromInt32(0xa7c9)),
	255: libc.Uint32FromInt32(libc.Int32FromInt32(0xa7d0)),
	256: libc.Uint32FromInt32(libc.Int32FromInt32(0xa7d6)),
	257: libc.Uint32FromInt32(libc.Int32FromInt32(0xa7d8)),
	258: libc.Uint32FromInt32(libc.Int32FromInt32(0xa7f5)),
	259: libc.Uint32FromInt32(libc.Int32FromInt32(0xab70) | libc.Int32FromInt32(0x40000000)),
	260: libc.Uint32FromInt32(libc.Int32FromInt32(0xabbf)) | libc.Uint32FromUint32(0x80000000),
	261: libc.Uint32FromInt32(libc.Int32FromInt32(0xff21) | libc.Int32FromInt32(0x40000000)),
	262: libc.Uint32FromInt32(libc.Int32FromInt32(0xff3a)) | libc.Uint32FromUint32(0x80000000),
	263: libc.Uint32FromInt32(libc.Int32FromInt32(0x10400) | libc.Int32FromInt32(0x40000000)),
	264: libc.Uint32FromInt32(libc.Int32FromInt32(0x10427)) | libc.Uint32FromUint32(0x80000000),
	265: libc.Uint32FromInt32(libc.Int32FromInt32(0x104b0) | libc.Int32FromInt32(0x40000000)),
	266: libc.Uint32FromInt32(libc.Int32FromInt32(0x104d3)) | libc.Uint32FromUint32(0x80000000),
	267: libc.Uint32FromInt32(libc.Int32FromInt32(0x10570) | libc.Int32FromInt32(0x40000000)),
	268: libc.Uint32FromInt32(libc.Int32FromInt32(0x1057a)) | libc.Uint32FromUint32(0x80000000),
	269: libc.Uint32FromInt32(libc.Int32FromInt32(0x1057c) | libc.Int32FromInt32(0x40000000)),
	270: libc.Uint32FromInt32(libc.Int32FromInt32(0x1058a)) | libc.Uint32FromUint32(0x80000000),
	271: libc.Uint32FromInt32(libc.Int32FromInt32(0x1058c) | libc.Int32FromInt32(0x40000000)),
	272: libc.Uint32FromInt32(libc.Int32FromInt32(0x10592)) | libc.Uint32FromUint32(0x80000000),
	273: libc.Uint32FromInt32(libc.Int32FromInt32(0x10594)),
	274: libc.Uint32FromInt32(libc.Int32FromInt32(0x10595)),
	275: libc.Uint32FromInt32(libc.Int32FromInt32(0x10c80) | libc.Int32FromInt32(0x40000000)),
	276: libc.Uint32FromInt32(libc.Int32FromInt32(0x10cb2)) | libc.Uint32FromUint32(0x80000000),
	277: libc.Uint32FromInt32(libc.Int32FromInt32(0x118a0) | libc.Int32FromInt32(0x40000000)),
	278: libc.Uint32FromInt32(libc.Int32FromInt32(0x118bf)) | libc.Uint32FromUint32(0x80000000),
	279: libc.Uint32FromInt32(libc.Int32FromInt32(0x16e40) | libc.Int32FromInt32(0x40000000)),
	280: libc.Uint32FromInt32(libc.Int32FromInt32(0x16e5f)) | libc.Uint32FromUint32(0x80000000),
	281: libc.Uint32FromInt32(libc.Int32FromInt32(0x1e900) | libc.Int32FromInt32(0x40000000)),
	282: libc.Uint32FromInt32(libc.Int32FromInt32(0x1e921)) | libc.Uint32FromUint32(0x80000000),
}

var FOLD_MAP_1_DATA = [283]uint32{
	0:   uint32(0x0061),
	1:   uint32(0x007a),
	2:   uint32(0x03bc),
	3:   uint32(0x00e0),
	4:   uint32(0x00f6),
	5:   uint32(0x00f8),
	6:   uint32(0x00fe),
	7:   uint32(0x0101),
	8:   uint32(0x012f),
	9:   uint32(0x0133),
	10:  uint32(0x0137),
	11:  uint32(0x013a),
	12:  uint32(0x0148),
	13:  uint32(0x014b),
	14:  uint32(0x0177),
	15:  uint32(0x00ff),
	16:  uint32(0x017a),
	17:  uint32(0x017e),
	18:  uint32(0x0073),
	19:  uint32(0x0253),
	20:  uint32(0x0183),
	21:  uint32(0x0185),
	22:  uint32(0x0254),
	23:  uint32(0x0188),
	24:  uint32(0x0256),
	25:  uint32(0x0257),
	26:  uint32(0x018c),
	27:  uint32(0x01dd),
	28:  uint32(0x0259),
	29:  uint32(0x025b),
	30:  uint32(0x0192),
	31:  uint32(0x0260),
	32:  uint32(0x0263),
	33:  uint32(0x0269),
	34:  uint32(0x0268),
	35:  uint32(0x0199),
	36:  uint32(0x026f),
	37:  uint32(0x0272),
	38:  uint32(0x0275),
	39:  uint32(0x01a1),
	40:  uint32(0x01a5),
	41:  uint32(0x0280),
	42:  uint32(0x01a8),
	43:  uint32(0x0283),
	44:  uint32(0x01ad),
	45:  uint32(0x0288),
	46:  uint32(0x01b0),
	47:  uint32(0x028a),
	48:  uint32(0x028b),
	49:  uint32(0x01b4),
	50:  uint32(0x01b6),
	51:  uint32(0x0292),
	52:  uint32(0x01b9),
	53:  uint32(0x01bd),
	54:  uint32(0x01c6),
	55:  uint32(0x01c6),
	56:  uint32(0x01c9),
	57:  uint32(0x01c9),
	58:  uint32(0x01cc),
	59:  uint32(0x01cc),
	60:  uint32(0x01dc),
	61:  uint32(0x01df),
	62:  uint32(0x01ef),
	63:  uint32(0x01f3),
	64:  uint32(0x01f3),
	65:  uint32(0x01f5),
	66:  uint32(0x0195),
	67:  uint32(0x01bf),
	68:  uint32(0x01f9),
	69:  uint32(0x021f),
	70:  uint32(0x019e),
	71:  uint32(0x0223),
	72:  uint32(0x0233),
	73:  uint32(0x2c65),
	74:  uint32(0x023c),
	75:  uint32(0x019a),
	76:  uint32(0x2c66),
	77:  uint32(0x0242),
	78:  uint32(0x0180),
	79:  uint32(0x0289),
	80:  uint32(0x028c),
	81:  uint32(0x0247),
	82:  uint32(0x024f),
	83:  uint32(0x03b9),
	84:  uint32(0x0371),
	85:  uint32(0x0373),
	86:  uint32(0x0377),
	87:  uint32(0x03f3),
	88:  uint32(0x03ac),
	89:  uint32(0x03ad),
	90:  uint32(0x03af),
	91:  uint32(0x03cc),
	92:  uint32(0x03cd),
	93:  uint32(0x03ce),
	94:  uint32(0x03b1),
	95:  uint32(0x03c1),
	96:  uint32(0x03c3),
	97:  uint32(0x03cb),
	98:  uint32(0x03c3),
	99:  uint32(0x03d7),
	100: uint32(0x03b2),
	101: uint32(0x03b8),
	102: uint32(0x03c6),
	103: uint32(0x03c0),
	104: uint32(0x03d9),
	105: uint32(0x03ef),
	106: uint32(0x03ba),
	107: uint32(0x03c1),
	108: uint32(0x03b8),
	109: uint32(0x03b5),
	110: uint32(0x03f8),
	111: uint32(0x03f2),
	112: uint32(0x03fb),
	113: uint32(0x037b),
	114: uint32(0x037d),
	115: uint32(0x0450),
	116: uint32(0x045f),
	117: uint32(0x0430),
	118: uint32(0x044f),
	119: uint32(0x0461),
	120: uint32(0x0481),
	121: uint32(0x048b),
	122: uint32(0x04bf),
	123: uint32(0x04cf),
	124: uint32(0x04c2),
	125: uint32(0x04ce),
	126: uint32(0x04d1),
	127: uint32(0x052f),
	128: uint32(0x0561),
	129: uint32(0x0586),
	130: uint32(0x2d00),
	131: uint32(0x2d25),
	132: uint32(0x2d27),
	133: uint32(0x2d2d),
	134: uint32(0x13f0),
	135: uint32(0x13f5),
	136: uint32(0x0432),
	137: uint32(0x0434),
	138: uint32(0x043e),
	139: uint32(0x0441),
	140: uint32(0x0442),
	141: uint32(0x0442),
	142: uint32(0x044a),
	143: uint32(0x0463),
	144: uint32(0xa64b),
	145: uint32(0x10d0),
	146: uint32(0x10fa),
	147: uint32(0x10fd),
	148: uint32(0x10ff),
	149: uint32(0x1e01),
	150: uint32(0x1e95),
	151: uint32(0x1e61),
	152: uint32(0x1ea1),
	153: uint32(0x1eff),
	154: uint32(0x1f00),
	155: uint32(0x1f07),
	156: uint32(0x1f10),
	157: uint32(0x1f15),
	158: uint32(0x1f20),
	159: uint32(0x1f27),
	160: uint32(0x1f30),
	161: uint32(0x1f37),
	162: uint32(0x1f40),
	163: uint32(0x1f45),
	164: uint32(0x1f51),
	165: uint32(0x1f53),
	166: uint32(0x1f55),
	167: uint32(0x1f57),
	168: uint32(0x1f60),
	169: uint32(0x1f67),
	170: uint32(0x1fb0),
	171: uint32(0x1fb1),
	172: uint32(0x1f70),
	173: uint32(0x1f71),
	174: uint32(0x03b9),
	175: uint32(0x1f72),
	176: uint32(0x1f75),
	177: uint32(0x1fd0),
	178: uint32(0x1fd1),
	179: uint32(0x1f76),
	180: uint32(0x1f77),
	181: uint32(0x1fe0),
	182: uint32(0x1fe1),
	183: uint32(0x1f7a),
	184: uint32(0x1f7b),
	185: uint32(0x1fe5),
	186: uint32(0x1f78),
	187: uint32(0x1f79),
	188: uint32(0x1f7c),
	189: uint32(0x1f7d),
	190: uint32(0x03c9),
	191: uint32(0x006b),
	192: uint32(0x00e5),
	193: uint32(0x214e),
	194: uint32(0x2170),
	195: uint32(0x217f),
	196: uint32(0x2184),
	197: uint32(0x24d0),
	198: uint32(0x24e9),
	199: uint32(0x2c30),
	200: uint32(0x2c5f),
	201: uint32(0x2c61),
	202: uint32(0x026b),
	203: uint32(0x1d7d),
	204: uint32(0x027d),
	205: uint32(0x2c68),
	206: uint32(0x2c6c),
	207: uint32(0x0251),
	208: uint32(0x0271),
	209: uint32(0x0250),
	210: uint32(0x0252),
	211: uint32(0x2c73),
	212: uint32(0x2c76),
	213: uint32(0x023f),
	214: uint32(0x0240),
	215: uint32(0x2c81),
	216: uint32(0x2ce3),
	217: uint32(0x2cec),
	218: uint32(0x2cee),
	219: uint32(0x2cf3),
	220: uint32(0xa641),
	221: uint32(0xa66d),
	222: uint32(0xa681),
	223: uint32(0xa69b),
	224: uint32(0xa723),
	225: uint32(0xa72f),
	226: uint32(0xa733),
	227: uint32(0xa76f),
	228: uint32(0xa77a),
	229: uint32(0xa77c),
	230: uint32(0x1d79),
	231: uint32(0xa77f),
	232: uint32(0xa787),
	233: uint32(0xa78c),
	234: uint32(0x0265),
	235: uint32(0xa791),
	236: uint32(0xa793),
	237: uint32(0xa797),
	238: uint32(0xa7a9),
	239: uint32(0x0266),
	240: uint32(0x025c),
	241: uint32(0x0261),
	242: uint32(0x026c),
	243: uint32(0x026a),
	244: uint32(0x029e),
	245: uint32(0x0287),
	246: uint32(0x029d),
	247: uint32(0xab53),
	248: uint32(0xa7b5),
	249: uint32(0xa7c3),
	250: uint32(0xa794),
	251: uint32(0x0282),
	252: uint32(0x1d8e),
	253: uint32(0xa7c8),
	254: uint32(0xa7ca),
	255: uint32(0xa7d1),
	256: uint32(0xa7d7),
	257: uint32(0xa7d9),
	258: uint32(0xa7f6),
	259: uint32(0x13a0),
	260: uint32(0x13ef),
	261: uint32(0xff41),
	262: uint32(0xff5a),
	263: uint32(0x10428),
	264: uint32(0x1044f),
	265: uint32(0x104d8),
	266: uint32(0x104fb),
	267: uint32(0x10597),
	268: uint32(0x105a1),
	269: uint32(0x105a3),
	270: uint32(0x105b1),
	271: uint32(0x105b3),
	272: uint32(0x105b9),
	273: uint32(0x105bb),
	274: uint32(0x105bc),
	275: uint32(0x10cc0),
	276: uint32(0x10cf2),
	277: uint32(0x118c0),
	278: uint32(0x118df),
	279: uint32(0x16e60),
	280: uint32(0x16e7f),
	281: uint32(0x1e922),
	282: uint32(0x1e943),
}

var FOLD_MAP_2 = [52]uint32{
	0:  libc.Uint32FromInt32(libc.Int32FromInt32(0x00df)),
	1:  libc.Uint32FromInt32(libc.Int32FromInt32(0x0130)),
	2:  libc.Uint32FromInt32(libc.Int32FromInt32(0x0149)),
	3:  libc.Uint32FromInt32(libc.Int32FromInt32(0x01f0)),
	4:  libc.Uint32FromInt32(libc.Int32FromInt32(0x0587)),
	5:  libc.Uint32FromInt32(libc.Int32FromInt32(0x1e96)),
	6:  libc.Uint32FromInt32(libc.Int32FromInt32(0x1e97)),
	7:  libc.Uint32FromInt32(libc.Int32FromInt32(0x1e98)),
	8:  libc.Uint32FromInt32(libc.Int32FromInt32(0x1e99)),
	9:  libc.Uint32FromInt32(libc.Int32FromInt32(0x1e9a)),
	10: libc.Uint32FromInt32(libc.Int32FromInt32(0x1e9e)),
	11: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f50)),
	12: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f80) | libc.Int32FromInt32(0x40000000)),
	13: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f87)) | libc.Uint32FromUint32(0x80000000),
	14: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f88) | libc.Int32FromInt32(0x40000000)),
	15: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f8f)) | libc.Uint32FromUint32(0x80000000),
	16: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f90) | libc.Int32FromInt32(0x40000000)),
	17: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f97)) | libc.Uint32FromUint32(0x80000000),
	18: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f98) | libc.Int32FromInt32(0x40000000)),
	19: libc.Uint32FromInt32(libc.Int32FromInt32(0x1f9f)) | libc.Uint32FromUint32(0x80000000),
	20: libc.Uint32FromInt32(libc.Int32FromInt32(0x1fa0) | libc.Int32FromInt32(0x40000000)),
	21: libc.Uint32FromInt32(libc.Int32FromInt32(0x1fa7)) | libc.Uint32FromUint32(0x80000000),
	22: libc.Uint32FromInt32(libc.Int32FromInt32(0x1fa8) | libc.Int32FromInt32(0x40000000)),
	23: libc.Uint32FromInt32(libc.Int32FromInt32(0x1faf)) | libc.Uint32FromUint32(0x80000000),
	24: libc.Uint32FromInt32(libc.Int32FromInt32(0x1fb2)),
	25: libc.Uint32FromInt32(libc.Int32FromInt32(0x1fb3)),
	26: libc.Uint32FromInt32(libc.Int32FromInt32(0x1fb4)),
	27: libc.Uint32FromInt32(libc.Int32FromInt32(0x1fb6)),
	28: libc.Uint32FromInt32(libc.Int32FromInt32(0x1fbc)),
	29: libc.Uint32FromInt32(libc.Int32FromInt32(0x1fc2)),
	30: libc.Uint32FromInt32(libc.Int32FromInt32(0x1fc3)),
	31: libc.Uint32FromInt32(libc.Int32FromInt32(0x1fc4)),
	32: libc.Uint32FromInt32(libc.Int32FromInt32(0x1fc6)),
	33: libc.Uint32FromInt32(libc.Int32FromInt32(0x1fcc)),
	34: libc.Uint32FromInt32(libc.Int32FromInt32(0x1fd6)),
	35: libc.Uint32FromInt32(libc.Int32FromInt32(0x1fe4)),
	36: libc.Uint32FromInt32(libc.Int32FromInt32(0x1fe6)),
	37: libc.Uint32FromInt32(libc.Int32FromInt32(0x1ff2)),
	38: libc.Uint32FromInt32(libc.Int32FromInt32(0x1ff3)),
	39: libc.Uint32FromInt32(libc.Int32FromInt32(0x1ff4)),
	40: libc.Uint32FromInt32(libc.Int32FromInt32(0x1ff6)),
	41: libc.Uint32FromInt32(libc.Int32FromInt32(0x1ffc)),
	42: libc.Uint32FromInt32(libc.Int32FromInt32(0xfb00)),
	43: libc.Uint32FromInt32(libc.Int32FromInt32(0xfb01)),
	44: libc.Uint32FromInt32(libc.Int32FromInt32(0xfb02)),
	45: libc.Uint32FromInt32(libc.Int32FromInt32(0xfb05)),
	46: libc.Uint32FromInt32(libc.Int32FromInt32(0xfb06)),
	47: libc.Uint32FromInt32(libc.Int32FromInt32(0xfb13)),
	48: libc.Uint32FromInt32(libc.Int32FromInt32(0xfb14)),
	49: libc.Uint32FromInt32(libc.Int32FromInt32(0xfb15)),
	50: libc.Uint32FromInt32(libc.Int32FromInt32(0xfb16)),
	51: libc.Uint32FromInt32(libc.Int32FromInt32(0xfb17)),
}

var FOLD_MAP_2_DATA = [104]uint32{
	0:   uint32(0x0073),
	1:   uint32(0x0073),
	2:   uint32(0x0069),
	3:   uint32(0x0307),
	4:   uint32(0x02bc),
	5:   uint32(0x006e),
	6:   uint32(0x006a),
	7:   uint32(0x030c),
	8:   uint32(0x0565),
	9:   uint32(0x0582),
	10:  uint32(0x0068),
	11:  uint32(0x0331),
	12:  uint32(0x0074),
	13:  uint32(0x0308),
	14:  uint32(0x0077),
	15:  uint32(0x030a),
	16:  uint32(0x0079),
	17:  uint32(0x030a),
	18:  uint32(0x0061),
	19:  uint32(0x02be),
	20:  uint32(0x0073),
	21:  uint32(0x0073),
	22:  uint32(0x03c5),
	23:  uint32(0x0313),
	24:  uint32(0x1f00),
	25:  uint32(0x03b9),
	26:  uint32(0x1f07),
	27:  uint32(0x03b9),
	28:  uint32(0x1f00),
	29:  uint32(0x03b9),
	30:  uint32(0x1f07),
	31:  uint32(0x03b9),
	32:  uint32(0x1f20),
	33:  uint32(0x03b9),
	34:  uint32(0x1f27),
	35:  uint32(0x03b9),
	36:  uint32(0x1f20),
	37:  uint32(0x03b9),
	38:  uint32(0x1f27),
	39:  uint32(0x03b9),
	40:  uint32(0x1f60),
	41:  uint32(0x03b9),
	42:  uint32(0x1f67),
	43:  uint32(0x03b9),
	44:  uint32(0x1f60),
	45:  uint32(0x03b9),
	46:  uint32(0x1f67),
	47:  uint32(0x03b9),
	48:  uint32(0x1f70),
	49:  uint32(0x03b9),
	50:  uint32(0x03b1),
	51:  uint32(0x03b9),
	52:  uint32(0x03ac),
	53:  uint32(0x03b9),
	54:  uint32(0x03b1),
	55:  uint32(0x0342),
	56:  uint32(0x03b1),
	57:  uint32(0x03b9),
	58:  uint32(0x1f74),
	59:  uint32(0x03b9),
	60:  uint32(0x03b7),
	61:  uint32(0x03b9),
	62:  uint32(0x03ae),
	63:  uint32(0x03b9),
	64:  uint32(0x03b7),
	65:  uint32(0x0342),
	66:  uint32(0x03b7),
	67:  uint32(0x03b9),
	68:  uint32(0x03b9),
	69:  uint32(0x0342),
	70:  uint32(0x03c1),
	71:  uint32(0x0313),
	72:  uint32(0x03c5),
	73:  uint32(0x0342),
	74:  uint32(0x1f7c),
	75:  uint32(0x03b9),
	76:  uint32(0x03c9),
	77:  uint32(0x03b9),
	78:  uint32(0x03ce),
	79:  uint32(0x03b9),
	80:  uint32(0x03c9),
	81:  uint32(0x0342),
	82:  uint32(0x03c9),
	83:  uint32(0x03b9),
	84:  uint32(0x0066),
	85:  uint32(0x0066),
	86:  uint32(0x0066),
	87:  uint32(0x0069),
	88:  uint32(0x0066),
	89:  uint32(0x006c),
	90:  uint32(0x0073),
	91:  uint32(0x0074),
	92:  uint32(0x0073),
	93:  uint32(0x0074),
	94:  uint32(0x0574),
	95:  uint32(0x0576),
	96:  uint32(0x0574),
	97:  uint32(0x0565),
	98:  uint32(0x0574),
	99:  uint32(0x056b),
	100: uint32(0x057e),
	101: uint32(0x0576),
	102: uint32(0x0574),
	103: uint32(0x056d),
}

var FOLD_MAP_3 = [16]uint32{
	0:  libc.Uint32FromInt32(libc.Int32FromInt32(0x0390)),
	1:  libc.Uint32FromInt32(libc.Int32FromInt32(0x03b0)),
	2:  libc.Uint32FromInt32(libc.Int32FromInt32(0x1f52)),
	3:  libc.Uint32FromInt32(libc.Int32FromInt32(0x1f54)),
	4:  libc.Uint32FromInt32(libc.Int32FromInt32(0x1f56)),
	5:  libc.Uint32FromInt32(libc.Int32FromInt32(0x1fb7)),
	6:  libc.Uint32FromInt32(libc.Int32FromInt32(0x1fc7)),
	7:  libc.Uint32FromInt32(libc.Int32FromInt32(0x1fd2)),
	8:  libc.Uint32FromInt32(libc.Int32FromInt32(0x1fd3)),
	9:  libc.Uint32FromInt32(libc.Int32FromInt32(0x1fd7)),
	10: libc.Uint32FromInt32(libc.Int32FromInt32(0x1fe2)),
	11: libc.Uint32FromInt32(libc.Int32FromInt32(0x1fe3)),
	12: libc.Uint32FromInt32(libc.Int32FromInt32(0x1fe7)),
	13: libc.Uint32FromInt32(libc.Int32FromInt32(0x1ff7)),
	14: libc.Uint32FromInt32(libc.Int32FromInt32(0xfb03)),
	15: libc.Uint32FromInt32(libc.Int32FromInt32(0xfb04)),
}

var FOLD_MAP_3_DATA = [48]uint32{
	0:  uint32(0x03b9),
	1:  uint32(0x0308),
	2:  uint32(0x0301),
	3:  uint32(0x03c5),
	4:  uint32(0x0308),
	5:  uint32(0x0301),
	6:  uint32(0x03c5),
	7:  uint32(0x0313),
	8:  uint32(0x0300),
	9:  uint32(0x03c5),
	10: uint32(0x0313),
	11: uint32(0x0301),
	12: uint32(0x03c5),
	13: uint32(0x0313),
	14: uint32(0x0342),
	15: uint32(0x03b1),
	16: uint32(0x0342),
	17: uint32(0x03b9),
	18: uint32(0x03b7),
	19: uint32(0x0342),
	20: uint32(0x03b9),
	21: uint32(0x03b9),
	22: uint32(0x0308),
	23: uint32(0x0300),
	24: uint32(0x03b9),
	25: uint32(0x0308),
	26: uint32(0x0301),
	27: uint32(0x03b9),
	28: uint32(0x0308),
	29: uint32(0x0342),
	30: uint32(0x03c5),
	31: uint32(0x0308),
	32: uint32(0x0300),
	33: uint32(0x03c5),
	34: uint32(0x0308),
	35: uint32(0x0301),
	36: uint32(0x03c5),
	37: uint32(0x0308),
	38: uint32(0x0342),
	39: uint32(0x03c9),
	40: uint32(0x0342),
	41: uint32(0x03b9),
	42: uint32(0x0066),
	43: uint32(0x0066),
	44: uint32(0x0069),
	45: uint32(0x0066),
	46: uint32(0x0066),
	47: uint32(0x006c),
}

var FOLD_MAP_LIST = [3]struct {
	Fmap1         uintptr
	Fdata         uintptr
	Fmap_size     size_t
	Fn_codepoints uint32
}{
	0: {
		Fmap1:         uintptr(unsafe.Pointer(&FOLD_MAP_1)),
		Fdata:         uintptr(unsafe.Pointer(&FOLD_MAP_1_DATA)),
		Fmap_size:     libc.Uint64FromInt64(1132) / libc.Uint64FromInt64(4),
		Fn_codepoints: uint32(1),
	},
	1: {
		Fmap1:         uintptr(unsafe.Pointer(&FOLD_MAP_2)),
		Fdata:         uintptr(unsafe.Pointer(&FOLD_MAP_2_DATA)),
		Fmap_size:     libc.Uint64FromInt64(208) / libc.Uint64FromInt64(4),
		Fn_codepoints: uint32(2),
	},
	2: {
		Fmap1:         uintptr(unsafe.Pointer(&FOLD_MAP_3)),
		Fdata:         uintptr(unsafe.Pointer(&FOLD_MAP_3_DATA)),
		Fmap_size:     libc.Uint64FromInt64(64) / libc.Uint64FromInt64(4),
		Fn_codepoints: uint32(3),
	},
}

func md_decode_utf8__(tls *libc.TLS, str uintptr, str_size MD_SIZE, p_size uintptr) (r uint32) {
	if !(libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr(str))) <= libc.Int32FromInt32(0x7f)) {
		if libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr(str)))&int32(0xe0) == int32(0xc0) {
			if uint32(1) < str_size && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr(str + 1)))&int32(0xc0) == int32(0x80) {
				if p_size != libc.UintptrFromInt32(0) {
					*(*MD_SIZE)(__ccgo_ptr(p_size)) = uint32(2)
				}
				return uint32(*(*MD_CHAR)(__ccgo_ptr(str)))&uint32(0x1f)<<int32(6) | uint32(*(*MD_CHAR)(__ccgo_ptr(str + 1)))&uint32(0x3f)<<0
			}
		} else {
			if libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr(str)))&int32(0xf0) == int32(0xe0) {
				if uint32(2) < str_size && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr(str + 1)))&int32(0xc0) == int32(0x80) && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr(str + 2)))&int32(0xc0) == int32(0x80) {
					if p_size != libc.UintptrFromInt32(0) {
						*(*MD_SIZE)(__ccgo_ptr(p_size)) = uint32(3)
					}
					return uint32(*(*MD_CHAR)(__ccgo_ptr(str)))&uint32(0x0f)<<int32(12) | uint32(*(*MD_CHAR)(__ccgo_ptr(str + 1)))&uint32(0x3f)<<int32(6) | uint32(*(*MD_CHAR)(__ccgo_ptr(str + 2)))&uint32(0x3f)<<0
				}
			} else {
				if libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr(str)))&int32(0xf8) == int32(0xf0) {
					if uint32(3) < str_size && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr(str + 1)))&int32(0xc0) == int32(0x80) && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr(str + 2)))&int32(0xc0) == int32(0x80) && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr(str + 3)))&int32(0xc0) == int32(0x80) {
						if p_size != libc.UintptrFromInt32(0) {
							*(*MD_SIZE)(__ccgo_ptr(p_size)) = uint32(4)
						}
						return uint32(*(*MD_CHAR)(__ccgo_ptr(str)))&uint32(0x07)<<int32(18) | uint32(*(*MD_CHAR)(__ccgo_ptr(str + 1)))&uint32(0x3f)<<int32(12) | uint32(*(*MD_CHAR)(__ccgo_ptr(str + 2)))&uint32(0x3f)<<int32(6) | uint32(*(*MD_CHAR)(__ccgo_ptr(str + 3)))&uint32(0x3f)<<0
					}
				}
			}
		}
	}
	if p_size != libc.UintptrFromInt32(0) {
		*(*MD_SIZE)(__ccgo_ptr(p_size)) = uint32(1)
	}
	return uint32(*(*MD_CHAR)(__ccgo_ptr(str)))
}

func md_decode_utf8_before__(tls *libc.TLS, ctx uintptr, off MD_OFFSET) (r uint32) {
	if !(libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off-libc.Uint32FromInt32(1))))) <= libc.Int32FromInt32(0x7f)) {
		if off > uint32(1) && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off-libc.Uint32FromInt32(2)))))&int32(0xe0) == int32(0xc0) && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off-libc.Uint32FromInt32(1)))))&int32(0xc0) == int32(0x80) {
			return uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off-libc.Uint32FromInt32(2)))))&uint32(0x1f)<<int32(6) | uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off-libc.Uint32FromInt32(1)))))&uint32(0x3f)<<0
		}
		if off > uint32(2) && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off-libc.Uint32FromInt32(3)))))&int32(0xf0) == int32(0xe0) && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off-libc.Uint32FromInt32(2)))))&int32(0xc0) == int32(0x80) && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off-libc.Uint32FromInt32(1)))))&int32(0xc0) == int32(0x80) {
			return uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off-libc.Uint32FromInt32(3)))))&uint32(0x0f)<<int32(12) | uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off-libc.Uint32FromInt32(2)))))&uint32(0x3f)<<int32(6) | uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off-libc.Uint32FromInt32(1)))))&uint32(0x3f)<<0
		}
		if off > uint32(3) && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off-libc.Uint32FromInt32(4)))))&int32(0xf8) == int32(0xf0) && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off-libc.Uint32FromInt32(3)))))&int32(0xc0) == int32(0x80) && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off-libc.Uint32FromInt32(2)))))&int32(0xc0) == int32(0x80) && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off-libc.Uint32FromInt32(1)))))&int32(0xc0) == int32(0x80) {
			return uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off-libc.Uint32FromInt32(4)))))&uint32(0x07)<<int32(18) | uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off-libc.Uint32FromInt32(3)))))&uint32(0x3f)<<int32(12) | uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off-libc.Uint32FromInt32(2)))))&uint32(0x3f)<<int32(6) | uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off-libc.Uint32FromInt32(1)))))&uint32(0x3f)<<0
		}
	}
	return uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off-libc.Uint32FromInt32(1)))))
}

func md_decode_unicode(tls *libc.TLS, str uintptr, off MD_OFFSET, str_size MD_SIZE, p_char_size uintptr) (r uint32) {
	return md_decode_utf8__(tls, str+uintptr(off), str_size-off, p_char_size)
}

/*************************************
 ***  Helper string manipulations  ***
 *************************************/

// C documentation
//
//	/* Fill buffer with copy of the string between 'beg' and 'end' but replace any
//	 * line breaks with given replacement character.
//	 *
//	 * NOTE: Caller is responsible to make sure the buffer is large enough.
//	 * (Given the output is always shorter then input, (end - beg) is good idea
//	 * what the caller should allocate.)
//	 */
func md_merge_lines(tls *libc.TLS, ctx uintptr, beg MD_OFFSET, end MD_OFFSET, lines uintptr, n_lines MD_SIZE, line_break_replacement_char MD_CHAR, buffer uintptr, p_size uintptr) {
	var line, ptr uintptr
	var line_end, off MD_OFFSET
	var line_index int32
	_, _, _, _, _ = line, line_end, line_index, off, ptr
	ptr = buffer
	line_index = 0
	off = beg
	_ = n_lines
	for int32(1) != 0 {
		line = lines + uintptr(line_index)*8
		line_end = (*MD_LINE)(unsafe.Pointer(line)).Fend
		if end < line_end {
			line_end = end
		}
		for off < line_end {
			*(*MD_CHAR)(__ccgo_ptr(ptr)) = *(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))
			ptr = ptr + 1
			off = off + 1
		}
		if off >= end {
			*(*MD_SIZE)(__ccgo_ptr(p_size)) = libc.Uint32FromInt64(int64(ptr) - int64(buffer))
			return
		}
		*(*MD_CHAR)(__ccgo_ptr(ptr)) = line_break_replacement_char
		ptr = ptr + 1
		line_index = line_index + 1
		off = (*(*MD_LINE)(__ccgo_ptr(lines + uintptr(line_index)*8))).Fbeg
	}
}

// C documentation
//
//	/* Wrapper of md_merge_lines() which allocates new buffer for the output string.
//	 */
func md_merge_lines_alloc(tls *libc.TLS, ctx uintptr, beg MD_OFFSET, end MD_OFFSET, lines uintptr, n_lines MD_SIZE, line_break_replacement_char MD_CHAR, p_str uintptr, p_size uintptr) (r int32) {
	var buffer uintptr
	_ = buffer
	buffer = libc.Xmalloc(tls, uint64(1)*uint64(end-beg))
	if buffer == libc.UintptrFromInt32(0) {
		if (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log != libc.UintptrFromInt32(0) {
			(*(*func(*libc.TLS, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log})))(tls, __ccgo_ts+1, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
		}
		return -int32(1)
	}
	md_merge_lines(tls, ctx, beg, end, lines, n_lines, line_break_replacement_char, buffer, p_size)
	*(*uintptr)(__ccgo_ptr(p_str)) = buffer
	return 0
}

func md_skip_unicode_whitespace(tls *libc.TLS, label uintptr, off MD_OFFSET, size MD_SIZE) (r MD_OFFSET) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var codepoint uint32
	var _ /* char_size at bp+0 */ MD_SIZE
	_ = codepoint
	for off < size {
		codepoint = md_decode_unicode(tls, label, off, size, bp)
		if !(md_is_unicode_whitespace__(tls, codepoint) != 0) && !(libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr(label + uintptr(off)))) == int32('\r') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr(label + uintptr(off)))) == int32('\n')) {
			break
		}
		off = off + *(*MD_SIZE)(__ccgo_ptr(bp))
	}
	return off
}

/******************************
 ***  Recognizing raw HTML  ***
 ******************************/

// C documentation
//
//	/* md_is_html_tag() may be called when processing inlines (inline raw HTML)
//	 * or when breaking document to blocks (checking for start of HTML block type 7).
//	 *
//	 * When breaking document to blocks, we do not yet know line boundaries, but
//	 * in that case the whole tag has to live on a single line. We distinguish this
//	 * by n_lines == 0.
//	 */
func md_is_html_tag(tls *libc.TLS, ctx uintptr, lines uintptr, n_lines MD_SIZE, beg MD_OFFSET, max_end MD_OFFSET, p_end uintptr) (r int32) {
	var attr_state int32
	var line_end, off MD_OFFSET
	var line_index MD_SIZE
	var v1 uint32
	_, _, _, _, _ = attr_state, line_end, line_index, off, v1
	off = beg
	if n_lines > uint32(0) {
		v1 = (*(*MD_LINE)(__ccgo_ptr(lines))).Fend
	} else {
		v1 = (*MD_CTX)(unsafe.Pointer(ctx)).Fsize
	}
	line_end = v1
	line_index = uint32(0)
	if !(libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(beg)))) == libc.Int32FromUint8('<')) {
		libc.X__builtin_unreachable(tls)
	}
	if off+uint32(1) >= line_end {
		return FALSE
	}
	off = off + 1
	/* For parsing attributes, we need a little state automaton below.
	 * State -1: no attributes are allowed.
	 * State 0: attribute could follow after some whitespace.
	 * State 1: after a whitespace (attribute name may follow).
	 * State 2: after attribute name ('=' MAY follow).
	 * State 3: after '=' (value specification MUST follow).
	 * State 41: in middle of unquoted attribute value.
	 * State 42: in middle of single-quoted attribute value.
	 * State 43: in middle of double-quoted attribute value.
	 */
	attr_state = 0
	if libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('/') {
		/* Closer tag "</ ... >". No attributes may be present. */
		attr_state = -int32(1)
		off = off + 1
	}
	/* Tag name */
	if off >= line_end || !(libc.Uint32FromInt32(libc.Int32FromUint8('A')) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) <= libc.Uint32FromInt32(libc.Int32FromUint8('Z')) || libc.Uint32FromInt32(libc.Int32FromUint8('a')) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) <= libc.Uint32FromInt32(libc.Int32FromUint8('z'))) {
		return FALSE
	}
	off = off + 1
	for off < line_end && (libc.Uint32FromInt32(libc.Int32FromUint8('A')) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) <= libc.Uint32FromInt32(libc.Int32FromUint8('Z')) || libc.Uint32FromInt32(libc.Int32FromUint8('a')) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) <= libc.Uint32FromInt32(libc.Int32FromUint8('z')) || libc.Uint32FromInt32(libc.Int32FromUint8('0')) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) <= libc.Uint32FromInt32(libc.Int32FromUint8('9')) || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('-')) {
		off = off + 1
	}
	/* (Optional) attributes (if not closer), (optional) '/' (if not closer)
	 * and final '>'. */
	for int32(1) != 0 {
		for off < line_end && !(libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('\r') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('\n')) {
			if attr_state > int32(40) {
				if attr_state == int32(41) && (libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32(' ') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('\t') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) != int32('\000') && libc.Xstrchr(tls, __ccgo_ts+18, libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off))))) != libc.UintptrFromInt32(0)) {
					attr_state = 0
					off = off - 1 /* Put the char back for re-inspection in the new state. */
				} else {
					if attr_state == int32(42) && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('\'') {
						attr_state = 0
					} else {
						if attr_state == int32(43) && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('"') {
							attr_state = 0
						}
					}
				}
				off = off + 1
			} else {
				if libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32(' ') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('\t') || (libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('\v') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('\f')) {
					if attr_state == 0 {
						attr_state = int32(1)
					}
					off = off + 1
				} else {
					if attr_state <= int32(2) && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('>') {
						/* End. */
						goto done
					} else {
						if attr_state <= int32(2) && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('/') && off+uint32(1) < line_end && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))) == int32('>') {
							/* End with digraph '/>' */
							off = off + 1
							goto done
						} else {
							if (attr_state == int32(1) || attr_state == int32(2)) && (libc.Uint32FromInt32(libc.Int32FromUint8('A')) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) <= libc.Uint32FromInt32(libc.Int32FromUint8('Z')) || libc.Uint32FromInt32(libc.Int32FromUint8('a')) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) <= libc.Uint32FromInt32(libc.Int32FromUint8('z')) || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('_') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32(':')) {
								off = off + 1
								/* Attribute name */
								for off < line_end && (libc.Uint32FromInt32(libc.Int32FromUint8('A')) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) <= libc.Uint32FromInt32(libc.Int32FromUint8('Z')) || libc.Uint32FromInt32(libc.Int32FromUint8('a')) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) <= libc.Uint32FromInt32(libc.Int32FromUint8('z')) || libc.Uint32FromInt32(libc.Int32FromUint8('0')) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) <= libc.Uint32FromInt32(libc.Int32FromUint8('9')) || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) != int32('\000') && libc.Xstrchr(tls, __ccgo_ts+25, libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off))))) != libc.UintptrFromInt32(0)) {
									off = off + 1
								}
								attr_state = int32(2)
							} else {
								if attr_state == int32(2) && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('=') {
									/* Attribute assignment sign */
									off = off + 1
									attr_state = int32(3)
								} else {
									if attr_state == int32(3) {
										/* Expecting start of attribute value. */
										if libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('"') {
											attr_state = int32(43)
										} else {
											if libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('\'') {
												attr_state = int32(42)
											} else {
												if !(libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) != int32('\000') && libc.Xstrchr(tls, __ccgo_ts+18, libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off))))) != libc.UintptrFromInt32(0)) && !(libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('\r') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('\n')) {
													attr_state = int32(41)
												} else {
													return FALSE
												}
											}
										}
										off = off + 1
									} else {
										/* Anything unexpected. */
										return FALSE
									}
								}
							}
						}
					}
				}
			}
		}
		/* We have to be on a single line. See definition of start condition
		 * of HTML block, type 7. */
		if n_lines == uint32(0) {
			return FALSE
		}
		line_index = line_index + 1
		if line_index >= n_lines {
			return FALSE
		}
		off = (*(*MD_LINE)(__ccgo_ptr(lines + uintptr(line_index)*8))).Fbeg
		line_end = (*(*MD_LINE)(__ccgo_ptr(lines + uintptr(line_index)*8))).Fend
		if attr_state == 0 || attr_state == int32(41) {
			attr_state = int32(1)
		}
		if off >= max_end {
			return FALSE
		}
	}
	goto done
done:
	;
	if off >= max_end {
		return FALSE
	}
	*(*MD_OFFSET)(__ccgo_ptr(p_end)) = off + uint32(1)
	return int32(TRUE)
}

func md_scan_for_html_closer(tls *libc.TLS, ctx uintptr, str uintptr, len1 MD_SIZE, lines uintptr, n_lines MD_SIZE, beg MD_OFFSET, max_end MD_OFFSET, p_end uintptr, p_scan_horizon uintptr) (r int32) {
	var line_index MD_SIZE
	var off MD_OFFSET
	_, _ = line_index, off
	off = beg
	line_index = uint32(0)
	if off < *(*MD_OFFSET)(__ccgo_ptr(p_scan_horizon)) && *(*MD_OFFSET)(__ccgo_ptr(p_scan_horizon)) >= max_end-len1 {
		/* We have already scanned the range up to the max_end so we know
		 * there is nothing to see. */
		return FALSE
	}
	for int32(TRUE) != 0 {
		for off+len1 <= (*(*MD_LINE)(__ccgo_ptr(lines + uintptr(line_index)*8))).Fend && off+len1 <= max_end {
			if md_ascii_eq(tls, (*MD_CTX)(unsafe.Pointer(ctx)).Ftext+uintptr(off), str, len1) != 0 {
				/* Success. */
				*(*MD_OFFSET)(__ccgo_ptr(p_end)) = off + len1
				return int32(TRUE)
			}
			off = off + 1
		}
		line_index = line_index + 1
		if off >= max_end || line_index >= n_lines {
			/* Failure. */
			*(*MD_OFFSET)(__ccgo_ptr(p_scan_horizon)) = off
			return FALSE
		}
		off = (*(*MD_LINE)(__ccgo_ptr(lines + uintptr(line_index)*8))).Fbeg
	}
	return r
}

func md_is_html_comment(tls *libc.TLS, ctx uintptr, lines uintptr, n_lines MD_SIZE, beg MD_OFFSET, max_end MD_OFFSET, p_end uintptr) (r int32) {
	var off MD_OFFSET
	_ = off
	off = beg
	if !(libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(beg)))) == libc.Int32FromUint8('<')) {
		libc.X__builtin_unreachable(tls)
	}
	if off+uint32(4) >= (*(*MD_LINE)(__ccgo_ptr(lines))).Fend {
		return FALSE
	}
	if libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))) != int32('!') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(2))))) != int32('-') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(3))))) != int32('-') {
		return FALSE
	}
	/* Skip only "<!" so that we accept also "<!-->" or "<!--->" */
	off = off + uint32(2)
	/* Scan for ordinary comment closer "-->". */
	return md_scan_for_html_closer(tls, ctx, __ccgo_ts+30, uint32(3), lines, n_lines, off, max_end, p_end, ctx+504)
}

func md_is_html_processing_instruction(tls *libc.TLS, ctx uintptr, lines uintptr, n_lines MD_SIZE, beg MD_OFFSET, max_end MD_OFFSET, p_end uintptr) (r int32) {
	var off MD_OFFSET
	_ = off
	off = beg
	if off+uint32(2) >= (*(*MD_LINE)(__ccgo_ptr(lines))).Fend {
		return FALSE
	}
	if libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))) != int32('?') {
		return FALSE
	}
	off = off + uint32(2)
	return md_scan_for_html_closer(tls, ctx, __ccgo_ts+34, uint32(2), lines, n_lines, off, max_end, p_end, ctx+508)
}

func md_is_html_declaration(tls *libc.TLS, ctx uintptr, lines uintptr, n_lines MD_SIZE, beg MD_OFFSET, max_end MD_OFFSET, p_end uintptr) (r int32) {
	var off MD_OFFSET
	_ = off
	off = beg
	if off+uint32(2) >= (*(*MD_LINE)(__ccgo_ptr(lines))).Fend {
		return FALSE
	}
	if libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))) != int32('!') {
		return FALSE
	}
	off = off + uint32(2)
	/* Declaration name. */
	if off >= (*(*MD_LINE)(__ccgo_ptr(lines))).Fend || !(libc.Uint32FromInt32(libc.Int32FromUint8('A')) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) <= libc.Uint32FromInt32(libc.Int32FromUint8('Z')) || libc.Uint32FromInt32(libc.Int32FromUint8('a')) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) <= libc.Uint32FromInt32(libc.Int32FromUint8('z'))) {
		return FALSE
	}
	off = off + 1
	for off < (*(*MD_LINE)(__ccgo_ptr(lines))).Fend && (libc.Uint32FromInt32(libc.Int32FromUint8('A')) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) <= libc.Uint32FromInt32(libc.Int32FromUint8('Z')) || libc.Uint32FromInt32(libc.Int32FromUint8('a')) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) <= libc.Uint32FromInt32(libc.Int32FromUint8('z'))) {
		off = off + 1
	}
	return md_scan_for_html_closer(tls, ctx, __ccgo_ts+37, uint32(1), lines, n_lines, off, max_end, p_end, ctx+512)
}

func md_is_html_cdata(tls *libc.TLS, ctx uintptr, lines uintptr, n_lines MD_SIZE, beg MD_OFFSET, max_end MD_OFFSET, p_end uintptr) (r int32) {
	var off MD_OFFSET
	_ = off
	off = beg
	if off+open_size >= (*(*MD_LINE)(__ccgo_ptr(lines))).Fend {
		return FALSE
	}
	if libc.Xmemcmp(tls, (*MD_CTX)(unsafe.Pointer(ctx)).Ftext+uintptr(off), uintptr(unsafe.Pointer(&open_str)), uint64(open_size)) != 0 {
		return FALSE
	}
	off = off + open_size
	return md_scan_for_html_closer(tls, ctx, __ccgo_ts+39, uint32(3), lines, n_lines, off, max_end, p_end, ctx+516)
}

var open_str = [10]MD_CHAR{'<', '!', '[', 'C', 'D', 'A', 'T', 'A', '['}

var open_size = uint32(libc.Uint64FromInt64(10)/libc.Uint64FromInt64(1) - libc.Uint64FromInt32(1))

func md_is_html_any(tls *libc.TLS, ctx uintptr, lines uintptr, n_lines MD_SIZE, beg MD_OFFSET, max_end MD_OFFSET, p_end uintptr) (r int32) {
	if !(libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(beg)))) == libc.Int32FromUint8('<')) {
		libc.X__builtin_unreachable(tls)
	}
	return libc.BoolInt32(md_is_html_tag(tls, ctx, lines, n_lines, beg, max_end, p_end) != 0 || md_is_html_comment(tls, ctx, lines, n_lines, beg, max_end, p_end) != 0 || md_is_html_processing_instruction(tls, ctx, lines, n_lines, beg, max_end, p_end) != 0 || md_is_html_declaration(tls, ctx, lines, n_lines, beg, max_end, p_end) != 0 || md_is_html_cdata(tls, ctx, lines, n_lines, beg, max_end, p_end) != 0)
}

/****************************
 ***  Recognizing Entity  ***
 ****************************/

func md_is_hex_entity_contents(tls *libc.TLS, ctx uintptr, text uintptr, beg MD_OFFSET, max_end MD_OFFSET, p_end uintptr) (r int32) {
	var off MD_OFFSET
	_ = off
	off = beg
	_ = ctx
	for off < max_end && (libc.Uint32FromInt32(libc.Int32FromUint8('0')) <= uint32(*(*MD_CHAR)(__ccgo_ptr(text + uintptr(off)))) && uint32(*(*MD_CHAR)(__ccgo_ptr(text + uintptr(off)))) <= libc.Uint32FromInt32(libc.Int32FromUint8('9')) || libc.Uint32FromInt32(libc.Int32FromUint8('A')) <= uint32(*(*MD_CHAR)(__ccgo_ptr(text + uintptr(off)))) && uint32(*(*MD_CHAR)(__ccgo_ptr(text + uintptr(off)))) <= libc.Uint32FromInt32(libc.Int32FromUint8('F')) || libc.Uint32FromInt32(libc.Int32FromUint8('a')) <= uint32(*(*MD_CHAR)(__ccgo_ptr(text + uintptr(off)))) && uint32(*(*MD_CHAR)(__ccgo_ptr(text + uintptr(off)))) <= libc.Uint32FromInt32(libc.Int32FromUint8('f'))) && off-beg <= uint32(8) {
		off = off + 1
	}
	if uint32(1) <= off-beg && off-beg <= uint32(6) {
		*(*MD_OFFSET)(__ccgo_ptr(p_end)) = off
		return int32(TRUE)
	} else {
		return FALSE
	}
	return r
}

func md_is_dec_entity_contents(tls *libc.TLS, ctx uintptr, text uintptr, beg MD_OFFSET, max_end MD_OFFSET, p_end uintptr) (r int32) {
	var off MD_OFFSET
	_ = off
	off = beg
	_ = ctx
	for off < max_end && (libc.Uint32FromInt32(libc.Int32FromUint8('0')) <= uint32(*(*MD_CHAR)(__ccgo_ptr(text + uintptr(off)))) && uint32(*(*MD_CHAR)(__ccgo_ptr(text + uintptr(off)))) <= libc.Uint32FromInt32(libc.Int32FromUint8('9'))) && off-beg <= uint32(8) {
		off = off + 1
	}
	if uint32(1) <= off-beg && off-beg <= uint32(7) {
		*(*MD_OFFSET)(__ccgo_ptr(p_end)) = off
		return int32(TRUE)
	} else {
		return FALSE
	}
	return r
}

func md_is_named_entity_contents(tls *libc.TLS, ctx uintptr, text uintptr, beg MD_OFFSET, max_end MD_OFFSET, p_end uintptr) (r int32) {
	var off MD_OFFSET
	_ = off
	off = beg
	_ = ctx
	if off < max_end && (libc.Uint32FromInt32(libc.Int32FromUint8('A')) <= uint32(*(*MD_CHAR)(__ccgo_ptr(text + uintptr(off)))) && uint32(*(*MD_CHAR)(__ccgo_ptr(text + uintptr(off)))) <= libc.Uint32FromInt32(libc.Int32FromUint8('Z')) || libc.Uint32FromInt32(libc.Int32FromUint8('a')) <= uint32(*(*MD_CHAR)(__ccgo_ptr(text + uintptr(off)))) && uint32(*(*MD_CHAR)(__ccgo_ptr(text + uintptr(off)))) <= libc.Uint32FromInt32(libc.Int32FromUint8('z'))) {
		off = off + 1
	} else {
		return FALSE
	}
	for off < max_end && (libc.Uint32FromInt32(libc.Int32FromUint8('A')) <= uint32(*(*MD_CHAR)(__ccgo_ptr(text + uintptr(off)))) && uint32(*(*MD_CHAR)(__ccgo_ptr(text + uintptr(off)))) <= libc.Uint32FromInt32(libc.Int32FromUint8('Z')) || libc.Uint32FromInt32(libc.Int32FromUint8('a')) <= uint32(*(*MD_CHAR)(__ccgo_ptr(text + uintptr(off)))) && uint32(*(*MD_CHAR)(__ccgo_ptr(text + uintptr(off)))) <= libc.Uint32FromInt32(libc.Int32FromUint8('z')) || libc.Uint32FromInt32(libc.Int32FromUint8('0')) <= uint32(*(*MD_CHAR)(__ccgo_ptr(text + uintptr(off)))) && uint32(*(*MD_CHAR)(__ccgo_ptr(text + uintptr(off)))) <= libc.Uint32FromInt32(libc.Int32FromUint8('9'))) && off-beg <= uint32(48) {
		off = off + 1
	}
	if uint32(2) <= off-beg && off-beg <= uint32(48) {
		*(*MD_OFFSET)(__ccgo_ptr(p_end)) = off
		return int32(TRUE)
	} else {
		return FALSE
	}
	return r
}

func md_is_entity_str(tls *libc.TLS, ctx uintptr, text uintptr, beg MD_OFFSET, max_end MD_OFFSET, p_end uintptr) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var is_contents int32
	var _ /* off at bp+0 */ MD_OFFSET
	_ = is_contents
	*(*MD_OFFSET)(__ccgo_ptr(bp)) = beg
	if !(libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr(text + uintptr(*(*MD_OFFSET)(__ccgo_ptr(bp)))))) == libc.Int32FromUint8('&')) {
		libc.X__builtin_unreachable(tls)
	}
	*(*MD_OFFSET)(__ccgo_ptr(bp)) = *(*MD_OFFSET)(__ccgo_ptr(bp)) + 1
	if *(*MD_OFFSET)(__ccgo_ptr(bp))+uint32(2) < max_end && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr(text + uintptr(*(*MD_OFFSET)(__ccgo_ptr(bp)))))) == int32('#') && (libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr(text + uintptr(*(*MD_OFFSET)(__ccgo_ptr(bp))+uint32(1))))) == int32('x') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr(text + uintptr(*(*MD_OFFSET)(__ccgo_ptr(bp))+uint32(1))))) == int32('X')) {
		is_contents = md_is_hex_entity_contents(tls, ctx, text, *(*MD_OFFSET)(__ccgo_ptr(bp))+uint32(2), max_end, bp)
	} else {
		if *(*MD_OFFSET)(__ccgo_ptr(bp))+uint32(1) < max_end && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr(text + uintptr(*(*MD_OFFSET)(__ccgo_ptr(bp)))))) == int32('#') {
			is_contents = md_is_dec_entity_contents(tls, ctx, text, *(*MD_OFFSET)(__ccgo_ptr(bp))+uint32(1), max_end, bp)
		} else {
			is_contents = md_is_named_entity_contents(tls, ctx, text, *(*MD_OFFSET)(__ccgo_ptr(bp)), max_end, bp)
		}
	}
	if is_contents != 0 && *(*MD_OFFSET)(__ccgo_ptr(bp)) < max_end && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr(text + uintptr(*(*MD_OFFSET)(__ccgo_ptr(bp)))))) == int32(';') {
		*(*MD_OFFSET)(__ccgo_ptr(p_end)) = *(*MD_OFFSET)(__ccgo_ptr(bp)) + uint32(1)
		return int32(TRUE)
	} else {
		return FALSE
	}
	return r
}

func md_is_entity(tls *libc.TLS, ctx uintptr, beg MD_OFFSET, max_end MD_OFFSET, p_end uintptr) (r int32) {
	return md_is_entity_str(tls, ctx, (*MD_CTX)(unsafe.Pointer(ctx)).Ftext, beg, max_end, p_end)
}

/******************************
 ***  Attribute Management  ***
 ******************************/

type MD_ATTRIBUTE_BUILD = struct {
	Ftext            uintptr
	Fsubstr_types    uintptr
	Fsubstr_offsets  uintptr
	Fsubstr_count    int32
	Fsubstr_alloc    int32
	Ftrivial_types   [1]MD_TEXTTYPE1
	Ftrivial_offsets [2]MD_OFFSET
}

/******************************
 ***  Attribute Management  ***
 ******************************/

type MD_ATTRIBUTE_BUILD_tag = MD_ATTRIBUTE_BUILD

func md_build_attr_append_substr(tls *libc.TLS, ctx uintptr, build uintptr, type1 MD_TEXTTYPE1, off MD_OFFSET) (r int32) {
	var new_substr_offsets, new_substr_types uintptr
	var v1 int32
	_, _, _ = new_substr_offsets, new_substr_types, v1
	if (*MD_ATTRIBUTE_BUILD)(unsafe.Pointer(build)).Fsubstr_count >= (*MD_ATTRIBUTE_BUILD)(unsafe.Pointer(build)).Fsubstr_alloc {
		if (*MD_ATTRIBUTE_BUILD)(unsafe.Pointer(build)).Fsubstr_alloc > 0 {
			v1 = (*MD_ATTRIBUTE_BUILD)(unsafe.Pointer(build)).Fsubstr_alloc + (*MD_ATTRIBUTE_BUILD)(unsafe.Pointer(build)).Fsubstr_alloc/int32(2)
		} else {
			v1 = int32(8)
		}
		(*MD_ATTRIBUTE_BUILD)(unsafe.Pointer(build)).Fsubstr_alloc = v1
		new_substr_types = libc.Xrealloc(tls, (*MD_ATTRIBUTE_BUILD)(unsafe.Pointer(build)).Fsubstr_types, libc.Uint64FromInt32((*MD_ATTRIBUTE_BUILD)(unsafe.Pointer(build)).Fsubstr_alloc)*uint64(4))
		if new_substr_types == libc.UintptrFromInt32(0) {
			if (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log != libc.UintptrFromInt32(0) {
				(*(*func(*libc.TLS, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log})))(tls, __ccgo_ts+43, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
			}
			return -int32(1)
		}
		/* Note +1 to reserve space for final offset (== raw_size). */
		new_substr_offsets = libc.Xrealloc(tls, (*MD_ATTRIBUTE_BUILD)(unsafe.Pointer(build)).Fsubstr_offsets, libc.Uint64FromInt32((*MD_ATTRIBUTE_BUILD)(unsafe.Pointer(build)).Fsubstr_alloc+libc.Int32FromInt32(1))*uint64(4))
		if new_substr_offsets == libc.UintptrFromInt32(0) {
			if (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log != libc.UintptrFromInt32(0) {
				(*(*func(*libc.TLS, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log})))(tls, __ccgo_ts+43, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
			}
			libc.Xfree(tls, new_substr_types)
			return -int32(1)
		}
		(*MD_ATTRIBUTE_BUILD)(unsafe.Pointer(build)).Fsubstr_types = new_substr_types
		(*MD_ATTRIBUTE_BUILD)(unsafe.Pointer(build)).Fsubstr_offsets = new_substr_offsets
	}
	*(*MD_TEXTTYPE1)(__ccgo_ptr((*MD_ATTRIBUTE_BUILD)(unsafe.Pointer(build)).Fsubstr_types + uintptr((*MD_ATTRIBUTE_BUILD)(unsafe.Pointer(build)).Fsubstr_count)*4)) = type1
	*(*MD_OFFSET)(__ccgo_ptr((*MD_ATTRIBUTE_BUILD)(unsafe.Pointer(build)).Fsubstr_offsets + uintptr((*MD_ATTRIBUTE_BUILD)(unsafe.Pointer(build)).Fsubstr_count)*4)) = off
	(*MD_ATTRIBUTE_BUILD)(unsafe.Pointer(build)).Fsubstr_count = (*MD_ATTRIBUTE_BUILD)(unsafe.Pointer(build)).Fsubstr_count + 1
	return 0
}

func md_free_attribute(tls *libc.TLS, ctx uintptr, build uintptr) {
	_ = ctx
	if (*MD_ATTRIBUTE_BUILD)(unsafe.Pointer(build)).Fsubstr_alloc > 0 {
		libc.Xfree(tls, (*MD_ATTRIBUTE_BUILD)(unsafe.Pointer(build)).Ftext)
		libc.Xfree(tls, (*MD_ATTRIBUTE_BUILD)(unsafe.Pointer(build)).Fsubstr_types)
		libc.Xfree(tls, (*MD_ATTRIBUTE_BUILD)(unsafe.Pointer(build)).Fsubstr_offsets)
	}
}

func md_build_attribute(tls *libc.TLS, ctx uintptr, raw_text uintptr, raw_size MD_SIZE, flags uint32, attr uintptr, build uintptr) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var is_trivial, ret int32
	var off, raw_off, v3, v4 MD_OFFSET
	var v2 uintptr
	var _ /* ent_end at bp+0 */ MD_OFFSET
	_, _, _, _, _, _, _ = is_trivial, off, raw_off, ret, v2, v3, v4
	ret = 0
	libc.Xmemset(tls, build, 0, uint64(48))
	/* If there is no backslash and no ampersand, build trivial attribute
	 * without any malloc(). */
	is_trivial = int32(TRUE)
	raw_off = uint32(0)
	for {
		if !(raw_off < raw_size) {
			break
		}
		if libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr(raw_text + uintptr(raw_off)))) == int32('\\') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr(raw_text + uintptr(raw_off)))) == int32('&') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr(raw_text + uintptr(raw_off)))) == int32('\000') {
			is_trivial = FALSE
			break
		}
		goto _1
	_1:
		;
		raw_off = raw_off + 1
	}
	if is_trivial != 0 {
		if raw_size != 0 {
			v2 = raw_text
		} else {
			v2 = libc.UintptrFromInt32(0)
		}
		(*MD_ATTRIBUTE_BUILD)(unsafe.Pointer(build)).Ftext = v2
		(*MD_ATTRIBUTE_BUILD)(unsafe.Pointer(build)).Fsubstr_types = build + 32
		(*MD_ATTRIBUTE_BUILD)(unsafe.Pointer(build)).Fsubstr_offsets = build + 36
		(*MD_ATTRIBUTE_BUILD)(unsafe.Pointer(build)).Fsubstr_count = int32(1)
		(*MD_ATTRIBUTE_BUILD)(unsafe.Pointer(build)).Fsubstr_alloc = 0
		*(*MD_TEXTTYPE1)(__ccgo_ptr(build + 32)) = int32(MD_TEXT_NORMAL)
		*(*MD_OFFSET)(__ccgo_ptr(build + 36)) = uint32(0)
		*(*MD_OFFSET)(__ccgo_ptr(build + 36 + 1*4)) = raw_size
		off = raw_size
	} else {
		(*MD_ATTRIBUTE_BUILD)(unsafe.Pointer(build)).Ftext = libc.Xmalloc(tls, uint64(raw_size)*uint64(1))
		if (*MD_ATTRIBUTE_BUILD)(unsafe.Pointer(build)).Ftext == libc.UintptrFromInt32(0) {
			if (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log != libc.UintptrFromInt32(0) {
				(*(*func(*libc.TLS, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log})))(tls, __ccgo_ts+1, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
			}
			goto abort
		}
		raw_off = uint32(0)
		off = uint32(0)
		for raw_off < raw_size {
			if libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr(raw_text + uintptr(raw_off)))) == int32('\000') {
				ret = md_build_attr_append_substr(tls, ctx, build, int32(MD_TEXT_NULLCHAR), off)
				if ret < 0 {
					goto abort
				}
				libc.Xmemcpy(tls, (*MD_ATTRIBUTE_BUILD)(unsafe.Pointer(build)).Ftext+uintptr(off), raw_text+uintptr(raw_off), uint64(1))
				off = off + 1
				raw_off = raw_off + 1
				continue
			}
			if libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr(raw_text + uintptr(raw_off)))) == int32('&') {
				if md_is_entity_str(tls, ctx, raw_text, raw_off, raw_size, bp) != 0 {
					ret = md_build_attr_append_substr(tls, ctx, build, int32(MD_TEXT_ENTITY), off)
					if ret < 0 {
						goto abort
					}
					libc.Xmemcpy(tls, (*MD_ATTRIBUTE_BUILD)(unsafe.Pointer(build)).Ftext+uintptr(off), raw_text+uintptr(raw_off), uint64(*(*MD_OFFSET)(__ccgo_ptr(bp))-raw_off))
					off = off + (*(*MD_OFFSET)(__ccgo_ptr(bp)) - raw_off)
					raw_off = *(*MD_OFFSET)(__ccgo_ptr(bp))
					continue
				}
			}
			if (*MD_ATTRIBUTE_BUILD)(unsafe.Pointer(build)).Fsubstr_count == 0 || *(*MD_TEXTTYPE1)(__ccgo_ptr((*MD_ATTRIBUTE_BUILD)(unsafe.Pointer(build)).Fsubstr_types + uintptr((*MD_ATTRIBUTE_BUILD)(unsafe.Pointer(build)).Fsubstr_count-int32(1))*4)) != int32(MD_TEXT_NORMAL) {
				ret = md_build_attr_append_substr(tls, ctx, build, int32(MD_TEXT_NORMAL), off)
				if ret < 0 {
					goto abort
				}
			}
			if !(flags&libc.Uint32FromInt32(MD_BUILD_ATTR_NO_ESCAPES) != 0) && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr(raw_text + uintptr(raw_off)))) == int32('\\') && raw_off+uint32(1) < raw_size && (libc.Uint32FromInt32(libc.Int32FromInt32(33)) <= uint32(*(*MD_CHAR)(__ccgo_ptr(raw_text + uintptr(raw_off+uint32(1))))) && uint32(*(*MD_CHAR)(__ccgo_ptr(raw_text + uintptr(raw_off+uint32(1))))) <= libc.Uint32FromInt32(libc.Int32FromInt32(47)) || libc.Uint32FromInt32(libc.Int32FromInt32(58)) <= uint32(*(*MD_CHAR)(__ccgo_ptr(raw_text + uintptr(raw_off+uint32(1))))) && uint32(*(*MD_CHAR)(__ccgo_ptr(raw_text + uintptr(raw_off+uint32(1))))) <= libc.Uint32FromInt32(libc.Int32FromInt32(64)) || libc.Uint32FromInt32(libc.Int32FromInt32(91)) <= uint32(*(*MD_CHAR)(__ccgo_ptr(raw_text + uintptr(raw_off+uint32(1))))) && uint32(*(*MD_CHAR)(__ccgo_ptr(raw_text + uintptr(raw_off+uint32(1))))) <= libc.Uint32FromInt32(libc.Int32FromInt32(96)) || libc.Uint32FromInt32(libc.Int32FromInt32(123)) <= uint32(*(*MD_CHAR)(__ccgo_ptr(raw_text + uintptr(raw_off+uint32(1))))) && uint32(*(*MD_CHAR)(__ccgo_ptr(raw_text + uintptr(raw_off+uint32(1))))) <= libc.Uint32FromInt32(libc.Int32FromInt32(126)) || (libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr(raw_text + uintptr(raw_off+uint32(1))))) == int32('\r') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr(raw_text + uintptr(raw_off+uint32(1))))) == int32('\n'))) {
				raw_off = raw_off + 1
			}
			v3 = off
			off = off + 1
			v4 = raw_off
			raw_off = raw_off + 1
			*(*MD_CHAR)(__ccgo_ptr((*MD_ATTRIBUTE_BUILD)(unsafe.Pointer(build)).Ftext + uintptr(v3))) = *(*MD_CHAR)(__ccgo_ptr(raw_text + uintptr(v4)))
		}
		*(*MD_OFFSET)(__ccgo_ptr((*MD_ATTRIBUTE_BUILD)(unsafe.Pointer(build)).Fsubstr_offsets + uintptr((*MD_ATTRIBUTE_BUILD)(unsafe.Pointer(build)).Fsubstr_count)*4)) = off
	}
	(*MD_ATTRIBUTE)(unsafe.Pointer(attr)).Ftext = (*MD_ATTRIBUTE_BUILD)(unsafe.Pointer(build)).Ftext
	(*MD_ATTRIBUTE)(unsafe.Pointer(attr)).Fsize = off
	(*MD_ATTRIBUTE)(unsafe.Pointer(attr)).Fsubstr_offsets = (*MD_ATTRIBUTE_BUILD)(unsafe.Pointer(build)).Fsubstr_offsets
	(*MD_ATTRIBUTE)(unsafe.Pointer(attr)).Fsubstr_types = (*MD_ATTRIBUTE_BUILD)(unsafe.Pointer(build)).Fsubstr_types
	return 0
	goto abort
abort:
	;
	md_free_attribute(tls, ctx, build)
	return -int32(1)
}

/*********************************************
 ***  Dictionary of Reference Definitions  ***
 *********************************************/

func md_fnv1a(tls *libc.TLS, base uint32, data uintptr, n size_t) (r uint32) {
	var buf uintptr
	var hash uint32
	var i size_t
	_, _, _ = buf, hash, i
	buf = data
	hash = base
	i = uint64(0)
	for {
		if !(i < n) {
			break
		}
		hash = hash ^ uint32(*(*uint8)(__ccgo_ptr(buf + uintptr(i))))
		hash = hash * uint32(16777619)
		goto _1
	_1:
		;
		i = i + 1
	}
	return hash
}

/* Label equivalence is quite complicated with regards to whitespace and case
 * folding. This complicates computing a hash of it as well as direct comparison
 * of two labels. */

func md_link_label_hash(tls *libc.TLS, label uintptr, size MD_SIZE) (r uint32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var hash uint32
	var is_whitespace int32
	var off MD_OFFSET
	var _ /* char_size at bp+4 */ MD_SIZE
	var _ /* codepoint at bp+0 */ uint32
	var _ /* fold_info at bp+8 */ MD_UNICODE_FOLD_INFO
	_, _, _ = hash, is_whitespace, off
	hash = uint32(2166136261)
	is_whitespace = FALSE
	off = md_skip_unicode_whitespace(tls, label, uint32(0), size)
	for off < size {
		*(*uint32)(__ccgo_ptr(bp)) = md_decode_unicode(tls, label, off, size, bp+4)
		is_whitespace = libc.BoolInt32(md_is_unicode_whitespace__(tls, *(*uint32)(__ccgo_ptr(bp))) != 0 || (libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr(label + uintptr(off)))) == int32('\r') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr(label + uintptr(off)))) == int32('\n')))
		if is_whitespace != 0 {
			*(*uint32)(__ccgo_ptr(bp)) = uint32(' ')
			hash = md_fnv1a(tls, hash, bp, uint64(4))
			off = md_skip_unicode_whitespace(tls, label, off, size)
		} else {
			md_get_unicode_fold_info(tls, *(*uint32)(__ccgo_ptr(bp)), bp+8)
			hash = md_fnv1a(tls, hash, bp+8, uint64((*(*MD_UNICODE_FOLD_INFO)(__ccgo_ptr(bp + 8))).Fn_codepoints)*uint64(4))
			off = off + *(*MD_SIZE)(__ccgo_ptr(bp + 4))
		}
	}
	return hash
}

func md_link_label_cmp_load_fold_info(tls *libc.TLS, label uintptr, off MD_OFFSET, size MD_SIZE, fold_info uintptr) (r MD_OFFSET) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var codepoint uint32
	var _ /* char_size at bp+0 */ MD_SIZE
	_ = codepoint
	if off >= size {
		/* Treat end of a link label as a whitespace. */
		goto whitespace
	}
	codepoint = md_decode_unicode(tls, label, off, size, bp)
	off = off + *(*MD_SIZE)(__ccgo_ptr(bp))
	if md_is_unicode_whitespace__(tls, codepoint) != 0 {
		/* Treat all whitespace as equivalent */
		goto whitespace
	}
	/* Get real folding info. */
	md_get_unicode_fold_info(tls, codepoint, fold_info)
	return off
	goto whitespace
whitespace:
	;
	*(*uint32)(__ccgo_ptr(fold_info)) = uint32(' ')
	(*MD_UNICODE_FOLD_INFO)(unsafe.Pointer(fold_info)).Fn_codepoints = uint32(1)
	return md_skip_unicode_whitespace(tls, label, off, size)
}

func md_link_label_cmp(tls *libc.TLS, a_label uintptr, a_size MD_SIZE, b_label uintptr, b_size MD_SIZE) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var a_fi_off, a_off, b_fi_off, b_off MD_OFFSET
	var cmp int32
	var _ /* a_fi at bp+0 */ MD_UNICODE_FOLD_INFO
	var _ /* b_fi at bp+16 */ MD_UNICODE_FOLD_INFO
	_, _, _, _, _ = a_fi_off, a_off, b_fi_off, b_off, cmp
	*(*MD_UNICODE_FOLD_INFO)(__ccgo_ptr(bp)) = MD_UNICODE_FOLD_INFO{}
	*(*MD_UNICODE_FOLD_INFO)(__ccgo_ptr(bp + 16)) = MD_UNICODE_FOLD_INFO{}
	a_fi_off = uint32(0)
	b_fi_off = uint32(0)
	a_off = md_skip_unicode_whitespace(tls, a_label, uint32(0), a_size)
	b_off = md_skip_unicode_whitespace(tls, b_label, uint32(0), b_size)
	for a_off < a_size || a_fi_off < (*(*MD_UNICODE_FOLD_INFO)(__ccgo_ptr(bp))).Fn_codepoints || b_off < b_size || b_fi_off < (*(*MD_UNICODE_FOLD_INFO)(__ccgo_ptr(bp + 16))).Fn_codepoints {
		/* If needed, load fold info for next char. */
		if a_fi_off >= (*(*MD_UNICODE_FOLD_INFO)(__ccgo_ptr(bp))).Fn_codepoints {
			a_fi_off = uint32(0)
			a_off = md_link_label_cmp_load_fold_info(tls, a_label, a_off, a_size, bp)
		}
		if b_fi_off >= (*(*MD_UNICODE_FOLD_INFO)(__ccgo_ptr(bp + 16))).Fn_codepoints {
			b_fi_off = uint32(0)
			b_off = md_link_label_cmp_load_fold_info(tls, b_label, b_off, b_size, bp+16)
		}
		cmp = libc.Int32FromUint32(*(*uint32)(__ccgo_ptr(bp + 16 + uintptr(b_fi_off)*4)) - *(*uint32)(__ccgo_ptr(bp + uintptr(a_fi_off)*4)))
		if cmp != 0 {
			return cmp
		}
		a_fi_off = a_fi_off + 1
		b_fi_off = b_fi_off + 1
	}
	return 0
}

type MD_REF_DEF_LIST = struct {
	F__ccgo_align   [0]uint64
	Fn_ref_defs     int32
	Falloc_ref_defs int32
}

type MD_REF_DEF_LIST_tag = MD_REF_DEF_LIST

func md_ref_def_cmp(tls *libc.TLS, a uintptr, b uintptr) (r int32) {
	var a_ref, b_ref uintptr
	_, _ = a_ref, b_ref
	a_ref = *(*uintptr)(__ccgo_ptr(a))
	b_ref = *(*uintptr)(__ccgo_ptr(b))
	if (*MD_REF_DEF)(unsafe.Pointer(a_ref)).Fhash < (*MD_REF_DEF)(unsafe.Pointer(b_ref)).Fhash {
		return -int32(1)
	} else {
		if (*MD_REF_DEF)(unsafe.Pointer(a_ref)).Fhash > (*MD_REF_DEF)(unsafe.Pointer(b_ref)).Fhash {
			return +libc.Int32FromInt32(1)
		} else {
			return md_link_label_cmp(tls, (*MD_REF_DEF)(unsafe.Pointer(a_ref)).Flabel, (*MD_REF_DEF)(unsafe.Pointer(a_ref)).Flabel_size, (*MD_REF_DEF)(unsafe.Pointer(b_ref)).Flabel, (*MD_REF_DEF)(unsafe.Pointer(b_ref)).Flabel_size)
		}
	}
	return r
}

func md_ref_def_cmp_for_sort(tls *libc.TLS, a uintptr, b uintptr) (r int32) {
	var a_ref, b_ref uintptr
	var cmp int32
	_, _, _ = a_ref, b_ref, cmp
	cmp = md_ref_def_cmp(tls, a, b)
	/* Ensure stability of the sorting. */
	if cmp == 0 {
		a_ref = *(*uintptr)(__ccgo_ptr(a))
		b_ref = *(*uintptr)(__ccgo_ptr(b))
		if a_ref < b_ref {
			cmp = -int32(1)
		} else {
			if a_ref > b_ref {
				cmp = +libc.Int32FromInt32(1)
			} else {
				cmp = 0
			}
		}
	}
	return cmp
}

func md_build_ref_def_hashtable(tls *libc.TLS, ctx uintptr) (r int32) {
	var alloc_ref_defs, i, j int32
	var bucket, bucket1, def, list, list1, list_tmp, old_def uintptr
	_, _, _, _, _, _, _, _, _, _ = alloc_ref_defs, bucket, bucket1, def, i, j, list, list1, list_tmp, old_def
	if (*MD_CTX)(unsafe.Pointer(ctx)).Fn_ref_defs == 0 {
		return 0
	}
	(*MD_CTX)(unsafe.Pointer(ctx)).Fref_def_hashtable_size = (*MD_CTX)(unsafe.Pointer(ctx)).Fn_ref_defs * int32(5) / int32(4)
	(*MD_CTX)(unsafe.Pointer(ctx)).Fref_def_hashtable = libc.Xmalloc(tls, libc.Uint64FromInt32((*MD_CTX)(unsafe.Pointer(ctx)).Fref_def_hashtable_size)*uint64(8))
	if (*MD_CTX)(unsafe.Pointer(ctx)).Fref_def_hashtable == libc.UintptrFromInt32(0) {
		if (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log != libc.UintptrFromInt32(0) {
			(*(*func(*libc.TLS, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log})))(tls, __ccgo_ts+1, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
		}
		goto abort
	}
	libc.Xmemset(tls, (*MD_CTX)(unsafe.Pointer(ctx)).Fref_def_hashtable, 0, libc.Uint64FromInt32((*MD_CTX)(unsafe.Pointer(ctx)).Fref_def_hashtable_size)*uint64(8))
	/* Each member of ctx->ref_def_hashtable[] can be:
	 *  -- NULL,
	 *  -- pointer to the MD_REF_DEF in ctx->ref_defs[], or
	 *  -- pointer to a MD_REF_DEF_LIST, which holds multiple pointers to
	 *     such MD_REF_DEFs.
	 */
	i = 0
	for {
		if !(i < (*MD_CTX)(unsafe.Pointer(ctx)).Fn_ref_defs) {
			break
		}
		def = (*MD_CTX)(unsafe.Pointer(ctx)).Fref_defs + uintptr(i)*40
		(*MD_REF_DEF)(unsafe.Pointer(def)).Fhash = md_link_label_hash(tls, (*MD_REF_DEF)(unsafe.Pointer(def)).Flabel, (*MD_REF_DEF)(unsafe.Pointer(def)).Flabel_size)
		bucket = *(*uintptr)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Fref_def_hashtable + uintptr((*MD_REF_DEF)(unsafe.Pointer(def)).Fhash%libc.Uint32FromInt32((*MD_CTX)(unsafe.Pointer(ctx)).Fref_def_hashtable_size))*8))
		if bucket == libc.UintptrFromInt32(0) {
			/* The bucket is empty. Make it just point to the def. */
			*(*uintptr)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Fref_def_hashtable + uintptr((*MD_REF_DEF)(unsafe.Pointer(def)).Fhash%libc.Uint32FromInt32((*MD_CTX)(unsafe.Pointer(ctx)).Fref_def_hashtable_size))*8)) = def
			goto _1
		}
		if (*MD_CTX)(unsafe.Pointer(ctx)).Fref_defs <= bucket && bucket < (*MD_CTX)(unsafe.Pointer(ctx)).Fref_defs+uintptr((*MD_CTX)(unsafe.Pointer(ctx)).Fn_ref_defs)*40 {
			/* The bucket already contains one ref. def. Lets see whether it
			 * is the same label (ref. def. duplicate) or different one
			 * (hash conflict). */
			old_def = bucket
			if md_link_label_cmp(tls, (*MD_REF_DEF)(unsafe.Pointer(def)).Flabel, (*MD_REF_DEF)(unsafe.Pointer(def)).Flabel_size, (*MD_REF_DEF)(unsafe.Pointer(old_def)).Flabel, (*MD_REF_DEF)(unsafe.Pointer(old_def)).Flabel_size) == 0 {
				/* Duplicate label: Ignore this ref. def. */
				goto _1
			}
			/* Make the bucket complex, i.e. able to hold more ref. defs. */
			list = libc.Xmalloc(tls, libc.Uint64FromInt64(8)+libc.Uint64FromInt32(2)*libc.Uint64FromInt64(8))
			if list == libc.UintptrFromInt32(0) {
				if (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log != libc.UintptrFromInt32(0) {
					(*(*func(*libc.TLS, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log})))(tls, __ccgo_ts+1, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
				}
				goto abort
			}
			*(*uintptr)(unsafe.Pointer(list + 8)) = old_def
			*(*uintptr)(unsafe.Pointer(list + 8 + 1*8)) = def
			(*MD_REF_DEF_LIST)(unsafe.Pointer(list)).Fn_ref_defs = int32(2)
			(*MD_REF_DEF_LIST)(unsafe.Pointer(list)).Falloc_ref_defs = int32(2)
			*(*uintptr)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Fref_def_hashtable + uintptr((*MD_REF_DEF)(unsafe.Pointer(def)).Fhash%libc.Uint32FromInt32((*MD_CTX)(unsafe.Pointer(ctx)).Fref_def_hashtable_size))*8)) = list
			goto _1
		}
		/* Append the def to the complex bucket list.
		 *
		 * Note in this case we ignore potential duplicates to avoid expensive
		 * iterating over the complex bucket. Below, we revisit all the complex
		 * buckets and handle it more cheaply after the complex bucket contents
		 * is sorted. */
		list = bucket
		if (*MD_REF_DEF_LIST)(unsafe.Pointer(list)).Fn_ref_defs >= (*MD_REF_DEF_LIST)(unsafe.Pointer(list)).Falloc_ref_defs {
			alloc_ref_defs = (*MD_REF_DEF_LIST)(unsafe.Pointer(list)).Falloc_ref_defs + (*MD_REF_DEF_LIST)(unsafe.Pointer(list)).Falloc_ref_defs/int32(2)
			list_tmp = libc.Xrealloc(tls, list, uint64(8)+libc.Uint64FromInt32(alloc_ref_defs)*uint64(8))
			if list_tmp == libc.UintptrFromInt32(0) {
				if (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log != libc.UintptrFromInt32(0) {
					(*(*func(*libc.TLS, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log})))(tls, __ccgo_ts+43, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
				}
				goto abort
			}
			list = list_tmp
			(*MD_REF_DEF_LIST)(unsafe.Pointer(list)).Falloc_ref_defs = alloc_ref_defs
			*(*uintptr)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Fref_def_hashtable + uintptr((*MD_REF_DEF)(unsafe.Pointer(def)).Fhash%libc.Uint32FromInt32((*MD_CTX)(unsafe.Pointer(ctx)).Fref_def_hashtable_size))*8)) = list
		}
		*(*uintptr)(unsafe.Pointer(list + 8 + uintptr((*MD_REF_DEF_LIST)(unsafe.Pointer(list)).Fn_ref_defs)*8)) = def
		(*MD_REF_DEF_LIST)(unsafe.Pointer(list)).Fn_ref_defs = (*MD_REF_DEF_LIST)(unsafe.Pointer(list)).Fn_ref_defs + 1
		goto _1
	_1:
		;
		i = i + 1
	}
	/* Sort the complex buckets so we can use bsearch() with them. */
	i = 0
	for {
		if !(i < (*MD_CTX)(unsafe.Pointer(ctx)).Fref_def_hashtable_size) {
			break
		}
		bucket1 = *(*uintptr)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Fref_def_hashtable + uintptr(i)*8))
		if bucket1 == libc.UintptrFromInt32(0) {
			goto _2
		}
		if (*MD_CTX)(unsafe.Pointer(ctx)).Fref_defs <= bucket1 && bucket1 < (*MD_CTX)(unsafe.Pointer(ctx)).Fref_defs+uintptr((*MD_CTX)(unsafe.Pointer(ctx)).Fn_ref_defs)*40 {
			goto _2
		}
		list1 = bucket1
		libc.Xqsort(tls, list1+8, libc.Uint64FromInt32((*MD_REF_DEF_LIST)(unsafe.Pointer(list1)).Fn_ref_defs), uint64(8), __ccgo_fp(md_ref_def_cmp_for_sort))
		/* Disable all duplicates in the complex bucket by forcing all such
		 * records to point to the 1st such ref. def. I.e. no matter which
		 * record is found during the lookup, it will always point to the right
		 * ref. def. in ctx->ref_defs[]. */
		j = int32(1)
		for {
			if !(j < (*MD_REF_DEF_LIST)(unsafe.Pointer(list1)).Fn_ref_defs) {
				break
			}
			if md_ref_def_cmp(tls, list1+8+uintptr(j-int32(1))*8, list1+8+uintptr(j)*8) == 0 {
				*(*uintptr)(unsafe.Pointer(list1 + 8 + uintptr(j)*8)) = *(*uintptr)(unsafe.Pointer(list1 + 8 + uintptr(j-int32(1))*8))
			}
			goto _3
		_3:
			;
			j = j + 1
		}
		goto _2
	_2:
		;
		i = i + 1
	}
	return 0
	goto abort
abort:
	;
	return -int32(1)
	return r
}

func md_free_ref_def_hashtable(tls *libc.TLS, ctx uintptr) {
	var bucket uintptr
	var i int32
	_, _ = bucket, i
	if (*MD_CTX)(unsafe.Pointer(ctx)).Fref_def_hashtable != libc.UintptrFromInt32(0) {
		i = 0
		for {
			if !(i < (*MD_CTX)(unsafe.Pointer(ctx)).Fref_def_hashtable_size) {
				break
			}
			bucket = *(*uintptr)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Fref_def_hashtable + uintptr(i)*8))
			if bucket == libc.UintptrFromInt32(0) {
				goto _1
			}
			if (*MD_CTX)(unsafe.Pointer(ctx)).Fref_defs <= bucket && bucket < (*MD_CTX)(unsafe.Pointer(ctx)).Fref_defs+uintptr((*MD_CTX)(unsafe.Pointer(ctx)).Fn_ref_defs)*40 {
				goto _1
			}
			libc.Xfree(tls, bucket)
			goto _1
		_1:
			;
			i = i + 1
		}
		libc.Xfree(tls, (*MD_CTX)(unsafe.Pointer(ctx)).Fref_def_hashtable)
	}
}

func md_lookup_ref_def(tls *libc.TLS, ctx uintptr, label uintptr, label_size MD_SIZE) (r uintptr) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var bucket, def, list, ret uintptr
	var hash uint32
	var _ /* key at bp+40 */ uintptr
	var _ /* key_buf at bp+0 */ MD_REF_DEF
	_, _, _, _, _ = bucket, def, hash, list, ret
	if (*MD_CTX)(unsafe.Pointer(ctx)).Fref_def_hashtable_size == 0 {
		return libc.UintptrFromInt32(0)
	}
	hash = md_link_label_hash(tls, label, label_size)
	bucket = *(*uintptr)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Fref_def_hashtable + uintptr(hash%libc.Uint32FromInt32((*MD_CTX)(unsafe.Pointer(ctx)).Fref_def_hashtable_size))*8))
	if bucket == libc.UintptrFromInt32(0) {
		return libc.UintptrFromInt32(0)
	} else {
		if (*MD_CTX)(unsafe.Pointer(ctx)).Fref_defs <= bucket && bucket < (*MD_CTX)(unsafe.Pointer(ctx)).Fref_defs+uintptr((*MD_CTX)(unsafe.Pointer(ctx)).Fn_ref_defs)*40 {
			def = bucket
			if md_link_label_cmp(tls, (*MD_REF_DEF)(unsafe.Pointer(def)).Flabel, (*MD_REF_DEF)(unsafe.Pointer(def)).Flabel_size, label, label_size) == 0 {
				return def
			} else {
				return libc.UintptrFromInt32(0)
			}
		} else {
			list = bucket
			*(*uintptr)(__ccgo_ptr(bp + 40)) = bp
			(*(*MD_REF_DEF)(__ccgo_ptr(bp))).Flabel = label
			(*(*MD_REF_DEF)(__ccgo_ptr(bp))).Flabel_size = label_size
			(*(*MD_REF_DEF)(__ccgo_ptr(bp))).Fhash = md_link_label_hash(tls, (*(*MD_REF_DEF)(__ccgo_ptr(bp))).Flabel, (*(*MD_REF_DEF)(__ccgo_ptr(bp))).Flabel_size)
			ret = libc.Xbsearch(tls, bp+40, list+8, libc.Uint64FromInt32((*MD_REF_DEF_LIST)(unsafe.Pointer(list)).Fn_ref_defs), uint64(8), __ccgo_fp(md_ref_def_cmp))
			if ret != libc.UintptrFromInt32(0) {
				return *(*uintptr)(__ccgo_ptr(ret))
			} else {
				return libc.UintptrFromInt32(0)
			}
		}
	}
	return r
}

/***************************
 ***  Recognizing Links  ***
 ***************************/

/* Note this code is partially shared between processing inlines and blocks
 * as reference definitions and links share some helper parser functions.
 */

type MD_LINK_ATTR = struct {
	Fdest_beg         MD_OFFSET
	Fdest_end         MD_OFFSET
	Ftitle            uintptr
	Ftitle_size       MD_SIZE
	Ftitle_needs_free int32
}

/***************************
 ***  Recognizing Links  ***
 ***************************/

/* Note this code is partially shared between processing inlines and blocks
 * as reference definitions and links share some helper parser functions.
 */

type MD_LINK_ATTR_tag = MD_LINK_ATTR

func md_is_link_label(tls *libc.TLS, ctx uintptr, lines uintptr, n_lines MD_SIZE, beg MD_OFFSET, p_end uintptr, p_beg_line_index uintptr, p_end_line_index uintptr, p_contents_beg uintptr, p_contents_end uintptr) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var codepoint uint32
	var contents_beg, contents_end, line_end, off MD_OFFSET
	var len1 int32
	var line_index MD_SIZE
	var _ /* char_size at bp+0 */ MD_SIZE
	_, _, _, _, _, _, _ = codepoint, contents_beg, contents_end, len1, line_end, line_index, off
	off = beg
	contents_beg = uint32(0)
	contents_end = uint32(0)
	line_index = uint32(0)
	len1 = 0
	if libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) != int32('[') {
		return FALSE
	}
	off = off + 1
	for int32(1) != 0 {
		line_end = (*(*MD_LINE)(__ccgo_ptr(lines + uintptr(line_index)*8))).Fend
		for off < line_end {
			if libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('\\') && off+uint32(1) < (*MD_CTX)(unsafe.Pointer(ctx)).Fsize && (libc.Uint32FromInt32(libc.Int32FromInt32(33)) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))) <= libc.Uint32FromInt32(libc.Int32FromInt32(47)) || libc.Uint32FromInt32(libc.Int32FromInt32(58)) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))) <= libc.Uint32FromInt32(libc.Int32FromInt32(64)) || libc.Uint32FromInt32(libc.Int32FromInt32(91)) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))) <= libc.Uint32FromInt32(libc.Int32FromInt32(96)) || libc.Uint32FromInt32(libc.Int32FromInt32(123)) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))) <= libc.Uint32FromInt32(libc.Int32FromInt32(126)) || (libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))) == int32('\r') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))) == int32('\n'))) {
				if contents_end == uint32(0) {
					contents_beg = off
					*(*MD_SIZE)(__ccgo_ptr(p_beg_line_index)) = line_index
				}
				contents_end = off + uint32(2)
				off = off + uint32(2)
			} else {
				if libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('[') {
					return FALSE
				} else {
					if libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32(']') {
						if contents_beg < contents_end {
							/* Success. */
							*(*MD_OFFSET)(__ccgo_ptr(p_contents_beg)) = contents_beg
							*(*MD_OFFSET)(__ccgo_ptr(p_contents_end)) = contents_end
							*(*MD_OFFSET)(__ccgo_ptr(p_end)) = off + uint32(1)
							*(*MD_SIZE)(__ccgo_ptr(p_end_line_index)) = line_index
							return int32(TRUE)
						} else {
							/* Link label must have some non-whitespace contents. */
							return FALSE
						}
					} else {
						codepoint = md_decode_unicode(tls, (*MD_CTX)(unsafe.Pointer(ctx)).Ftext, off, (*MD_CTX)(unsafe.Pointer(ctx)).Fsize, bp)
						if !(md_is_unicode_whitespace__(tls, codepoint) != 0) {
							if contents_end == uint32(0) {
								contents_beg = off
								*(*MD_SIZE)(__ccgo_ptr(p_beg_line_index)) = line_index
							}
							contents_end = off + *(*MD_SIZE)(__ccgo_ptr(bp))
						}
						off = off + *(*MD_SIZE)(__ccgo_ptr(bp))
					}
				}
			}
			len1 = len1 + 1
			if len1 > int32(999) {
				return FALSE
			}
		}
		line_index = line_index + 1
		len1 = len1 + 1
		if line_index < n_lines {
			off = (*(*MD_LINE)(__ccgo_ptr(lines + uintptr(line_index)*8))).Fbeg
		} else {
			break
		}
	}
	return FALSE
}

func md_is_link_destination_A(tls *libc.TLS, ctx uintptr, beg MD_OFFSET, max_end MD_OFFSET, p_end uintptr, p_contents_beg uintptr, p_contents_end uintptr) (r int32) {
	var off MD_OFFSET
	_ = off
	off = beg
	if off >= max_end || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) != int32('<') {
		return FALSE
	}
	off = off + 1
	for off < max_end {
		if libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('\\') && off+uint32(1) < max_end && (libc.Uint32FromInt32(libc.Int32FromInt32(33)) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))) <= libc.Uint32FromInt32(libc.Int32FromInt32(47)) || libc.Uint32FromInt32(libc.Int32FromInt32(58)) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))) <= libc.Uint32FromInt32(libc.Int32FromInt32(64)) || libc.Uint32FromInt32(libc.Int32FromInt32(91)) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))) <= libc.Uint32FromInt32(libc.Int32FromInt32(96)) || libc.Uint32FromInt32(libc.Int32FromInt32(123)) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))) <= libc.Uint32FromInt32(libc.Int32FromInt32(126))) {
			off = off + uint32(2)
			continue
		}
		if libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('\r') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('\n') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('<') {
			return FALSE
		}
		if libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('>') {
			/* Success. */
			*(*MD_OFFSET)(__ccgo_ptr(p_contents_beg)) = beg + uint32(1)
			*(*MD_OFFSET)(__ccgo_ptr(p_contents_end)) = off
			*(*MD_OFFSET)(__ccgo_ptr(p_end)) = off + uint32(1)
			return int32(TRUE)
		}
		off = off + 1
	}
	return FALSE
}

func md_is_link_destination_B(tls *libc.TLS, ctx uintptr, beg MD_OFFSET, max_end MD_OFFSET, p_end uintptr, p_contents_beg uintptr, p_contents_end uintptr) (r int32) {
	var off MD_OFFSET
	var parenthesis_level int32
	_, _ = off, parenthesis_level
	off = beg
	parenthesis_level = 0
	for off < max_end {
		if libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('\\') && off+uint32(1) < max_end && (libc.Uint32FromInt32(libc.Int32FromInt32(33)) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))) <= libc.Uint32FromInt32(libc.Int32FromInt32(47)) || libc.Uint32FromInt32(libc.Int32FromInt32(58)) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))) <= libc.Uint32FromInt32(libc.Int32FromInt32(64)) || libc.Uint32FromInt32(libc.Int32FromInt32(91)) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))) <= libc.Uint32FromInt32(libc.Int32FromInt32(96)) || libc.Uint32FromInt32(libc.Int32FromInt32(123)) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))) <= libc.Uint32FromInt32(libc.Int32FromInt32(126))) {
			off = off + uint32(2)
			continue
		}
		if libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32(' ') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('\t') || (libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('\v') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('\f')) || (uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) <= uint32(31) || uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == uint32(127)) {
			break
		}
		/* Link destination may include balanced pairs of unescaped '(' ')'.
		 * Note we limit the maximal nesting level by 32 to protect us from
		 * https://github.com/jgm/cmark/issues/214 */
		if libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('(') {
			parenthesis_level = parenthesis_level + 1
			if parenthesis_level > int32(32) {
				return FALSE
			}
		} else {
			if libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32(')') {
				if parenthesis_level == 0 {
					break
				}
				parenthesis_level = parenthesis_level - 1
			}
		}
		off = off + 1
	}
	if parenthesis_level != 0 || off == beg {
		return FALSE
	}
	/* Success. */
	*(*MD_OFFSET)(__ccgo_ptr(p_contents_beg)) = beg
	*(*MD_OFFSET)(__ccgo_ptr(p_contents_end)) = off
	*(*MD_OFFSET)(__ccgo_ptr(p_end)) = off
	return int32(TRUE)
}

func md_is_link_destination(tls *libc.TLS, ctx uintptr, beg MD_OFFSET, max_end MD_OFFSET, p_end uintptr, p_contents_beg uintptr, p_contents_end uintptr) (r int32) {
	if libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(beg)))) == int32('<') {
		return md_is_link_destination_A(tls, ctx, beg, max_end, p_end, p_contents_beg, p_contents_end)
	} else {
		return md_is_link_destination_B(tls, ctx, beg, max_end, p_end, p_contents_beg, p_contents_end)
	}
	return r
}

func md_is_link_title(tls *libc.TLS, ctx uintptr, lines uintptr, n_lines MD_SIZE, beg MD_OFFSET, p_end uintptr, p_beg_line_index uintptr, p_end_line_index uintptr, p_contents_beg uintptr, p_contents_end uintptr) (r int32) {
	var closer_char MD_CHAR
	var line_end, off MD_OFFSET
	var line_index MD_SIZE
	_, _, _, _ = closer_char, line_end, line_index, off
	off = beg
	line_index = uint32(0)
	/* White space with up to one line break. */
	for off < (*(*MD_LINE)(__ccgo_ptr(lines + uintptr(line_index)*8))).Fend && (libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32(' ') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('\t') || (libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('\v') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('\f'))) {
		off = off + 1
	}
	if off >= (*(*MD_LINE)(__ccgo_ptr(lines + uintptr(line_index)*8))).Fend {
		line_index = line_index + 1
		if line_index >= n_lines {
			return FALSE
		}
		off = (*(*MD_LINE)(__ccgo_ptr(lines + uintptr(line_index)*8))).Fbeg
	}
	if off == beg {
		return FALSE
	}
	*(*MD_SIZE)(__ccgo_ptr(p_beg_line_index)) = line_index
	/* First char determines how to detect end of it. */
	switch libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) {
	case int32('"'):
		closer_char = uint8('"')
	case int32('\''):
		closer_char = uint8('\'')
	case int32('('):
		closer_char = uint8(')')
	default:
		return FALSE
	}
	off = off + 1
	*(*MD_OFFSET)(__ccgo_ptr(p_contents_beg)) = off
	for line_index < n_lines {
		line_end = (*(*MD_LINE)(__ccgo_ptr(lines + uintptr(line_index)*8))).Fend
		for off < line_end {
			if libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('\\') && off+uint32(1) < (*MD_CTX)(unsafe.Pointer(ctx)).Fsize && (libc.Uint32FromInt32(libc.Int32FromInt32(33)) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))) <= libc.Uint32FromInt32(libc.Int32FromInt32(47)) || libc.Uint32FromInt32(libc.Int32FromInt32(58)) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))) <= libc.Uint32FromInt32(libc.Int32FromInt32(64)) || libc.Uint32FromInt32(libc.Int32FromInt32(91)) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))) <= libc.Uint32FromInt32(libc.Int32FromInt32(96)) || libc.Uint32FromInt32(libc.Int32FromInt32(123)) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))) <= libc.Uint32FromInt32(libc.Int32FromInt32(126)) || (libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))) == int32('\r') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))) == int32('\n'))) {
				off = off + 1
			} else {
				if libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == libc.Int32FromUint8(closer_char) {
					/* Success. */
					*(*MD_OFFSET)(__ccgo_ptr(p_contents_end)) = off
					*(*MD_OFFSET)(__ccgo_ptr(p_end)) = off + uint32(1)
					*(*MD_SIZE)(__ccgo_ptr(p_end_line_index)) = line_index
					return int32(TRUE)
				} else {
					if libc.Int32FromUint8(closer_char) == int32(')') && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('(') {
						/* ()-style title cannot contain (unescaped '(')) */
						return FALSE
					}
				}
			}
			off = off + 1
		}
		line_index = line_index + 1
	}
	return FALSE
}

// C documentation
//
//	/* Returns 0 if it is not a reference definition.
//	 *
//	 * Returns N > 0 if it is a reference definition. N then corresponds to the
//	 * number of lines forming it). In this case the definition is stored for
//	 * resolving any links referring to it.
//	 *
//	 * Returns -1 in case of an error (out of memory).
//	 */
func md_is_link_reference_definition(tls *libc.TLS, ctx uintptr, lines uintptr, n_lines MD_SIZE) (r int32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var def, new_defs uintptr
	var label_is_multiline, ret, title_is_multiline, v1 int32
	var _ /* dest_contents_beg at bp+12 */ MD_OFFSET
	var _ /* dest_contents_end at bp+16 */ MD_OFFSET
	var _ /* label_contents_beg at bp+0 */ MD_OFFSET
	var _ /* label_contents_end at bp+4 */ MD_OFFSET
	var _ /* label_contents_line_index at bp+8 */ MD_SIZE
	var _ /* line_index at bp+36 */ MD_SIZE
	var _ /* off at bp+32 */ MD_OFFSET
	var _ /* title_contents_beg at bp+20 */ MD_OFFSET
	var _ /* title_contents_end at bp+24 */ MD_OFFSET
	var _ /* title_contents_line_index at bp+28 */ MD_SIZE
	var _ /* tmp_line_index at bp+40 */ MD_SIZE
	_, _, _, _, _, _ = def, label_is_multiline, new_defs, ret, title_is_multiline, v1
	label_is_multiline = FALSE
	title_is_multiline = FALSE
	*(*MD_SIZE)(__ccgo_ptr(bp + 36)) = uint32(0)
	def = libc.UintptrFromInt32(0)
	ret = 0
	/* Link label. */
	if !(md_is_link_label(tls, ctx, lines, n_lines, (*(*MD_LINE)(__ccgo_ptr(lines))).Fbeg, bp+32, bp+8, bp+36, bp, bp+4) != 0) {
		return FALSE
	}
	label_is_multiline = libc.BoolInt32(*(*MD_SIZE)(__ccgo_ptr(bp + 8)) != *(*MD_SIZE)(__ccgo_ptr(bp + 36)))
	/* Colon. */
	if *(*MD_OFFSET)(__ccgo_ptr(bp + 32)) >= (*(*MD_LINE)(__ccgo_ptr(lines + uintptr(*(*MD_SIZE)(__ccgo_ptr(bp + 36)))*8))).Fend || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(*(*MD_OFFSET)(__ccgo_ptr(bp + 32)))))) != int32(':') {
		return FALSE
	}
	*(*MD_OFFSET)(__ccgo_ptr(bp + 32)) = *(*MD_OFFSET)(__ccgo_ptr(bp + 32)) + 1
	/* Optional white space with up to one line break. */
	for *(*MD_OFFSET)(__ccgo_ptr(bp + 32)) < (*(*MD_LINE)(__ccgo_ptr(lines + uintptr(*(*MD_SIZE)(__ccgo_ptr(bp + 36)))*8))).Fend && (libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(*(*MD_OFFSET)(__ccgo_ptr(bp + 32)))))) == int32(' ') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(*(*MD_OFFSET)(__ccgo_ptr(bp + 32)))))) == int32('\t') || (libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(*(*MD_OFFSET)(__ccgo_ptr(bp + 32)))))) == int32('\v') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(*(*MD_OFFSET)(__ccgo_ptr(bp + 32)))))) == int32('\f'))) {
		*(*MD_OFFSET)(__ccgo_ptr(bp + 32)) = *(*MD_OFFSET)(__ccgo_ptr(bp + 32)) + 1
	}
	if *(*MD_OFFSET)(__ccgo_ptr(bp + 32)) >= (*(*MD_LINE)(__ccgo_ptr(lines + uintptr(*(*MD_SIZE)(__ccgo_ptr(bp + 36)))*8))).Fend {
		*(*MD_SIZE)(__ccgo_ptr(bp + 36)) = *(*MD_SIZE)(__ccgo_ptr(bp + 36)) + 1
		if *(*MD_SIZE)(__ccgo_ptr(bp + 36)) >= n_lines {
			return FALSE
		}
		*(*MD_OFFSET)(__ccgo_ptr(bp + 32)) = (*(*MD_LINE)(__ccgo_ptr(lines + uintptr(*(*MD_SIZE)(__ccgo_ptr(bp + 36)))*8))).Fbeg
	}
	/* Link destination. */
	if !(md_is_link_destination(tls, ctx, *(*MD_OFFSET)(__ccgo_ptr(bp + 32)), (*(*MD_LINE)(__ccgo_ptr(lines + uintptr(*(*MD_SIZE)(__ccgo_ptr(bp + 36)))*8))).Fend, bp+32, bp+12, bp+16) != 0) {
		return FALSE
	}
	/* (Optional) title. Note we interpret it as an title only if nothing
	 * more follows on its last line. */
	if md_is_link_title(tls, ctx, lines+uintptr(*(*MD_SIZE)(__ccgo_ptr(bp + 36)))*8, n_lines-*(*MD_SIZE)(__ccgo_ptr(bp + 36)), *(*MD_OFFSET)(__ccgo_ptr(bp + 32)), bp+32, bp+28, bp+40, bp+20, bp+24) != 0 && *(*MD_OFFSET)(__ccgo_ptr(bp + 32)) >= (*(*MD_LINE)(__ccgo_ptr(lines + uintptr(*(*MD_SIZE)(__ccgo_ptr(bp + 36))+*(*MD_SIZE)(__ccgo_ptr(bp + 40)))*8))).Fend {
		title_is_multiline = libc.BoolInt32(*(*MD_SIZE)(__ccgo_ptr(bp + 40)) != *(*MD_SIZE)(__ccgo_ptr(bp + 28)))
		*(*MD_SIZE)(__ccgo_ptr(bp + 28)) = *(*MD_SIZE)(__ccgo_ptr(bp + 28)) + *(*MD_SIZE)(__ccgo_ptr(bp + 36))
		*(*MD_SIZE)(__ccgo_ptr(bp + 36)) = *(*MD_SIZE)(__ccgo_ptr(bp + 36)) + *(*MD_SIZE)(__ccgo_ptr(bp + 40))
	} else {
		/* Not a title. */
		title_is_multiline = FALSE
		*(*MD_OFFSET)(__ccgo_ptr(bp + 20)) = *(*MD_OFFSET)(__ccgo_ptr(bp + 32))
		*(*MD_OFFSET)(__ccgo_ptr(bp + 24)) = *(*MD_OFFSET)(__ccgo_ptr(bp + 32))
		*(*MD_SIZE)(__ccgo_ptr(bp + 28)) = uint32(0)
	}
	/* Nothing more can follow on the last line. */
	if *(*MD_OFFSET)(__ccgo_ptr(bp + 32)) < (*(*MD_LINE)(__ccgo_ptr(lines + uintptr(*(*MD_SIZE)(__ccgo_ptr(bp + 36)))*8))).Fend {
		return FALSE
	}
	/* So, it _is_ a reference definition. Remember it. */
	if (*MD_CTX)(unsafe.Pointer(ctx)).Fn_ref_defs >= (*MD_CTX)(unsafe.Pointer(ctx)).Falloc_ref_defs {
		if (*MD_CTX)(unsafe.Pointer(ctx)).Falloc_ref_defs > 0 {
			v1 = (*MD_CTX)(unsafe.Pointer(ctx)).Falloc_ref_defs + (*MD_CTX)(unsafe.Pointer(ctx)).Falloc_ref_defs/int32(2)
		} else {
			v1 = int32(16)
		}
		(*MD_CTX)(unsafe.Pointer(ctx)).Falloc_ref_defs = v1
		new_defs = libc.Xrealloc(tls, (*MD_CTX)(unsafe.Pointer(ctx)).Fref_defs, libc.Uint64FromInt32((*MD_CTX)(unsafe.Pointer(ctx)).Falloc_ref_defs)*uint64(40))
		if new_defs == libc.UintptrFromInt32(0) {
			if (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log != libc.UintptrFromInt32(0) {
				(*(*func(*libc.TLS, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log})))(tls, __ccgo_ts+43, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
			}
			goto abort
		}
		(*MD_CTX)(unsafe.Pointer(ctx)).Fref_defs = new_defs
	}
	def = (*MD_CTX)(unsafe.Pointer(ctx)).Fref_defs + uintptr((*MD_CTX)(unsafe.Pointer(ctx)).Fn_ref_defs)*40
	libc.Xmemset(tls, def, 0, uint64(40))
	if label_is_multiline != 0 {
		ret = md_merge_lines_alloc(tls, ctx, *(*MD_OFFSET)(__ccgo_ptr(bp)), *(*MD_OFFSET)(__ccgo_ptr(bp + 4)), lines+uintptr(*(*MD_SIZE)(__ccgo_ptr(bp + 8)))*8, n_lines-*(*MD_SIZE)(__ccgo_ptr(bp + 8)), uint8(' '), def, def+20)
		if ret < 0 {
			goto abort
		}
		libc.SetBitFieldPtr8Uint8(def+36, libc.Uint8FromInt32(TRUE), 0, 0x1)
	} else {
		(*MD_REF_DEF)(unsafe.Pointer(def)).Flabel = (*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(*(*MD_OFFSET)(__ccgo_ptr(bp)))
		(*MD_REF_DEF)(unsafe.Pointer(def)).Flabel_size = *(*MD_OFFSET)(__ccgo_ptr(bp + 4)) - *(*MD_OFFSET)(__ccgo_ptr(bp))
	}
	if title_is_multiline != 0 {
		ret = md_merge_lines_alloc(tls, ctx, *(*MD_OFFSET)(__ccgo_ptr(bp + 20)), *(*MD_OFFSET)(__ccgo_ptr(bp + 24)), lines+uintptr(*(*MD_SIZE)(__ccgo_ptr(bp + 28)))*8, n_lines-*(*MD_SIZE)(__ccgo_ptr(bp + 28)), uint8('\n'), def+8, def+24)
		if ret < 0 {
			goto abort
		}
		libc.SetBitFieldPtr8Uint8(def+36, libc.Uint8FromInt32(TRUE), 1, 0x2)
	} else {
		(*MD_REF_DEF)(unsafe.Pointer(def)).Ftitle = (*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(*(*MD_OFFSET)(__ccgo_ptr(bp + 20)))
		(*MD_REF_DEF)(unsafe.Pointer(def)).Ftitle_size = *(*MD_OFFSET)(__ccgo_ptr(bp + 24)) - *(*MD_OFFSET)(__ccgo_ptr(bp + 20))
	}
	(*MD_REF_DEF)(unsafe.Pointer(def)).Fdest_beg = *(*MD_OFFSET)(__ccgo_ptr(bp + 12))
	(*MD_REF_DEF)(unsafe.Pointer(def)).Fdest_end = *(*MD_OFFSET)(__ccgo_ptr(bp + 16))
	/* Success. */
	(*MD_CTX)(unsafe.Pointer(ctx)).Fn_ref_defs = (*MD_CTX)(unsafe.Pointer(ctx)).Fn_ref_defs + 1
	return libc.Int32FromUint32(*(*MD_SIZE)(__ccgo_ptr(bp + 36)) + uint32(1))
	goto abort
abort:
	;
	/* Failure. */
	if def != libc.UintptrFromInt32(0) && int32(*(*uint8)(unsafe.Pointer(def + 36))&0x1>>0) != 0 {
		libc.Xfree(tls, (*MD_REF_DEF)(unsafe.Pointer(def)).Flabel)
	}
	if def != libc.UintptrFromInt32(0) && int32(*(*uint8)(unsafe.Pointer(def + 36))&0x2>>1) != 0 {
		libc.Xfree(tls, (*MD_REF_DEF)(unsafe.Pointer(def)).Ftitle)
	}
	return ret
}

func md_is_link_reference(tls *libc.TLS, ctx uintptr, lines uintptr, n_lines MD_SIZE, beg MD_OFFSET, end MD_OFFSET, attr uintptr) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var beg_line, def uintptr
	var is_multiline, ret, v1 int32
	var _ /* label at bp+0 */ uintptr
	var _ /* label_size at bp+8 */ MD_SIZE
	_, _, _, _, _ = beg_line, def, is_multiline, ret, v1
	if !(libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(beg)))) == int32('[') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(beg)))) == int32('!')) {
		libc.X__builtin_unreachable(tls)
	}
	if !(libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(end-libc.Uint32FromInt32(1))))) == libc.Int32FromUint8(']')) {
		libc.X__builtin_unreachable(tls)
	}
	if libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(beg)))) == int32('!') {
		v1 = int32(2)
	} else {
		v1 = int32(1)
	}
	beg = beg + libc.Uint32FromInt32(v1)
	end = end - 1
	/* Find lines corresponding to the beg and end positions. */
	beg_line = md_lookup_line(tls, beg, lines, n_lines, libc.UintptrFromInt32(0))
	is_multiline = libc.BoolInt32(end > (*MD_LINE)(unsafe.Pointer(beg_line)).Fend)
	if is_multiline != 0 {
		ret = md_merge_lines_alloc(tls, ctx, beg, end, beg_line, libc.Uint32FromInt32(int32(libc.Int64FromUint32(n_lines)-(int64(beg_line)-int64(lines))/8)), uint8(' '), bp, bp+8)
		if ret < 0 {
			goto abort
		}
	} else {
		*(*uintptr)(__ccgo_ptr(bp)) = (*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(beg)
		*(*MD_SIZE)(__ccgo_ptr(bp + 8)) = end - beg
	}
	def = md_lookup_ref_def(tls, ctx, *(*uintptr)(__ccgo_ptr(bp)), *(*MD_SIZE)(__ccgo_ptr(bp + 8)))
	if def != libc.UintptrFromInt32(0) {
		(*MD_LINK_ATTR)(unsafe.Pointer(attr)).Fdest_beg = (*MD_REF_DEF)(unsafe.Pointer(def)).Fdest_beg
		(*MD_LINK_ATTR)(unsafe.Pointer(attr)).Fdest_end = (*MD_REF_DEF)(unsafe.Pointer(def)).Fdest_end
		(*MD_LINK_ATTR)(unsafe.Pointer(attr)).Ftitle = (*MD_REF_DEF)(unsafe.Pointer(def)).Ftitle
		(*MD_LINK_ATTR)(unsafe.Pointer(attr)).Ftitle_size = (*MD_REF_DEF)(unsafe.Pointer(def)).Ftitle_size
		(*MD_LINK_ATTR)(unsafe.Pointer(attr)).Ftitle_needs_free = FALSE
	}
	if is_multiline != 0 {
		libc.Xfree(tls, *(*uintptr)(__ccgo_ptr(bp)))
	}
	ret = libc.BoolInt32(def != libc.UintptrFromInt32(0))
	goto abort
abort:
	;
	return ret
	return r
}

func md_is_inline_link_spec(tls *libc.TLS, ctx uintptr, lines uintptr, n_lines MD_SIZE, beg MD_OFFSET, p_end uintptr, attr uintptr) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var line_index MD_SIZE
	var ret, title_is_multiline int32
	var _ /* off at bp+16 */ MD_OFFSET
	var _ /* title_contents_beg at bp+4 */ MD_OFFSET
	var _ /* title_contents_end at bp+8 */ MD_OFFSET
	var _ /* title_contents_line_index at bp+12 */ MD_SIZE
	var _ /* tmp_line_index at bp+0 */ MD_SIZE
	_, _, _ = line_index, ret, title_is_multiline
	line_index = uint32(0)
	*(*MD_OFFSET)(__ccgo_ptr(bp + 16)) = beg
	ret = FALSE
	for *(*MD_OFFSET)(__ccgo_ptr(bp + 16)) >= (*(*MD_LINE)(__ccgo_ptr(lines + uintptr(line_index)*8))).Fend {
		line_index = line_index + 1
	}
	if !(libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(*(*MD_OFFSET)(__ccgo_ptr(bp + 16)))))) == libc.Int32FromUint8('(')) {
		libc.X__builtin_unreachable(tls)
	}
	*(*MD_OFFSET)(__ccgo_ptr(bp + 16)) = *(*MD_OFFSET)(__ccgo_ptr(bp + 16)) + 1
	/* Optional white space with up to one line break. */
	for *(*MD_OFFSET)(__ccgo_ptr(bp + 16)) < (*(*MD_LINE)(__ccgo_ptr(lines + uintptr(line_index)*8))).Fend && (libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(*(*MD_OFFSET)(__ccgo_ptr(bp + 16)))))) == int32(' ') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(*(*MD_OFFSET)(__ccgo_ptr(bp + 16)))))) == int32('\t') || (libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(*(*MD_OFFSET)(__ccgo_ptr(bp + 16)))))) == int32('\v') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(*(*MD_OFFSET)(__ccgo_ptr(bp + 16)))))) == int32('\f'))) {
		*(*MD_OFFSET)(__ccgo_ptr(bp + 16)) = *(*MD_OFFSET)(__ccgo_ptr(bp + 16)) + 1
	}
	if *(*MD_OFFSET)(__ccgo_ptr(bp + 16)) >= (*(*MD_LINE)(__ccgo_ptr(lines + uintptr(line_index)*8))).Fend && (*(*MD_OFFSET)(__ccgo_ptr(bp + 16)) >= (*MD_CTX)(unsafe.Pointer(ctx)).Fsize || (libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(*(*MD_OFFSET)(__ccgo_ptr(bp + 16)))))) == int32('\r') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(*(*MD_OFFSET)(__ccgo_ptr(bp + 16)))))) == int32('\n'))) {
		line_index = line_index + 1
		if line_index >= n_lines {
			return FALSE
		}
		*(*MD_OFFSET)(__ccgo_ptr(bp + 16)) = (*(*MD_LINE)(__ccgo_ptr(lines + uintptr(line_index)*8))).Fbeg
	}
	/* Link destination may be omitted, but only when not also having a title. */
	if *(*MD_OFFSET)(__ccgo_ptr(bp + 16)) < (*MD_CTX)(unsafe.Pointer(ctx)).Fsize && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(*(*MD_OFFSET)(__ccgo_ptr(bp + 16)))))) == int32(')') {
		(*MD_LINK_ATTR)(unsafe.Pointer(attr)).Fdest_beg = *(*MD_OFFSET)(__ccgo_ptr(bp + 16))
		(*MD_LINK_ATTR)(unsafe.Pointer(attr)).Fdest_end = *(*MD_OFFSET)(__ccgo_ptr(bp + 16))
		(*MD_LINK_ATTR)(unsafe.Pointer(attr)).Ftitle = libc.UintptrFromInt32(0)
		(*MD_LINK_ATTR)(unsafe.Pointer(attr)).Ftitle_size = uint32(0)
		(*MD_LINK_ATTR)(unsafe.Pointer(attr)).Ftitle_needs_free = FALSE
		*(*MD_OFFSET)(__ccgo_ptr(bp + 16)) = *(*MD_OFFSET)(__ccgo_ptr(bp + 16)) + 1
		*(*MD_OFFSET)(__ccgo_ptr(p_end)) = *(*MD_OFFSET)(__ccgo_ptr(bp + 16))
		return int32(TRUE)
	}
	/* Link destination. */
	if !(md_is_link_destination(tls, ctx, *(*MD_OFFSET)(__ccgo_ptr(bp + 16)), (*(*MD_LINE)(__ccgo_ptr(lines + uintptr(line_index)*8))).Fend, bp+16, attr, attr+4) != 0) {
		return FALSE
	}
	/* (Optional) title. */
	if md_is_link_title(tls, ctx, lines+uintptr(line_index)*8, n_lines-line_index, *(*MD_OFFSET)(__ccgo_ptr(bp + 16)), bp+16, bp+12, bp, bp+4, bp+8) != 0 {
		title_is_multiline = libc.BoolInt32(*(*MD_SIZE)(__ccgo_ptr(bp)) != *(*MD_SIZE)(__ccgo_ptr(bp + 12)))
		*(*MD_SIZE)(__ccgo_ptr(bp + 12)) = *(*MD_SIZE)(__ccgo_ptr(bp + 12)) + line_index
		line_index = line_index + *(*MD_SIZE)(__ccgo_ptr(bp))
	} else {
		/* Not a title. */
		title_is_multiline = FALSE
		*(*MD_OFFSET)(__ccgo_ptr(bp + 4)) = *(*MD_OFFSET)(__ccgo_ptr(bp + 16))
		*(*MD_OFFSET)(__ccgo_ptr(bp + 8)) = *(*MD_OFFSET)(__ccgo_ptr(bp + 16))
		*(*MD_SIZE)(__ccgo_ptr(bp + 12)) = uint32(0)
	}
	/* Optional whitespace followed with final ')'. */
	for *(*MD_OFFSET)(__ccgo_ptr(bp + 16)) < (*(*MD_LINE)(__ccgo_ptr(lines + uintptr(line_index)*8))).Fend && (libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(*(*MD_OFFSET)(__ccgo_ptr(bp + 16)))))) == int32(' ') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(*(*MD_OFFSET)(__ccgo_ptr(bp + 16)))))) == int32('\t') || (libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(*(*MD_OFFSET)(__ccgo_ptr(bp + 16)))))) == int32('\v') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(*(*MD_OFFSET)(__ccgo_ptr(bp + 16)))))) == int32('\f'))) {
		*(*MD_OFFSET)(__ccgo_ptr(bp + 16)) = *(*MD_OFFSET)(__ccgo_ptr(bp + 16)) + 1
	}
	if *(*MD_OFFSET)(__ccgo_ptr(bp + 16)) >= (*(*MD_LINE)(__ccgo_ptr(lines + uintptr(line_index)*8))).Fend {
		line_index = line_index + 1
		if line_index >= n_lines {
			return FALSE
		}
		*(*MD_OFFSET)(__ccgo_ptr(bp + 16)) = (*(*MD_LINE)(__ccgo_ptr(lines + uintptr(line_index)*8))).Fbeg
	}
	if libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(*(*MD_OFFSET)(__ccgo_ptr(bp + 16)))))) != int32(')') {
		goto abort
	}
	*(*MD_OFFSET)(__ccgo_ptr(bp + 16)) = *(*MD_OFFSET)(__ccgo_ptr(bp + 16)) + 1
	if *(*MD_OFFSET)(__ccgo_ptr(bp + 4)) >= *(*MD_OFFSET)(__ccgo_ptr(bp + 8)) {
		(*MD_LINK_ATTR)(unsafe.Pointer(attr)).Ftitle = libc.UintptrFromInt32(0)
		(*MD_LINK_ATTR)(unsafe.Pointer(attr)).Ftitle_size = uint32(0)
		(*MD_LINK_ATTR)(unsafe.Pointer(attr)).Ftitle_needs_free = FALSE
	} else {
		if !(title_is_multiline != 0) {
			(*MD_LINK_ATTR)(unsafe.Pointer(attr)).Ftitle = (*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(*(*MD_OFFSET)(__ccgo_ptr(bp + 4)))
			(*MD_LINK_ATTR)(unsafe.Pointer(attr)).Ftitle_size = *(*MD_OFFSET)(__ccgo_ptr(bp + 8)) - *(*MD_OFFSET)(__ccgo_ptr(bp + 4))
			(*MD_LINK_ATTR)(unsafe.Pointer(attr)).Ftitle_needs_free = FALSE
		} else {
			ret = md_merge_lines_alloc(tls, ctx, *(*MD_OFFSET)(__ccgo_ptr(bp + 4)), *(*MD_OFFSET)(__ccgo_ptr(bp + 8)), lines+uintptr(*(*MD_SIZE)(__ccgo_ptr(bp + 12)))*8, n_lines-*(*MD_SIZE)(__ccgo_ptr(bp + 12)), uint8('\n'), attr+8, attr+16)
			if ret < 0 {
				goto abort
			}
			(*MD_LINK_ATTR)(unsafe.Pointer(attr)).Ftitle_needs_free = int32(TRUE)
		}
	}
	*(*MD_OFFSET)(__ccgo_ptr(p_end)) = *(*MD_OFFSET)(__ccgo_ptr(bp + 16))
	ret = int32(TRUE)
	goto abort
abort:
	;
	return ret
	return r
}

func md_free_ref_defs(tls *libc.TLS, ctx uintptr) {
	var def uintptr
	var i int32
	_, _ = def, i
	i = 0
	for {
		if !(i < (*MD_CTX)(unsafe.Pointer(ctx)).Fn_ref_defs) {
			break
		}
		def = (*MD_CTX)(unsafe.Pointer(ctx)).Fref_defs + uintptr(i)*40
		if int32(*(*uint8)(unsafe.Pointer(def + 36))&0x1>>0) != 0 {
			libc.Xfree(tls, (*MD_REF_DEF)(unsafe.Pointer(def)).Flabel)
		}
		if int32(*(*uint8)(unsafe.Pointer(def + 36))&0x2>>1) != 0 {
			libc.Xfree(tls, (*MD_REF_DEF)(unsafe.Pointer(def)).Ftitle)
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	libc.Xfree(tls, (*MD_CTX)(unsafe.Pointer(ctx)).Fref_defs)
}

/* Mark flags (these apply to ALL mark types). */

/* Mark flags specific for various mark types (so they can share bits). */

func md_emph_stack(tls *libc.TLS, ctx uintptr, ch MD_CHAR, flags uint32) (r uintptr) {
	var stack uintptr
	_ = stack
	switch libc.Int32FromUint8(ch) {
	case int32('*'):
		goto _1
	case int32('_'):
		goto _2
	default:
		goto _3
	}
	goto _4
_1:
	;
	stack = ctx + 416
	goto _4
_2:
	;
	stack = ctx + 416 + 6*4
	goto _4
_3:
	;
_7:
	;
	libc.X__builtin_unreachable(tls)
	goto _6
_6:
	;
	if 0 != 0 {
		goto _7
	}
	goto _5
_5:
	;
_4:
	;
	if flags&uint32(MD_MARK_EMPH_OC) != 0 {
		stack = stack + uintptr(3)*4
	}
	switch flags & libc.Uint32FromInt32(libc.Int32FromInt32(0x40)|libc.Int32FromInt32(0x80)) {
	case uint32(MD_MARK_EMPH_MOD3_0):
		goto _8
	case uint32(MD_MARK_EMPH_MOD3_1):
		goto _9
	case libc.Uint32FromInt32(libc.Int32FromInt32(0x40) | libc.Int32FromInt32(0x80)):
		goto _10
	default:
		goto _11
	}
	goto _12
_8:
	;
	stack = stack + uintptr(0)*4
	goto _12
_9:
	;
	stack = stack + uintptr(1)*4
	goto _12
_10:
	;
	stack = stack + uintptr(2)*4
	goto _12
_11:
	;
_15:
	;
	libc.X__builtin_unreachable(tls)
	goto _14
_14:
	;
	if 0 != 0 {
		goto _15
	}
	goto _13
_13:
	;
_12:
	;
	return stack
}

func md_opener_stack(tls *libc.TLS, ctx uintptr, mark_index int32) (r uintptr) {
	var mark, v8 uintptr
	_, _ = mark, v8
	mark = (*MD_CTX)(unsafe.Pointer(ctx)).Fmarks + uintptr(mark_index)*20
	switch libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(mark)).Fch) {
	case int32('_'):
		goto _1
	case int32('*'):
		goto _2
	case int32('~'):
		goto _3
	case int32('['):
		goto _4
	case int32('!'):
		goto _5
	default:
		goto _6
	}
	goto _7
_2:
	;
_1:
	;
	return md_emph_stack(tls, ctx, (*MD_MARK)(unsafe.Pointer(mark)).Fch, uint32((*MD_MARK)(unsafe.Pointer(mark)).Fflags))
_3:
	;
	if (*MD_MARK)(unsafe.Pointer(mark)).Fend-(*MD_MARK)(unsafe.Pointer(mark)).Fbeg == uint32(1) {
		v8 = ctx + 416 + 12*4
	} else {
		v8 = ctx + 416 + 13*4
	}
	return v8
_5:
	;
_4:
	;
	return ctx + 416 + 14*4
_6:
	;
_11:
	;
	libc.X__builtin_unreachable(tls)
	goto _10
_10:
	;
	if 0 != 0 {
		goto _11
	}
	goto _9
_9:
	;
_7:
	;
	return r
}

func md_add_mark(tls *libc.TLS, ctx uintptr) (r uintptr) {
	var new_marks, v3 uintptr
	var v1 int32
	_, _, _ = new_marks, v1, v3
	if (*MD_CTX)(unsafe.Pointer(ctx)).Fn_marks >= (*MD_CTX)(unsafe.Pointer(ctx)).Falloc_marks {
		if (*MD_CTX)(unsafe.Pointer(ctx)).Falloc_marks > 0 {
			v1 = (*MD_CTX)(unsafe.Pointer(ctx)).Falloc_marks + (*MD_CTX)(unsafe.Pointer(ctx)).Falloc_marks/int32(2)
		} else {
			v1 = int32(64)
		}
		(*MD_CTX)(unsafe.Pointer(ctx)).Falloc_marks = v1
		new_marks = libc.Xrealloc(tls, (*MD_CTX)(unsafe.Pointer(ctx)).Fmarks, libc.Uint64FromInt32((*MD_CTX)(unsafe.Pointer(ctx)).Falloc_marks)*uint64(20))
		if new_marks == libc.UintptrFromInt32(0) {
			if (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log != libc.UintptrFromInt32(0) {
				(*(*func(*libc.TLS, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log})))(tls, __ccgo_ts+43, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
			}
			return libc.UintptrFromInt32(0)
		}
		(*MD_CTX)(unsafe.Pointer(ctx)).Fmarks = new_marks
	}
	v3 = ctx + 152
	v1 = *(*int32)(unsafe.Pointer(v3))
	*(*int32)(unsafe.Pointer(v3)) = *(*int32)(unsafe.Pointer(v3)) + 1
	return (*MD_CTX)(unsafe.Pointer(ctx)).Fmarks + uintptr(v1)*20
}

func md_mark_stack_push(tls *libc.TLS, ctx uintptr, stack uintptr, mark_index int32) {
	(*(*MD_MARK)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Fmarks + uintptr(mark_index)*20))).Fnext = (*MD_MARKSTACK)(unsafe.Pointer(stack)).Ftop
	(*MD_MARKSTACK)(unsafe.Pointer(stack)).Ftop = mark_index
}

func md_mark_stack_pop(tls *libc.TLS, ctx uintptr, stack uintptr) (r int32) {
	var top int32
	_ = top
	top = (*MD_MARKSTACK)(unsafe.Pointer(stack)).Ftop
	if top >= 0 {
		(*MD_MARKSTACK)(unsafe.Pointer(stack)).Ftop = (*(*MD_MARK)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Fmarks + uintptr(top)*20))).Fnext
	}
	return top
}

// C documentation
//
//	/* Sometimes, we need to store a pointer into the mark. It is quite rare
//	 * so we do not bother to make MD_MARK use union, and it can only happen
//	 * for dummy marks. */
func md_mark_store_ptr(tls *libc.TLS, ctx uintptr, mark_index int32, _ptr uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*uintptr)(unsafe.Pointer(bp)) = _ptr
	var mark uintptr
	_ = mark
	mark = (*MD_CTX)(unsafe.Pointer(ctx)).Fmarks + uintptr(mark_index)*20
	if !(libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(mark)).Fch) == libc.Int32FromUint8('D')) {
		libc.X__builtin_unreachable(tls)
	}
	/* Check only members beg and end are misused for this. */
	if !(libc.Uint64FromInt64(8) <= libc.Uint64FromInt32(2)*libc.Uint64FromInt64(4)) {
		libc.X__builtin_unreachable(tls)
	}
	libc.Xmemcpy(tls, mark, bp, uint64(8))
}

func md_mark_get_ptr(tls *libc.TLS, ctx uintptr, mark_index int32) (r uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var mark uintptr
	var _ /* ptr at bp+0 */ uintptr
	_ = mark
	mark = (*MD_CTX)(unsafe.Pointer(ctx)).Fmarks + uintptr(mark_index)*20
	if !(libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(mark)).Fch) == libc.Int32FromUint8('D')) {
		libc.X__builtin_unreachable(tls)
	}
	libc.Xmemcpy(tls, bp, mark, uint64(8))
	return *(*uintptr)(__ccgo_ptr(bp))
}

func md_resolve_range(tls *libc.TLS, ctx uintptr, opener_index int32, closer_index int32) {
	var closer, opener, v1 uintptr
	_, _, _ = closer, opener, v1
	opener = (*MD_CTX)(unsafe.Pointer(ctx)).Fmarks + uintptr(opener_index)*20
	closer = (*MD_CTX)(unsafe.Pointer(ctx)).Fmarks + uintptr(closer_index)*20
	/* Interconnect opener and closer and mark both as resolved. */
	(*MD_MARK)(unsafe.Pointer(opener)).Fnext = closer_index
	(*MD_MARK)(unsafe.Pointer(closer)).Fprev = opener_index
	v1 = opener + 17
	*(*uint8)(unsafe.Pointer(v1)) = uint8(int32(*(*uint8)(unsafe.Pointer(v1))) | (libc.Int32FromInt32(MD_MARK_OPENER) | libc.Int32FromInt32(MD_MARK_RESOLVED)))
	v1 = closer + 17
	*(*uint8)(unsafe.Pointer(v1)) = uint8(int32(*(*uint8)(unsafe.Pointer(v1))) | (libc.Int32FromInt32(MD_MARK_CLOSER) | libc.Int32FromInt32(MD_MARK_RESOLVED)))
}

// C documentation
//
//	/* In the range ctx->marks[opener_index] ... [closer_index], undo some or all
//	 * resolvings accordingly to these rules:
//	 *
//	 * (1) All stacks of openers are cut so that any pending potential openers
//	 *     are discarded from future consideration.
//	 *
//	 * (2) If 'how' is MD_ROLLBACK_ALL, then ALL resolved marks inside the range
//	 *     are thrown away and turned into dummy marks ('D').
//	 *
//	 * WARNING: Do not call for arbitrary range of opener and closer.
//	 * This must form (potentially) valid range not crossing nesting boundaries
//	 * of already resolved ranges.
//	 */
func md_rollback(tls *libc.TLS, ctx uintptr, opener_index int32, closer_index int32, how int32) {
	var i int32
	var stack uintptr
	_, _ = i, stack
	i = 0
	for {
		if !(i < libc.Int32FromUint64(libc.Uint64FromInt64(64)/libc.Uint64FromInt64(4))) {
			break
		}
		stack = ctx + 416 + uintptr(i)*4
		for (*MD_MARKSTACK)(unsafe.Pointer(stack)).Ftop >= opener_index {
			md_mark_stack_pop(tls, ctx, stack)
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	if how == int32(MD_ROLLBACK_ALL) {
		i = opener_index + int32(1)
		for {
			if !(i < closer_index) {
				break
			}
			(*(*MD_MARK)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Fmarks + uintptr(i)*20))).Fch = uint8('D')
			(*(*MD_MARK)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Fmarks + uintptr(i)*20))).Fflags = uint8(0)
			goto _2
		_2:
			;
			i = i + 1
		}
	}
}

func md_build_mark_char_map(tls *libc.TLS, ctx uintptr) {
	var i int32
	_ = i
	libc.Xmemset(tls, ctx+160, 0, uint64(256))
	*(*uint8)(__ccgo_ptr(ctx + 160 + uintptr('\\'))) = uint8(1)
	*(*uint8)(__ccgo_ptr(ctx + 160 + uintptr('*'))) = uint8(1)
	*(*uint8)(__ccgo_ptr(ctx + 160 + uintptr('_'))) = uint8(1)
	*(*uint8)(__ccgo_ptr(ctx + 160 + uintptr('`'))) = uint8(1)
	*(*uint8)(__ccgo_ptr(ctx + 160 + uintptr('&'))) = uint8(1)
	*(*uint8)(__ccgo_ptr(ctx + 160 + uintptr(';'))) = uint8(1)
	*(*uint8)(__ccgo_ptr(ctx + 160 + uintptr('<'))) = uint8(1)
	*(*uint8)(__ccgo_ptr(ctx + 160 + uintptr('>'))) = uint8(1)
	*(*uint8)(__ccgo_ptr(ctx + 160 + uintptr('['))) = uint8(1)
	*(*uint8)(__ccgo_ptr(ctx + 160 + uintptr('!'))) = uint8(1)
	*(*uint8)(__ccgo_ptr(ctx + 160 + uintptr(']'))) = uint8(1)
	*(*uint8)(__ccgo_ptr(ctx + 160 + uintptr('\000'))) = uint8(1)
	if (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fflags&uint32(MD_FLAG_STRIKETHROUGH) != 0 {
		*(*uint8)(__ccgo_ptr(ctx + 160 + uintptr('~'))) = uint8(1)
	}
	if (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fflags&uint32(MD_FLAG_LATEXMATHSPANS) != 0 {
		*(*uint8)(__ccgo_ptr(ctx + 160 + uintptr('$'))) = uint8(1)
	}
	if (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fflags&uint32(MD_FLAG_PERMISSIVEEMAILAUTOLINKS) != 0 {
		*(*uint8)(__ccgo_ptr(ctx + 160 + uintptr('@'))) = uint8(1)
	}
	if (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fflags&uint32(MD_FLAG_PERMISSIVEURLAUTOLINKS) != 0 {
		*(*uint8)(__ccgo_ptr(ctx + 160 + uintptr(':'))) = uint8(1)
	}
	if (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fflags&uint32(MD_FLAG_PERMISSIVEWWWAUTOLINKS) != 0 {
		*(*uint8)(__ccgo_ptr(ctx + 160 + uintptr('.'))) = uint8(1)
	}
	if (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fflags&uint32(MD_FLAG_TABLES) != 0 || (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fflags&uint32(MD_FLAG_WIKILINKS) != 0 {
		*(*uint8)(__ccgo_ptr(ctx + 160 + uintptr('|'))) = uint8(1)
	}
	if (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fflags&uint32(MD_FLAG_COLLAPSEWHITESPACE) != 0 {
		i = 0
		for {
			if !(i < libc.Int32FromInt64(256)) {
				break
			}
			if i == int32(' ') || i == int32('\t') || (i == int32('\v') || i == int32('\f')) {
				*(*uint8)(__ccgo_ptr(ctx + 160 + uintptr(i))) = uint8(1)
			}
			goto _1
		_1:
			;
			i = i + 1
		}
	}
}

func md_is_code_span(tls *libc.TLS, ctx uintptr, lines uintptr, n_lines MD_SIZE, beg MD_OFFSET, opener uintptr, closer uintptr, last_potential_closers uintptr, p_reached_paragraph_end uintptr) (r int32) {
	var closer_beg, closer_end, line_end, opener_beg, opener_end MD_OFFSET
	var has_eol_after_opener, has_eol_before_closer, has_only_space, has_space_after_opener, has_space_before_closer int32
	var line_index, mark_len MD_SIZE
	_, _, _, _, _, _, _, _, _, _, _, _ = closer_beg, closer_end, has_eol_after_opener, has_eol_before_closer, has_only_space, has_space_after_opener, has_space_before_closer, line_end, line_index, mark_len, opener_beg, opener_end
	opener_beg = beg
	has_space_after_opener = FALSE
	has_eol_after_opener = FALSE
	has_space_before_closer = FALSE
	has_eol_before_closer = FALSE
	has_only_space = int32(TRUE)
	line_index = uint32(0)
	line_end = (*(*MD_LINE)(__ccgo_ptr(lines))).Fend
	opener_end = opener_beg
	for opener_end < line_end && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(opener_end)))) == int32('`') {
		opener_end = opener_end + 1
	}
	has_space_after_opener = libc.BoolInt32(opener_end < line_end && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(opener_end)))) == int32(' '))
	has_eol_after_opener = libc.BoolInt32(opener_end == line_end)
	/* The caller needs to know end of the opening mark even if we fail. */
	(*MD_MARK)(unsafe.Pointer(opener)).Fend = opener_end
	mark_len = opener_end - opener_beg
	if mark_len > uint32(CODESPAN_MARK_MAXLEN) {
		return FALSE
	}
	/* Check whether we already know there is no closer of this length.
	 * If so, re-scan does no sense. This fixes issue #59. */
	if *(*MD_OFFSET)(__ccgo_ptr(last_potential_closers + uintptr(mark_len-uint32(1))*4)) >= (*(*MD_LINE)(__ccgo_ptr(lines + uintptr(n_lines-uint32(1))*8))).Fend || *(*int32)(__ccgo_ptr(p_reached_paragraph_end)) != 0 && *(*MD_OFFSET)(__ccgo_ptr(last_potential_closers + uintptr(mark_len-uint32(1))*4)) < opener_end {
		return FALSE
	}
	closer_beg = opener_end
	closer_end = opener_end
	/* Find closer mark. */
	for int32(TRUE) != 0 {
		for closer_beg < line_end && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(closer_beg)))) != int32('`') {
			if libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(closer_beg)))) != int32(' ') {
				has_only_space = FALSE
			}
			closer_beg = closer_beg + 1
		}
		closer_end = closer_beg
		for closer_end < line_end && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(closer_end)))) == int32('`') {
			closer_end = closer_end + 1
		}
		if closer_end-closer_beg == mark_len {
			/* Success. */
			has_space_before_closer = libc.BoolInt32(closer_beg > (*(*MD_LINE)(__ccgo_ptr(lines + uintptr(line_index)*8))).Fbeg && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(closer_beg-libc.Uint32FromInt32(1))))) == int32(' '))
			has_eol_before_closer = libc.BoolInt32(closer_beg == (*(*MD_LINE)(__ccgo_ptr(lines + uintptr(line_index)*8))).Fbeg)
			break
		}
		if closer_end-closer_beg > uint32(0) {
			/* We have found a back-tick which is not part of the closer. */
			has_only_space = FALSE
			/* But if we eventually fail, remember it as a potential closer
			 * of its own length for future attempts. This mitigates needs for
			 * rescans. */
			if closer_end-closer_beg < uint32(CODESPAN_MARK_MAXLEN) {
				if closer_beg > *(*MD_OFFSET)(__ccgo_ptr(last_potential_closers + uintptr(closer_end-closer_beg-uint32(1))*4)) {
					*(*MD_OFFSET)(__ccgo_ptr(last_potential_closers + uintptr(closer_end-closer_beg-uint32(1))*4)) = closer_beg
				}
			}
		}
		if closer_end >= line_end {
			line_index = line_index + 1
			if line_index >= n_lines {
				/* Reached end of the paragraph and still nothing. */
				*(*int32)(__ccgo_ptr(p_reached_paragraph_end)) = int32(TRUE)
				return FALSE
			}
			/* Try on the next line. */
			line_end = (*(*MD_LINE)(__ccgo_ptr(lines + uintptr(line_index)*8))).Fend
			closer_beg = (*(*MD_LINE)(__ccgo_ptr(lines + uintptr(line_index)*8))).Fbeg
		} else {
			closer_beg = closer_end
		}
	}
	/* If there is a space or a new line both after and before the opener
	 * (and if the code span is not made of spaces only), consume one initial
	 * and one trailing space as part of the marks. */
	if !(has_only_space != 0) && (has_space_after_opener != 0 || has_eol_after_opener != 0) && (has_space_before_closer != 0 || has_eol_before_closer != 0) {
		if has_space_after_opener != 0 {
			opener_end = opener_end + 1
		} else {
			opener_end = (*(*MD_LINE)(__ccgo_ptr(lines + 1*8))).Fbeg
		}
		if has_space_before_closer != 0 {
			closer_beg = closer_beg - 1
		} else {
			/* Go back to the end of prev line */
			closer_beg = (*(*MD_LINE)(__ccgo_ptr(lines + uintptr(line_index-uint32(1))*8))).Fend
			/* But restore any trailing whitespace */
			for closer_beg < (*MD_CTX)(unsafe.Pointer(ctx)).Fsize && (libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(closer_beg)))) == int32(' ') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(closer_beg)))) == int32('\t')) {
				closer_beg = closer_beg + 1
			}
		}
	}
	(*MD_MARK)(unsafe.Pointer(opener)).Fch = uint8('`')
	(*MD_MARK)(unsafe.Pointer(opener)).Fbeg = opener_beg
	(*MD_MARK)(unsafe.Pointer(opener)).Fend = opener_end
	(*MD_MARK)(unsafe.Pointer(opener)).Fflags = uint8(MD_MARK_POTENTIAL_OPENER)
	(*MD_MARK)(unsafe.Pointer(closer)).Fch = uint8('`')
	(*MD_MARK)(unsafe.Pointer(closer)).Fbeg = closer_beg
	(*MD_MARK)(unsafe.Pointer(closer)).Fend = closer_end
	(*MD_MARK)(unsafe.Pointer(closer)).Fflags = uint8(MD_MARK_POTENTIAL_CLOSER)
	return int32(TRUE)
}

func md_is_autolink_uri(tls *libc.TLS, ctx uintptr, beg MD_OFFSET, max_end MD_OFFSET, p_end uintptr) (r int32) {
	var off MD_OFFSET
	_ = off
	off = beg + uint32(1)
	if !(libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(beg)))) == libc.Int32FromUint8('<')) {
		libc.X__builtin_unreachable(tls)
	}
	/* Check for scheme. */
	if off >= max_end || !(uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) <= libc.Uint32FromInt32(127)) {
		return FALSE
	}
	off = off + 1
	for int32(1) != 0 {
		if off >= max_end {
			return FALSE
		}
		if off-beg > uint32(32) {
			return FALSE
		}
		if libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32(':') && off-beg >= uint32(3) {
			break
		}
		if !(libc.Uint32FromInt32(libc.Int32FromUint8('A')) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) <= libc.Uint32FromInt32(libc.Int32FromUint8('Z')) || libc.Uint32FromInt32(libc.Int32FromUint8('a')) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) <= libc.Uint32FromInt32(libc.Int32FromUint8('z')) || libc.Uint32FromInt32(libc.Int32FromUint8('0')) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) <= libc.Uint32FromInt32(libc.Int32FromUint8('9'))) && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) != int32('+') && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) != int32('-') && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) != int32('.') {
			return FALSE
		}
		off = off + 1
	}
	/* Check the path after the scheme. */
	for off < max_end && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) != int32('>') {
		if libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32(' ') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('\t') || (libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('\v') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('\f')) || (uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) <= uint32(31) || uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == uint32(127)) || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('<') {
			return FALSE
		}
		off = off + 1
	}
	if off >= max_end {
		return FALSE
	}
	if !(libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == libc.Int32FromUint8('>')) {
		libc.X__builtin_unreachable(tls)
	}
	*(*MD_OFFSET)(__ccgo_ptr(p_end)) = off + uint32(1)
	return int32(TRUE)
}

func md_is_autolink_email(tls *libc.TLS, ctx uintptr, beg MD_OFFSET, max_end MD_OFFSET, p_end uintptr) (r int32) {
	var label_len int32
	var off MD_OFFSET
	_, _ = label_len, off
	off = beg + uint32(1)
	if !(libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(beg)))) == libc.Int32FromUint8('<')) {
		libc.X__builtin_unreachable(tls)
	}
	/* The code should correspond to this regexp:
	   /^[a-zA-Z0-9.!#$%&'*+\/=?^_`{|}~-]+
	   @[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?
	   (?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$/
	*/
	/* Username (before '@'). */
	for off < max_end && (libc.Uint32FromInt32(libc.Int32FromUint8('A')) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) <= libc.Uint32FromInt32(libc.Int32FromUint8('Z')) || libc.Uint32FromInt32(libc.Int32FromUint8('a')) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) <= libc.Uint32FromInt32(libc.Int32FromUint8('z')) || libc.Uint32FromInt32(libc.Int32FromUint8('0')) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) <= libc.Uint32FromInt32(libc.Int32FromUint8('9')) || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) != int32('\000') && libc.Xstrchr(tls, __ccgo_ts+61, libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off))))) != libc.UintptrFromInt32(0)) {
		off = off + 1
	}
	if off <= beg+uint32(1) {
		return FALSE
	}
	/* '@' */
	if off >= max_end || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) != int32('@') {
		return FALSE
	}
	off = off + 1
	/* Labels delimited with '.'; each label is sequence of 1 - 63 alnum
	 * characters or '-', but '-' is not allowed as first or last char. */
	label_len = 0
	for off < max_end {
		if libc.Uint32FromInt32(libc.Int32FromUint8('A')) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) <= libc.Uint32FromInt32(libc.Int32FromUint8('Z')) || libc.Uint32FromInt32(libc.Int32FromUint8('a')) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) <= libc.Uint32FromInt32(libc.Int32FromUint8('z')) || libc.Uint32FromInt32(libc.Int32FromUint8('0')) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) <= libc.Uint32FromInt32(libc.Int32FromUint8('9')) {
			label_len = label_len + 1
		} else {
			if libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('-') && label_len > 0 {
				label_len = label_len + 1
			} else {
				if libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('.') && label_len > 0 && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off-libc.Uint32FromInt32(1))))) != int32('-') {
					label_len = 0
				} else {
					break
				}
			}
		}
		if label_len > int32(63) {
			return FALSE
		}
		off = off + 1
	}
	if label_len <= 0 || off >= max_end || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) != int32('>') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off-libc.Uint32FromInt32(1))))) == int32('-') {
		return FALSE
	}
	*(*MD_OFFSET)(__ccgo_ptr(p_end)) = off + uint32(1)
	return int32(TRUE)
}

func md_is_autolink(tls *libc.TLS, ctx uintptr, beg MD_OFFSET, max_end MD_OFFSET, p_end uintptr, p_missing_mailto uintptr) (r int32) {
	if md_is_autolink_uri(tls, ctx, beg, max_end, p_end) != 0 {
		*(*int32)(__ccgo_ptr(p_missing_mailto)) = FALSE
		return int32(TRUE)
	}
	if md_is_autolink_email(tls, ctx, beg, max_end, p_end) != 0 {
		*(*int32)(__ccgo_ptr(p_missing_mailto)) = int32(TRUE)
		return int32(TRUE)
	}
	return FALSE
}

func md_collect_marks(tls *libc.TLS, ctx uintptr, lines uintptr, n_lines MD_SIZE, table_mode int32) (r int32) {
	bp := tls.Alloc(192)
	defer tls.Free(192)
	var ch MD_CHAR
	var flags, flags1, flags2, flags3, v2 uint32
	var is_autolink, is_code_span, is_html, left_level, ret, right_level, scheme_index int32
	var line, mark, scheme, suffix uintptr
	var off, tmp, tmp1, tmp2, tmp3, tmp4 MD_OFFSET
	var scheme_size, suffix_size MD_SIZE
	var _ /* autolink_end at bp+176 */ MD_OFFSET
	var _ /* closer at bp+156 */ MD_MARK
	var _ /* codespan_last_potential_closers at bp+4 */ [32]MD_OFFSET
	var _ /* codespan_scanned_till_paragraph_end at bp+132 */ int32
	var _ /* html_end at bp+184 */ MD_OFFSET
	var _ /* line_index at bp+0 */ MD_SIZE
	var _ /* missing_mailto at bp+180 */ int32
	var _ /* opener at bp+136 */ MD_MARK
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = ch, flags, flags1, flags2, flags3, is_autolink, is_code_span, is_html, left_level, line, mark, off, ret, right_level, scheme, scheme_index, scheme_size, suffix, suffix_size, tmp, tmp1, tmp2, tmp3, tmp4, v2
	ret = 0
	*(*[32]MD_OFFSET)(__ccgo_ptr(bp + 4)) = [32]MD_OFFSET{}
	*(*int32)(__ccgo_ptr(bp + 132)) = FALSE
	*(*MD_SIZE)(__ccgo_ptr(bp)) = uint32(0)
	for {
		if !(*(*MD_SIZE)(__ccgo_ptr(bp)) < n_lines) {
			break
		}
		line = lines + uintptr(*(*MD_SIZE)(__ccgo_ptr(bp)))*8
		off = (*MD_LINE)(unsafe.Pointer(line)).Fbeg
		for int32(TRUE) != 0 {
			/* For 8-bit encodings, mark_char_map[] covers all 256 elements. */
			/* Optimization: Use some loop unrolling. */
			for off+uint32(3) < (*MD_LINE)(unsafe.Pointer(line)).Fend && !(*(*uint8)(__ccgo_ptr(ctx + 160 + uintptr(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(0))))))) != 0) && !(*(*uint8)(__ccgo_ptr(ctx + 160 + uintptr(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))))) != 0) && !(*(*uint8)(__ccgo_ptr(ctx + 160 + uintptr(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(2))))))) != 0) && !(*(*uint8)(__ccgo_ptr(ctx + 160 + uintptr(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(3))))))) != 0) {
				off = off + uint32(4)
			}
			for off < (*MD_LINE)(unsafe.Pointer(line)).Fend && !(*(*uint8)(__ccgo_ptr(ctx + 160 + uintptr(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(0))))))) != 0) {
				off = off + 1
			}
			if off >= (*MD_LINE)(unsafe.Pointer(line)).Fend {
				break
			}
			ch = *(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))
			/* A backslash escape.
			 * It can go beyond line->end as it may involve escaped new
			 * line to form a hard break. */
			if libc.Int32FromUint8(ch) == int32('\\') && off+uint32(1) < (*MD_CTX)(unsafe.Pointer(ctx)).Fsize && (libc.Uint32FromInt32(libc.Int32FromInt32(33)) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))) <= libc.Uint32FromInt32(libc.Int32FromInt32(47)) || libc.Uint32FromInt32(libc.Int32FromInt32(58)) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))) <= libc.Uint32FromInt32(libc.Int32FromInt32(64)) || libc.Uint32FromInt32(libc.Int32FromInt32(91)) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))) <= libc.Uint32FromInt32(libc.Int32FromInt32(96)) || libc.Uint32FromInt32(libc.Int32FromInt32(123)) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))) <= libc.Uint32FromInt32(libc.Int32FromInt32(126)) || (libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))) == int32('\r') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))) == int32('\n'))) {
				/* Hard-break cannot be on the last line of the block. */
				if !(libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))) == int32('\r') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))) == int32('\n')) || *(*MD_SIZE)(__ccgo_ptr(bp))+uint32(1) < n_lines {
					mark = md_add_mark(tls, ctx)
					if mark == libc.UintptrFromInt32(0) {
						ret = -int32(1)
						goto abort
					}
					(*MD_MARK)(unsafe.Pointer(mark)).Fbeg = off
					(*MD_MARK)(unsafe.Pointer(mark)).Fend = off + libc.Uint32FromInt32(2)
					(*MD_MARK)(unsafe.Pointer(mark)).Fprev = -int32(1)
					(*MD_MARK)(unsafe.Pointer(mark)).Fnext = -int32(1)
					(*MD_MARK)(unsafe.Pointer(mark)).Fch = ch
					(*MD_MARK)(unsafe.Pointer(mark)).Fflags = libc.Uint8FromInt32(libc.Int32FromInt32(MD_MARK_RESOLVED))
				}
				off = off + uint32(2)
				continue
			}
			/* A potential (string) emphasis start/end. */
			if libc.Int32FromUint8(ch) == int32('*') || libc.Int32FromUint8(ch) == int32('_') {
				tmp = off + uint32(1) /* What follows: 0 = whitespace; 1 = punctuation; 2 = other char. */
				for tmp < (*MD_LINE)(unsafe.Pointer(line)).Fend && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(tmp)))) == libc.Int32FromUint8(ch) {
					tmp = tmp + 1
				}
				if off == (*MD_LINE)(unsafe.Pointer(line)).Fbeg || md_is_unicode_whitespace__(tls, md_decode_utf8_before__(tls, ctx, off)) != 0 {
					left_level = 0
				} else {
					if md_is_unicode_punct__(tls, md_decode_utf8_before__(tls, ctx, off)) != 0 {
						left_level = int32(1)
					} else {
						left_level = int32(2)
					}
				}
				if tmp == (*MD_LINE)(unsafe.Pointer(line)).Fend || md_is_unicode_whitespace__(tls, md_decode_utf8__(tls, (*MD_CTX)(unsafe.Pointer(ctx)).Ftext+uintptr(tmp), (*MD_CTX)(unsafe.Pointer(ctx)).Fsize-tmp, libc.UintptrFromInt32(0))) != 0 {
					right_level = 0
				} else {
					if md_is_unicode_punct__(tls, md_decode_utf8__(tls, (*MD_CTX)(unsafe.Pointer(ctx)).Ftext+uintptr(tmp), (*MD_CTX)(unsafe.Pointer(ctx)).Fsize-tmp, libc.UintptrFromInt32(0))) != 0 {
						right_level = int32(1)
					} else {
						right_level = int32(2)
					}
				}
				/* Intra-word underscore doesn't have special meaning. */
				if libc.Int32FromUint8(ch) == int32('_') && left_level == int32(2) && right_level == int32(2) {
					left_level = 0
					right_level = 0
				}
				if left_level != 0 || right_level != 0 {
					flags = uint32(0)
					if left_level > 0 && left_level >= right_level {
						flags = flags | uint32(MD_MARK_POTENTIAL_CLOSER)
					}
					if right_level > 0 && right_level >= left_level {
						flags = flags | uint32(MD_MARK_POTENTIAL_OPENER)
					}
					if flags == libc.Uint32FromInt32(libc.Int32FromInt32(MD_MARK_POTENTIAL_OPENER)|libc.Int32FromInt32(MD_MARK_POTENTIAL_CLOSER)) {
						flags = flags | uint32(MD_MARK_EMPH_OC)
					}
					/* For "the rule of three" we need to remember the original
					 * size of the mark (modulo three), before we potentially
					 * split the mark when being later resolved partially by some
					 * shorter closer. */
					switch (tmp - off) % libc.Uint32FromInt32(3) {
					case uint32(0):
						flags = flags | uint32(MD_MARK_EMPH_MOD3_0)
					case uint32(1):
						flags = flags | uint32(MD_MARK_EMPH_MOD3_1)
					case uint32(2):
						flags = flags | libc.Uint32FromInt32(libc.Int32FromInt32(0x40)|libc.Int32FromInt32(0x80))
						break
					}
					mark = md_add_mark(tls, ctx)
					if mark == libc.UintptrFromInt32(0) {
						ret = -int32(1)
						goto abort
					}
					(*MD_MARK)(unsafe.Pointer(mark)).Fbeg = off
					(*MD_MARK)(unsafe.Pointer(mark)).Fend = tmp
					(*MD_MARK)(unsafe.Pointer(mark)).Fprev = -int32(1)
					(*MD_MARK)(unsafe.Pointer(mark)).Fnext = -int32(1)
					(*MD_MARK)(unsafe.Pointer(mark)).Fch = ch
					(*MD_MARK)(unsafe.Pointer(mark)).Fflags = uint8(flags)
					/* During resolving, multiple asterisks may have to be
					 * split into independent span start/ends. Consider e.g.
					 * "**foo* bar*". Therefore we push also some empty dummy
					 * marks to have enough space for that. */
					off = off + 1
					for off < tmp {
						mark = md_add_mark(tls, ctx)
						if mark == libc.UintptrFromInt32(0) {
							ret = -int32(1)
							goto abort
						}
						(*MD_MARK)(unsafe.Pointer(mark)).Fbeg = off
						(*MD_MARK)(unsafe.Pointer(mark)).Fend = off
						(*MD_MARK)(unsafe.Pointer(mark)).Fprev = -int32(1)
						(*MD_MARK)(unsafe.Pointer(mark)).Fnext = -int32(1)
						(*MD_MARK)(unsafe.Pointer(mark)).Fch = libc.Uint8FromInt32(libc.Int32FromUint8('D'))
						(*MD_MARK)(unsafe.Pointer(mark)).Fflags = libc.Uint8FromInt32(libc.Int32FromInt32(0))
						off = off + 1
					}
					continue
				}
				off = tmp
				continue
			}
			/* A potential code span start/end. */
			if libc.Int32FromUint8(ch) == int32('`') {
				is_code_span = md_is_code_span(tls, ctx, line, n_lines-*(*MD_SIZE)(__ccgo_ptr(bp)), off, bp+136, bp+156, bp+4, bp+132)
				if is_code_span != 0 {
					mark = md_add_mark(tls, ctx)
					if mark == libc.UintptrFromInt32(0) {
						ret = -int32(1)
						goto abort
					}
					(*MD_MARK)(unsafe.Pointer(mark)).Fbeg = (*(*MD_MARK)(__ccgo_ptr(bp + 136))).Fbeg
					(*MD_MARK)(unsafe.Pointer(mark)).Fend = (*(*MD_MARK)(__ccgo_ptr(bp + 136))).Fend
					(*MD_MARK)(unsafe.Pointer(mark)).Fprev = -int32(1)
					(*MD_MARK)(unsafe.Pointer(mark)).Fnext = -int32(1)
					(*MD_MARK)(unsafe.Pointer(mark)).Fch = (*(*MD_MARK)(__ccgo_ptr(bp + 136))).Fch
					(*MD_MARK)(unsafe.Pointer(mark)).Fflags = (*(*MD_MARK)(__ccgo_ptr(bp + 136))).Fflags
					mark = md_add_mark(tls, ctx)
					if mark == libc.UintptrFromInt32(0) {
						ret = -int32(1)
						goto abort
					}
					(*MD_MARK)(unsafe.Pointer(mark)).Fbeg = (*(*MD_MARK)(__ccgo_ptr(bp + 156))).Fbeg
					(*MD_MARK)(unsafe.Pointer(mark)).Fend = (*(*MD_MARK)(__ccgo_ptr(bp + 156))).Fend
					(*MD_MARK)(unsafe.Pointer(mark)).Fprev = -int32(1)
					(*MD_MARK)(unsafe.Pointer(mark)).Fnext = -int32(1)
					(*MD_MARK)(unsafe.Pointer(mark)).Fch = (*(*MD_MARK)(__ccgo_ptr(bp + 156))).Fch
					(*MD_MARK)(unsafe.Pointer(mark)).Fflags = (*(*MD_MARK)(__ccgo_ptr(bp + 156))).Fflags
					md_resolve_range(tls, ctx, (*MD_CTX)(unsafe.Pointer(ctx)).Fn_marks-int32(2), (*MD_CTX)(unsafe.Pointer(ctx)).Fn_marks-int32(1))
					off = (*(*MD_MARK)(__ccgo_ptr(bp + 156))).Fend
					/* Advance the current line accordingly. */
					if off > (*MD_LINE)(unsafe.Pointer(line)).Fend {
						line = md_lookup_line(tls, off, lines, n_lines, bp)
					}
					continue
				}
				off = (*(*MD_MARK)(__ccgo_ptr(bp + 136))).Fend
				continue
			}
			/* A potential entity start. */
			if libc.Int32FromUint8(ch) == int32('&') {
				mark = md_add_mark(tls, ctx)
				if mark == libc.UintptrFromInt32(0) {
					ret = -int32(1)
					goto abort
				}
				(*MD_MARK)(unsafe.Pointer(mark)).Fbeg = off
				(*MD_MARK)(unsafe.Pointer(mark)).Fend = off + libc.Uint32FromInt32(1)
				(*MD_MARK)(unsafe.Pointer(mark)).Fprev = -int32(1)
				(*MD_MARK)(unsafe.Pointer(mark)).Fnext = -int32(1)
				(*MD_MARK)(unsafe.Pointer(mark)).Fch = ch
				(*MD_MARK)(unsafe.Pointer(mark)).Fflags = libc.Uint8FromInt32(libc.Int32FromInt32(MD_MARK_POTENTIAL_OPENER))
				off = off + 1
				continue
			}
			/* A potential entity end. */
			if libc.Int32FromUint8(ch) == int32(';') {
				/* We surely cannot be entity unless the previous mark is '&'. */
				if (*MD_CTX)(unsafe.Pointer(ctx)).Fn_marks > 0 && libc.Int32FromUint8((*(*MD_MARK)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Fmarks + uintptr((*MD_CTX)(unsafe.Pointer(ctx)).Fn_marks-int32(1))*20))).Fch) == int32('&') {
					mark = md_add_mark(tls, ctx)
					if mark == libc.UintptrFromInt32(0) {
						ret = -int32(1)
						goto abort
					}
					(*MD_MARK)(unsafe.Pointer(mark)).Fbeg = off
					(*MD_MARK)(unsafe.Pointer(mark)).Fend = off + libc.Uint32FromInt32(1)
					(*MD_MARK)(unsafe.Pointer(mark)).Fprev = -int32(1)
					(*MD_MARK)(unsafe.Pointer(mark)).Fnext = -int32(1)
					(*MD_MARK)(unsafe.Pointer(mark)).Fch = ch
					(*MD_MARK)(unsafe.Pointer(mark)).Fflags = libc.Uint8FromInt32(libc.Int32FromInt32(MD_MARK_POTENTIAL_CLOSER))
				}
				off = off + 1
				continue
			}
			/* A potential autolink or raw HTML start/end. */
			if libc.Int32FromUint8(ch) == int32('<') {
				if !((*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fflags&libc.Uint32FromInt32(MD_FLAG_NOHTMLSPANS) != 0) {
					/* Given the nature of the raw HTML, we have to recognize
					 * it here. Doing so later in md_analyze_lt_gt() could
					 * open can of worms of quadratic complexity. */
					is_html = md_is_html_any(tls, ctx, line, n_lines-*(*MD_SIZE)(__ccgo_ptr(bp)), off, (*(*MD_LINE)(__ccgo_ptr(lines + uintptr(n_lines-uint32(1))*8))).Fend, bp+184)
					if is_html != 0 {
						mark = md_add_mark(tls, ctx)
						if mark == libc.UintptrFromInt32(0) {
							ret = -int32(1)
							goto abort
						}
						(*MD_MARK)(unsafe.Pointer(mark)).Fbeg = off
						(*MD_MARK)(unsafe.Pointer(mark)).Fend = off
						(*MD_MARK)(unsafe.Pointer(mark)).Fprev = -int32(1)
						(*MD_MARK)(unsafe.Pointer(mark)).Fnext = -int32(1)
						(*MD_MARK)(unsafe.Pointer(mark)).Fch = libc.Uint8FromInt32(libc.Int32FromUint8('<'))
						(*MD_MARK)(unsafe.Pointer(mark)).Fflags = libc.Uint8FromInt32(libc.Int32FromInt32(MD_MARK_OPENER) | libc.Int32FromInt32(MD_MARK_RESOLVED))
						mark = md_add_mark(tls, ctx)
						if mark == libc.UintptrFromInt32(0) {
							ret = -int32(1)
							goto abort
						}
						(*MD_MARK)(unsafe.Pointer(mark)).Fbeg = *(*MD_OFFSET)(__ccgo_ptr(bp + 184))
						(*MD_MARK)(unsafe.Pointer(mark)).Fend = *(*MD_OFFSET)(__ccgo_ptr(bp + 184))
						(*MD_MARK)(unsafe.Pointer(mark)).Fprev = -int32(1)
						(*MD_MARK)(unsafe.Pointer(mark)).Fnext = -int32(1)
						(*MD_MARK)(unsafe.Pointer(mark)).Fch = libc.Uint8FromInt32(libc.Int32FromUint8('>'))
						(*MD_MARK)(unsafe.Pointer(mark)).Fflags = libc.Uint8FromInt32(libc.Int32FromInt32(MD_MARK_CLOSER) | libc.Int32FromInt32(MD_MARK_RESOLVED))
						(*(*MD_MARK)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Fmarks + uintptr((*MD_CTX)(unsafe.Pointer(ctx)).Fn_marks-int32(2))*20))).Fnext = (*MD_CTX)(unsafe.Pointer(ctx)).Fn_marks - int32(1)
						(*(*MD_MARK)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Fmarks + uintptr((*MD_CTX)(unsafe.Pointer(ctx)).Fn_marks-int32(1))*20))).Fprev = (*MD_CTX)(unsafe.Pointer(ctx)).Fn_marks - int32(2)
						off = *(*MD_OFFSET)(__ccgo_ptr(bp + 184))
						/* Advance the current line accordingly. */
						if off > (*MD_LINE)(unsafe.Pointer(line)).Fend {
							line = md_lookup_line(tls, off, lines, n_lines, bp)
						}
						continue
					}
				}
				is_autolink = md_is_autolink(tls, ctx, off, (*(*MD_LINE)(__ccgo_ptr(lines + uintptr(n_lines-uint32(1))*8))).Fend, bp+176, bp+180)
				if is_autolink != 0 {
					flags1 = libc.Uint32FromInt32(libc.Int32FromInt32(MD_MARK_RESOLVED) | libc.Int32FromInt32(MD_MARK_AUTOLINK))
					if *(*int32)(__ccgo_ptr(bp + 180)) != 0 {
						flags1 = flags1 | uint32(MD_MARK_AUTOLINK_MISSING_MAILTO)
					}
					mark = md_add_mark(tls, ctx)
					if mark == libc.UintptrFromInt32(0) {
						ret = -int32(1)
						goto abort
					}
					(*MD_MARK)(unsafe.Pointer(mark)).Fbeg = off
					(*MD_MARK)(unsafe.Pointer(mark)).Fend = off + libc.Uint32FromInt32(1)
					(*MD_MARK)(unsafe.Pointer(mark)).Fprev = -int32(1)
					(*MD_MARK)(unsafe.Pointer(mark)).Fnext = -int32(1)
					(*MD_MARK)(unsafe.Pointer(mark)).Fch = libc.Uint8FromInt32(libc.Int32FromUint8('<'))
					(*MD_MARK)(unsafe.Pointer(mark)).Fflags = uint8(libc.Uint32FromInt32(MD_MARK_OPENER) | flags1)
					mark = md_add_mark(tls, ctx)
					if mark == libc.UintptrFromInt32(0) {
						ret = -int32(1)
						goto abort
					}
					(*MD_MARK)(unsafe.Pointer(mark)).Fbeg = *(*MD_OFFSET)(__ccgo_ptr(bp + 176)) - libc.Uint32FromInt32(1)
					(*MD_MARK)(unsafe.Pointer(mark)).Fend = *(*MD_OFFSET)(__ccgo_ptr(bp + 176))
					(*MD_MARK)(unsafe.Pointer(mark)).Fprev = -int32(1)
					(*MD_MARK)(unsafe.Pointer(mark)).Fnext = -int32(1)
					(*MD_MARK)(unsafe.Pointer(mark)).Fch = libc.Uint8FromInt32(libc.Int32FromUint8('>'))
					(*MD_MARK)(unsafe.Pointer(mark)).Fflags = uint8(libc.Uint32FromInt32(MD_MARK_CLOSER) | flags1)
					(*(*MD_MARK)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Fmarks + uintptr((*MD_CTX)(unsafe.Pointer(ctx)).Fn_marks-int32(2))*20))).Fnext = (*MD_CTX)(unsafe.Pointer(ctx)).Fn_marks - int32(1)
					(*(*MD_MARK)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Fmarks + uintptr((*MD_CTX)(unsafe.Pointer(ctx)).Fn_marks-int32(1))*20))).Fprev = (*MD_CTX)(unsafe.Pointer(ctx)).Fn_marks - int32(2)
					off = *(*MD_OFFSET)(__ccgo_ptr(bp + 176))
					continue
				}
				off = off + 1
				continue
			}
			/* A potential link or its part. */
			if libc.Int32FromUint8(ch) == int32('[') || libc.Int32FromUint8(ch) == int32('!') && off+uint32(1) < (*MD_LINE)(unsafe.Pointer(line)).Fend && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))) == int32('[') {
				if libc.Int32FromUint8(ch) == int32('[') {
					v2 = off + uint32(1)
				} else {
					v2 = off + uint32(2)
				}
				tmp1 = v2
				mark = md_add_mark(tls, ctx)
				if mark == libc.UintptrFromInt32(0) {
					ret = -int32(1)
					goto abort
				}
				(*MD_MARK)(unsafe.Pointer(mark)).Fbeg = off
				(*MD_MARK)(unsafe.Pointer(mark)).Fend = tmp1
				(*MD_MARK)(unsafe.Pointer(mark)).Fprev = -int32(1)
				(*MD_MARK)(unsafe.Pointer(mark)).Fnext = -int32(1)
				(*MD_MARK)(unsafe.Pointer(mark)).Fch = ch
				(*MD_MARK)(unsafe.Pointer(mark)).Fflags = libc.Uint8FromInt32(libc.Int32FromInt32(MD_MARK_POTENTIAL_OPENER))
				off = tmp1
				/* Two dummies to make enough place for data we need if it is
				 * a link. */
				mark = md_add_mark(tls, ctx)
				if mark == libc.UintptrFromInt32(0) {
					ret = -int32(1)
					goto abort
				}
				(*MD_MARK)(unsafe.Pointer(mark)).Fbeg = off
				(*MD_MARK)(unsafe.Pointer(mark)).Fend = off
				(*MD_MARK)(unsafe.Pointer(mark)).Fprev = -int32(1)
				(*MD_MARK)(unsafe.Pointer(mark)).Fnext = -int32(1)
				(*MD_MARK)(unsafe.Pointer(mark)).Fch = libc.Uint8FromInt32(libc.Int32FromUint8('D'))
				(*MD_MARK)(unsafe.Pointer(mark)).Fflags = libc.Uint8FromInt32(libc.Int32FromInt32(0))
				mark = md_add_mark(tls, ctx)
				if mark == libc.UintptrFromInt32(0) {
					ret = -int32(1)
					goto abort
				}
				(*MD_MARK)(unsafe.Pointer(mark)).Fbeg = off
				(*MD_MARK)(unsafe.Pointer(mark)).Fend = off
				(*MD_MARK)(unsafe.Pointer(mark)).Fprev = -int32(1)
				(*MD_MARK)(unsafe.Pointer(mark)).Fnext = -int32(1)
				(*MD_MARK)(unsafe.Pointer(mark)).Fch = libc.Uint8FromInt32(libc.Int32FromUint8('D'))
				(*MD_MARK)(unsafe.Pointer(mark)).Fflags = libc.Uint8FromInt32(libc.Int32FromInt32(0))
				continue
			}
			if libc.Int32FromUint8(ch) == int32(']') {
				mark = md_add_mark(tls, ctx)
				if mark == libc.UintptrFromInt32(0) {
					ret = -int32(1)
					goto abort
				}
				(*MD_MARK)(unsafe.Pointer(mark)).Fbeg = off
				(*MD_MARK)(unsafe.Pointer(mark)).Fend = off + libc.Uint32FromInt32(1)
				(*MD_MARK)(unsafe.Pointer(mark)).Fprev = -int32(1)
				(*MD_MARK)(unsafe.Pointer(mark)).Fnext = -int32(1)
				(*MD_MARK)(unsafe.Pointer(mark)).Fch = ch
				(*MD_MARK)(unsafe.Pointer(mark)).Fflags = libc.Uint8FromInt32(libc.Int32FromInt32(MD_MARK_POTENTIAL_CLOSER))
				off = off + 1
				continue
			}
			/* A potential permissive e-mail autolink. */
			if libc.Int32FromUint8(ch) == int32('@') {
				if (*MD_LINE)(unsafe.Pointer(line)).Fbeg+uint32(1) <= off && (libc.Uint32FromInt32(libc.Int32FromUint8('A')) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off-libc.Uint32FromInt32(1))))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off-libc.Uint32FromInt32(1))))) <= libc.Uint32FromInt32(libc.Int32FromUint8('Z')) || libc.Uint32FromInt32(libc.Int32FromUint8('a')) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off-libc.Uint32FromInt32(1))))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off-libc.Uint32FromInt32(1))))) <= libc.Uint32FromInt32(libc.Int32FromUint8('z')) || libc.Uint32FromInt32(libc.Int32FromUint8('0')) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off-libc.Uint32FromInt32(1))))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off-libc.Uint32FromInt32(1))))) <= libc.Uint32FromInt32(libc.Int32FromUint8('9'))) && off+uint32(3) < (*MD_LINE)(unsafe.Pointer(line)).Fend && (libc.Uint32FromInt32(libc.Int32FromUint8('A')) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))) <= libc.Uint32FromInt32(libc.Int32FromUint8('Z')) || libc.Uint32FromInt32(libc.Int32FromUint8('a')) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))) <= libc.Uint32FromInt32(libc.Int32FromUint8('z')) || libc.Uint32FromInt32(libc.Int32FromUint8('0')) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))) <= libc.Uint32FromInt32(libc.Int32FromUint8('9'))) {
					mark = md_add_mark(tls, ctx)
					if mark == libc.UintptrFromInt32(0) {
						ret = -int32(1)
						goto abort
					}
					(*MD_MARK)(unsafe.Pointer(mark)).Fbeg = off
					(*MD_MARK)(unsafe.Pointer(mark)).Fend = off + libc.Uint32FromInt32(1)
					(*MD_MARK)(unsafe.Pointer(mark)).Fprev = -int32(1)
					(*MD_MARK)(unsafe.Pointer(mark)).Fnext = -int32(1)
					(*MD_MARK)(unsafe.Pointer(mark)).Fch = ch
					(*MD_MARK)(unsafe.Pointer(mark)).Fflags = libc.Uint8FromInt32(libc.Int32FromInt32(MD_MARK_POTENTIAL_OPENER))
					/* Push a dummy as a reserve for a closer. */
					mark = md_add_mark(tls, ctx)
					if mark == libc.UintptrFromInt32(0) {
						ret = -int32(1)
						goto abort
					}
					(*MD_MARK)(unsafe.Pointer(mark)).Fbeg = (*MD_LINE)(unsafe.Pointer(line)).Fbeg
					(*MD_MARK)(unsafe.Pointer(mark)).Fend = (*MD_LINE)(unsafe.Pointer(line)).Fend
					(*MD_MARK)(unsafe.Pointer(mark)).Fprev = -int32(1)
					(*MD_MARK)(unsafe.Pointer(mark)).Fnext = -int32(1)
					(*MD_MARK)(unsafe.Pointer(mark)).Fch = libc.Uint8FromInt32(libc.Int32FromUint8('D'))
					(*MD_MARK)(unsafe.Pointer(mark)).Fflags = libc.Uint8FromInt32(libc.Int32FromInt32(0))
				}
				off = off + 1
				continue
			}
			/* A potential permissive URL autolink. */
			if libc.Int32FromUint8(ch) == int32(':') {
				scheme_index = 0
				for {
					if !(scheme_index < libc.Int32FromUint64(libc.Uint64FromInt64(96)/libc.Uint64FromInt64(32))) {
						break
					}
					scheme = scheme_map[scheme_index].Fscheme
					scheme_size = scheme_map[scheme_index].Fscheme_size
					suffix = scheme_map[scheme_index].Fsuffix
					suffix_size = scheme_map[scheme_index].Fsuffix_size
					if (*MD_LINE)(unsafe.Pointer(line)).Fbeg+scheme_size <= off && md_ascii_eq(tls, (*MD_CTX)(unsafe.Pointer(ctx)).Ftext+uintptr(off-scheme_size), scheme, scheme_size) != 0 && off+uint32(1)+suffix_size < (*MD_LINE)(unsafe.Pointer(line)).Fend && md_ascii_eq(tls, (*MD_CTX)(unsafe.Pointer(ctx)).Ftext+uintptr(off+libc.Uint32FromInt32(1)), suffix, suffix_size) != 0 {
						mark = md_add_mark(tls, ctx)
						if mark == libc.UintptrFromInt32(0) {
							ret = -int32(1)
							goto abort
						}
						(*MD_MARK)(unsafe.Pointer(mark)).Fbeg = off - scheme_size
						(*MD_MARK)(unsafe.Pointer(mark)).Fend = off + libc.Uint32FromInt32(1) + suffix_size
						(*MD_MARK)(unsafe.Pointer(mark)).Fprev = -int32(1)
						(*MD_MARK)(unsafe.Pointer(mark)).Fnext = -int32(1)
						(*MD_MARK)(unsafe.Pointer(mark)).Fch = ch
						(*MD_MARK)(unsafe.Pointer(mark)).Fflags = libc.Uint8FromInt32(libc.Int32FromInt32(MD_MARK_POTENTIAL_OPENER))
						/* Push a dummy as a reserve for a closer. */
						mark = md_add_mark(tls, ctx)
						if mark == libc.UintptrFromInt32(0) {
							ret = -int32(1)
							goto abort
						}
						(*MD_MARK)(unsafe.Pointer(mark)).Fbeg = (*MD_LINE)(unsafe.Pointer(line)).Fbeg
						(*MD_MARK)(unsafe.Pointer(mark)).Fend = (*MD_LINE)(unsafe.Pointer(line)).Fend
						(*MD_MARK)(unsafe.Pointer(mark)).Fprev = -int32(1)
						(*MD_MARK)(unsafe.Pointer(mark)).Fnext = -int32(1)
						(*MD_MARK)(unsafe.Pointer(mark)).Fch = libc.Uint8FromInt32(libc.Int32FromUint8('D'))
						(*MD_MARK)(unsafe.Pointer(mark)).Fflags = libc.Uint8FromInt32(libc.Int32FromInt32(0))
						off = off + (uint32(1) + suffix_size)
						break
					}
					goto _3
				_3:
					;
					scheme_index = scheme_index + 1
				}
				off = off + 1
				continue
			}
			/* A potential permissive WWW autolink. */
			if libc.Int32FromUint8(ch) == int32('.') {
				if (*MD_LINE)(unsafe.Pointer(line)).Fbeg+uint32(3) <= off && md_ascii_eq(tls, (*MD_CTX)(unsafe.Pointer(ctx)).Ftext+uintptr(off-libc.Uint32FromInt32(3)), __ccgo_ts+100, uint32(3)) != 0 && (off-uint32(3) == (*MD_LINE)(unsafe.Pointer(line)).Fbeg || md_is_unicode_whitespace__(tls, md_decode_utf8_before__(tls, ctx, off-uint32(3))) != 0 || md_is_unicode_punct__(tls, md_decode_utf8_before__(tls, ctx, off-uint32(3))) != 0) {
					mark = md_add_mark(tls, ctx)
					if mark == libc.UintptrFromInt32(0) {
						ret = -int32(1)
						goto abort
					}
					(*MD_MARK)(unsafe.Pointer(mark)).Fbeg = off - libc.Uint32FromInt32(3)
					(*MD_MARK)(unsafe.Pointer(mark)).Fend = off + libc.Uint32FromInt32(1)
					(*MD_MARK)(unsafe.Pointer(mark)).Fprev = -int32(1)
					(*MD_MARK)(unsafe.Pointer(mark)).Fnext = -int32(1)
					(*MD_MARK)(unsafe.Pointer(mark)).Fch = ch
					(*MD_MARK)(unsafe.Pointer(mark)).Fflags = libc.Uint8FromInt32(libc.Int32FromInt32(MD_MARK_POTENTIAL_OPENER))
					/* Push a dummy as a reserve for a closer. */
					mark = md_add_mark(tls, ctx)
					if mark == libc.UintptrFromInt32(0) {
						ret = -int32(1)
						goto abort
					}
					(*MD_MARK)(unsafe.Pointer(mark)).Fbeg = (*MD_LINE)(unsafe.Pointer(line)).Fbeg
					(*MD_MARK)(unsafe.Pointer(mark)).Fend = (*MD_LINE)(unsafe.Pointer(line)).Fend
					(*MD_MARK)(unsafe.Pointer(mark)).Fprev = -int32(1)
					(*MD_MARK)(unsafe.Pointer(mark)).Fnext = -int32(1)
					(*MD_MARK)(unsafe.Pointer(mark)).Fch = libc.Uint8FromInt32(libc.Int32FromUint8('D'))
					(*MD_MARK)(unsafe.Pointer(mark)).Fflags = libc.Uint8FromInt32(libc.Int32FromInt32(0))
					off = off + 1
					continue
				}
				off = off + 1
				continue
			}
			/* A potential table cell boundary or wiki link label delimiter. */
			if (table_mode != 0 || (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fflags&uint32(MD_FLAG_WIKILINKS) != 0) && libc.Int32FromUint8(ch) == int32('|') {
				mark = md_add_mark(tls, ctx)
				if mark == libc.UintptrFromInt32(0) {
					ret = -int32(1)
					goto abort
				}
				(*MD_MARK)(unsafe.Pointer(mark)).Fbeg = off
				(*MD_MARK)(unsafe.Pointer(mark)).Fend = off + libc.Uint32FromInt32(1)
				(*MD_MARK)(unsafe.Pointer(mark)).Fprev = -int32(1)
				(*MD_MARK)(unsafe.Pointer(mark)).Fnext = -int32(1)
				(*MD_MARK)(unsafe.Pointer(mark)).Fch = ch
				(*MD_MARK)(unsafe.Pointer(mark)).Fflags = libc.Uint8FromInt32(libc.Int32FromInt32(0))
				off = off + 1
				continue
			}
			/* A potential strikethrough start/end. */
			if libc.Int32FromUint8(ch) == int32('~') {
				tmp2 = off + uint32(1)
				for tmp2 < (*MD_LINE)(unsafe.Pointer(line)).Fend && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(tmp2)))) == int32('~') {
					tmp2 = tmp2 + 1
				}
				if tmp2-off < uint32(3) {
					flags2 = uint32(0)
					if tmp2 < (*MD_LINE)(unsafe.Pointer(line)).Fend && !(md_is_unicode_whitespace__(tls, md_decode_utf8__(tls, (*MD_CTX)(unsafe.Pointer(ctx)).Ftext+uintptr(tmp2), (*MD_CTX)(unsafe.Pointer(ctx)).Fsize-tmp2, libc.UintptrFromInt32(0))) != 0) {
						flags2 = flags2 | uint32(MD_MARK_POTENTIAL_OPENER)
					}
					if off > (*MD_LINE)(unsafe.Pointer(line)).Fbeg && !(md_is_unicode_whitespace__(tls, md_decode_utf8_before__(tls, ctx, off)) != 0) {
						flags2 = flags2 | uint32(MD_MARK_POTENTIAL_CLOSER)
					}
					if flags2 != uint32(0) {
						mark = md_add_mark(tls, ctx)
						if mark == libc.UintptrFromInt32(0) {
							ret = -int32(1)
							goto abort
						}
						(*MD_MARK)(unsafe.Pointer(mark)).Fbeg = off
						(*MD_MARK)(unsafe.Pointer(mark)).Fend = tmp2
						(*MD_MARK)(unsafe.Pointer(mark)).Fprev = -int32(1)
						(*MD_MARK)(unsafe.Pointer(mark)).Fnext = -int32(1)
						(*MD_MARK)(unsafe.Pointer(mark)).Fch = ch
						(*MD_MARK)(unsafe.Pointer(mark)).Fflags = uint8(flags2)
					}
				}
				off = tmp2
				continue
			}
			/* A potential equation start/end */
			if libc.Int32FromUint8(ch) == int32('$') {
				/* We can have at most two consecutive $ signs,
				 * where two dollar signs signify a display equation. */
				tmp3 = off + uint32(1)
				for tmp3 < (*MD_LINE)(unsafe.Pointer(line)).Fend && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(tmp3)))) == int32('$') {
					tmp3 = tmp3 + 1
				}
				if tmp3-off <= uint32(2) {
					flags3 = libc.Uint32FromInt32(libc.Int32FromInt32(MD_MARK_POTENTIAL_OPENER) | libc.Int32FromInt32(MD_MARK_POTENTIAL_CLOSER))
					if off > (*MD_LINE)(unsafe.Pointer(line)).Fbeg && !(md_is_unicode_whitespace__(tls, md_decode_utf8_before__(tls, ctx, off)) != 0) && !(md_is_unicode_punct__(tls, md_decode_utf8_before__(tls, ctx, off)) != 0) {
						flags3 = flags3 & libc.Uint32FromInt32(^libc.Int32FromInt32(MD_MARK_POTENTIAL_OPENER))
					}
					if tmp3 < (*MD_LINE)(unsafe.Pointer(line)).Fend && !(md_is_unicode_whitespace__(tls, md_decode_utf8__(tls, (*MD_CTX)(unsafe.Pointer(ctx)).Ftext+uintptr(tmp3), (*MD_CTX)(unsafe.Pointer(ctx)).Fsize-tmp3, libc.UintptrFromInt32(0))) != 0) && !(md_is_unicode_punct__(tls, md_decode_utf8__(tls, (*MD_CTX)(unsafe.Pointer(ctx)).Ftext+uintptr(tmp3), (*MD_CTX)(unsafe.Pointer(ctx)).Fsize-tmp3, libc.UintptrFromInt32(0))) != 0) {
						flags3 = flags3 & libc.Uint32FromInt32(^libc.Int32FromInt32(MD_MARK_POTENTIAL_CLOSER))
					}
					if flags3 != uint32(0) {
						mark = md_add_mark(tls, ctx)
						if mark == libc.UintptrFromInt32(0) {
							ret = -int32(1)
							goto abort
						}
						(*MD_MARK)(unsafe.Pointer(mark)).Fbeg = off
						(*MD_MARK)(unsafe.Pointer(mark)).Fend = tmp3
						(*MD_MARK)(unsafe.Pointer(mark)).Fprev = -int32(1)
						(*MD_MARK)(unsafe.Pointer(mark)).Fnext = -int32(1)
						(*MD_MARK)(unsafe.Pointer(mark)).Fch = ch
						(*MD_MARK)(unsafe.Pointer(mark)).Fflags = uint8(flags3)
					}
				}
				off = tmp3
				continue
			}
			/* Turn non-trivial whitespace into single space. */
			if libc.Int32FromUint8(ch) == int32(' ') || libc.Int32FromUint8(ch) == int32('\t') || (libc.Int32FromUint8(ch) == int32('\v') || libc.Int32FromUint8(ch) == int32('\f')) {
				tmp4 = off + uint32(1)
				for tmp4 < (*MD_LINE)(unsafe.Pointer(line)).Fend && (libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(tmp4)))) == int32(' ') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(tmp4)))) == int32('\t') || (libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(tmp4)))) == int32('\v') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(tmp4)))) == int32('\f'))) {
					tmp4 = tmp4 + 1
				}
				if tmp4-off > uint32(1) || libc.Int32FromUint8(ch) != int32(' ') {
					mark = md_add_mark(tls, ctx)
					if mark == libc.UintptrFromInt32(0) {
						ret = -int32(1)
						goto abort
					}
					(*MD_MARK)(unsafe.Pointer(mark)).Fbeg = off
					(*MD_MARK)(unsafe.Pointer(mark)).Fend = tmp4
					(*MD_MARK)(unsafe.Pointer(mark)).Fprev = -int32(1)
					(*MD_MARK)(unsafe.Pointer(mark)).Fnext = -int32(1)
					(*MD_MARK)(unsafe.Pointer(mark)).Fch = ch
					(*MD_MARK)(unsafe.Pointer(mark)).Fflags = libc.Uint8FromInt32(libc.Int32FromInt32(MD_MARK_RESOLVED))
				}
				off = tmp4
				continue
			}
			/* NULL character. */
			if libc.Int32FromUint8(ch) == int32('\000') {
				mark = md_add_mark(tls, ctx)
				if mark == libc.UintptrFromInt32(0) {
					ret = -int32(1)
					goto abort
				}
				(*MD_MARK)(unsafe.Pointer(mark)).Fbeg = off
				(*MD_MARK)(unsafe.Pointer(mark)).Fend = off + libc.Uint32FromInt32(1)
				(*MD_MARK)(unsafe.Pointer(mark)).Fprev = -int32(1)
				(*MD_MARK)(unsafe.Pointer(mark)).Fnext = -int32(1)
				(*MD_MARK)(unsafe.Pointer(mark)).Fch = ch
				(*MD_MARK)(unsafe.Pointer(mark)).Fflags = libc.Uint8FromInt32(libc.Int32FromInt32(MD_MARK_RESOLVED))
				off = off + 1
				continue
			}
			off = off + 1
		}
		goto _1
	_1:
		;
		*(*MD_SIZE)(__ccgo_ptr(bp)) = *(*MD_SIZE)(__ccgo_ptr(bp)) + 1
	}
	/* Add a dummy mark at the end of the mark vector to simplify
	 * process_inlines(). */
	mark = md_add_mark(tls, ctx)
	if mark == libc.UintptrFromInt32(0) {
		ret = -int32(1)
		goto abort
	}
	(*MD_MARK)(unsafe.Pointer(mark)).Fbeg = (*MD_CTX)(unsafe.Pointer(ctx)).Fsize
	(*MD_MARK)(unsafe.Pointer(mark)).Fend = (*MD_CTX)(unsafe.Pointer(ctx)).Fsize
	(*MD_MARK)(unsafe.Pointer(mark)).Fprev = -int32(1)
	(*MD_MARK)(unsafe.Pointer(mark)).Fnext = -int32(1)
	(*MD_MARK)(unsafe.Pointer(mark)).Fch = libc.Uint8FromInt32(libc.Int32FromInt32(127))
	(*MD_MARK)(unsafe.Pointer(mark)).Fflags = libc.Uint8FromInt32(libc.Int32FromInt32(MD_MARK_RESOLVED))
	goto abort
abort:
	;
	return ret
	return r
}

var scheme_map = [3]struct {
	Fscheme      uintptr
	Fscheme_size MD_SIZE
	Fsuffix      uintptr
	Fsuffix_size MD_SIZE
}{
	0: {
		Fscheme:      __ccgo_ts + 82,
		Fscheme_size: uint32(4),
		Fsuffix:      __ccgo_ts + 87,
		Fsuffix_size: uint32(2),
	},
	1: {
		Fscheme:      __ccgo_ts + 90,
		Fscheme_size: uint32(5),
		Fsuffix:      __ccgo_ts + 87,
		Fsuffix_size: uint32(2),
	},
	2: {
		Fscheme:      __ccgo_ts + 96,
		Fscheme_size: uint32(3),
		Fsuffix:      __ccgo_ts + 87,
		Fsuffix_size: uint32(2),
	},
}

func md_analyze_bracket(tls *libc.TLS, ctx uintptr, mark_index int32) {
	var mark, opener, v1 uintptr
	var opener_index int32
	_, _, _, _ = mark, opener, opener_index, v1
	/* We cannot really resolve links here as for that we would need
	 * more context. E.g. a following pair of brackets (reference link),
	 * or enclosing pair of brackets (if the inner is the link, the outer
	 * one cannot be.)
	 *
	 * Therefore we here only construct a list of '[' ']' pairs ordered by
	 * position of the closer. This allows us to analyze what is or is not
	 * link in the right order, from inside to outside in case of nested
	 * brackets.
	 *
	 * The resolving itself is deferred to md_resolve_links().
	 */
	mark = (*MD_CTX)(unsafe.Pointer(ctx)).Fmarks + uintptr(mark_index)*20
	if libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(mark)).Fflags)&int32(MD_MARK_POTENTIAL_OPENER) != 0 {
		if (*(*MD_MARKSTACK)(__ccgo_ptr(ctx + 416 + 14*4))).Ftop >= 0 {
			v1 = (*MD_CTX)(unsafe.Pointer(ctx)).Fmarks + uintptr((*(*MD_MARKSTACK)(__ccgo_ptr(ctx + 416 + 14*4))).Ftop)*20 + 17
			*(*uint8)(unsafe.Pointer(v1)) = uint8(int32(*(*uint8)(unsafe.Pointer(v1))) | libc.Int32FromInt32(MD_MARK_HASNESTEDBRACKETS))
		}
		md_mark_stack_push(tls, ctx, ctx+416+14*4, mark_index)
		return
	}
	if (*(*MD_MARKSTACK)(__ccgo_ptr(ctx + 416 + 14*4))).Ftop >= 0 {
		opener_index = md_mark_stack_pop(tls, ctx, ctx+416+14*4)
		opener = (*MD_CTX)(unsafe.Pointer(ctx)).Fmarks + uintptr(opener_index)*20
		/* Interconnect the opener and closer. */
		(*MD_MARK)(unsafe.Pointer(opener)).Fnext = mark_index
		(*MD_MARK)(unsafe.Pointer(mark)).Fprev = opener_index
		/* Add the pair into a list of potential links for md_resolve_links().
		 * Note we misuse opener->prev for this as opener->next points to its
		 * closer. */
		if (*MD_CTX)(unsafe.Pointer(ctx)).Funresolved_link_tail >= 0 {
			(*(*MD_MARK)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Fmarks + uintptr((*MD_CTX)(unsafe.Pointer(ctx)).Funresolved_link_tail)*20))).Fprev = opener_index
		} else {
			(*MD_CTX)(unsafe.Pointer(ctx)).Funresolved_link_head = opener_index
		}
		(*MD_CTX)(unsafe.Pointer(ctx)).Funresolved_link_tail = opener_index
		(*MD_MARK)(unsafe.Pointer(opener)).Fprev = -int32(1)
	}
}

func md_resolve_links(tls *libc.TLS, ctx uintptr, lines uintptr, n_lines MD_SIZE) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var closer, delim, first_nested, last_nested, m, mark, next_closer, next_opener, opener, v3 uintptr
	var closer_index, delim_index, i, is_link, next_index, opener_index int32
	var dest_beg, dest_end, last_img_beg, last_img_end, last_link_beg, last_link_end, off MD_OFFSET
	var v1 uint32
	var _ /* attr at bp+0 */ MD_LINK_ATTR
	var _ /* inline_link_end at bp+24 */ MD_OFFSET
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = closer, closer_index, delim, delim_index, dest_beg, dest_end, first_nested, i, is_link, last_img_beg, last_img_end, last_link_beg, last_link_end, last_nested, m, mark, next_closer, next_index, next_opener, off, opener, opener_index, v1, v3
	opener_index = (*MD_CTX)(unsafe.Pointer(ctx)).Funresolved_link_head
	last_link_beg = uint32(0)
	last_link_end = uint32(0)
	last_img_beg = uint32(0)
	last_img_end = uint32(0)
	for opener_index >= 0 {
		opener = (*MD_CTX)(unsafe.Pointer(ctx)).Fmarks + uintptr(opener_index)*20
		closer_index = (*MD_MARK)(unsafe.Pointer(opener)).Fnext
		closer = (*MD_CTX)(unsafe.Pointer(ctx)).Fmarks + uintptr(closer_index)*20
		next_index = (*MD_MARK)(unsafe.Pointer(opener)).Fprev
		is_link = FALSE
		if next_index >= 0 {
			next_opener = (*MD_CTX)(unsafe.Pointer(ctx)).Fmarks + uintptr(next_index)*20
			next_closer = (*MD_CTX)(unsafe.Pointer(ctx)).Fmarks + uintptr((*MD_MARK)(unsafe.Pointer(next_opener)).Fnext)*20
		} else {
			next_opener = libc.UintptrFromInt32(0)
			next_closer = libc.UintptrFromInt32(0)
		}
		/* If nested ("[ [ ] ]"), we need to make sure that:
		 *   - The outer does not end inside of (...) belonging to the inner.
		 *   - The outer cannot be link if the inner is link (i.e. not image).
		 *
		 * (Note we here analyze from inner to outer as the marks are ordered
		 * by closer->beg.)
		 */
		if (*MD_MARK)(unsafe.Pointer(opener)).Fbeg < last_link_beg && (*MD_MARK)(unsafe.Pointer(closer)).Fend < last_link_end || (*MD_MARK)(unsafe.Pointer(opener)).Fbeg < last_img_beg && (*MD_MARK)(unsafe.Pointer(closer)).Fend < last_img_end || (*MD_MARK)(unsafe.Pointer(opener)).Fbeg < last_link_end && libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(opener)).Fch) == int32('[') {
			opener_index = next_index
			continue
		}
		/* Recognize and resolve wiki links.
		 * Wiki-links maybe '[[destination]]' or '[[destination|label]]'.
		 */
		if (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fflags&uint32(MD_FLAG_WIKILINKS) != 0 && (*MD_MARK)(unsafe.Pointer(opener)).Fend-(*MD_MARK)(unsafe.Pointer(opener)).Fbeg == uint32(1) && next_opener != libc.UintptrFromInt32(0) && libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(next_opener)).Fch) == int32('[') && (*MD_MARK)(unsafe.Pointer(next_opener)).Fbeg == (*MD_MARK)(unsafe.Pointer(opener)).Fbeg-uint32(1) && (*MD_MARK)(unsafe.Pointer(next_opener)).Fend-(*MD_MARK)(unsafe.Pointer(next_opener)).Fbeg == uint32(1) && next_closer != libc.UintptrFromInt32(0) && libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(next_closer)).Fch) == int32(']') && (*MD_MARK)(unsafe.Pointer(next_closer)).Fbeg == (*MD_MARK)(unsafe.Pointer(closer)).Fbeg+uint32(1) && (*MD_MARK)(unsafe.Pointer(next_closer)).Fend-(*MD_MARK)(unsafe.Pointer(next_closer)).Fbeg == uint32(1) {
			delim = libc.UintptrFromInt32(0)
			is_link = int32(TRUE)
			/* We don't allow destination to be longer than 100 characters.
			 * Lets scan to see whether there is '|'. (If not then the whole
			 * wiki-link has to be below the 100 characters.) */
			delim_index = opener_index + int32(1)
			for delim_index < closer_index {
				m = (*MD_CTX)(unsafe.Pointer(ctx)).Fmarks + uintptr(delim_index)*20
				if libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(m)).Fch) == int32('|') {
					delim = m
					break
				}
				if libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(m)).Fch) != int32('D') {
					if (*MD_MARK)(unsafe.Pointer(m)).Fbeg-(*MD_MARK)(unsafe.Pointer(opener)).Fend > uint32(100) {
						break
					}
					if libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(m)).Fch) != int32('D') && libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(m)).Fflags)&int32(MD_MARK_OPENER) != 0 {
						delim_index = (*MD_MARK)(unsafe.Pointer(m)).Fnext
					}
				}
				delim_index = delim_index + 1
			}
			dest_beg = (*MD_MARK)(unsafe.Pointer(opener)).Fend
			if delim != libc.UintptrFromInt32(0) {
				v1 = (*MD_MARK)(unsafe.Pointer(delim)).Fbeg
			} else {
				v1 = (*MD_MARK)(unsafe.Pointer(closer)).Fbeg
			}
			dest_end = v1
			if dest_end-dest_beg == uint32(0) || dest_end-dest_beg > uint32(100) {
				is_link = FALSE
			}
			/* There may not be any new line in the destination. */
			if is_link != 0 {
				off = dest_beg
				for {
					if !(off < dest_end) {
						break
					}
					if libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('\r') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('\n') {
						is_link = FALSE
						break
					}
					goto _2
				_2:
					;
					off = off + 1
				}
			}
			if is_link != 0 {
				if delim != libc.UintptrFromInt32(0) {
					if (*MD_MARK)(unsafe.Pointer(delim)).Fend < (*MD_MARK)(unsafe.Pointer(closer)).Fbeg {
						md_rollback(tls, ctx, opener_index, delim_index, int32(MD_ROLLBACK_ALL))
						md_rollback(tls, ctx, delim_index, closer_index, MD_ROLLBACK_CROSSING)
						v3 = delim + 17
						*(*uint8)(unsafe.Pointer(v3)) = uint8(int32(*(*uint8)(unsafe.Pointer(v3))) | libc.Int32FromInt32(MD_MARK_RESOLVED))
						(*MD_MARK)(unsafe.Pointer(opener)).Fend = (*MD_MARK)(unsafe.Pointer(delim)).Fbeg
					} else {
						/* The pipe is just before the closer: [[foo|]] */
						md_rollback(tls, ctx, opener_index, closer_index, int32(MD_ROLLBACK_ALL))
						(*MD_MARK)(unsafe.Pointer(closer)).Fbeg = (*MD_MARK)(unsafe.Pointer(delim)).Fbeg
						delim = libc.UintptrFromInt32(0)
					}
				}
				(*MD_MARK)(unsafe.Pointer(opener)).Fbeg = (*MD_MARK)(unsafe.Pointer(next_opener)).Fbeg
				(*MD_MARK)(unsafe.Pointer(opener)).Fnext = closer_index
				v3 = opener + 17
				*(*uint8)(unsafe.Pointer(v3)) = uint8(int32(*(*uint8)(unsafe.Pointer(v3))) | (libc.Int32FromInt32(MD_MARK_OPENER) | libc.Int32FromInt32(MD_MARK_RESOLVED)))
				(*MD_MARK)(unsafe.Pointer(closer)).Fend = (*MD_MARK)(unsafe.Pointer(next_closer)).Fend
				(*MD_MARK)(unsafe.Pointer(closer)).Fprev = opener_index
				v3 = closer + 17
				*(*uint8)(unsafe.Pointer(v3)) = uint8(int32(*(*uint8)(unsafe.Pointer(v3))) | (libc.Int32FromInt32(MD_MARK_CLOSER) | libc.Int32FromInt32(MD_MARK_RESOLVED)))
				last_link_beg = (*MD_MARK)(unsafe.Pointer(opener)).Fbeg
				last_link_end = (*MD_MARK)(unsafe.Pointer(closer)).Fend
				if delim != libc.UintptrFromInt32(0) {
					md_analyze_link_contents(tls, ctx, lines, n_lines, delim_index+int32(1), closer_index)
				}
				opener_index = (*MD_MARK)(unsafe.Pointer(next_opener)).Fprev
				continue
			}
		}
		if next_opener != libc.UintptrFromInt32(0) && (*MD_MARK)(unsafe.Pointer(next_opener)).Fbeg == (*MD_MARK)(unsafe.Pointer(closer)).Fend {
			if (*MD_MARK)(unsafe.Pointer(next_closer)).Fbeg > (*MD_MARK)(unsafe.Pointer(closer)).Fend+uint32(1) {
				/* Might be full reference link. */
				if !(libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(next_opener)).Fflags)&libc.Int32FromInt32(MD_MARK_HASNESTEDBRACKETS) != 0) {
					is_link = md_is_link_reference(tls, ctx, lines, n_lines, (*MD_MARK)(unsafe.Pointer(next_opener)).Fbeg, (*MD_MARK)(unsafe.Pointer(next_closer)).Fend, bp)
				}
			} else {
				/* Might be shortcut reference link. */
				if !(libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(opener)).Fflags)&libc.Int32FromInt32(MD_MARK_HASNESTEDBRACKETS) != 0) {
					is_link = md_is_link_reference(tls, ctx, lines, n_lines, (*MD_MARK)(unsafe.Pointer(opener)).Fbeg, (*MD_MARK)(unsafe.Pointer(closer)).Fend, bp)
				}
			}
			if is_link < 0 {
				return -int32(1)
			}
			if is_link != 0 {
				/* Eat the 2nd "[...]". */
				(*MD_MARK)(unsafe.Pointer(closer)).Fend = (*MD_MARK)(unsafe.Pointer(next_closer)).Fend
				/* Do not analyze the label as a standalone link in the next
				 * iteration. */
				next_index = (*(*MD_MARK)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Fmarks + uintptr(next_index)*20))).Fprev
			}
		} else {
			if (*MD_MARK)(unsafe.Pointer(closer)).Fend < (*MD_CTX)(unsafe.Pointer(ctx)).Fsize && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr((*MD_MARK)(unsafe.Pointer(closer)).Fend)))) == int32('(') {
				/* Might be inline link. */
				*(*MD_OFFSET)(__ccgo_ptr(bp + 24)) = uint32(0xffffffff)
				is_link = md_is_inline_link_spec(tls, ctx, lines, n_lines, (*MD_MARK)(unsafe.Pointer(closer)).Fend, bp+24, bp)
				if is_link < 0 {
					return -int32(1)
				}
				/* Check the closing ')' is not inside an already resolved range
				 * (i.e. a range with a higher priority), e.g. a code span. */
				if is_link != 0 {
					i = closer_index + int32(1)
					for i < (*MD_CTX)(unsafe.Pointer(ctx)).Fn_marks {
						mark = (*MD_CTX)(unsafe.Pointer(ctx)).Fmarks + uintptr(i)*20
						if (*MD_MARK)(unsafe.Pointer(mark)).Fbeg >= *(*MD_OFFSET)(__ccgo_ptr(bp + 24)) {
							break
						}
						if libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(mark)).Fflags)&(libc.Int32FromInt32(MD_MARK_OPENER)|libc.Int32FromInt32(MD_MARK_RESOLVED)) == libc.Int32FromInt32(MD_MARK_OPENER)|libc.Int32FromInt32(MD_MARK_RESOLVED) {
							if (*(*MD_MARK)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Fmarks + uintptr((*MD_MARK)(unsafe.Pointer(mark)).Fnext)*20))).Fbeg >= *(*MD_OFFSET)(__ccgo_ptr(bp + 24)) {
								/* Cancel the link status. */
								if (*(*MD_LINK_ATTR)(__ccgo_ptr(bp))).Ftitle_needs_free != 0 {
									libc.Xfree(tls, (*(*MD_LINK_ATTR)(__ccgo_ptr(bp))).Ftitle)
								}
								is_link = FALSE
								break
							}
							i = (*MD_MARK)(unsafe.Pointer(mark)).Fnext + int32(1)
						} else {
							i = i + 1
						}
					}
				}
				if is_link != 0 {
					/* Eat the "(...)" */
					(*MD_MARK)(unsafe.Pointer(closer)).Fend = *(*MD_OFFSET)(__ccgo_ptr(bp + 24))
				}
			}
			if !(is_link != 0) {
				/* Might be collapsed reference link. */
				if !(libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(opener)).Fflags)&libc.Int32FromInt32(MD_MARK_HASNESTEDBRACKETS) != 0) {
					is_link = md_is_link_reference(tls, ctx, lines, n_lines, (*MD_MARK)(unsafe.Pointer(opener)).Fbeg, (*MD_MARK)(unsafe.Pointer(closer)).Fend, bp)
				}
				if is_link < 0 {
					return -int32(1)
				}
			}
		}
		if is_link != 0 {
			/* Resolve the brackets as a link. */
			v3 = opener + 17
			*(*uint8)(unsafe.Pointer(v3)) = uint8(int32(*(*uint8)(unsafe.Pointer(v3))) | (libc.Int32FromInt32(MD_MARK_OPENER) | libc.Int32FromInt32(MD_MARK_RESOLVED)))
			v3 = closer + 17
			*(*uint8)(unsafe.Pointer(v3)) = uint8(int32(*(*uint8)(unsafe.Pointer(v3))) | (libc.Int32FromInt32(MD_MARK_CLOSER) | libc.Int32FromInt32(MD_MARK_RESOLVED)))
			/* If it is a link, we store the destination and title in the two
			 * dummy marks after the opener. */
			if !(libc.Int32FromUint8((*(*MD_MARK)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Fmarks + uintptr(opener_index+int32(1))*20))).Fch) == libc.Int32FromUint8('D')) {
				libc.X__builtin_unreachable(tls)
			}
			(*(*MD_MARK)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Fmarks + uintptr(opener_index+int32(1))*20))).Fbeg = (*(*MD_LINK_ATTR)(__ccgo_ptr(bp))).Fdest_beg
			(*(*MD_MARK)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Fmarks + uintptr(opener_index+int32(1))*20))).Fend = (*(*MD_LINK_ATTR)(__ccgo_ptr(bp))).Fdest_end
			if !(libc.Int32FromUint8((*(*MD_MARK)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Fmarks + uintptr(opener_index+int32(2))*20))).Fch) == libc.Int32FromUint8('D')) {
				libc.X__builtin_unreachable(tls)
			}
			md_mark_store_ptr(tls, ctx, opener_index+int32(2), (*(*MD_LINK_ATTR)(__ccgo_ptr(bp))).Ftitle)
			/* The title might or might not have been allocated for us. */
			if (*(*MD_LINK_ATTR)(__ccgo_ptr(bp))).Ftitle_needs_free != 0 {
				md_mark_stack_push(tls, ctx, ctx+480, opener_index+int32(2))
			}
			(*(*MD_MARK)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Fmarks + uintptr(opener_index+int32(2))*20))).Fprev = libc.Int32FromUint32((*(*MD_LINK_ATTR)(__ccgo_ptr(bp))).Ftitle_size)
			if libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(opener)).Fch) == int32('[') {
				last_link_beg = (*MD_MARK)(unsafe.Pointer(opener)).Fbeg
				last_link_end = (*MD_MARK)(unsafe.Pointer(closer)).Fend
			} else {
				last_img_beg = (*MD_MARK)(unsafe.Pointer(opener)).Fbeg
				last_img_end = (*MD_MARK)(unsafe.Pointer(closer)).Fend
			}
			md_analyze_link_contents(tls, ctx, lines, n_lines, opener_index+int32(1), closer_index)
			/* If the link text is formed by nothing but permissive autolink,
			 * suppress the autolink.
			 * See https://github.com/mity/md4c/issues/152 for more info. */
			if (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fflags&libc.Uint32FromInt32(libc.Int32FromInt32(MD_FLAG_PERMISSIVEEMAILAUTOLINKS)|libc.Int32FromInt32(MD_FLAG_PERMISSIVEURLAUTOLINKS)|libc.Int32FromInt32(MD_FLAG_PERMISSIVEWWWAUTOLINKS)) != 0 {
				first_nested = opener + uintptr(1)*20
				for libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(first_nested)).Fch) == int32('D') && first_nested < closer {
					first_nested += 20
				}
				last_nested = closer - uintptr(1)*20
				for libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(first_nested)).Fch) == int32('D') && last_nested > opener {
					last_nested -= 20
				}
				if libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(first_nested)).Fflags)&int32(MD_MARK_RESOLVED) != 0 && (*MD_MARK)(unsafe.Pointer(first_nested)).Fbeg == (*MD_MARK)(unsafe.Pointer(opener)).Fend && (libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(first_nested)).Fch) != int32('\000') && libc.Xstrchr(tls, __ccgo_ts+104, libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(first_nested)).Fch)) != libc.UintptrFromInt32(0)) && int64((*MD_MARK)(unsafe.Pointer(first_nested)).Fnext) == (int64(last_nested)-int64((*MD_CTX)(unsafe.Pointer(ctx)).Fmarks))/20 && (*MD_MARK)(unsafe.Pointer(last_nested)).Fend == (*MD_MARK)(unsafe.Pointer(closer)).Fbeg {
					(*MD_MARK)(unsafe.Pointer(first_nested)).Fch = uint8('D')
					v3 = first_nested + 17
					*(*uint8)(unsafe.Pointer(v3)) = uint8(int32(*(*uint8)(unsafe.Pointer(v3))) & ^libc.Int32FromInt32(MD_MARK_RESOLVED))
					(*MD_MARK)(unsafe.Pointer(last_nested)).Fch = uint8('D')
					v3 = last_nested + 17
					*(*uint8)(unsafe.Pointer(v3)) = uint8(int32(*(*uint8)(unsafe.Pointer(v3))) & ^libc.Int32FromInt32(MD_MARK_RESOLVED))
				}
			}
		}
		opener_index = next_index
	}
	return 0
}

// C documentation
//
//	/* Analyze whether the mark '&' starts a HTML entity.
//	 * If so, update its flags as well as flags of corresponding closer ';'. */
func md_analyze_entity(tls *libc.TLS, ctx uintptr, mark_index int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var closer, opener uintptr
	var _ /* off at bp+0 */ MD_OFFSET
	_, _ = closer, opener
	opener = (*MD_CTX)(unsafe.Pointer(ctx)).Fmarks + uintptr(mark_index)*20
	/* Cannot be entity if there is no closer as the next mark.
	 * (Any other mark between would mean strange character which cannot be
	 * part of the entity.
	 *
	 * So we can do all the work on '&' and do not call this later for the
	 * closing mark ';'.
	 */
	if mark_index+int32(1) >= (*MD_CTX)(unsafe.Pointer(ctx)).Fn_marks {
		return
	}
	closer = (*MD_CTX)(unsafe.Pointer(ctx)).Fmarks + uintptr(mark_index+int32(1))*20
	if libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(closer)).Fch) != int32(';') {
		return
	}
	if md_is_entity(tls, ctx, (*MD_MARK)(unsafe.Pointer(opener)).Fbeg, (*MD_MARK)(unsafe.Pointer(closer)).Fend, bp) != 0 {
		if !(*(*MD_OFFSET)(__ccgo_ptr(bp)) == (*MD_MARK)(unsafe.Pointer(closer)).Fend) {
			libc.X__builtin_unreachable(tls)
		}
		md_resolve_range(tls, ctx, mark_index, mark_index+int32(1))
		(*MD_MARK)(unsafe.Pointer(opener)).Fend = (*MD_MARK)(unsafe.Pointer(closer)).Fend
	}
}

func md_analyze_table_cell_boundary(tls *libc.TLS, ctx uintptr, mark_index int32) {
	var mark, v1 uintptr
	_, _ = mark, v1
	mark = (*MD_CTX)(unsafe.Pointer(ctx)).Fmarks + uintptr(mark_index)*20
	v1 = mark + 17
	*(*uint8)(unsafe.Pointer(v1)) = uint8(int32(*(*uint8)(unsafe.Pointer(v1))) | libc.Int32FromInt32(MD_MARK_RESOLVED))
	(*MD_MARK)(unsafe.Pointer(mark)).Fnext = -int32(1)
	if (*MD_CTX)(unsafe.Pointer(ctx)).Ftable_cell_boundaries_head < 0 {
		(*MD_CTX)(unsafe.Pointer(ctx)).Ftable_cell_boundaries_head = mark_index
	} else {
		(*(*MD_MARK)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Fmarks + uintptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftable_cell_boundaries_tail)*20))).Fnext = mark_index
	}
	(*MD_CTX)(unsafe.Pointer(ctx)).Ftable_cell_boundaries_tail = mark_index
	(*MD_CTX)(unsafe.Pointer(ctx)).Fn_table_cell_boundaries = (*MD_CTX)(unsafe.Pointer(ctx)).Fn_table_cell_boundaries + 1
}

// C documentation
//
//	/* Split a longer mark into two. The new mark takes the given count of
//	 * characters. May only be called if an adequate number of dummy 'D' marks
//	 * follows.
//	 */
func md_split_emph_mark(tls *libc.TLS, ctx uintptr, mark_index int32, n MD_SIZE) (r int32) {
	var dummy, mark uintptr
	var new_mark_index int32
	_, _, _ = dummy, mark, new_mark_index
	mark = (*MD_CTX)(unsafe.Pointer(ctx)).Fmarks + uintptr(mark_index)*20
	new_mark_index = libc.Int32FromUint32(libc.Uint32FromInt32(mark_index) + ((*MD_MARK)(unsafe.Pointer(mark)).Fend - (*MD_MARK)(unsafe.Pointer(mark)).Fbeg - n))
	dummy = (*MD_CTX)(unsafe.Pointer(ctx)).Fmarks + uintptr(new_mark_index)*20
	if !((*MD_MARK)(unsafe.Pointer(mark)).Fend-(*MD_MARK)(unsafe.Pointer(mark)).Fbeg > n) {
		libc.X__builtin_unreachable(tls)
	}
	if !(libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(dummy)).Fch) == libc.Int32FromUint8('D')) {
		libc.X__builtin_unreachable(tls)
	}
	libc.Xmemcpy(tls, dummy, mark, uint64(20))
	*(*MD_OFFSET)(__ccgo_ptr(mark + 4)) -= n
	(*MD_MARK)(unsafe.Pointer(dummy)).Fbeg = (*MD_MARK)(unsafe.Pointer(mark)).Fend
	return new_mark_index
}

func md_analyze_emph(tls *libc.TLS, ctx uintptr, mark_index int32) {
	var closer_size, opener_size MD_SIZE
	var flags uint32
	var i, m_index, n_opener_stacks, opener_index, v1 int32
	var m, mark, opener, stack uintptr
	var opener_stacks [6]uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _ = closer_size, flags, i, m, m_index, mark, n_opener_stacks, opener, opener_index, opener_size, opener_stacks, stack, v1
	mark = (*MD_CTX)(unsafe.Pointer(ctx)).Fmarks + uintptr(mark_index)*20
	/* If we can be a closer, try to resolve with the preceding opener. */
	if libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(mark)).Fflags)&int32(MD_MARK_POTENTIAL_CLOSER) != 0 {
		opener = libc.UintptrFromInt32(0)
		opener_index = 0
		flags = uint32((*MD_MARK)(unsafe.Pointer(mark)).Fflags)
		n_opener_stacks = 0
		/* Apply the rule of 3 */
		v1 = n_opener_stacks
		n_opener_stacks = n_opener_stacks + 1
		opener_stacks[v1] = md_emph_stack(tls, ctx, (*MD_MARK)(unsafe.Pointer(mark)).Fch, libc.Uint32FromInt32(libc.Int32FromInt32(MD_MARK_EMPH_MOD3_0)|libc.Int32FromInt32(MD_MARK_EMPH_OC)))
		if flags&libc.Uint32FromInt32(libc.Int32FromInt32(0x40)|libc.Int32FromInt32(0x80)) != libc.Uint32FromInt32(libc.Int32FromInt32(0x40)|libc.Int32FromInt32(0x80)) {
			v1 = n_opener_stacks
			n_opener_stacks = n_opener_stacks + 1
			opener_stacks[v1] = md_emph_stack(tls, ctx, (*MD_MARK)(unsafe.Pointer(mark)).Fch, libc.Uint32FromInt32(libc.Int32FromInt32(MD_MARK_EMPH_MOD3_1)|libc.Int32FromInt32(MD_MARK_EMPH_OC)))
		}
		if flags&libc.Uint32FromInt32(libc.Int32FromInt32(0x40)|libc.Int32FromInt32(0x80)) != uint32(MD_MARK_EMPH_MOD3_1) {
			v1 = n_opener_stacks
			n_opener_stacks = n_opener_stacks + 1
			opener_stacks[v1] = md_emph_stack(tls, ctx, (*MD_MARK)(unsafe.Pointer(mark)).Fch, libc.Uint32FromInt32(libc.Int32FromInt32(0x40)|libc.Int32FromInt32(0x80)|libc.Int32FromInt32(MD_MARK_EMPH_OC)))
		}
		v1 = n_opener_stacks
		n_opener_stacks = n_opener_stacks + 1
		opener_stacks[v1] = md_emph_stack(tls, ctx, (*MD_MARK)(unsafe.Pointer(mark)).Fch, uint32(MD_MARK_EMPH_MOD3_0))
		if !(flags&libc.Uint32FromInt32(MD_MARK_EMPH_OC) != 0) || flags&libc.Uint32FromInt32(libc.Int32FromInt32(0x40)|libc.Int32FromInt32(0x80)) != libc.Uint32FromInt32(libc.Int32FromInt32(0x40)|libc.Int32FromInt32(0x80)) {
			v1 = n_opener_stacks
			n_opener_stacks = n_opener_stacks + 1
			opener_stacks[v1] = md_emph_stack(tls, ctx, (*MD_MARK)(unsafe.Pointer(mark)).Fch, uint32(MD_MARK_EMPH_MOD3_1))
		}
		if !(flags&libc.Uint32FromInt32(MD_MARK_EMPH_OC) != 0) || flags&libc.Uint32FromInt32(libc.Int32FromInt32(0x40)|libc.Int32FromInt32(0x80)) != uint32(MD_MARK_EMPH_MOD3_1) {
			v1 = n_opener_stacks
			n_opener_stacks = n_opener_stacks + 1
			opener_stacks[v1] = md_emph_stack(tls, ctx, (*MD_MARK)(unsafe.Pointer(mark)).Fch, libc.Uint32FromInt32(libc.Int32FromInt32(0x40)|libc.Int32FromInt32(0x80)))
		}
		/* Opener is the most recent mark from the allowed stacks. */
		i = 0
		for {
			if !(i < n_opener_stacks) {
				break
			}
			if (*MD_MARKSTACK)(unsafe.Pointer(opener_stacks[i])).Ftop >= 0 {
				m_index = (*MD_MARKSTACK)(unsafe.Pointer(opener_stacks[i])).Ftop
				m = (*MD_CTX)(unsafe.Pointer(ctx)).Fmarks + uintptr(m_index)*20
				if opener == libc.UintptrFromInt32(0) || (*MD_MARK)(unsafe.Pointer(m)).Fend > (*MD_MARK)(unsafe.Pointer(opener)).Fend {
					opener_index = m_index
					opener = m
				}
			}
			goto _7
		_7:
			;
			i = i + 1
		}
		/* Resolve, if we have found matching opener. */
		if opener != libc.UintptrFromInt32(0) {
			opener_size = (*MD_MARK)(unsafe.Pointer(opener)).Fend - (*MD_MARK)(unsafe.Pointer(opener)).Fbeg
			closer_size = (*MD_MARK)(unsafe.Pointer(mark)).Fend - (*MD_MARK)(unsafe.Pointer(mark)).Fbeg
			stack = md_opener_stack(tls, ctx, opener_index)
			if opener_size > closer_size {
				opener_index = md_split_emph_mark(tls, ctx, opener_index, closer_size)
				md_mark_stack_push(tls, ctx, stack, opener_index)
			} else {
				if opener_size < closer_size {
					md_split_emph_mark(tls, ctx, mark_index, closer_size-opener_size)
				}
			}
			/* Above we were only peeking. */
			md_mark_stack_pop(tls, ctx, stack)
			md_rollback(tls, ctx, opener_index, mark_index, MD_ROLLBACK_CROSSING)
			md_resolve_range(tls, ctx, opener_index, mark_index)
			return
		}
	}
	/* If we could not resolve as closer, we may be yet be an opener. */
	if libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(mark)).Fflags)&int32(MD_MARK_POTENTIAL_OPENER) != 0 {
		md_mark_stack_push(tls, ctx, md_emph_stack(tls, ctx, (*MD_MARK)(unsafe.Pointer(mark)).Fch, uint32((*MD_MARK)(unsafe.Pointer(mark)).Fflags)), mark_index)
	}
}

func md_analyze_tilde(tls *libc.TLS, ctx uintptr, mark_index int32) {
	var mark, stack uintptr
	var opener_index int32
	_, _, _ = mark, opener_index, stack
	mark = (*MD_CTX)(unsafe.Pointer(ctx)).Fmarks + uintptr(mark_index)*20
	stack = md_opener_stack(tls, ctx, mark_index)
	/* We attempt to be Github Flavored Markdown compatible here. GFM accepts
	 * only tildes sequences of length 1 and 2, and the length of the opener
	 * and closer has to match. */
	if libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(mark)).Fflags)&int32(MD_MARK_POTENTIAL_CLOSER) != 0 && (*MD_MARKSTACK)(unsafe.Pointer(stack)).Ftop >= 0 {
		opener_index = (*MD_MARKSTACK)(unsafe.Pointer(stack)).Ftop
		md_mark_stack_pop(tls, ctx, stack)
		md_rollback(tls, ctx, opener_index, mark_index, MD_ROLLBACK_CROSSING)
		md_resolve_range(tls, ctx, opener_index, mark_index)
		return
	}
	if libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(mark)).Fflags)&int32(MD_MARK_POTENTIAL_OPENER) != 0 {
		md_mark_stack_push(tls, ctx, stack, mark_index)
	}
}

func md_analyze_dollar(tls *libc.TLS, ctx uintptr, mark_index int32) {
	var closer, mark, opener uintptr
	var closer_index, opener_index int32
	_, _, _, _, _ = closer, closer_index, mark, opener, opener_index
	mark = (*MD_CTX)(unsafe.Pointer(ctx)).Fmarks + uintptr(mark_index)*20
	if libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(mark)).Fflags)&int32(MD_MARK_POTENTIAL_CLOSER) != 0 && (*(*MD_MARKSTACK)(__ccgo_ptr(ctx + 416 + 15*4))).Ftop >= 0 {
		/* If the potential closer has a non-matching number of $, discard */
		opener = (*MD_CTX)(unsafe.Pointer(ctx)).Fmarks + uintptr((*(*MD_MARKSTACK)(__ccgo_ptr(ctx + 416 + 15*4))).Ftop)*20
		opener_index = (*(*MD_MARKSTACK)(__ccgo_ptr(ctx + 416 + 15*4))).Ftop
		closer = mark
		closer_index = mark_index
		if (*MD_MARK)(unsafe.Pointer(opener)).Fend-(*MD_MARK)(unsafe.Pointer(opener)).Fbeg == (*MD_MARK)(unsafe.Pointer(closer)).Fend-(*MD_MARK)(unsafe.Pointer(closer)).Fbeg {
			/* We are the matching closer */
			md_mark_stack_pop(tls, ctx, ctx+416+15*4)
			md_rollback(tls, ctx, opener_index, closer_index, int32(MD_ROLLBACK_ALL))
			md_resolve_range(tls, ctx, opener_index, closer_index)
			/* Discard all pending openers: Latex math span do not allow
			 * nesting. */
			(*(*MD_MARKSTACK)(__ccgo_ptr(ctx + 416 + 15*4))).Ftop = -int32(1)
			return
		}
	}
	if libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(mark)).Fflags)&int32(MD_MARK_POTENTIAL_OPENER) != 0 {
		md_mark_stack_push(tls, ctx, ctx+416+15*4, mark_index)
	}
}

func md_scan_left_for_resolved_mark(tls *libc.TLS, ctx uintptr, mark_from uintptr, off MD_OFFSET, p_cursor uintptr) (r uintptr) {
	var mark uintptr
	_ = mark
	mark = mark_from
	for {
		if !(mark >= (*MD_CTX)(unsafe.Pointer(ctx)).Fmarks) {
			break
		}
		if libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(mark)).Fch) == int32('D') || (*MD_MARK)(unsafe.Pointer(mark)).Fbeg > off {
			goto _1
		}
		if (*MD_MARK)(unsafe.Pointer(mark)).Fbeg <= off && off < (*MD_MARK)(unsafe.Pointer(mark)).Fend && libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(mark)).Fflags)&int32(MD_MARK_RESOLVED) != 0 {
			if p_cursor != libc.UintptrFromInt32(0) {
				*(*uintptr)(__ccgo_ptr(p_cursor)) = mark
			}
			return mark
		}
		if (*MD_MARK)(unsafe.Pointer(mark)).Fend <= off {
			break
		}
		goto _1
	_1:
		;
		mark -= 20
	}
	if p_cursor != libc.UintptrFromInt32(0) {
		*(*uintptr)(__ccgo_ptr(p_cursor)) = mark
	}
	return libc.UintptrFromInt32(0)
}

func md_scan_right_for_resolved_mark(tls *libc.TLS, ctx uintptr, mark_from uintptr, off MD_OFFSET, p_cursor uintptr) (r uintptr) {
	var mark uintptr
	_ = mark
	mark = mark_from
	for {
		if !(mark < (*MD_CTX)(unsafe.Pointer(ctx)).Fmarks+uintptr((*MD_CTX)(unsafe.Pointer(ctx)).Fn_marks)*20) {
			break
		}
		if libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(mark)).Fch) == int32('D') || (*MD_MARK)(unsafe.Pointer(mark)).Fend <= off {
			goto _1
		}
		if (*MD_MARK)(unsafe.Pointer(mark)).Fbeg <= off && off < (*MD_MARK)(unsafe.Pointer(mark)).Fend && libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(mark)).Fflags)&int32(MD_MARK_RESOLVED) != 0 {
			if p_cursor != libc.UintptrFromInt32(0) {
				*(*uintptr)(__ccgo_ptr(p_cursor)) = mark
			}
			return mark
		}
		if (*MD_MARK)(unsafe.Pointer(mark)).Fbeg > off {
			break
		}
		goto _1
	_1:
		;
		mark += 20
	}
	if p_cursor != libc.UintptrFromInt32(0) {
		*(*uintptr)(__ccgo_ptr(p_cursor)) = mark
	}
	return libc.UintptrFromInt32(0)
}

func md_analyze_permissive_autolink(tls *libc.TLS, ctx uintptr, mark_index int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var beg, end, line_beg, line_end MD_OFFSET
	var closer, left_mark, opener, right_mark uintptr
	var i uint32
	var left_boundary_ok, n_components, n_open_brackets, right_boundary_ok int32
	var _ /* left_cursor at bp+0 */ uintptr
	var _ /* right_cursor at bp+8 */ uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _ = beg, closer, end, i, left_boundary_ok, left_mark, line_beg, line_end, n_components, n_open_brackets, opener, right_boundary_ok, right_mark
	opener = (*MD_CTX)(unsafe.Pointer(ctx)).Fmarks + uintptr(mark_index)*20
	closer = (*MD_CTX)(unsafe.Pointer(ctx)).Fmarks + uintptr(mark_index+int32(1))*20 /* The dummy. */
	line_beg = (*MD_MARK)(unsafe.Pointer(closer)).Fbeg                               /* md_collect_mark() set this for us */
	line_end = (*MD_MARK)(unsafe.Pointer(closer)).Fend                               /* ditto */
	beg = (*MD_MARK)(unsafe.Pointer(opener)).Fbeg
	end = (*MD_MARK)(unsafe.Pointer(opener)).Fend
	*(*uintptr)(__ccgo_ptr(bp)) = opener
	left_boundary_ok = FALSE
	*(*uintptr)(__ccgo_ptr(bp + 8)) = opener
	right_boundary_ok = FALSE
	if !(libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(closer)).Fch) == libc.Int32FromUint8('D')) {
		libc.X__builtin_unreachable(tls)
	}
	if libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(opener)).Fch) == int32('@') {
		if !(libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr((*MD_MARK)(unsafe.Pointer(opener)).Fbeg)))) == libc.Int32FromUint8('@')) {
			libc.X__builtin_unreachable(tls)
		}
		/* Scan backwards for the user name (before '@'). */
		for beg > line_beg {
			if libc.Uint32FromInt32(libc.Int32FromUint8('A')) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(beg-libc.Uint32FromInt32(1))))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(beg-libc.Uint32FromInt32(1))))) <= libc.Uint32FromInt32(libc.Int32FromUint8('Z')) || libc.Uint32FromInt32(libc.Int32FromUint8('a')) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(beg-libc.Uint32FromInt32(1))))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(beg-libc.Uint32FromInt32(1))))) <= libc.Uint32FromInt32(libc.Int32FromUint8('z')) || libc.Uint32FromInt32(libc.Int32FromUint8('0')) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(beg-libc.Uint32FromInt32(1))))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(beg-libc.Uint32FromInt32(1))))) <= libc.Uint32FromInt32(libc.Int32FromUint8('9')) {
				beg = beg - 1
			} else {
				if beg >= line_beg+uint32(2) && (libc.Uint32FromInt32(libc.Int32FromUint8('A')) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(beg-libc.Uint32FromInt32(2))))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(beg-libc.Uint32FromInt32(2))))) <= libc.Uint32FromInt32(libc.Int32FromUint8('Z')) || libc.Uint32FromInt32(libc.Int32FromUint8('a')) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(beg-libc.Uint32FromInt32(2))))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(beg-libc.Uint32FromInt32(2))))) <= libc.Uint32FromInt32(libc.Int32FromUint8('z')) || libc.Uint32FromInt32(libc.Int32FromUint8('0')) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(beg-libc.Uint32FromInt32(2))))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(beg-libc.Uint32FromInt32(2))))) <= libc.Uint32FromInt32(libc.Int32FromUint8('9'))) && (libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(beg-libc.Uint32FromInt32(1))))) != int32('\000') && libc.Xstrchr(tls, __ccgo_ts+131, libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(beg-libc.Uint32FromInt32(1)))))) != libc.UintptrFromInt32(0)) && md_scan_left_for_resolved_mark(tls, ctx, *(*uintptr)(__ccgo_ptr(bp)), beg-uint32(1), bp) == libc.UintptrFromInt32(0) && (libc.Uint32FromInt32(libc.Int32FromUint8('A')) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(beg)))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(beg)))) <= libc.Uint32FromInt32(libc.Int32FromUint8('Z')) || libc.Uint32FromInt32(libc.Int32FromUint8('a')) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(beg)))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(beg)))) <= libc.Uint32FromInt32(libc.Int32FromUint8('z')) || libc.Uint32FromInt32(libc.Int32FromUint8('0')) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(beg)))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(beg)))) <= libc.Uint32FromInt32(libc.Int32FromUint8('9'))) {
					beg = beg - 1
				} else {
					break
				}
			}
		}
		if beg == (*MD_MARK)(unsafe.Pointer(opener)).Fbeg { /* empty user name */
			return
		}
	}
	/* Verify there's line boundary, whitespace, allowed punctuation or
	 * resolved emphasis mark just before the suspected autolink. */
	if beg == line_beg || md_is_unicode_whitespace__(tls, md_decode_utf8_before__(tls, ctx, beg)) != 0 || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(beg-libc.Uint32FromInt32(1))))) != int32('\000') && libc.Xstrchr(tls, __ccgo_ts+136, libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(beg-libc.Uint32FromInt32(1)))))) != libc.UintptrFromInt32(0) {
		left_boundary_ok = int32(TRUE)
	} else {
		if libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(beg-libc.Uint32FromInt32(1))))) != int32('\000') && libc.Xstrchr(tls, __ccgo_ts+140, libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(beg-libc.Uint32FromInt32(1)))))) != libc.UintptrFromInt32(0) {
			left_mark = md_scan_left_for_resolved_mark(tls, ctx, *(*uintptr)(__ccgo_ptr(bp)), beg-uint32(1), bp)
			if left_mark != libc.UintptrFromInt32(0) && libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(left_mark)).Fflags)&int32(MD_MARK_OPENER) != 0 {
				left_boundary_ok = int32(TRUE)
			}
		}
	}
	if !(left_boundary_ok != 0) {
		return
	}
	i = uint32(0)
	for {
		if !(uint64(i) < libc.Uint64FromInt64(96)/libc.Uint64FromInt64(24)) {
			break
		}
		n_components = 0
		n_open_brackets = 0
		if libc.Int32FromUint8(URL_MAP[i].Fstart_char) != int32('\000') {
			if end >= line_end || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(end)))) != libc.Int32FromUint8(URL_MAP[i].Fstart_char) {
				goto _1
			}
			if URL_MAP[i].Fmin_components > 0 && (end+uint32(1) >= line_end || !(libc.Uint32FromInt32(libc.Int32FromUint8('A')) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(end+libc.Uint32FromInt32(1))))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(end+libc.Uint32FromInt32(1))))) <= libc.Uint32FromInt32(libc.Int32FromUint8('Z')) || libc.Uint32FromInt32(libc.Int32FromUint8('a')) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(end+libc.Uint32FromInt32(1))))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(end+libc.Uint32FromInt32(1))))) <= libc.Uint32FromInt32(libc.Int32FromUint8('z')) || libc.Uint32FromInt32(libc.Int32FromUint8('0')) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(end+libc.Uint32FromInt32(1))))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(end+libc.Uint32FromInt32(1))))) <= libc.Uint32FromInt32(libc.Int32FromUint8('9')))) {
				goto _1
			}
			end = end + 1
		}
		for end < line_end {
			if libc.Uint32FromInt32(libc.Int32FromUint8('A')) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(end)))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(end)))) <= libc.Uint32FromInt32(libc.Int32FromUint8('Z')) || libc.Uint32FromInt32(libc.Int32FromUint8('a')) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(end)))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(end)))) <= libc.Uint32FromInt32(libc.Int32FromUint8('z')) || libc.Uint32FromInt32(libc.Int32FromUint8('0')) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(end)))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(end)))) <= libc.Uint32FromInt32(libc.Int32FromUint8('9')) {
				if n_components == 0 {
					n_components = n_components + 1
				}
				end = end + 1
			} else {
				if end < line_end && (libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(end)))) != int32('\000') && libc.Xstrchr(tls, URL_MAP[i].Fallowed_nonalnum_chars, libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(end))))) != libc.UintptrFromInt32(0)) && md_scan_right_for_resolved_mark(tls, ctx, *(*uintptr)(__ccgo_ptr(bp + 8)), end, bp+8) == libc.UintptrFromInt32(0) && (end > line_beg && (libc.Uint32FromInt32(libc.Int32FromUint8('A')) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(end-libc.Uint32FromInt32(1))))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(end-libc.Uint32FromInt32(1))))) <= libc.Uint32FromInt32(libc.Int32FromUint8('Z')) || libc.Uint32FromInt32(libc.Int32FromUint8('a')) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(end-libc.Uint32FromInt32(1))))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(end-libc.Uint32FromInt32(1))))) <= libc.Uint32FromInt32(libc.Int32FromUint8('z')) || libc.Uint32FromInt32(libc.Int32FromUint8('0')) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(end-libc.Uint32FromInt32(1))))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(end-libc.Uint32FromInt32(1))))) <= libc.Uint32FromInt32(libc.Int32FromUint8('9')) || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(end-libc.Uint32FromInt32(1))))) == int32(')')) || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(end)))) == int32('(')) && (end+uint32(1) < line_end && (libc.Uint32FromInt32(libc.Int32FromUint8('A')) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(end+libc.Uint32FromInt32(1))))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(end+libc.Uint32FromInt32(1))))) <= libc.Uint32FromInt32(libc.Int32FromUint8('Z')) || libc.Uint32FromInt32(libc.Int32FromUint8('a')) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(end+libc.Uint32FromInt32(1))))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(end+libc.Uint32FromInt32(1))))) <= libc.Uint32FromInt32(libc.Int32FromUint8('z')) || libc.Uint32FromInt32(libc.Int32FromUint8('0')) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(end+libc.Uint32FromInt32(1))))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(end+libc.Uint32FromInt32(1))))) <= libc.Uint32FromInt32(libc.Int32FromUint8('9')) || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(end+libc.Uint32FromInt32(1))))) == int32('(')) || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(end)))) == int32(')')) {
					if libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(end)))) == libc.Int32FromUint8(URL_MAP[i].Fdelim_char) {
						n_components = n_components + 1
					}
					/* brackets have to be balanced. */
					if libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(end)))) == int32('(') {
						n_open_brackets = n_open_brackets + 1
					} else {
						if libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(end)))) == int32(')') {
							if n_open_brackets <= 0 {
								break
							}
							n_open_brackets = n_open_brackets - 1
						}
					}
					end = end + 1
				} else {
					break
				}
			}
		}
		if end < line_end && libc.Int32FromUint8(URL_MAP[i].Foptional_end_char) != int32('\000') && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(end)))) == libc.Int32FromUint8(URL_MAP[i].Foptional_end_char) {
			end = end + 1
		}
		if n_components < URL_MAP[i].Fmin_components || n_open_brackets != 0 {
			return
		}
		if libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(opener)).Fch) == int32('@') { /* E-mail autolinks wants only the host. */
			break
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	/* Verify there's line boundary, whitespace, allowed punctuation or
	 * resolved emphasis mark just after the suspected autolink. */
	if end == line_end || md_is_unicode_whitespace__(tls, md_decode_utf8__(tls, (*MD_CTX)(unsafe.Pointer(ctx)).Ftext+uintptr(end), (*MD_CTX)(unsafe.Pointer(ctx)).Fsize-end, libc.UintptrFromInt32(0))) != 0 || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(end)))) != int32('\000') && libc.Xstrchr(tls, __ccgo_ts+144, libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(end))))) != libc.UintptrFromInt32(0) {
		right_boundary_ok = int32(TRUE)
	} else {
		right_mark = md_scan_right_for_resolved_mark(tls, ctx, *(*uintptr)(__ccgo_ptr(bp + 8)), end, bp+8)
		if right_mark != libc.UintptrFromInt32(0) && libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(right_mark)).Fflags)&int32(MD_MARK_CLOSER) != 0 {
			right_boundary_ok = int32(TRUE)
		}
	}
	if !(right_boundary_ok != 0) {
		return
	}
	/* Success, we are an autolink. */
	(*MD_MARK)(unsafe.Pointer(opener)).Fbeg = beg
	(*MD_MARK)(unsafe.Pointer(opener)).Fend = beg
	(*MD_MARK)(unsafe.Pointer(closer)).Fbeg = end
	(*MD_MARK)(unsafe.Pointer(closer)).Fend = end
	(*MD_MARK)(unsafe.Pointer(closer)).Fch = (*MD_MARK)(unsafe.Pointer(opener)).Fch
	md_resolve_range(tls, ctx, mark_index, mark_index+int32(1))
}

var URL_MAP = [4]struct {
	Fstart_char             MD_CHAR
	Fdelim_char             MD_CHAR
	Fallowed_nonalnum_chars uintptr
	Fmin_components         int32
	Foptional_end_char      MD_CHAR
}{
	0: {
		Fdelim_char:             uint8('.'),
		Fallowed_nonalnum_chars: __ccgo_ts + 108,
		Fmin_components:         int32(2),
	},
	1: {
		Fstart_char:             uint8('/'),
		Fdelim_char:             uint8('/'),
		Fallowed_nonalnum_chars: __ccgo_ts + 112,
		Foptional_end_char:      uint8('/'),
	},
	2: {
		Fstart_char:             uint8('?'),
		Fdelim_char:             uint8('&'),
		Fallowed_nonalnum_chars: __ccgo_ts + 117,
		Fmin_components:         int32(1),
	},
	3: {
		Fstart_char:             uint8('#'),
		Fallowed_nonalnum_chars: __ccgo_ts + 126,
		Fmin_components:         int32(1),
	},
}

func md_analyze_marks(tls *libc.TLS, ctx uintptr, lines uintptr, n_lines MD_SIZE, mark_beg int32, mark_end int32, mark_chars uintptr, flags uint32) {
	var i int32
	var last_end MD_OFFSET
	var mark uintptr
	_, _, _ = i, last_end, mark
	i = mark_beg
	last_end = (*(*MD_LINE)(__ccgo_ptr(lines))).Fbeg
	_ = lines
	_ = n_lines
	for i < mark_end {
		mark = (*MD_CTX)(unsafe.Pointer(ctx)).Fmarks + uintptr(i)*20
		/* Skip resolved spans. */
		if libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(mark)).Fflags)&int32(MD_MARK_RESOLVED) != 0 {
			if libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(mark)).Fflags)&int32(MD_MARK_OPENER) != 0 && !(flags&uint32(MD_ANALYZE_NOSKIP_EMPH) != 0 && (libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(mark)).Fch) != int32('\000') && libc.Xstrchr(tls, __ccgo_ts+140, libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(mark)).Fch)) != libc.UintptrFromInt32(0))) {
				if !(i < (*MD_MARK)(unsafe.Pointer(mark)).Fnext) {
					libc.X__builtin_unreachable(tls)
				}
				i = (*MD_MARK)(unsafe.Pointer(mark)).Fnext + int32(1)
			} else {
				i = i + 1
			}
			continue
		}
		/* Skip marks we do not want to deal with. */
		if !(libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(mark)).Fch) != int32('\000') && libc.Xstrchr(tls, mark_chars, libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(mark)).Fch)) != libc.UintptrFromInt32(0)) {
			i = i + 1
			continue
		}
		/* The resolving in previous step could have expanded a mark. */
		if (*MD_MARK)(unsafe.Pointer(mark)).Fbeg < last_end {
			i = i + 1
			continue
		}
		/* Analyze the mark. */
		switch libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(mark)).Fch) {
		case int32('['): /* Pass through. */
			fallthrough
		case int32('!'): /* Pass through. */
			fallthrough
		case int32(']'):
			md_analyze_bracket(tls, ctx, i)
		case int32('&'):
			md_analyze_entity(tls, ctx, i)
		case int32('|'):
			md_analyze_table_cell_boundary(tls, ctx, i)
		case int32('_'): /* Pass through. */
			fallthrough
		case int32('*'):
			md_analyze_emph(tls, ctx, i)
		case int32('~'):
			md_analyze_tilde(tls, ctx, i)
		case int32('$'):
			md_analyze_dollar(tls, ctx, i)
		case int32('.'): /* Pass through. */
			fallthrough
		case int32(':'): /* Pass through. */
			fallthrough
		case int32('@'):
			md_analyze_permissive_autolink(tls, ctx, i)
			break
		}
		if libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(mark)).Fflags)&int32(MD_MARK_RESOLVED) != 0 {
			if libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(mark)).Fflags)&int32(MD_MARK_OPENER) != 0 {
				last_end = (*(*MD_MARK)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Fmarks + uintptr((*MD_MARK)(unsafe.Pointer(mark)).Fnext)*20))).Fend
			} else {
				last_end = (*MD_MARK)(unsafe.Pointer(mark)).Fend
			}
		}
		i = i + 1
	}
}

// C documentation
//
//	/* Analyze marks (build ctx->marks). */
func md_analyze_inlines(tls *libc.TLS, ctx uintptr, lines uintptr, n_lines MD_SIZE, table_mode int32) (r int32) {
	var ret int32
	_ = ret
	/* Reset the previously collected stack of marks. */
	(*MD_CTX)(unsafe.Pointer(ctx)).Fn_marks = 0
	/* Collect all marks. */
	ret = md_collect_marks(tls, ctx, lines, n_lines, table_mode)
	if ret < 0 {
		goto abort
	}
	/* (1) Links. */
	md_analyze_marks(tls, ctx, lines, n_lines, 0, (*MD_CTX)(unsafe.Pointer(ctx)).Fn_marks, __ccgo_ts+153, uint32(0))
	ret = md_resolve_links(tls, ctx, lines, n_lines)
	if ret < 0 {
		goto abort
	}
	(*(*MD_MARKSTACK)(__ccgo_ptr(ctx + 416 + 14*4))).Ftop = -int32(1)
	(*MD_CTX)(unsafe.Pointer(ctx)).Funresolved_link_head = -int32(1)
	(*MD_CTX)(unsafe.Pointer(ctx)).Funresolved_link_tail = -int32(1)
	if table_mode != 0 {
		/* (2) Analyze table cell boundaries. */
		if !(n_lines == libc.Uint32FromInt32(1)) {
			libc.X__builtin_unreachable(tls)
		}
		(*MD_CTX)(unsafe.Pointer(ctx)).Fn_table_cell_boundaries = 0
		md_analyze_marks(tls, ctx, lines, n_lines, 0, (*MD_CTX)(unsafe.Pointer(ctx)).Fn_marks, __ccgo_ts+157, uint32(0))
		return ret
	}
	/* (3) Emphasis and strong emphasis; permissive autolinks. */
	md_analyze_link_contents(tls, ctx, lines, n_lines, 0, (*MD_CTX)(unsafe.Pointer(ctx)).Fn_marks)
	goto abort
abort:
	;
	return ret
	return r
}

func md_analyze_link_contents(tls *libc.TLS, ctx uintptr, lines uintptr, n_lines MD_SIZE, mark_beg int32, mark_end int32) {
	var i int32
	_ = i
	md_analyze_marks(tls, ctx, lines, n_lines, mark_beg, mark_end, __ccgo_ts+159, uint32(0))
	md_analyze_marks(tls, ctx, lines, n_lines, mark_beg, mark_end, __ccgo_ts+161, uint32(0))
	if (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fflags&libc.Uint32FromInt32(libc.Int32FromInt32(MD_FLAG_PERMISSIVEEMAILAUTOLINKS)|libc.Int32FromInt32(MD_FLAG_PERMISSIVEURLAUTOLINKS)|libc.Int32FromInt32(MD_FLAG_PERMISSIVEWWWAUTOLINKS)) != uint32(0) {
		/* These have to be processed last, as they may be greedy and expand
		 * from their original mark. Also their implementation must be careful
		 * not to cross any (previously) resolved marks when doing so. */
		md_analyze_marks(tls, ctx, lines, n_lines, mark_beg, mark_end, __ccgo_ts+104, uint32(MD_ANALYZE_NOSKIP_EMPH))
	}
	i = 0
	for {
		if !(i < libc.Int32FromUint64(libc.Uint64FromInt64(64)/libc.Uint64FromInt64(4))) {
			break
		}
		(*(*MD_MARKSTACK)(__ccgo_ptr(ctx + 416 + uintptr(i)*4))).Ftop = -int32(1)
		goto _1
	_1:
		;
		i = i + 1
	}
}

func md_enter_leave_span_a(tls *libc.TLS, ctx uintptr, enter int32, type1 MD_SPANTYPE1, dest uintptr, dest_size MD_SIZE, is_autolink int32, title uintptr, title_size MD_SIZE) (r int32) {
	bp := tls.Alloc(176)
	defer tls.Free(176)
	var ret, v1 int32
	var _ /* det at bp+96 */ MD_SPAN_A_DETAIL
	var _ /* href_build at bp+0 */ MD_ATTRIBUTE_BUILD
	var _ /* title_build at bp+48 */ MD_ATTRIBUTE_BUILD
	_, _ = ret, v1
	*(*MD_ATTRIBUTE_BUILD)(__ccgo_ptr(bp)) = MD_ATTRIBUTE_BUILD{}
	*(*MD_ATTRIBUTE_BUILD)(__ccgo_ptr(bp + 48)) = MD_ATTRIBUTE_BUILD{}
	ret = 0
	/* Note we here rely on fact that MD_SPAN_A_DETAIL and
	 * MD_SPAN_IMG_DETAIL are binary-compatible. */
	libc.Xmemset(tls, bp+96, 0, uint64(72))
	if is_autolink != 0 {
		v1 = int32(MD_BUILD_ATTR_NO_ESCAPES)
	} else {
		v1 = 0
	}
	ret = md_build_attribute(tls, ctx, dest, dest_size, libc.Uint32FromInt32(v1), bp+96, bp)
	if ret < 0 {
		goto abort
	}
	ret = md_build_attribute(tls, ctx, title, title_size, uint32(0), bp+96+32, bp+48)
	if ret < 0 {
		goto abort
	}
	(*(*MD_SPAN_A_DETAIL)(__ccgo_ptr(bp + 96))).Fis_autolink = is_autolink
	if enter != 0 {
		ret = (*(*func(*libc.TLS, MD_SPANTYPE1, uintptr, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fenter_span})))(tls, type1, bp+96, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
		if ret != 0 {
			if (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log != libc.UintptrFromInt32(0) {
				(*(*func(*libc.TLS, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log})))(tls, __ccgo_ts+166, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
			}
			goto abort
		}
	} else {
		ret = (*(*func(*libc.TLS, MD_SPANTYPE1, uintptr, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fleave_span})))(tls, type1, bp+96, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
		if ret != 0 {
			if (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log != libc.UintptrFromInt32(0) {
				(*(*func(*libc.TLS, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log})))(tls, __ccgo_ts+202, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
			}
			goto abort
		}
	}
	goto abort
abort:
	;
	md_free_attribute(tls, ctx, bp)
	md_free_attribute(tls, ctx, bp+48)
	return ret
}

func md_enter_leave_span_wikilink(tls *libc.TLS, ctx uintptr, enter int32, target uintptr, target_size MD_SIZE) (r int32) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	var ret int32
	var _ /* det at bp+48 */ MD_SPAN_WIKILINK_DETAIL
	var _ /* target_build at bp+0 */ MD_ATTRIBUTE_BUILD
	_ = ret
	*(*MD_ATTRIBUTE_BUILD)(__ccgo_ptr(bp)) = MD_ATTRIBUTE_BUILD{}
	ret = 0
	libc.Xmemset(tls, bp+48, 0, uint64(32))
	ret = md_build_attribute(tls, ctx, target, target_size, uint32(0), bp+48, bp)
	if ret < 0 {
		goto abort
	}
	if enter != 0 {
		ret = (*(*func(*libc.TLS, MD_SPANTYPE1, uintptr, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fenter_span})))(tls, int32(MD_SPAN_WIKILINK1), bp+48, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
		if ret != 0 {
			if (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log != libc.UintptrFromInt32(0) {
				(*(*func(*libc.TLS, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log})))(tls, __ccgo_ts+166, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
			}
			goto abort
		}
	} else {
		ret = (*(*func(*libc.TLS, MD_SPANTYPE1, uintptr, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fleave_span})))(tls, int32(MD_SPAN_WIKILINK1), bp+48, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
		if ret != 0 {
			if (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log != libc.UintptrFromInt32(0) {
				(*(*func(*libc.TLS, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log})))(tls, __ccgo_ts+202, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
			}
			goto abort
		}
	}
	goto abort
abort:
	;
	md_free_attribute(tls, ctx, bp)
	return ret
}

// C documentation
//
//	/* Render the output, accordingly to the analyzed ctx->marks. */
func md_process_inlines(tls *libc.TLS, ctx uintptr, lines uintptr, n_lines MD_SIZE) (r int32) {
	var break_type, text_type MD_TEXTTYPE1
	var closer, closer1, dest, dest_mark, line, mark, new_buffer, opener, opener1, prev_mark, title_mark, v26, v27 uintptr
	var dest_size, new_size, target_sz MD_SIZE
	var end, off, tmp MD_OFFSET
	var enforce_hardbreak, has_label, ret, v24 int32
	var v1 uint32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = break_type, closer, closer1, dest, dest_mark, dest_size, end, enforce_hardbreak, has_label, line, mark, new_buffer, new_size, off, opener, opener1, prev_mark, ret, target_sz, text_type, title_mark, tmp, v1, v24, v26, v27
	line = lines
	prev_mark = libc.UintptrFromInt32(0)
	off = (*(*MD_LINE)(__ccgo_ptr(lines))).Fbeg
	end = (*(*MD_LINE)(__ccgo_ptr(lines + uintptr(n_lines-uint32(1))*8))).Fend
	enforce_hardbreak = 0
	ret = 0
	/* Find first resolved mark. Note there is always at least one resolved
	 * mark,  the dummy last one after the end of the latest line we actually
	 * never really reach. This saves us of a lot of special checks and cases
	 * in this function. */
	mark = (*MD_CTX)(unsafe.Pointer(ctx)).Fmarks
	for !(libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(mark)).Fflags)&libc.Int32FromInt32(MD_MARK_RESOLVED) != 0) {
		mark += 20
	}
	text_type = int32(MD_TEXT_NORMAL)
	for int32(1) != 0 {
		/* Process the text up to the next mark or end-of-line. */
		if (*MD_LINE)(unsafe.Pointer(line)).Fend < (*MD_MARK)(unsafe.Pointer(mark)).Fbeg {
			v1 = (*MD_LINE)(unsafe.Pointer(line)).Fend
		} else {
			v1 = (*MD_MARK)(unsafe.Pointer(mark)).Fbeg
		}
		tmp = v1
		if tmp > off {
			if tmp-off > uint32(0) {
				ret = (*(*func(*libc.TLS, MD_TEXTTYPE1, uintptr, MD_SIZE, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Ftext})))(tls, text_type, (*MD_CTX)(unsafe.Pointer(ctx)).Ftext+uintptr(off), tmp-off, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
				if ret != 0 {
					if (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log != libc.UintptrFromInt32(0) {
						(*(*func(*libc.TLS, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log})))(tls, __ccgo_ts+238, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
					}
					goto abort
				}
			}
			off = tmp
		}
		/* If reached the mark, process it and move to next one. */
		if off >= (*MD_MARK)(unsafe.Pointer(mark)).Fbeg {
			switch libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(mark)).Fch) {
			case int32('\\'):
				goto _2
			case int32(' '):
				goto _3
			case int32('`'):
				goto _4
			case int32('_'):
				goto _5
			case int32('*'):
				goto _6
			case int32('~'):
				goto _7
			case int32('$'):
				goto _8
			case int32(']'):
				goto _9
			case int32('!'):
				goto _10
			case int32('['):
				goto _11
			case int32('>'):
				goto _12
			case int32('<'):
				goto _13
			case int32('.'):
				goto _14
			case int32(':'):
				goto _15
			case int32('@'):
				goto _16
			case int32('&'):
				goto _17
			case int32('\000'):
				goto _18
			case int32(127):
				goto _19
			}
			goto _20
		_2:
			; /* Backslash escape. */
			if libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr((*MD_MARK)(unsafe.Pointer(mark)).Fbeg+libc.Uint32FromInt32(1))))) == int32('\r') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr((*MD_MARK)(unsafe.Pointer(mark)).Fbeg+libc.Uint32FromInt32(1))))) == int32('\n') {
				enforce_hardbreak = int32(1)
			} else {
				if int32(1) > 0 {
					ret = (*(*func(*libc.TLS, MD_TEXTTYPE1, uintptr, MD_SIZE, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Ftext})))(tls, text_type, (*MD_CTX)(unsafe.Pointer(ctx)).Ftext+uintptr((*MD_MARK)(unsafe.Pointer(mark)).Fbeg+libc.Uint32FromInt32(1)), libc.Uint32FromInt32(libc.Int32FromInt32(1)), (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
					if ret != 0 {
						if (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log != libc.UintptrFromInt32(0) {
							(*(*func(*libc.TLS, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log})))(tls, __ccgo_ts+238, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
						}
						goto abort
					}
				}
			}
			goto _20
		_3:
			; /* Non-trivial space. */
		_23:
			;
			if int32(1) > 0 {
				ret = (*(*func(*libc.TLS, MD_TEXTTYPE1, uintptr, MD_SIZE, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Ftext})))(tls, text_type, __ccgo_ts+268, libc.Uint32FromInt32(libc.Int32FromInt32(1)), (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
				if ret != 0 {
					if (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log != libc.UintptrFromInt32(0) {
						(*(*func(*libc.TLS, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log})))(tls, __ccgo_ts+238, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
					}
					goto abort
				}
			}
			goto _22
		_22:
			;
			if 0 != 0 {
				goto _23
			}
			goto _21
		_21:
			;
			goto _20
		_4:
			; /* Code span. */
			if libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(mark)).Fflags)&int32(MD_MARK_OPENER) != 0 {
				ret = (*(*func(*libc.TLS, MD_SPANTYPE1, uintptr, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fenter_span})))(tls, int32(MD_SPAN_CODE), libc.UintptrFromInt32(0), (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
				if ret != 0 {
					if (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log != libc.UintptrFromInt32(0) {
						(*(*func(*libc.TLS, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log})))(tls, __ccgo_ts+166, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
					}
					goto abort
				}
				text_type = int32(MD_TEXT_CODE)
			} else {
				ret = (*(*func(*libc.TLS, MD_SPANTYPE1, uintptr, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fleave_span})))(tls, int32(MD_SPAN_CODE), libc.UintptrFromInt32(0), (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
				if ret != 0 {
					if (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log != libc.UintptrFromInt32(0) {
						(*(*func(*libc.TLS, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log})))(tls, __ccgo_ts+202, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
					}
					goto abort
				}
				text_type = int32(MD_TEXT_NORMAL)
			}
			goto _20
		_5:
			; /* Underline (or emphasis if we fall through). */
			if (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fflags&uint32(MD_FLAG_UNDERLINE) != 0 {
				if libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(mark)).Fflags)&int32(MD_MARK_OPENER) != 0 {
					for off < (*MD_MARK)(unsafe.Pointer(mark)).Fend {
						ret = (*(*func(*libc.TLS, MD_SPANTYPE1, uintptr, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fenter_span})))(tls, int32(MD_SPAN_U), libc.UintptrFromInt32(0), (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
						if ret != 0 {
							if (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log != libc.UintptrFromInt32(0) {
								(*(*func(*libc.TLS, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log})))(tls, __ccgo_ts+166, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
							}
							goto abort
						}
						off = off + 1
					}
				} else {
					for off < (*MD_MARK)(unsafe.Pointer(mark)).Fend {
						ret = (*(*func(*libc.TLS, MD_SPANTYPE1, uintptr, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fleave_span})))(tls, int32(MD_SPAN_U), libc.UintptrFromInt32(0), (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
						if ret != 0 {
							if (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log != libc.UintptrFromInt32(0) {
								(*(*func(*libc.TLS, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log})))(tls, __ccgo_ts+202, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
							}
							goto abort
						}
						off = off + 1
					}
				}
				goto _20
			}
		_6:
			; /* Emphasis, strong emphasis. */
			if libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(mark)).Fflags)&int32(MD_MARK_OPENER) != 0 {
				if ((*MD_MARK)(unsafe.Pointer(mark)).Fend-off)%uint32(2) != 0 {
					ret = (*(*func(*libc.TLS, MD_SPANTYPE1, uintptr, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fenter_span})))(tls, int32(MD_SPAN_EM), libc.UintptrFromInt32(0), (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
					if ret != 0 {
						if (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log != libc.UintptrFromInt32(0) {
							(*(*func(*libc.TLS, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log})))(tls, __ccgo_ts+166, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
						}
						goto abort
					}
					off = off + 1
				}
				for off+uint32(1) < (*MD_MARK)(unsafe.Pointer(mark)).Fend {
					ret = (*(*func(*libc.TLS, MD_SPANTYPE1, uintptr, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fenter_span})))(tls, int32(MD_SPAN_STRONG), libc.UintptrFromInt32(0), (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
					if ret != 0 {
						if (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log != libc.UintptrFromInt32(0) {
							(*(*func(*libc.TLS, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log})))(tls, __ccgo_ts+166, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
						}
						goto abort
					}
					off = off + uint32(2)
				}
			} else {
				for off+uint32(1) < (*MD_MARK)(unsafe.Pointer(mark)).Fend {
					ret = (*(*func(*libc.TLS, MD_SPANTYPE1, uintptr, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fleave_span})))(tls, int32(MD_SPAN_STRONG), libc.UintptrFromInt32(0), (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
					if ret != 0 {
						if (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log != libc.UintptrFromInt32(0) {
							(*(*func(*libc.TLS, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log})))(tls, __ccgo_ts+202, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
						}
						goto abort
					}
					off = off + uint32(2)
				}
				if ((*MD_MARK)(unsafe.Pointer(mark)).Fend-off)%uint32(2) != 0 {
					ret = (*(*func(*libc.TLS, MD_SPANTYPE1, uintptr, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fleave_span})))(tls, int32(MD_SPAN_EM), libc.UintptrFromInt32(0), (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
					if ret != 0 {
						if (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log != libc.UintptrFromInt32(0) {
							(*(*func(*libc.TLS, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log})))(tls, __ccgo_ts+202, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
						}
						goto abort
					}
					off = off + 1
				}
			}
			goto _20
		_7:
			;
			if libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(mark)).Fflags)&int32(MD_MARK_OPENER) != 0 {
				ret = (*(*func(*libc.TLS, MD_SPANTYPE1, uintptr, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fenter_span})))(tls, int32(MD_SPAN_DEL), libc.UintptrFromInt32(0), (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
				if ret != 0 {
					if (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log != libc.UintptrFromInt32(0) {
						(*(*func(*libc.TLS, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log})))(tls, __ccgo_ts+166, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
					}
					goto abort
				}
			} else {
				ret = (*(*func(*libc.TLS, MD_SPANTYPE1, uintptr, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fleave_span})))(tls, int32(MD_SPAN_DEL), libc.UintptrFromInt32(0), (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
				if ret != 0 {
					if (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log != libc.UintptrFromInt32(0) {
						(*(*func(*libc.TLS, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log})))(tls, __ccgo_ts+202, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
					}
					goto abort
				}
			}
			goto _20
		_8:
			;
			if libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(mark)).Fflags)&int32(MD_MARK_OPENER) != 0 {
				if ((*MD_MARK)(unsafe.Pointer(mark)).Fend-off)%uint32(2) != 0 {
					v24 = int32(MD_SPAN_LATEXMATH)
				} else {
					v24 = int32(MD_SPAN_LATEXMATH_DISPLAY)
				}
				ret = (*(*func(*libc.TLS, MD_SPANTYPE1, uintptr, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fenter_span})))(tls, v24, libc.UintptrFromInt32(0), (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
				if ret != 0 {
					if (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log != libc.UintptrFromInt32(0) {
						(*(*func(*libc.TLS, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log})))(tls, __ccgo_ts+166, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
					}
					goto abort
				}
				text_type = int32(MD_TEXT_LATEXMATH)
			} else {
				if ((*MD_MARK)(unsafe.Pointer(mark)).Fend-off)%uint32(2) != 0 {
					v24 = int32(MD_SPAN_LATEXMATH)
				} else {
					v24 = int32(MD_SPAN_LATEXMATH_DISPLAY)
				}
				ret = (*(*func(*libc.TLS, MD_SPANTYPE1, uintptr, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fleave_span})))(tls, v24, libc.UintptrFromInt32(0), (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
				if ret != 0 {
					if (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log != libc.UintptrFromInt32(0) {
						(*(*func(*libc.TLS, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log})))(tls, __ccgo_ts+202, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
					}
					goto abort
				}
				text_type = int32(MD_TEXT_NORMAL)
			}
			goto _20
		_11:
			; /* Link, wiki link, image. */
		_10:
			;
		_9:
			;
			if libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(mark)).Fch) != int32(']') {
				v26 = mark
			} else {
				v26 = (*MD_CTX)(unsafe.Pointer(ctx)).Fmarks + uintptr((*MD_MARK)(unsafe.Pointer(mark)).Fprev)*20
			}
			opener = v26
			closer = (*MD_CTX)(unsafe.Pointer(ctx)).Fmarks + uintptr((*MD_MARK)(unsafe.Pointer(opener)).Fnext)*20
			if libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(opener)).Fch) == int32('[') && libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(closer)).Fch) == int32(']') && (*MD_MARK)(unsafe.Pointer(opener)).Fend-(*MD_MARK)(unsafe.Pointer(opener)).Fbeg >= uint32(2) && (*MD_MARK)(unsafe.Pointer(closer)).Fend-(*MD_MARK)(unsafe.Pointer(closer)).Fbeg >= uint32(2) {
				has_label = libc.BoolInt32((*MD_MARK)(unsafe.Pointer(opener)).Fend-(*MD_MARK)(unsafe.Pointer(opener)).Fbeg > uint32(2))
				if has_label != 0 {
					target_sz = (*MD_MARK)(unsafe.Pointer(opener)).Fend - ((*MD_MARK)(unsafe.Pointer(opener)).Fbeg + uint32(2))
				} else {
					target_sz = (*MD_MARK)(unsafe.Pointer(closer)).Fbeg - (*MD_MARK)(unsafe.Pointer(opener)).Fend
				}
				if has_label != 0 {
					v26 = (*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr((*MD_MARK)(unsafe.Pointer(opener)).Fbeg+libc.Uint32FromInt32(2))
				} else {
					v26 = (*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr((*MD_MARK)(unsafe.Pointer(opener)).Fend)
				}
				ret = md_enter_leave_span_wikilink(tls, ctx, libc.BoolInt32(libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(mark)).Fch) != int32(']')), v26, target_sz)
				if ret < 0 {
					goto abort
				}
				goto _20
			}
			dest_mark = opener + uintptr(1)*20
			if !(libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(dest_mark)).Fch) == libc.Int32FromUint8('D')) {
				libc.X__builtin_unreachable(tls)
			}
			title_mark = opener + uintptr(2)*20
			if !(libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(title_mark)).Fch) == libc.Int32FromUint8('D')) {
				libc.X__builtin_unreachable(tls)
			}
			if libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(opener)).Fch) == int32('!') {
				v24 = int32(MD_SPAN_IMG)
			} else {
				v24 = int32(MD_SPAN_A)
			}
			ret = md_enter_leave_span_a(tls, ctx, libc.BoolInt32(libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(mark)).Fch) != int32(']')), v24, (*MD_CTX)(unsafe.Pointer(ctx)).Ftext+uintptr((*MD_MARK)(unsafe.Pointer(dest_mark)).Fbeg), (*MD_MARK)(unsafe.Pointer(dest_mark)).Fend-(*MD_MARK)(unsafe.Pointer(dest_mark)).Fbeg, FALSE, md_mark_get_ptr(tls, ctx, int32((int64(title_mark)-int64((*MD_CTX)(unsafe.Pointer(ctx)).Fmarks))/20)), libc.Uint32FromInt32((*MD_MARK)(unsafe.Pointer(title_mark)).Fprev))
			if ret < 0 {
				goto abort
			}
			/* link/image closer may span multiple lines. */
			if libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(mark)).Fch) == int32(']') {
				for (*MD_MARK)(unsafe.Pointer(mark)).Fend > (*MD_LINE)(unsafe.Pointer(line)).Fend {
					line += 8
				}
			}
			goto _20
		_13:
			;
		_12:
			; /* Autolink or raw HTML. */
			if !(libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(mark)).Fflags)&libc.Int32FromInt32(MD_MARK_AUTOLINK) != 0) {
				/* Raw HTML. */
				if libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(mark)).Fflags)&int32(MD_MARK_OPENER) != 0 {
					text_type = int32(MD_TEXT_HTML)
				} else {
					text_type = int32(MD_TEXT_NORMAL)
				}
				goto _20
			}
		_16:
			; /* Permissive e-mail autolink. */
		_15:
			; /* Permissive URL autolink. */
		_14:
			; /* Permissive WWW autolink. */
			if libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(mark)).Fflags)&int32(MD_MARK_OPENER) != 0 {
				v26 = mark
			} else {
				v26 = (*MD_CTX)(unsafe.Pointer(ctx)).Fmarks + uintptr((*MD_MARK)(unsafe.Pointer(mark)).Fprev)*20
			}
			opener1 = v26
			closer1 = (*MD_CTX)(unsafe.Pointer(ctx)).Fmarks + uintptr((*MD_MARK)(unsafe.Pointer(opener1)).Fnext)*20
			dest = (*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr((*MD_MARK)(unsafe.Pointer(opener1)).Fend)
			dest_size = (*MD_MARK)(unsafe.Pointer(closer1)).Fbeg - (*MD_MARK)(unsafe.Pointer(opener1)).Fend
			/* For permissive auto-links we do not know closer mark
			 * position at the time of md_collect_marks(), therefore
			 * it can be out-of-order in ctx->marks[].
			 *
			 * With this flag, we make sure that we output the closer
			 * only if we processed the opener. */
			if libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(mark)).Fflags)&int32(MD_MARK_OPENER) != 0 {
				v27 = closer1 + 17
				*(*uint8)(unsafe.Pointer(v27)) = uint8(int32(*(*uint8)(unsafe.Pointer(v27))) | libc.Int32FromInt32(MD_MARK_VALIDPERMISSIVEAUTOLINK))
			}
			if libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(opener1)).Fch) == int32('@') || libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(opener1)).Fch) == int32('.') || libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(opener1)).Fch) == int32('<') && libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(opener1)).Fflags)&int32(MD_MARK_AUTOLINK_MISSING_MAILTO) != 0 {
				dest_size = dest_size + uint32(7)
				if uint64(dest_size)*uint64(1) > uint64((*MD_CTX)(unsafe.Pointer(ctx)).Falloc_buffer) {
					new_size = uint32((uint64(dest_size)*uint64(1) + uint64(dest_size)*uint64(1)/uint64(2) + uint64(128)) & libc.Uint64FromInt32(^libc.Int32FromInt32(127)))
					new_buffer = libc.Xrealloc(tls, (*MD_CTX)(unsafe.Pointer(ctx)).Fbuffer, uint64(new_size))
					if new_buffer == libc.UintptrFromInt32(0) {
						if (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log != libc.UintptrFromInt32(0) {
							(*(*func(*libc.TLS, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log})))(tls, __ccgo_ts+43, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
						}
						ret = -int32(1)
						goto abort
					}
					(*MD_CTX)(unsafe.Pointer(ctx)).Fbuffer = new_buffer
					(*MD_CTX)(unsafe.Pointer(ctx)).Falloc_buffer = new_size
				}
				if libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(opener1)).Fch) == int32('.') {
					v26 = __ccgo_ts + 270
				} else {
					v26 = __ccgo_ts + 278
				}
				libc.Xmemcpy(tls, (*MD_CTX)(unsafe.Pointer(ctx)).Fbuffer, v26, libc.Uint64FromInt32(7)*libc.Uint64FromInt64(1))
				libc.Xmemcpy(tls, (*MD_CTX)(unsafe.Pointer(ctx)).Fbuffer+uintptr(7), dest, uint64(dest_size-libc.Uint32FromInt32(7))*uint64(1))
				dest = (*MD_CTX)(unsafe.Pointer(ctx)).Fbuffer
			}
			if libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(closer1)).Fflags)&int32(MD_MARK_VALIDPERMISSIVEAUTOLINK) != 0 {
				ret = md_enter_leave_span_a(tls, ctx, libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(mark)).Fflags)&int32(MD_MARK_OPENER), int32(MD_SPAN_A), dest, dest_size, int32(TRUE), libc.UintptrFromInt32(0), uint32(0))
				if ret < 0 {
					goto abort
				}
			}
			goto _20
		_17:
			; /* Entity. */
			if (*MD_MARK)(unsafe.Pointer(mark)).Fend-(*MD_MARK)(unsafe.Pointer(mark)).Fbeg > uint32(0) {
				ret = (*(*func(*libc.TLS, MD_TEXTTYPE1, uintptr, MD_SIZE, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Ftext})))(tls, int32(MD_TEXT_ENTITY), (*MD_CTX)(unsafe.Pointer(ctx)).Ftext+uintptr((*MD_MARK)(unsafe.Pointer(mark)).Fbeg), (*MD_MARK)(unsafe.Pointer(mark)).Fend-(*MD_MARK)(unsafe.Pointer(mark)).Fbeg, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
				if ret != 0 {
					if (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log != libc.UintptrFromInt32(0) {
						(*(*func(*libc.TLS, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log})))(tls, __ccgo_ts+238, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
					}
					goto abort
				}
			}
			goto _20
		_18:
			;
			if int32(1) > 0 {
				ret = (*(*func(*libc.TLS, MD_TEXTTYPE1, uintptr, MD_SIZE, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Ftext})))(tls, int32(MD_TEXT_NULLCHAR), __ccgo_ts, libc.Uint32FromInt32(libc.Int32FromInt32(1)), (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
				if ret != 0 {
					if (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log != libc.UintptrFromInt32(0) {
						(*(*func(*libc.TLS, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log})))(tls, __ccgo_ts+238, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
					}
					goto abort
				}
			}
			goto _20
		_19:
			;
			goto abort
		_20:
			;
			off = (*MD_MARK)(unsafe.Pointer(mark)).Fend
			/* Move to next resolved mark. */
			prev_mark = mark
			mark += 20
			for !(libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(mark)).Fflags)&libc.Int32FromInt32(MD_MARK_RESOLVED) != 0) || (*MD_MARK)(unsafe.Pointer(mark)).Fbeg < off {
				mark += 20
			}
		}
		/* If reached end of line, move to next one. */
		if off >= (*MD_LINE)(unsafe.Pointer(line)).Fend {
			/* If it is the last line, we are done. */
			if off >= end {
				break
			}
			if text_type == int32(MD_TEXT_CODE) || text_type == int32(MD_TEXT_LATEXMATH) {
				if !(prev_mark != libc.UintptrFromInt32(0)) {
					libc.X__builtin_unreachable(tls)
				}
				if !((libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(prev_mark)).Fch) == int32('`') || libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(prev_mark)).Fch) == int32('$')) && libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(prev_mark)).Fflags)&int32(MD_MARK_OPENER) != 0) {
					libc.X__builtin_unreachable(tls)
				}
				if !((libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(mark)).Fch) == int32('`') || libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(mark)).Fch) == int32('$')) && libc.Int32FromUint8((*MD_MARK)(unsafe.Pointer(mark)).Fflags)&int32(MD_MARK_CLOSER) != 0) {
					libc.X__builtin_unreachable(tls)
				}
				/* Inside a code span, trailing line whitespace has to be
				 * outputted. */
				tmp = off
				for off < (*MD_CTX)(unsafe.Pointer(ctx)).Fsize && (libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32(' ') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('\t')) {
					off = off + 1
				}
				if off > tmp {
					if off-tmp > uint32(0) {
						ret = (*(*func(*libc.TLS, MD_TEXTTYPE1, uintptr, MD_SIZE, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Ftext})))(tls, text_type, (*MD_CTX)(unsafe.Pointer(ctx)).Ftext+uintptr(tmp), off-tmp, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
						if ret != 0 {
							if (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log != libc.UintptrFromInt32(0) {
								(*(*func(*libc.TLS, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log})))(tls, __ccgo_ts+238, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
							}
							goto abort
						}
					}
				}
				/* and new lines are transformed into single spaces. */
				if off == (*MD_LINE)(unsafe.Pointer(line)).Fend {
					if int32(1) > 0 {
						ret = (*(*func(*libc.TLS, MD_TEXTTYPE1, uintptr, MD_SIZE, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Ftext})))(tls, text_type, __ccgo_ts+268, libc.Uint32FromInt32(libc.Int32FromInt32(1)), (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
						if ret != 0 {
							if (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log != libc.UintptrFromInt32(0) {
								(*(*func(*libc.TLS, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log})))(tls, __ccgo_ts+238, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
							}
							goto abort
						}
					}
				}
			} else {
				if text_type == int32(MD_TEXT_HTML) {
					/* Inside raw HTML, we output the new line verbatim, including
					 * any trailing spaces. */
					tmp = off
					for tmp < end && (libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(tmp)))) == int32(' ') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(tmp)))) == int32('\t')) {
						tmp = tmp + 1
					}
					if tmp > off {
						if tmp-off > uint32(0) {
							ret = (*(*func(*libc.TLS, MD_TEXTTYPE1, uintptr, MD_SIZE, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Ftext})))(tls, int32(MD_TEXT_HTML), (*MD_CTX)(unsafe.Pointer(ctx)).Ftext+uintptr(off), tmp-off, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
							if ret != 0 {
								if (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log != libc.UintptrFromInt32(0) {
									(*(*func(*libc.TLS, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log})))(tls, __ccgo_ts+238, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
								}
								goto abort
							}
						}
					}
					if int32(1) > 0 {
						ret = (*(*func(*libc.TLS, MD_TEXTTYPE1, uintptr, MD_SIZE, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Ftext})))(tls, int32(MD_TEXT_HTML), __ccgo_ts+286, libc.Uint32FromInt32(libc.Int32FromInt32(1)), (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
						if ret != 0 {
							if (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log != libc.UintptrFromInt32(0) {
								(*(*func(*libc.TLS, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log})))(tls, __ccgo_ts+238, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
							}
							goto abort
						}
					}
				} else {
					/* Output soft or hard line break. */
					break_type = int32(MD_TEXT_SOFTBR)
					if text_type == int32(MD_TEXT_NORMAL) {
						if (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fflags&uint32(MD_FLAG_HARD_SOFT_BREAKS) != 0 {
							break_type = int32(MD_TEXT_BR)
						} else {
							if enforce_hardbreak != 0 {
								break_type = int32(MD_TEXT_BR)
							} else {
								if libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr((*MD_LINE)(unsafe.Pointer(line)).Fend)))) == int32(' ') && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr((*MD_LINE)(unsafe.Pointer(line)).Fend+libc.Uint32FromInt32(1))))) == int32(' ') {
									break_type = int32(MD_TEXT_BR)
								}
							}
						}
					}
					if int32(1) > 0 {
						ret = (*(*func(*libc.TLS, MD_TEXTTYPE1, uintptr, MD_SIZE, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Ftext})))(tls, break_type, __ccgo_ts+286, libc.Uint32FromInt32(libc.Int32FromInt32(1)), (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
						if ret != 0 {
							if (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log != libc.UintptrFromInt32(0) {
								(*(*func(*libc.TLS, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log})))(tls, __ccgo_ts+238, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
							}
							goto abort
						}
					}
				}
			}
			/* Move to the next line. */
			line += 8
			off = (*MD_LINE)(unsafe.Pointer(line)).Fbeg
			enforce_hardbreak = 0
		}
	}
	goto abort
abort:
	;
	return ret
	return r
}

/***************************
 ***  Processing Tables  ***
 ***************************/

func md_analyze_table_alignment(tls *libc.TLS, ctx uintptr, beg MD_OFFSET, end MD_OFFSET, align uintptr, n_align int32) {
	var index int32
	var off MD_OFFSET
	_, _ = index, off
	off = beg
	for n_align > 0 {
		index = 0 /* index into align_map[] */
		for libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) != int32('-') {
			off = off + 1
		}
		if off > beg && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off-libc.Uint32FromInt32(1))))) == int32(':') {
			index = index | int32(1)
		}
		for off < end && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('-') {
			off = off + 1
		}
		if off < end && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32(':') {
			index = index | int32(2)
		}
		*(*MD_ALIGN1)(__ccgo_ptr(align)) = align_map[index]
		align += 4
		n_align = n_align - 1
	}
}

var align_map = [4]MD_ALIGN1{
	1: int32(MD_ALIGN_LEFT),
	2: int32(MD_ALIGN_RIGHT),
	3: int32(MD_ALIGN_CENTER),
}

func md_process_table_cell(tls *libc.TLS, ctx uintptr, cell_type MD_BLOCKTYPE1, align MD_ALIGN1, beg MD_OFFSET, end MD_OFFSET) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var ret int32
	var _ /* det at bp+8 */ MD_BLOCK_TD_DETAIL
	var _ /* line at bp+0 */ MD_LINE
	_ = ret
	ret = 0
	for beg < end && (libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(beg)))) == int32(' ') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(beg)))) == int32('\t') || (libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(beg)))) == int32('\v') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(beg)))) == int32('\f'))) {
		beg = beg + 1
	}
	for end > beg && (libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(end-libc.Uint32FromInt32(1))))) == int32(' ') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(end-libc.Uint32FromInt32(1))))) == int32('\t') || (libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(end-libc.Uint32FromInt32(1))))) == int32('\v') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(end-libc.Uint32FromInt32(1))))) == int32('\f'))) {
		end = end - 1
	}
	(*(*MD_BLOCK_TD_DETAIL)(__ccgo_ptr(bp + 8))).Falign = align
	(*(*MD_LINE)(__ccgo_ptr(bp))).Fbeg = beg
	(*(*MD_LINE)(__ccgo_ptr(bp))).Fend = end
	ret = (*(*func(*libc.TLS, MD_BLOCKTYPE1, uintptr, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fenter_block})))(tls, cell_type, bp+8, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
	if ret != 0 {
		if (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log != libc.UintptrFromInt32(0) {
			(*(*func(*libc.TLS, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log})))(tls, __ccgo_ts+288, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
		}
		goto abort
	}
	ret = md_process_normal_block_contents(tls, ctx, bp, uint32(1))
	if ret < 0 {
		goto abort
	}
	ret = (*(*func(*libc.TLS, MD_BLOCKTYPE1, uintptr, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fleave_block})))(tls, cell_type, bp+8, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
	if ret != 0 {
		if (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log != libc.UintptrFromInt32(0) {
			(*(*func(*libc.TLS, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log})))(tls, __ccgo_ts+325, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
		}
		goto abort
	}
	goto abort
abort:
	;
	return ret
	return r
}

func md_process_table_row(tls *libc.TLS, ctx uintptr, cell_type MD_BLOCKTYPE1, beg MD_OFFSET, end MD_OFFSET, align uintptr, col_count int32) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var i, j, k, n, ret, v1 int32
	var mark, pipe_offs uintptr
	var _ /* line at bp+0 */ MD_LINE
	_, _, _, _, _, _, _, _ = i, j, k, mark, n, pipe_offs, ret, v1
	pipe_offs = libc.UintptrFromInt32(0)
	ret = 0
	(*(*MD_LINE)(__ccgo_ptr(bp))).Fbeg = beg
	(*(*MD_LINE)(__ccgo_ptr(bp))).Fend = end
	/* Break the line into table cells by identifying pipe characters who
	 * form the cell boundary. */
	ret = md_analyze_inlines(tls, ctx, bp, uint32(1), int32(TRUE))
	if ret < 0 {
		goto abort
	}
	/* We have to remember the cell boundaries in local buffer because
	 * ctx->marks[] shall be reused during cell contents processing. */
	n = (*MD_CTX)(unsafe.Pointer(ctx)).Fn_table_cell_boundaries + int32(2)
	pipe_offs = libc.Xmalloc(tls, libc.Uint64FromInt32(n)*uint64(4))
	if pipe_offs == libc.UintptrFromInt32(0) {
		if (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log != libc.UintptrFromInt32(0) {
			(*(*func(*libc.TLS, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log})))(tls, __ccgo_ts+1, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
		}
		ret = -int32(1)
		goto abort
	}
	j = 0
	v1 = j
	j = j + 1
	*(*MD_OFFSET)(__ccgo_ptr(pipe_offs + uintptr(v1)*4)) = beg
	i = (*MD_CTX)(unsafe.Pointer(ctx)).Ftable_cell_boundaries_head
	for {
		if !(i >= 0) {
			break
		}
		mark = (*MD_CTX)(unsafe.Pointer(ctx)).Fmarks + uintptr(i)*20
		v1 = j
		j = j + 1
		*(*MD_OFFSET)(__ccgo_ptr(pipe_offs + uintptr(v1)*4)) = (*MD_MARK)(unsafe.Pointer(mark)).Fend
		goto _2
	_2:
		;
		i = (*(*MD_MARK)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Fmarks + uintptr(i)*20))).Fnext
	}
	v1 = j
	j = j + 1
	*(*MD_OFFSET)(__ccgo_ptr(pipe_offs + uintptr(v1)*4)) = end + uint32(1)
	/* Process cells. */
	ret = (*(*func(*libc.TLS, MD_BLOCKTYPE1, uintptr, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fenter_block})))(tls, int32(MD_BLOCK_TR), libc.UintptrFromInt32(0), (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
	if ret != 0 {
		if (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log != libc.UintptrFromInt32(0) {
			(*(*func(*libc.TLS, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log})))(tls, __ccgo_ts+288, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
		}
		goto abort
	}
	k = 0
	i = 0
	for {
		if !(i < j-int32(1) && k < col_count) {
			break
		}
		if *(*MD_OFFSET)(__ccgo_ptr(pipe_offs + uintptr(i)*4)) < *(*MD_OFFSET)(__ccgo_ptr(pipe_offs + uintptr(i+int32(1))*4))-uint32(1) {
			v1 = k
			k = k + 1
			ret = md_process_table_cell(tls, ctx, cell_type, *(*MD_ALIGN1)(__ccgo_ptr(align + uintptr(v1)*4)), *(*MD_OFFSET)(__ccgo_ptr(pipe_offs + uintptr(i)*4)), *(*MD_OFFSET)(__ccgo_ptr(pipe_offs + uintptr(i+int32(1))*4))-uint32(1))
			if ret < 0 {
				goto abort
			}
		}
		goto _5
	_5:
		;
		i = i + 1
	}
	/* Make sure we call enough table cells even if the current table contains
	 * too few of them. */
	for k < col_count {
		v1 = k
		k = k + 1
		ret = md_process_table_cell(tls, ctx, cell_type, *(*MD_ALIGN1)(__ccgo_ptr(align + uintptr(v1)*4)), uint32(0), uint32(0))
		if ret < 0 {
			goto abort
		}
	}
	ret = (*(*func(*libc.TLS, MD_BLOCKTYPE1, uintptr, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fleave_block})))(tls, int32(MD_BLOCK_TR), libc.UintptrFromInt32(0), (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
	if ret != 0 {
		if (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log != libc.UintptrFromInt32(0) {
			(*(*func(*libc.TLS, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log})))(tls, __ccgo_ts+325, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
		}
		goto abort
	}
	goto abort
abort:
	;
	libc.Xfree(tls, pipe_offs)
	(*MD_CTX)(unsafe.Pointer(ctx)).Ftable_cell_boundaries_head = -int32(1)
	(*MD_CTX)(unsafe.Pointer(ctx)).Ftable_cell_boundaries_tail = -int32(1)
	return ret
}

func md_process_table_block_contents(tls *libc.TLS, ctx uintptr, col_count int32, lines uintptr, n_lines MD_SIZE) (r int32) {
	var align uintptr
	var line_index MD_SIZE
	var ret int32
	_, _, _ = align, line_index, ret
	ret = 0
	/* At least two lines have to be present: The column headers and the line
	 * with the underlines. */
	if !(n_lines >= libc.Uint32FromInt32(2)) {
		libc.X__builtin_unreachable(tls)
	}
	align = libc.Xmalloc(tls, libc.Uint64FromInt32(col_count)*uint64(4))
	if align == libc.UintptrFromInt32(0) {
		if (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log != libc.UintptrFromInt32(0) {
			(*(*func(*libc.TLS, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log})))(tls, __ccgo_ts+1, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
		}
		ret = -int32(1)
		goto abort
	}
	md_analyze_table_alignment(tls, ctx, (*(*MD_LINE)(__ccgo_ptr(lines + 1*8))).Fbeg, (*(*MD_LINE)(__ccgo_ptr(lines + 1*8))).Fend, align, col_count)
	ret = (*(*func(*libc.TLS, MD_BLOCKTYPE1, uintptr, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fenter_block})))(tls, int32(MD_BLOCK_THEAD), libc.UintptrFromInt32(0), (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
	if ret != 0 {
		if (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log != libc.UintptrFromInt32(0) {
			(*(*func(*libc.TLS, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log})))(tls, __ccgo_ts+288, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
		}
		goto abort
	}
	ret = md_process_table_row(tls, ctx, int32(MD_BLOCK_TH), (*(*MD_LINE)(__ccgo_ptr(lines))).Fbeg, (*(*MD_LINE)(__ccgo_ptr(lines))).Fend, align, col_count)
	if ret < 0 {
		goto abort
	}
	ret = (*(*func(*libc.TLS, MD_BLOCKTYPE1, uintptr, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fleave_block})))(tls, int32(MD_BLOCK_THEAD), libc.UintptrFromInt32(0), (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
	if ret != 0 {
		if (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log != libc.UintptrFromInt32(0) {
			(*(*func(*libc.TLS, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log})))(tls, __ccgo_ts+325, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
		}
		goto abort
	}
	if n_lines > uint32(2) {
		ret = (*(*func(*libc.TLS, MD_BLOCKTYPE1, uintptr, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fenter_block})))(tls, int32(MD_BLOCK_TBODY), libc.UintptrFromInt32(0), (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
		if ret != 0 {
			if (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log != libc.UintptrFromInt32(0) {
				(*(*func(*libc.TLS, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log})))(tls, __ccgo_ts+288, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
			}
			goto abort
		}
		line_index = uint32(2)
		for {
			if !(line_index < n_lines) {
				break
			}
			ret = md_process_table_row(tls, ctx, int32(MD_BLOCK_TD), (*(*MD_LINE)(__ccgo_ptr(lines + uintptr(line_index)*8))).Fbeg, (*(*MD_LINE)(__ccgo_ptr(lines + uintptr(line_index)*8))).Fend, align, col_count)
			if ret < 0 {
				goto abort
			}
			goto _1
		_1:
			;
			line_index = line_index + 1
		}
		ret = (*(*func(*libc.TLS, MD_BLOCKTYPE1, uintptr, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fleave_block})))(tls, int32(MD_BLOCK_TBODY), libc.UintptrFromInt32(0), (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
		if ret != 0 {
			if (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log != libc.UintptrFromInt32(0) {
				(*(*func(*libc.TLS, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log})))(tls, __ccgo_ts+325, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
			}
			goto abort
		}
	}
	goto abort
abort:
	;
	libc.Xfree(tls, align)
	return ret
}

func md_process_normal_block_contents(tls *libc.TLS, ctx uintptr, lines uintptr, n_lines MD_SIZE) (r int32) {
	var i, ret int32
	_, _ = i, ret
	ret = md_analyze_inlines(tls, ctx, lines, n_lines, FALSE)
	if ret < 0 {
		goto abort
	}
	ret = md_process_inlines(tls, ctx, lines, n_lines)
	if ret < 0 {
		goto abort
	}
	goto abort
abort:
	;
	/* Free any temporary memory blocks stored within some dummy marks. */
	i = (*MD_CTX)(unsafe.Pointer(ctx)).Fptr_stack.Ftop
	for {
		if !(i >= 0) {
			break
		}
		libc.Xfree(tls, md_mark_get_ptr(tls, ctx, i))
		goto _1
	_1:
		;
		i = (*(*MD_MARK)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Fmarks + uintptr(i)*20))).Fnext
	}
	(*MD_CTX)(unsafe.Pointer(ctx)).Fptr_stack.Ftop = -int32(1)
	return ret
}

func md_process_verbatim_block_contents(tls *libc.TLS, ctx uintptr, text_type MD_TEXTTYPE1, lines uintptr, n_lines MD_SIZE) (r int32) {
	var indent, ret int32
	var line uintptr
	var line_index MD_SIZE
	_, _, _, _ = indent, line, line_index, ret
	ret = 0
	line_index = uint32(0)
	for {
		if !(line_index < n_lines) {
			break
		}
		line = lines + uintptr(line_index)*12
		indent = libc.Int32FromUint32((*MD_VERBATIMLINE)(unsafe.Pointer(line)).Findent)
		if !(indent >= libc.Int32FromInt32(0)) {
			libc.X__builtin_unreachable(tls)
		}
		/* Output code indentation. */
		for indent > libc.Int32FromUint32(indent_chunk_size) {
			if indent_chunk_size > uint32(0) {
				ret = (*(*func(*libc.TLS, MD_TEXTTYPE1, uintptr, MD_SIZE, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Ftext})))(tls, text_type, uintptr(unsafe.Pointer(&indent_chunk_str)), indent_chunk_size, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
				if ret != 0 {
					if (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log != libc.UintptrFromInt32(0) {
						(*(*func(*libc.TLS, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log})))(tls, __ccgo_ts+238, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
					}
					goto abort
				}
			}
			indent = libc.Int32FromUint32(uint32(indent) - indent_chunk_size)
		}
		if indent > 0 {
			if indent > 0 {
				ret = (*(*func(*libc.TLS, MD_TEXTTYPE1, uintptr, MD_SIZE, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Ftext})))(tls, text_type, uintptr(unsafe.Pointer(&indent_chunk_str)), libc.Uint32FromInt32(indent), (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
				if ret != 0 {
					if (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log != libc.UintptrFromInt32(0) {
						(*(*func(*libc.TLS, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log})))(tls, __ccgo_ts+238, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
					}
					goto abort
				}
			}
		}
		/* Output the code line itself. */
		if (*MD_VERBATIMLINE)(unsafe.Pointer(line)).Fend-(*MD_VERBATIMLINE)(unsafe.Pointer(line)).Fbeg > uint32(0) {
			ret = md_text_with_null_replacement(tls, ctx, text_type, (*MD_CTX)(unsafe.Pointer(ctx)).Ftext+uintptr((*MD_VERBATIMLINE)(unsafe.Pointer(line)).Fbeg), (*MD_VERBATIMLINE)(unsafe.Pointer(line)).Fend-(*MD_VERBATIMLINE)(unsafe.Pointer(line)).Fbeg)
			if ret != 0 {
				if (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log != libc.UintptrFromInt32(0) {
					(*(*func(*libc.TLS, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log})))(tls, __ccgo_ts+238, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
				}
				goto abort
			}
		}
		/* Enforce end-of-line. */
		if int32(1) > 0 {
			ret = (*(*func(*libc.TLS, MD_TEXTTYPE1, uintptr, MD_SIZE, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Ftext})))(tls, text_type, __ccgo_ts+286, libc.Uint32FromInt32(libc.Int32FromInt32(1)), (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
			if ret != 0 {
				if (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log != libc.UintptrFromInt32(0) {
					(*(*func(*libc.TLS, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log})))(tls, __ccgo_ts+238, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
				}
				goto abort
			}
		}
		goto _1
	_1:
		;
		line_index = line_index + 1
	}
	goto abort
abort:
	;
	return ret
	return r
}

var indent_chunk_str = [17]MD_CHAR{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '}

var indent_chunk_size = uint32(libc.Uint64FromInt64(17)/libc.Uint64FromInt64(1) - libc.Uint64FromInt32(1))

func md_process_code_block_contents(tls *libc.TLS, ctx uintptr, is_fenced int32, lines uintptr, n_lines MD_SIZE) (r int32) {
	if is_fenced != 0 {
		/* Skip the first line in case of fenced code: It is the fence.
		 * (Only the starting fence is present due to logic in md_analyze_line().) */
		lines += 12
		n_lines = n_lines - 1
	} else {
		/* Ignore blank lines at start/end of indented code block. */
		for n_lines > uint32(0) && (*(*MD_VERBATIMLINE)(__ccgo_ptr(lines))).Fbeg == (*(*MD_VERBATIMLINE)(__ccgo_ptr(lines))).Fend {
			lines += 12
			n_lines = n_lines - 1
		}
		for n_lines > uint32(0) && (*(*MD_VERBATIMLINE)(__ccgo_ptr(lines + uintptr(n_lines-uint32(1))*12))).Fbeg == (*(*MD_VERBATIMLINE)(__ccgo_ptr(lines + uintptr(n_lines-uint32(1))*12))).Fend {
			n_lines = n_lines - 1
		}
	}
	if n_lines == uint32(0) {
		return 0
	}
	return md_process_verbatim_block_contents(tls, ctx, int32(MD_TEXT_CODE), lines, n_lines)
}

func md_setup_fenced_code_detail(tls *libc.TLS, ctx uintptr, block uintptr, det uintptr, info_build uintptr, lang_build uintptr) (r int32) {
	var beg, end, lang_end MD_OFFSET
	var fence_ch MD_CHAR
	var fence_line uintptr
	var ret int32
	_, _, _, _, _, _ = beg, end, fence_ch, fence_line, lang_end, ret
	fence_line = block + libc.UintptrFromInt32(1)*8
	beg = (*MD_VERBATIMLINE)(unsafe.Pointer(fence_line)).Fbeg
	end = (*MD_VERBATIMLINE)(unsafe.Pointer(fence_line)).Fend
	fence_ch = *(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr((*MD_VERBATIMLINE)(unsafe.Pointer(fence_line)).Fbeg)))
	ret = 0
	/* Skip the fence itself. */
	for beg < (*MD_CTX)(unsafe.Pointer(ctx)).Fsize && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(beg)))) == libc.Int32FromUint8(fence_ch) {
		beg = beg + 1
	}
	/* Trim initial spaces. */
	for beg < (*MD_CTX)(unsafe.Pointer(ctx)).Fsize && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(beg)))) == int32(' ') {
		beg = beg + 1
	}
	/* Trim trailing spaces. */
	for end > beg && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(end-libc.Uint32FromInt32(1))))) == int32(' ') {
		end = end - 1
	}
	/* Build info string attribute. */
	ret = md_build_attribute(tls, ctx, (*MD_CTX)(unsafe.Pointer(ctx)).Ftext+uintptr(beg), end-beg, uint32(0), det, info_build)
	if ret < 0 {
		goto abort
	}
	/* Build info string attribute. */
	lang_end = beg
	for lang_end < end && !(libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(lang_end)))) == int32(' ') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(lang_end)))) == int32('\t') || (libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(lang_end)))) == int32('\v') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(lang_end)))) == int32('\f'))) {
		lang_end = lang_end + 1
	}
	ret = md_build_attribute(tls, ctx, (*MD_CTX)(unsafe.Pointer(ctx)).Ftext+uintptr(beg), lang_end-beg, uint32(0), det+32, lang_build)
	if ret < 0 {
		goto abort
	}
	(*MD_BLOCK_CODE_DETAIL)(unsafe.Pointer(det)).Ffence_char = fence_ch
	goto abort
abort:
	;
	return ret
	return r
}

func md_process_leaf_block(tls *libc.TLS, ctx uintptr, block uintptr) (r int32) {
	bp := tls.Alloc(176)
	defer tls.Free(176)
	var clean_fence_code_detail, is_in_tight_list, ret int32
	var _ /* det at bp+0 */ struct {
		Fcode        [0]MD_BLOCK_CODE_DETAIL
		Ftable       [0]MD_BLOCK_TABLE_DETAIL
		Fheader      MD_BLOCK_H_DETAIL
		F__ccgo_pad3 [68]byte
	}
	var _ /* info_build at bp+72 */ MD_ATTRIBUTE_BUILD
	var _ /* lang_build at bp+120 */ MD_ATTRIBUTE_BUILD
	_, _, _ = clean_fence_code_detail, is_in_tight_list, ret
	clean_fence_code_detail = FALSE
	ret = 0
	libc.Xmemset(tls, bp, 0, uint64(72))
	if (*MD_CTX)(unsafe.Pointer(ctx)).Fn_containers == 0 {
		is_in_tight_list = FALSE
	} else {
		is_in_tight_list = libc.BoolInt32(!(int32(uint32(*(*uint16)(unsafe.Pointer((*MD_CTX)(unsafe.Pointer(ctx)).Fcontainers + uintptr((*MD_CTX)(unsafe.Pointer(ctx)).Fn_containers-int32(1))*28 + 4))&0xff>>0)) != 0))
	}
	switch MD_BLOCKTYPE1(*(*uint32)(unsafe.Pointer(block + 0)) & 0xff >> 0) {
	case int32(MD_BLOCK_H):
		(*(*struct {
			Fcode        [0]MD_BLOCK_CODE_DETAIL
			Ftable       [0]MD_BLOCK_TABLE_DETAIL
			Fheader      MD_BLOCK_H_DETAIL
			F__ccgo_pad3 [68]byte
		})(__ccgo_ptr(bp))).Fheader.Flevel = libc.Uint32FromInt32(int32(*(*uint32)(unsafe.Pointer(block + 0)) & 0xffff0000 >> 16))
	case int32(MD_BLOCK_CODE):
		/* For fenced code block, we may need to set the info string. */
		if int32(*(*uint32)(unsafe.Pointer(block + 0))&0xffff0000>>16) != 0 {
			libc.Xmemset(tls, bp, 0, uint64(72))
			clean_fence_code_detail = int32(TRUE)
			ret = md_setup_fenced_code_detail(tls, ctx, block, bp, bp+72, bp+120)
			if ret < 0 {
				goto abort
			}
		}
	case int32(MD_BLOCK_TABLE):
		(*(*MD_BLOCK_TABLE_DETAIL)(unsafe.Pointer(&*(*struct {
			Fcode        [0]MD_BLOCK_CODE_DETAIL
			Ftable       [0]MD_BLOCK_TABLE_DETAIL
			Fheader      MD_BLOCK_H_DETAIL
			F__ccgo_pad3 [68]byte
		})(__ccgo_ptr(bp))))).Fcol_count = libc.Uint32FromInt32(int32(*(*uint32)(unsafe.Pointer(block + 0)) & 0xffff0000 >> 16))
		(*(*MD_BLOCK_TABLE_DETAIL)(unsafe.Pointer(&*(*struct {
			Fcode        [0]MD_BLOCK_CODE_DETAIL
			Ftable       [0]MD_BLOCK_TABLE_DETAIL
			Fheader      MD_BLOCK_H_DETAIL
			F__ccgo_pad3 [68]byte
		})(__ccgo_ptr(bp))))).Fhead_row_count = uint32(1)
		(*(*MD_BLOCK_TABLE_DETAIL)(unsafe.Pointer(&*(*struct {
			Fcode        [0]MD_BLOCK_CODE_DETAIL
			Ftable       [0]MD_BLOCK_TABLE_DETAIL
			Fheader      MD_BLOCK_H_DETAIL
			F__ccgo_pad3 [68]byte
		})(__ccgo_ptr(bp))))).Fbody_row_count = (*MD_BLOCK)(unsafe.Pointer(block)).Fn_lines - uint32(2)
	default:
		/* Noop. */
		break
	}
	if !(is_in_tight_list != 0) || MD_BLOCKTYPE1(*(*uint32)(unsafe.Pointer(block + 0))&0xff>>0) != int32(MD_BLOCK_P) {
		ret = (*(*func(*libc.TLS, MD_BLOCKTYPE1, uintptr, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fenter_block})))(tls, MD_BLOCKTYPE1(*(*uint32)(unsafe.Pointer(block + 0))&0xff>>0), bp, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
		if ret != 0 {
			if (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log != libc.UintptrFromInt32(0) {
				(*(*func(*libc.TLS, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log})))(tls, __ccgo_ts+288, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
			}
			goto abort
		}
	}
	/* Process the block contents accordingly to is type. */
	switch MD_BLOCKTYPE1(*(*uint32)(unsafe.Pointer(block + 0)) & 0xff >> 0) {
	case int32(MD_BLOCK_HR):
		goto _1
	case int32(MD_BLOCK_CODE):
		goto _2
	case int32(MD_BLOCK_HTML):
		goto _3
	case int32(MD_BLOCK_TABLE):
		goto _4
	default:
		goto _5
	}
	goto _6
_1:
	;
	/* noop */
	goto _6
_2:
	;
_9:
	;
	ret = md_process_code_block_contents(tls, ctx, libc.BoolInt32(int32(*(*uint32)(unsafe.Pointer(block + 0))&0xffff0000>>16) != 0), block+libc.UintptrFromInt32(1)*8, (*MD_BLOCK)(unsafe.Pointer(block)).Fn_lines)
	if ret < 0 {
		goto abort
	}
	goto _8
_8:
	;
	if 0 != 0 {
		goto _9
	}
	goto _7
_7:
	;
	goto _6
_3:
	;
	ret = md_process_verbatim_block_contents(tls, ctx, int32(MD_TEXT_HTML), block+libc.UintptrFromInt32(1)*8, (*MD_BLOCK)(unsafe.Pointer(block)).Fn_lines)
	if ret < 0 {
		goto abort
	}
	goto _6
_4:
	;
	ret = md_process_table_block_contents(tls, ctx, int32(*(*uint32)(unsafe.Pointer(block + 0))&0xffff0000>>16), block+libc.UintptrFromInt32(1)*8, (*MD_BLOCK)(unsafe.Pointer(block)).Fn_lines)
	if ret < 0 {
		goto abort
	}
	goto _6
_5:
	;
	ret = md_process_normal_block_contents(tls, ctx, block+libc.UintptrFromInt32(1)*8, (*MD_BLOCK)(unsafe.Pointer(block)).Fn_lines)
	if ret < 0 {
		goto abort
	}
	goto _6
_6:
	;
	if !(is_in_tight_list != 0) || MD_BLOCKTYPE1(*(*uint32)(unsafe.Pointer(block + 0))&0xff>>0) != int32(MD_BLOCK_P) {
		ret = (*(*func(*libc.TLS, MD_BLOCKTYPE1, uintptr, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fleave_block})))(tls, MD_BLOCKTYPE1(*(*uint32)(unsafe.Pointer(block + 0))&0xff>>0), bp, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
		if ret != 0 {
			if (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log != libc.UintptrFromInt32(0) {
				(*(*func(*libc.TLS, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log})))(tls, __ccgo_ts+325, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
			}
			goto abort
		}
	}
	goto abort
abort:
	;
	if clean_fence_code_detail != 0 {
		md_free_attribute(tls, ctx, bp+72)
		md_free_attribute(tls, ctx, bp+120)
	}
	return ret
}

func md_process_all_blocks(tls *libc.TLS, ctx uintptr) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var block uintptr
	var byte_off, ret, v1 int32
	var _ /* det at bp+0 */ struct {
		Fol          [0]MD_BLOCK_OL_DETAIL
		Fli          [0]MD_BLOCK_LI_DETAIL
		Ful          MD_BLOCK_UL_DETAIL
		F__ccgo_pad3 [4]byte
	}
	_, _, _, _ = block, byte_off, ret, v1
	byte_off = 0
	ret = 0
	/* ctx->containers now is not needed for detection of lists and list items
	 * so we reuse it for tracking what lists are loose or tight. We rely
	 * on the fact the vector is large enough to hold the deepest nesting
	 * level of lists. */
	(*MD_CTX)(unsafe.Pointer(ctx)).Fn_containers = 0
	for byte_off < (*MD_CTX)(unsafe.Pointer(ctx)).Fn_block_bytes {
		block = (*MD_CTX)(unsafe.Pointer(ctx)).Fblock_bytes + uintptr(byte_off)
		switch MD_BLOCKTYPE1(*(*uint32)(unsafe.Pointer(block + 0)) & 0xff >> 0) {
		case int32(MD_BLOCK_UL):
			if int32(*(*uint32)(unsafe.Pointer(block + 0))&0xff00>>8)&int32(MD_BLOCK_LOOSE_LIST) != 0 {
				v1 = FALSE
			} else {
				v1 = int32(TRUE)
			}
			(*(*struct {
				Fol          [0]MD_BLOCK_OL_DETAIL
				Fli          [0]MD_BLOCK_LI_DETAIL
				Ful          MD_BLOCK_UL_DETAIL
				F__ccgo_pad3 [4]byte
			})(__ccgo_ptr(bp))).Ful.Fis_tight = v1
			(*(*struct {
				Fol          [0]MD_BLOCK_OL_DETAIL
				Fli          [0]MD_BLOCK_LI_DETAIL
				Ful          MD_BLOCK_UL_DETAIL
				F__ccgo_pad3 [4]byte
			})(__ccgo_ptr(bp))).Ful.Fmark = libc.Uint8FromInt32(int32(*(*uint32)(unsafe.Pointer(block + 0)) & 0xffff0000 >> 16))
		case int32(MD_BLOCK_OL):
			(*(*MD_BLOCK_OL_DETAIL)(unsafe.Pointer(&*(*struct {
				Fol          [0]MD_BLOCK_OL_DETAIL
				Fli          [0]MD_BLOCK_LI_DETAIL
				Ful          MD_BLOCK_UL_DETAIL
				F__ccgo_pad3 [4]byte
			})(__ccgo_ptr(bp))))).Fstart = (*MD_BLOCK)(unsafe.Pointer(block)).Fn_lines
			if int32(*(*uint32)(unsafe.Pointer(block + 0))&0xff00>>8)&int32(MD_BLOCK_LOOSE_LIST) != 0 {
				v1 = FALSE
			} else {
				v1 = int32(TRUE)
			}
			(*(*MD_BLOCK_OL_DETAIL)(unsafe.Pointer(&*(*struct {
				Fol          [0]MD_BLOCK_OL_DETAIL
				Fli          [0]MD_BLOCK_LI_DETAIL
				Ful          MD_BLOCK_UL_DETAIL
				F__ccgo_pad3 [4]byte
			})(__ccgo_ptr(bp))))).Fis_tight = v1
			(*(*MD_BLOCK_OL_DETAIL)(unsafe.Pointer(&*(*struct {
				Fol          [0]MD_BLOCK_OL_DETAIL
				Fli          [0]MD_BLOCK_LI_DETAIL
				Ful          MD_BLOCK_UL_DETAIL
				F__ccgo_pad3 [4]byte
			})(__ccgo_ptr(bp))))).Fmark_delimiter = libc.Uint8FromInt32(int32(*(*uint32)(unsafe.Pointer(block + 0)) & 0xffff0000 >> 16))
		case int32(MD_BLOCK_LI):
			(*(*MD_BLOCK_LI_DETAIL)(unsafe.Pointer(&*(*struct {
				Fol          [0]MD_BLOCK_OL_DETAIL
				Fli          [0]MD_BLOCK_LI_DETAIL
				Ful          MD_BLOCK_UL_DETAIL
				F__ccgo_pad3 [4]byte
			})(__ccgo_ptr(bp))))).Fis_task = libc.BoolInt32(int32(*(*uint32)(unsafe.Pointer(block + 0))&0xffff0000>>16) != 0)
			(*(*MD_BLOCK_LI_DETAIL)(unsafe.Pointer(&*(*struct {
				Fol          [0]MD_BLOCK_OL_DETAIL
				Fli          [0]MD_BLOCK_LI_DETAIL
				Ful          MD_BLOCK_UL_DETAIL
				F__ccgo_pad3 [4]byte
			})(__ccgo_ptr(bp))))).Ftask_mark = libc.Uint8FromInt32(int32(*(*uint32)(unsafe.Pointer(block + 0)) & 0xffff0000 >> 16))
			(*(*MD_BLOCK_LI_DETAIL)(unsafe.Pointer(&*(*struct {
				Fol          [0]MD_BLOCK_OL_DETAIL
				Fli          [0]MD_BLOCK_LI_DETAIL
				Ful          MD_BLOCK_UL_DETAIL
				F__ccgo_pad3 [4]byte
			})(__ccgo_ptr(bp))))).Ftask_mark_offset = (*MD_BLOCK)(unsafe.Pointer(block)).Fn_lines
		default:
			/* noop */
			break
		}
		if int32(*(*uint32)(unsafe.Pointer(block + 0))&0xff00>>8)&(libc.Int32FromInt32(MD_BLOCK_CONTAINER_OPENER)|libc.Int32FromInt32(MD_BLOCK_CONTAINER_CLOSER)) != 0 {
			if int32(*(*uint32)(unsafe.Pointer(block + 0))&0xff00>>8)&int32(MD_BLOCK_CONTAINER_CLOSER) != 0 {
				ret = (*(*func(*libc.TLS, MD_BLOCKTYPE1, uintptr, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fleave_block})))(tls, MD_BLOCKTYPE1(*(*uint32)(unsafe.Pointer(block + 0))&0xff>>0), bp, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
				if ret != 0 {
					if (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log != libc.UintptrFromInt32(0) {
						(*(*func(*libc.TLS, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log})))(tls, __ccgo_ts+325, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
					}
					goto abort
				}
				if MD_BLOCKTYPE1(*(*uint32)(unsafe.Pointer(block + 0))&0xff>>0) == int32(MD_BLOCK_UL) || MD_BLOCKTYPE1(*(*uint32)(unsafe.Pointer(block + 0))&0xff>>0) == int32(MD_BLOCK_OL) || MD_BLOCKTYPE1(*(*uint32)(unsafe.Pointer(block + 0))&0xff>>0) == int32(MD_BLOCK_QUOTE) {
					(*MD_CTX)(unsafe.Pointer(ctx)).Fn_containers = (*MD_CTX)(unsafe.Pointer(ctx)).Fn_containers - 1
				}
			}
			if int32(*(*uint32)(unsafe.Pointer(block + 0))&0xff00>>8)&int32(MD_BLOCK_CONTAINER_OPENER) != 0 {
				ret = (*(*func(*libc.TLS, MD_BLOCKTYPE1, uintptr, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fenter_block})))(tls, MD_BLOCKTYPE1(*(*uint32)(unsafe.Pointer(block + 0))&0xff>>0), bp, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
				if ret != 0 {
					if (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log != libc.UintptrFromInt32(0) {
						(*(*func(*libc.TLS, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log})))(tls, __ccgo_ts+288, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
					}
					goto abort
				}
				if MD_BLOCKTYPE1(*(*uint32)(unsafe.Pointer(block + 0))&0xff>>0) == int32(MD_BLOCK_UL) || MD_BLOCKTYPE1(*(*uint32)(unsafe.Pointer(block + 0))&0xff>>0) == int32(MD_BLOCK_OL) {
					libc.SetBitFieldPtr16Uint32((*MD_CTX)(unsafe.Pointer(ctx)).Fcontainers+uintptr((*MD_CTX)(unsafe.Pointer(ctx)).Fn_containers)*28+4, libc.Uint32FromInt32(int32(*(*uint32)(unsafe.Pointer(block + 0))&0xff00>>8)&libc.Int32FromInt32(MD_BLOCK_LOOSE_LIST)), 0, 0xff)
					(*MD_CTX)(unsafe.Pointer(ctx)).Fn_containers = (*MD_CTX)(unsafe.Pointer(ctx)).Fn_containers + 1
				} else {
					if MD_BLOCKTYPE1(*(*uint32)(unsafe.Pointer(block + 0))&0xff>>0) == int32(MD_BLOCK_QUOTE) {
						/* This causes that any text in a block quote, even if
						 * nested inside a tight list item, is wrapped with
						 * <p>...</p>. */
						libc.SetBitFieldPtr16Uint32((*MD_CTX)(unsafe.Pointer(ctx)).Fcontainers+uintptr((*MD_CTX)(unsafe.Pointer(ctx)).Fn_containers)*28+4, libc.Uint32FromInt32(TRUE), 0, 0xff)
						(*MD_CTX)(unsafe.Pointer(ctx)).Fn_containers = (*MD_CTX)(unsafe.Pointer(ctx)).Fn_containers + 1
					}
				}
			}
		} else {
			ret = md_process_leaf_block(tls, ctx, block)
			if ret < 0 {
				goto abort
			}
			if MD_BLOCKTYPE1(*(*uint32)(unsafe.Pointer(block + 0))&0xff>>0) == int32(MD_BLOCK_CODE) || MD_BLOCKTYPE1(*(*uint32)(unsafe.Pointer(block + 0))&0xff>>0) == int32(MD_BLOCK_HTML) {
				byte_off = libc.Int32FromUint64(uint64(byte_off) + uint64((*MD_BLOCK)(unsafe.Pointer(block)).Fn_lines)*libc.Uint64FromInt64(12))
			} else {
				byte_off = libc.Int32FromUint64(uint64(byte_off) + uint64((*MD_BLOCK)(unsafe.Pointer(block)).Fn_lines)*libc.Uint64FromInt64(8))
			}
		}
		byte_off = libc.Int32FromUint64(uint64(byte_off) + libc.Uint64FromInt64(8))
	}
	(*MD_CTX)(unsafe.Pointer(ctx)).Fn_block_bytes = 0
	goto abort
abort:
	;
	return ret
	return r
}

/************************************
 ***  Grouping Lines into Blocks  ***
 ************************************/

func md_push_block_bytes(tls *libc.TLS, ctx uintptr, n_bytes int32) (r uintptr) {
	var new_block_bytes, ptr uintptr
	var off_current_block MD_OFFSET
	var v1 int32
	_, _, _, _ = new_block_bytes, off_current_block, ptr, v1
	if (*MD_CTX)(unsafe.Pointer(ctx)).Fn_block_bytes+n_bytes > (*MD_CTX)(unsafe.Pointer(ctx)).Falloc_block_bytes {
		if (*MD_CTX)(unsafe.Pointer(ctx)).Falloc_block_bytes > 0 {
			v1 = (*MD_CTX)(unsafe.Pointer(ctx)).Falloc_block_bytes + (*MD_CTX)(unsafe.Pointer(ctx)).Falloc_block_bytes/int32(2)
		} else {
			v1 = int32(512)
		}
		(*MD_CTX)(unsafe.Pointer(ctx)).Falloc_block_bytes = v1
		new_block_bytes = libc.Xrealloc(tls, (*MD_CTX)(unsafe.Pointer(ctx)).Fblock_bytes, libc.Uint64FromInt32((*MD_CTX)(unsafe.Pointer(ctx)).Falloc_block_bytes))
		if new_block_bytes == libc.UintptrFromInt32(0) {
			if (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log != libc.UintptrFromInt32(0) {
				(*(*func(*libc.TLS, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log})))(tls, __ccgo_ts+43, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
			}
			return libc.UintptrFromInt32(0)
		}
		/* Fix the ->current_block after the reallocation. */
		if (*MD_CTX)(unsafe.Pointer(ctx)).Fcurrent_block != libc.UintptrFromInt32(0) {
			off_current_block = libc.Uint32FromInt64(int64((*MD_CTX)(unsafe.Pointer(ctx)).Fcurrent_block) - int64((*MD_CTX)(unsafe.Pointer(ctx)).Fblock_bytes))
			(*MD_CTX)(unsafe.Pointer(ctx)).Fcurrent_block = new_block_bytes + uintptr(off_current_block)
		}
		(*MD_CTX)(unsafe.Pointer(ctx)).Fblock_bytes = new_block_bytes
	}
	ptr = (*MD_CTX)(unsafe.Pointer(ctx)).Fblock_bytes + uintptr((*MD_CTX)(unsafe.Pointer(ctx)).Fn_block_bytes)
	*(*int32)(__ccgo_ptr(ctx + 536)) += n_bytes
	return ptr
}

func md_start_new_block(tls *libc.TLS, ctx uintptr, line uintptr) (r int32) {
	var block uintptr
	_ = block
	if !((*MD_CTX)(unsafe.Pointer(ctx)).Fcurrent_block == libc.UintptrFromInt32(0)) {
		libc.X__builtin_unreachable(tls)
	}
	block = md_push_block_bytes(tls, ctx, int32(8))
	if block == libc.UintptrFromInt32(0) {
		return -int32(1)
	}
	switch (*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Ftype1 {
	case int32(MD_LINE_HR):
		goto _1
	case int32(MD_LINE_SETEXTHEADER):
		goto _2
	case int32(MD_LINE_ATXHEADER):
		goto _3
	case int32(MD_LINE_INDENTEDCODE):
		goto _4
	case int32(MD_LINE_FENCEDCODE):
		goto _5
	case int32(MD_LINE_TEXT):
		goto _6
	case int32(MD_LINE_HTML):
		goto _7
	default:
		goto _8
	case int32(MD_LINE_TABLEUNDERLINE):
		goto _9
	case int32(MD_LINE_SETEXTUNDERLINE):
		goto _10
	case int32(MD_LINE_BLANK):
		goto _11
	}
	goto _12
_1:
	;
	libc.SetBitFieldPtr32Int32(block+0, int32(MD_BLOCK_HR), 0, 0xff)
	goto _12
_3:
	;
_2:
	;
	libc.SetBitFieldPtr32Int32(block+0, int32(MD_BLOCK_H), 0, 0xff)
	goto _12
_5:
	;
_4:
	;
	libc.SetBitFieldPtr32Int32(block+0, int32(MD_BLOCK_CODE), 0, 0xff)
	goto _12
_6:
	;
	libc.SetBitFieldPtr32Int32(block+0, int32(MD_BLOCK_P), 0, 0xff)
	goto _12
_7:
	;
	libc.SetBitFieldPtr32Int32(block+0, int32(MD_BLOCK_HTML), 0, 0xff)
	goto _12
_11:
	;
_10:
	;
_9:
	;
_8:
	;
_15:
	;
	libc.X__builtin_unreachable(tls)
	goto _14
_14:
	;
	if 0 != 0 {
		goto _15
	}
	goto _13
_13:
	;
	goto _12
_12:
	;
	libc.SetBitFieldPtr32Uint32(block+0, libc.Uint32FromInt32(0), 8, 0xff00)
	libc.SetBitFieldPtr32Uint32(block+0, (*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Fdata, 16, 0xffff0000)
	(*MD_BLOCK)(unsafe.Pointer(block)).Fn_lines = uint32(0)
	(*MD_CTX)(unsafe.Pointer(ctx)).Fcurrent_block = block
	return 0
}

// C documentation
//
//	/* Eat from start of current (textual) block any reference definitions and
//	 * remember them so we can resolve any links referring to them.
//	 *
//	 * (Reference definitions can only be at start of it as they cannot break
//	 * a paragraph.)
//	 */
func md_consume_link_reference_definitions(tls *libc.TLS, ctx uintptr) (r int32) {
	var lines, v1 uintptr
	var n, n_lines MD_SIZE
	var n_link_ref_lines int32
	_, _, _, _, _ = lines, n, n_lines, n_link_ref_lines, v1
	lines = (*MD_CTX)(unsafe.Pointer(ctx)).Fcurrent_block + libc.UintptrFromInt32(1)*8
	n_lines = (*MD_BLOCK)(unsafe.Pointer((*MD_CTX)(unsafe.Pointer(ctx)).Fcurrent_block)).Fn_lines
	n = uint32(0)
	/* Compute how many lines at the start of the block form one or more
	 * reference definitions. */
	for n < n_lines {
		n_link_ref_lines = md_is_link_reference_definition(tls, ctx, lines+uintptr(n)*8, n_lines-n)
		/* Not a reference definition? */
		if n_link_ref_lines == 0 {
			break
		}
		/* We fail if it is the ref. def. but it could not be stored due
		 * a memory allocation error. */
		if n_link_ref_lines < 0 {
			return -int32(1)
		}
		n = n + libc.Uint32FromInt32(n_link_ref_lines)
	}
	/* If there was at least one reference definition, we need to remove
	 * its lines from the block, or perhaps even the whole block. */
	if n > uint32(0) {
		if n == n_lines {
			/* Remove complete block. */
			v1 = ctx + 536
			*(*int32)(unsafe.Pointer(v1)) = int32(uint64(*(*int32)(unsafe.Pointer(v1))) - uint64(n)*libc.Uint64FromInt64(8))
			v1 = ctx + 536
			*(*int32)(unsafe.Pointer(v1)) = int32(uint64(*(*int32)(unsafe.Pointer(v1))) - libc.Uint64FromInt64(8))
			(*MD_CTX)(unsafe.Pointer(ctx)).Fcurrent_block = libc.UintptrFromInt32(0)
		} else {
			/* Remove just some initial lines from the block. */
			libc.Xmemmove(tls, lines, lines+uintptr(n)*8, uint64(n_lines-n)*uint64(8))
			*(*MD_SIZE)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Fcurrent_block + 4)) -= n
			v1 = ctx + 536
			*(*int32)(unsafe.Pointer(v1)) = int32(uint64(*(*int32)(unsafe.Pointer(v1))) - uint64(n)*libc.Uint64FromInt64(8))
		}
	}
	return 0
}

func md_end_current_block(tls *libc.TLS, ctx uintptr) (r int32) {
	var lines, v1 uintptr
	var n_lines MD_SIZE
	var ret int32
	_, _, _, _ = lines, n_lines, ret, v1
	ret = 0
	if (*MD_CTX)(unsafe.Pointer(ctx)).Fcurrent_block == libc.UintptrFromInt32(0) {
		return ret
	}
	/* Check whether there is a reference definition. (We do this here instead
	 * of in md_analyze_line() because reference definition can take multiple
	 * lines.) */
	if MD_BLOCKTYPE1(*(*uint32)(unsafe.Pointer((*MD_CTX)(unsafe.Pointer(ctx)).Fcurrent_block + 0))&0xff>>0) == int32(MD_BLOCK_P) || MD_BLOCKTYPE1(*(*uint32)(unsafe.Pointer((*MD_CTX)(unsafe.Pointer(ctx)).Fcurrent_block + 0))&0xff>>0) == int32(MD_BLOCK_H) && int32(*(*uint32)(unsafe.Pointer((*MD_CTX)(unsafe.Pointer(ctx)).Fcurrent_block + 0))&0xff00>>8)&int32(MD_BLOCK_SETEXT_HEADER) != 0 {
		lines = (*MD_CTX)(unsafe.Pointer(ctx)).Fcurrent_block + libc.UintptrFromInt32(1)*8
		if (*(*MD_LINE)(__ccgo_ptr(lines))).Fbeg < (*MD_CTX)(unsafe.Pointer(ctx)).Fsize && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr((*(*MD_LINE)(__ccgo_ptr(lines))).Fbeg)))) == int32('[') {
			ret = md_consume_link_reference_definitions(tls, ctx)
			if ret < 0 {
				goto abort
			}
			if (*MD_CTX)(unsafe.Pointer(ctx)).Fcurrent_block == libc.UintptrFromInt32(0) {
				return ret
			}
		}
	}
	if MD_BLOCKTYPE1(*(*uint32)(unsafe.Pointer((*MD_CTX)(unsafe.Pointer(ctx)).Fcurrent_block + 0))&0xff>>0) == int32(MD_BLOCK_H) && int32(*(*uint32)(unsafe.Pointer((*MD_CTX)(unsafe.Pointer(ctx)).Fcurrent_block + 0))&0xff00>>8)&int32(MD_BLOCK_SETEXT_HEADER) != 0 {
		n_lines = (*MD_BLOCK)(unsafe.Pointer((*MD_CTX)(unsafe.Pointer(ctx)).Fcurrent_block)).Fn_lines
		if n_lines > uint32(1) {
			/* Get rid of the underline. */
			(*MD_BLOCK)(unsafe.Pointer((*MD_CTX)(unsafe.Pointer(ctx)).Fcurrent_block)).Fn_lines = (*MD_BLOCK)(unsafe.Pointer((*MD_CTX)(unsafe.Pointer(ctx)).Fcurrent_block)).Fn_lines - 1
			v1 = ctx + 536
			*(*int32)(unsafe.Pointer(v1)) = int32(uint64(*(*int32)(unsafe.Pointer(v1))) - libc.Uint64FromInt64(8))
		} else {
			/* Only the underline has left after eating the ref. defs.
			 * Keep the line as beginning of a new ordinary paragraph. */
			libc.SetBitFieldPtr32Int32((*MD_CTX)(unsafe.Pointer(ctx)).Fcurrent_block+0, int32(MD_BLOCK_P), 0, 0xff)
			return 0
		}
	}
	/* Mark we are not building any block anymore. */
	(*MD_CTX)(unsafe.Pointer(ctx)).Fcurrent_block = libc.UintptrFromInt32(0)
	goto abort
abort:
	;
	return ret
	return r
}

func md_add_line_into_current_block(tls *libc.TLS, ctx uintptr, analysis uintptr) (r int32) {
	var line, line1 uintptr
	_, _ = line, line1
	if !((*MD_CTX)(unsafe.Pointer(ctx)).Fcurrent_block != libc.UintptrFromInt32(0)) {
		libc.X__builtin_unreachable(tls)
	}
	if MD_BLOCKTYPE1(*(*uint32)(unsafe.Pointer((*MD_CTX)(unsafe.Pointer(ctx)).Fcurrent_block + 0))&0xff>>0) == int32(MD_BLOCK_CODE) || MD_BLOCKTYPE1(*(*uint32)(unsafe.Pointer((*MD_CTX)(unsafe.Pointer(ctx)).Fcurrent_block + 0))&0xff>>0) == int32(MD_BLOCK_HTML) {
		line = md_push_block_bytes(tls, ctx, int32(12))
		if line == libc.UintptrFromInt32(0) {
			return -int32(1)
		}
		(*MD_VERBATIMLINE)(unsafe.Pointer(line)).Findent = (*MD_LINE_ANALYSIS)(unsafe.Pointer(analysis)).Findent
		(*MD_VERBATIMLINE)(unsafe.Pointer(line)).Fbeg = (*MD_LINE_ANALYSIS)(unsafe.Pointer(analysis)).Fbeg
		(*MD_VERBATIMLINE)(unsafe.Pointer(line)).Fend = (*MD_LINE_ANALYSIS)(unsafe.Pointer(analysis)).Fend
	} else {
		line1 = md_push_block_bytes(tls, ctx, int32(8))
		if line1 == libc.UintptrFromInt32(0) {
			return -int32(1)
		}
		(*MD_LINE)(unsafe.Pointer(line1)).Fbeg = (*MD_LINE_ANALYSIS)(unsafe.Pointer(analysis)).Fbeg
		(*MD_LINE)(unsafe.Pointer(line1)).Fend = (*MD_LINE_ANALYSIS)(unsafe.Pointer(analysis)).Fend
	}
	(*MD_BLOCK)(unsafe.Pointer((*MD_CTX)(unsafe.Pointer(ctx)).Fcurrent_block)).Fn_lines = (*MD_BLOCK)(unsafe.Pointer((*MD_CTX)(unsafe.Pointer(ctx)).Fcurrent_block)).Fn_lines + 1
	return 0
}

func md_push_container_bytes(tls *libc.TLS, ctx uintptr, type1 MD_BLOCKTYPE1, start uint32, data uint32, flags uint32) (r int32) {
	var block uintptr
	var ret int32
	_, _ = block, ret
	ret = 0
	ret = md_end_current_block(tls, ctx)
	if ret < 0 {
		goto abort
	}
	block = md_push_block_bytes(tls, ctx, int32(8))
	if block == libc.UintptrFromInt32(0) {
		return -int32(1)
	}
	libc.SetBitFieldPtr32Int32(block+0, type1, 0, 0xff)
	libc.SetBitFieldPtr32Uint32(block+0, flags, 8, 0xff00)
	libc.SetBitFieldPtr32Uint32(block+0, data, 16, 0xffff0000)
	(*MD_BLOCK)(unsafe.Pointer(block)).Fn_lines = start
	goto abort
abort:
	;
	return ret
	return r
}

/***********************
 ***  Line Analysis  ***
 ***********************/

func md_is_hr_line(tls *libc.TLS, ctx uintptr, beg MD_OFFSET, p_end uintptr, p_killer uintptr) (r int32) {
	var n int32
	var off MD_OFFSET
	_, _ = n, off
	off = beg + uint32(1)
	n = int32(1)
	for off < (*MD_CTX)(unsafe.Pointer(ctx)).Fsize && (libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(beg)))) || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32(' ') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('\t')) {
		if libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(beg)))) {
			n = n + 1
		}
		off = off + 1
	}
	if n < int32(3) {
		*(*MD_OFFSET)(__ccgo_ptr(p_killer)) = off
		return FALSE
	}
	/* Nothing else can be present on the line. */
	if off < (*MD_CTX)(unsafe.Pointer(ctx)).Fsize && !(libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('\r') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('\n')) {
		*(*MD_OFFSET)(__ccgo_ptr(p_killer)) = off
		return FALSE
	}
	*(*MD_OFFSET)(__ccgo_ptr(p_end)) = off
	return int32(TRUE)
}

func md_is_atxheader_line(tls *libc.TLS, ctx uintptr, beg MD_OFFSET, p_beg uintptr, p_end uintptr, p_level uintptr) (r int32) {
	var n int32
	var off MD_OFFSET
	_, _ = n, off
	off = beg + uint32(1)
	for off < (*MD_CTX)(unsafe.Pointer(ctx)).Fsize && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('#') && off-beg < uint32(7) {
		off = off + 1
	}
	n = libc.Int32FromUint32(off - beg)
	if n > int32(6) {
		return FALSE
	}
	*(*uint32)(__ccgo_ptr(p_level)) = libc.Uint32FromInt32(n)
	if !((*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fflags&libc.Uint32FromInt32(MD_FLAG_PERMISSIVEATXHEADERS) != 0) && off < (*MD_CTX)(unsafe.Pointer(ctx)).Fsize && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) != int32(' ') && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) != int32('\t') && !(libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('\r') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('\n')) {
		return FALSE
	}
	for off < (*MD_CTX)(unsafe.Pointer(ctx)).Fsize && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32(' ') {
		off = off + 1
	}
	*(*MD_OFFSET)(__ccgo_ptr(p_beg)) = off
	*(*MD_OFFSET)(__ccgo_ptr(p_end)) = off
	return int32(TRUE)
}

func md_is_setext_underline(tls *libc.TLS, ctx uintptr, beg MD_OFFSET, p_end uintptr, p_level uintptr) (r int32) {
	var off MD_OFFSET
	var v1 int32
	_, _ = off, v1
	off = beg + uint32(1)
	for off < (*MD_CTX)(unsafe.Pointer(ctx)).Fsize && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(beg)))) {
		off = off + 1
	}
	/* Optionally, space(s) or tabs can follow. */
	for off < (*MD_CTX)(unsafe.Pointer(ctx)).Fsize && (libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32(' ') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('\t')) {
		off = off + 1
	}
	/* But nothing more is allowed on the line. */
	if off < (*MD_CTX)(unsafe.Pointer(ctx)).Fsize && !(libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('\r') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('\n')) {
		return FALSE
	}
	if libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(beg)))) == int32('=') {
		v1 = int32(1)
	} else {
		v1 = int32(2)
	}
	*(*uint32)(__ccgo_ptr(p_level)) = libc.Uint32FromInt32(v1)
	*(*MD_OFFSET)(__ccgo_ptr(p_end)) = off
	return int32(TRUE)
}

func md_is_table_underline(tls *libc.TLS, ctx uintptr, beg MD_OFFSET, p_end uintptr, p_col_count uintptr) (r int32) {
	var col_count uint32
	var delimited, found_pipe int32
	var off MD_OFFSET
	_, _, _, _ = col_count, delimited, found_pipe, off
	off = beg
	found_pipe = FALSE
	col_count = uint32(0)
	if off < (*MD_CTX)(unsafe.Pointer(ctx)).Fsize && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('|') {
		found_pipe = int32(TRUE)
		off = off + 1
		for off < (*MD_CTX)(unsafe.Pointer(ctx)).Fsize && (libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32(' ') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('\t') || (libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('\v') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('\f'))) {
			off = off + 1
		}
	}
	for int32(1) != 0 {
		delimited = FALSE
		/* Cell underline ("-----", ":----", "----:" or ":----:") */
		if off < (*MD_CTX)(unsafe.Pointer(ctx)).Fsize && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32(':') {
			off = off + 1
		}
		if off >= (*MD_CTX)(unsafe.Pointer(ctx)).Fsize || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) != int32('-') {
			return FALSE
		}
		for off < (*MD_CTX)(unsafe.Pointer(ctx)).Fsize && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('-') {
			off = off + 1
		}
		if off < (*MD_CTX)(unsafe.Pointer(ctx)).Fsize && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32(':') {
			off = off + 1
		}
		col_count = col_count + 1
		if col_count > uint32(TABLE_MAXCOLCOUNT) {
			if (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log != libc.UintptrFromInt32(0) {
				(*(*func(*libc.TLS, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log})))(tls, __ccgo_ts+362, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
			}
			return FALSE
		}
		/* Pipe delimiter (optional at the end of line). */
		for off < (*MD_CTX)(unsafe.Pointer(ctx)).Fsize && (libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32(' ') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('\t') || (libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('\v') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('\f'))) {
			off = off + 1
		}
		if off < (*MD_CTX)(unsafe.Pointer(ctx)).Fsize && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('|') {
			delimited = int32(TRUE)
			found_pipe = int32(TRUE)
			off = off + 1
			for off < (*MD_CTX)(unsafe.Pointer(ctx)).Fsize && (libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32(' ') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('\t') || (libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('\v') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('\f'))) {
				off = off + 1
			}
		}
		/* Success, if we reach end of line. */
		if off >= (*MD_CTX)(unsafe.Pointer(ctx)).Fsize || (libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('\r') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('\n')) {
			break
		}
		if !(delimited != 0) {
			return FALSE
		}
	}
	if !(found_pipe != 0) {
		return FALSE
	}
	*(*MD_OFFSET)(__ccgo_ptr(p_end)) = off
	*(*uint32)(__ccgo_ptr(p_col_count)) = col_count
	return int32(TRUE)
}

func md_is_opening_code_fence(tls *libc.TLS, ctx uintptr, beg MD_OFFSET, p_end uintptr) (r int32) {
	var off MD_OFFSET
	_ = off
	off = beg
	for off < (*MD_CTX)(unsafe.Pointer(ctx)).Fsize && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(beg)))) {
		off = off + 1
	}
	/* Fence must have at least three characters. */
	if off-beg < uint32(3) {
		return FALSE
	}
	(*MD_CTX)(unsafe.Pointer(ctx)).Fcode_fence_length = off - beg
	/* Optionally, space(s) can follow. */
	for off < (*MD_CTX)(unsafe.Pointer(ctx)).Fsize && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32(' ') {
		off = off + 1
	}
	/* Optionally, an info string can follow. */
	for off < (*MD_CTX)(unsafe.Pointer(ctx)).Fsize && !(libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('\r') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('\n')) {
		/* Backtick-based fence must not contain '`' in the info string. */
		if libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(beg)))) == int32('`') && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('`') {
			return FALSE
		}
		off = off + 1
	}
	*(*MD_OFFSET)(__ccgo_ptr(p_end)) = off
	return int32(TRUE)
}

func md_is_closing_code_fence(tls *libc.TLS, ctx uintptr, ch MD_CHAR, beg MD_OFFSET, p_end uintptr) (r int32) {
	var off MD_OFFSET
	var ret int32
	_, _ = off, ret
	off = beg
	ret = FALSE
	/* Closing fence must have at least the same length and use same char as
	 * opening one. */
	for off < (*MD_CTX)(unsafe.Pointer(ctx)).Fsize && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == libc.Int32FromUint8(ch) {
		off = off + 1
	}
	if off-beg < (*MD_CTX)(unsafe.Pointer(ctx)).Fcode_fence_length {
		goto out
	}
	/* Optionally, space(s) can follow */
	for off < (*MD_CTX)(unsafe.Pointer(ctx)).Fsize && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32(' ') {
		off = off + 1
	}
	/* But nothing more is allowed on the line. */
	if off < (*MD_CTX)(unsafe.Pointer(ctx)).Fsize && !(libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('\r') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('\n')) {
		goto out
	}
	ret = int32(TRUE)
	goto out
out:
	;
	/* Note we set *p_end even on failure: If we are not closing fence, caller
	 * would eat the line anyway without any parsing. */
	*(*MD_OFFSET)(__ccgo_ptr(p_end)) = off
	return ret
}

// C documentation
//
//	/* Helper data for md_is_html_block_start_condition() and
//	 * md_is_html_block_end_condition() */
type TAG = struct {
	Fname    uintptr
	F__ccgo8 uint8
}

// C documentation
//
//	/* Helper data for md_is_html_block_start_condition() and
//	 * md_is_html_block_end_condition() */
type TAG_tag = TAG

var t1 = [5]TAG{
	0: {
		Fname:    __ccgo_ts + 400,
		F__ccgo8: uint8((libc.Uint64FromInt64(4)-libc.Uint64FromInt32(1))/libc.Uint64FromInt64(1)) & 0xff << 0,
	},
	1: {
		Fname:    __ccgo_ts + 404,
		F__ccgo8: uint8((libc.Uint64FromInt64(7)-libc.Uint64FromInt32(1))/libc.Uint64FromInt64(1)) & 0xff << 0,
	},
	2: {
		Fname:    __ccgo_ts + 411,
		F__ccgo8: uint8((libc.Uint64FromInt64(6)-libc.Uint64FromInt32(1))/libc.Uint64FromInt64(1)) & 0xff << 0,
	},
	3: {
		Fname:    __ccgo_ts + 417,
		F__ccgo8: uint8((libc.Uint64FromInt64(9)-libc.Uint64FromInt32(1))/libc.Uint64FromInt64(1)) & 0xff << 0,
	},
	4: {},
}

var a6 = [4]TAG{
	0: {
		Fname:    __ccgo_ts + 426,
		F__ccgo8: uint8((libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))/libc.Uint64FromInt64(1)) & 0xff << 0,
	},
	1: {
		Fname:    __ccgo_ts + 434,
		F__ccgo8: uint8((libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))/libc.Uint64FromInt64(1)) & 0xff << 0,
	},
	2: {
		Fname:    __ccgo_ts + 442,
		F__ccgo8: uint8((libc.Uint64FromInt64(6)-libc.Uint64FromInt32(1))/libc.Uint64FromInt64(1)) & 0xff << 0,
	},
	3: {},
}
var b6 = [5]TAG{
	0: {
		Fname:    __ccgo_ts + 448,
		F__ccgo8: uint8((libc.Uint64FromInt64(5)-libc.Uint64FromInt32(1))/libc.Uint64FromInt64(1)) & 0xff << 0,
	},
	1: {
		Fname:    __ccgo_ts + 453,
		F__ccgo8: uint8((libc.Uint64FromInt64(9)-libc.Uint64FromInt32(1))/libc.Uint64FromInt64(1)) & 0xff << 0,
	},
	2: {
		Fname:    __ccgo_ts + 462,
		F__ccgo8: uint8((libc.Uint64FromInt64(11)-libc.Uint64FromInt32(1))/libc.Uint64FromInt64(1)) & 0xff << 0,
	},
	3: {
		Fname:    __ccgo_ts + 473,
		F__ccgo8: uint8((libc.Uint64FromInt64(5)-libc.Uint64FromInt32(1))/libc.Uint64FromInt64(1)) & 0xff << 0,
	},
	4: {},
}
var c6 = [5]TAG{
	0: {
		Fname:    __ccgo_ts + 478,
		F__ccgo8: uint8((libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))/libc.Uint64FromInt64(1)) & 0xff << 0,
	},
	1: {
		Fname:    __ccgo_ts + 486,
		F__ccgo8: uint8((libc.Uint64FromInt64(7)-libc.Uint64FromInt32(1))/libc.Uint64FromInt64(1)) & 0xff << 0,
	},
	2: {
		Fname:    __ccgo_ts + 493,
		F__ccgo8: uint8((libc.Uint64FromInt64(4)-libc.Uint64FromInt32(1))/libc.Uint64FromInt64(1)) & 0xff << 0,
	},
	3: {
		Fname:    __ccgo_ts + 497,
		F__ccgo8: uint8((libc.Uint64FromInt64(9)-libc.Uint64FromInt32(1))/libc.Uint64FromInt64(1)) & 0xff << 0,
	},
	4: {},
}
var d6 = [8]TAG{
	0: {
		Fname:    __ccgo_ts + 506,
		F__ccgo8: uint8((libc.Uint64FromInt64(3)-libc.Uint64FromInt32(1))/libc.Uint64FromInt64(1)) & 0xff << 0,
	},
	1: {
		Fname:    __ccgo_ts + 509,
		F__ccgo8: uint8((libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))/libc.Uint64FromInt64(1)) & 0xff << 0,
	},
	2: {
		Fname:    __ccgo_ts + 517,
		F__ccgo8: uint8((libc.Uint64FromInt64(7)-libc.Uint64FromInt32(1))/libc.Uint64FromInt64(1)) & 0xff << 0,
	},
	3: {
		Fname:    __ccgo_ts + 524,
		F__ccgo8: uint8((libc.Uint64FromInt64(4)-libc.Uint64FromInt32(1))/libc.Uint64FromInt64(1)) & 0xff << 0,
	},
	4: {
		Fname:    __ccgo_ts + 528,
		F__ccgo8: uint8((libc.Uint64FromInt64(4)-libc.Uint64FromInt32(1))/libc.Uint64FromInt64(1)) & 0xff << 0,
	},
	5: {
		Fname:    __ccgo_ts + 532,
		F__ccgo8: uint8((libc.Uint64FromInt64(3)-libc.Uint64FromInt32(1))/libc.Uint64FromInt64(1)) & 0xff << 0,
	},
	6: {
		Fname:    __ccgo_ts + 535,
		F__ccgo8: uint8((libc.Uint64FromInt64(3)-libc.Uint64FromInt32(1))/libc.Uint64FromInt64(1)) & 0xff << 0,
	},
	7: {},
}
var f6 = [8]TAG{
	0: {
		Fname:    __ccgo_ts + 538,
		F__ccgo8: uint8((libc.Uint64FromInt64(9)-libc.Uint64FromInt32(1))/libc.Uint64FromInt64(1)) & 0xff << 0,
	},
	1: {
		Fname:    __ccgo_ts + 547,
		F__ccgo8: uint8((libc.Uint64FromInt64(11)-libc.Uint64FromInt32(1))/libc.Uint64FromInt64(1)) & 0xff << 0,
	},
	2: {
		Fname:    __ccgo_ts + 558,
		F__ccgo8: uint8((libc.Uint64FromInt64(7)-libc.Uint64FromInt32(1))/libc.Uint64FromInt64(1)) & 0xff << 0,
	},
	3: {
		Fname:    __ccgo_ts + 565,
		F__ccgo8: uint8((libc.Uint64FromInt64(7)-libc.Uint64FromInt32(1))/libc.Uint64FromInt64(1)) & 0xff << 0,
	},
	4: {
		Fname:    __ccgo_ts + 572,
		F__ccgo8: uint8((libc.Uint64FromInt64(5)-libc.Uint64FromInt32(1))/libc.Uint64FromInt64(1)) & 0xff << 0,
	},
	5: {
		Fname:    __ccgo_ts + 577,
		F__ccgo8: uint8((libc.Uint64FromInt64(6)-libc.Uint64FromInt32(1))/libc.Uint64FromInt64(1)) & 0xff << 0,
	},
	6: {
		Fname:    __ccgo_ts + 583,
		F__ccgo8: uint8((libc.Uint64FromInt64(9)-libc.Uint64FromInt32(1))/libc.Uint64FromInt64(1)) & 0xff << 0,
	},
	7: {},
}
var h6 = [11]TAG{
	0: {
		Fname:    __ccgo_ts + 592,
		F__ccgo8: uint8((libc.Uint64FromInt64(3)-libc.Uint64FromInt32(1))/libc.Uint64FromInt64(1)) & 0xff << 0,
	},
	1: {
		Fname:    __ccgo_ts + 595,
		F__ccgo8: uint8((libc.Uint64FromInt64(3)-libc.Uint64FromInt32(1))/libc.Uint64FromInt64(1)) & 0xff << 0,
	},
	2: {
		Fname:    __ccgo_ts + 598,
		F__ccgo8: uint8((libc.Uint64FromInt64(3)-libc.Uint64FromInt32(1))/libc.Uint64FromInt64(1)) & 0xff << 0,
	},
	3: {
		Fname:    __ccgo_ts + 601,
		F__ccgo8: uint8((libc.Uint64FromInt64(3)-libc.Uint64FromInt32(1))/libc.Uint64FromInt64(1)) & 0xff << 0,
	},
	4: {
		Fname:    __ccgo_ts + 604,
		F__ccgo8: uint8((libc.Uint64FromInt64(3)-libc.Uint64FromInt32(1))/libc.Uint64FromInt64(1)) & 0xff << 0,
	},
	5: {
		Fname:    __ccgo_ts + 607,
		F__ccgo8: uint8((libc.Uint64FromInt64(3)-libc.Uint64FromInt32(1))/libc.Uint64FromInt64(1)) & 0xff << 0,
	},
	6: {
		Fname:    __ccgo_ts + 610,
		F__ccgo8: uint8((libc.Uint64FromInt64(5)-libc.Uint64FromInt32(1))/libc.Uint64FromInt64(1)) & 0xff << 0,
	},
	7: {
		Fname:    __ccgo_ts + 615,
		F__ccgo8: uint8((libc.Uint64FromInt64(7)-libc.Uint64FromInt32(1))/libc.Uint64FromInt64(1)) & 0xff << 0,
	},
	8: {
		Fname:    __ccgo_ts + 622,
		F__ccgo8: uint8((libc.Uint64FromInt64(3)-libc.Uint64FromInt32(1))/libc.Uint64FromInt64(1)) & 0xff << 0,
	},
	9: {
		Fname:    __ccgo_ts + 625,
		F__ccgo8: uint8((libc.Uint64FromInt64(5)-libc.Uint64FromInt32(1))/libc.Uint64FromInt64(1)) & 0xff << 0,
	},
	10: {},
}
var i6 = [2]TAG{
	0: {
		Fname:    __ccgo_ts + 630,
		F__ccgo8: uint8((libc.Uint64FromInt64(7)-libc.Uint64FromInt32(1))/libc.Uint64FromInt64(1)) & 0xff << 0,
	},
	1: {},
}
var l6 = [4]TAG{
	0: {
		Fname:    __ccgo_ts + 637,
		F__ccgo8: uint8((libc.Uint64FromInt64(7)-libc.Uint64FromInt32(1))/libc.Uint64FromInt64(1)) & 0xff << 0,
	},
	1: {
		Fname:    __ccgo_ts + 644,
		F__ccgo8: uint8((libc.Uint64FromInt64(3)-libc.Uint64FromInt32(1))/libc.Uint64FromInt64(1)) & 0xff << 0,
	},
	2: {
		Fname:    __ccgo_ts + 647,
		F__ccgo8: uint8((libc.Uint64FromInt64(5)-libc.Uint64FromInt32(1))/libc.Uint64FromInt64(1)) & 0xff << 0,
	},
	3: {},
}
var m6 = [4]TAG{
	0: {
		Fname:    __ccgo_ts + 652,
		F__ccgo8: uint8((libc.Uint64FromInt64(5)-libc.Uint64FromInt32(1))/libc.Uint64FromInt64(1)) & 0xff << 0,
	},
	1: {
		Fname:    __ccgo_ts + 657,
		F__ccgo8: uint8((libc.Uint64FromInt64(5)-libc.Uint64FromInt32(1))/libc.Uint64FromInt64(1)) & 0xff << 0,
	},
	2: {
		Fname:    __ccgo_ts + 662,
		F__ccgo8: uint8((libc.Uint64FromInt64(9)-libc.Uint64FromInt32(1))/libc.Uint64FromInt64(1)) & 0xff << 0,
	},
	3: {},
}
var n6 = [3]TAG{
	0: {
		Fname:    __ccgo_ts + 671,
		F__ccgo8: uint8((libc.Uint64FromInt64(4)-libc.Uint64FromInt32(1))/libc.Uint64FromInt64(1)) & 0xff << 0,
	},
	1: {
		Fname:    __ccgo_ts + 675,
		F__ccgo8: uint8((libc.Uint64FromInt64(9)-libc.Uint64FromInt32(1))/libc.Uint64FromInt64(1)) & 0xff << 0,
	},
	2: {},
}
var o6 = [4]TAG{
	0: {
		Fname:    __ccgo_ts + 684,
		F__ccgo8: uint8((libc.Uint64FromInt64(3)-libc.Uint64FromInt32(1))/libc.Uint64FromInt64(1)) & 0xff << 0,
	},
	1: {
		Fname:    __ccgo_ts + 687,
		F__ccgo8: uint8((libc.Uint64FromInt64(9)-libc.Uint64FromInt32(1))/libc.Uint64FromInt64(1)) & 0xff << 0,
	},
	2: {
		Fname:    __ccgo_ts + 696,
		F__ccgo8: uint8((libc.Uint64FromInt64(7)-libc.Uint64FromInt32(1))/libc.Uint64FromInt64(1)) & 0xff << 0,
	},
	3: {},
}
var p6 = [3]TAG{
	0: {
		Fname:    __ccgo_ts + 703,
		F__ccgo8: uint8((libc.Uint64FromInt64(2)-libc.Uint64FromInt32(1))/libc.Uint64FromInt64(1)) & 0xff << 0,
	},
	1: {
		Fname:    __ccgo_ts + 705,
		F__ccgo8: uint8((libc.Uint64FromInt64(6)-libc.Uint64FromInt32(1))/libc.Uint64FromInt64(1)) & 0xff << 0,
	},
	2: {},
}
var s6 = [4]TAG{
	0: {
		Fname:    __ccgo_ts + 711,
		F__ccgo8: uint8((libc.Uint64FromInt64(7)-libc.Uint64FromInt32(1))/libc.Uint64FromInt64(1)) & 0xff << 0,
	},
	1: {
		Fname:    __ccgo_ts + 718,
		F__ccgo8: uint8((libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))/libc.Uint64FromInt64(1)) & 0xff << 0,
	},
	2: {
		Fname:    __ccgo_ts + 726,
		F__ccgo8: uint8((libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))/libc.Uint64FromInt64(1)) & 0xff << 0,
	},
	3: {},
}
var t6 = [10]TAG{
	0: {
		Fname:    __ccgo_ts + 734,
		F__ccgo8: uint8((libc.Uint64FromInt64(6)-libc.Uint64FromInt32(1))/libc.Uint64FromInt64(1)) & 0xff << 0,
	},
	1: {
		Fname:    __ccgo_ts + 740,
		F__ccgo8: uint8((libc.Uint64FromInt64(6)-libc.Uint64FromInt32(1))/libc.Uint64FromInt64(1)) & 0xff << 0,
	},
	2: {
		Fname:    __ccgo_ts + 746,
		F__ccgo8: uint8((libc.Uint64FromInt64(3)-libc.Uint64FromInt32(1))/libc.Uint64FromInt64(1)) & 0xff << 0,
	},
	3: {
		Fname:    __ccgo_ts + 749,
		F__ccgo8: uint8((libc.Uint64FromInt64(6)-libc.Uint64FromInt32(1))/libc.Uint64FromInt64(1)) & 0xff << 0,
	},
	4: {
		Fname:    __ccgo_ts + 755,
		F__ccgo8: uint8((libc.Uint64FromInt64(3)-libc.Uint64FromInt32(1))/libc.Uint64FromInt64(1)) & 0xff << 0,
	},
	5: {
		Fname:    __ccgo_ts + 758,
		F__ccgo8: uint8((libc.Uint64FromInt64(6)-libc.Uint64FromInt32(1))/libc.Uint64FromInt64(1)) & 0xff << 0,
	},
	6: {
		Fname:    __ccgo_ts + 764,
		F__ccgo8: uint8((libc.Uint64FromInt64(6)-libc.Uint64FromInt32(1))/libc.Uint64FromInt64(1)) & 0xff << 0,
	},
	7: {
		Fname:    __ccgo_ts + 770,
		F__ccgo8: uint8((libc.Uint64FromInt64(3)-libc.Uint64FromInt32(1))/libc.Uint64FromInt64(1)) & 0xff << 0,
	},
	8: {
		Fname:    __ccgo_ts + 773,
		F__ccgo8: uint8((libc.Uint64FromInt64(6)-libc.Uint64FromInt32(1))/libc.Uint64FromInt64(1)) & 0xff << 0,
	},
	9: {},
}
var u6 = [2]TAG{
	0: {
		Fname:    __ccgo_ts + 779,
		F__ccgo8: uint8((libc.Uint64FromInt64(3)-libc.Uint64FromInt32(1))/libc.Uint64FromInt64(1)) & 0xff << 0,
	},
	1: {},
}
var xx = [1]TAG{}

// C documentation
//
//	/* Returns type of the raw HTML block, or FALSE if it is not HTML block.
//	 * (Refer to CommonMark specification for details about the types.)
//	 */
func md_is_html_block_start_condition(tls *libc.TLS, ctx uintptr, beg MD_OFFSET) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var i, slot, v2 int32
	var off, tmp MD_OFFSET
	var tags uintptr
	var _ /* end at bp+0 */ MD_OFFSET
	_, _, _, _, _, _ = i, off, slot, tags, tmp, v2
	off = beg + uint32(1)
	/* Check for type 1: <script, <pre, or <style */
	i = 0
	for {
		if !(t1[i].Fname != libc.UintptrFromInt32(0)) {
			break
		}
		if off+libc.Uint32FromInt32(int32(uint32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&t1)) + uintptr(i)*16 + 8))&0xff>>0))) <= (*MD_CTX)(unsafe.Pointer(ctx)).Fsize {
			if md_ascii_case_eq(tls, (*MD_CTX)(unsafe.Pointer(ctx)).Ftext+uintptr(off), t1[i].Fname, libc.Uint32FromInt32(int32(uint32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&t1)) + uintptr(i)*16 + 8))&0xff>>0)))) != 0 {
				return int32(1)
			}
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	/* Check for type 2: <!-- */
	if off+uint32(3) < (*MD_CTX)(unsafe.Pointer(ctx)).Fsize && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('!') && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))) == int32('-') && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(2))))) == int32('-') {
		return int32(2)
	}
	/* Check for type 3: <? */
	if off < (*MD_CTX)(unsafe.Pointer(ctx)).Fsize && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('?') {
		return int32(3)
	}
	/* Check for type 4 or 5: <! */
	if off < (*MD_CTX)(unsafe.Pointer(ctx)).Fsize && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('!') {
		/* Check for type 4: <! followed by uppercase letter. */
		if off+uint32(1) < (*MD_CTX)(unsafe.Pointer(ctx)).Fsize && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))) <= uint32(127) {
			return int32(4)
		}
		/* Check for type 5: <![CDATA[ */
		if off+uint32(8) < (*MD_CTX)(unsafe.Pointer(ctx)).Fsize {
			if md_ascii_eq(tls, (*MD_CTX)(unsafe.Pointer(ctx)).Ftext+uintptr(off), __ccgo_ts+782, uint32(8)) != 0 {
				return int32(5)
			}
		}
	}
	/* Check for type 6: Many possible starting tags listed above. */
	if off+uint32(1) < (*MD_CTX)(unsafe.Pointer(ctx)).Fsize && (libc.Uint32FromInt32(libc.Int32FromUint8('A')) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) <= libc.Uint32FromInt32(libc.Int32FromUint8('Z')) || libc.Uint32FromInt32(libc.Int32FromUint8('a')) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) <= libc.Uint32FromInt32(libc.Int32FromUint8('z')) || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('/') && (libc.Uint32FromInt32(libc.Int32FromUint8('A')) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))) <= libc.Uint32FromInt32(libc.Int32FromUint8('Z')) || libc.Uint32FromInt32(libc.Int32FromUint8('a')) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))) <= libc.Uint32FromInt32(libc.Int32FromUint8('z')))) {
		if libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('/') {
			off = off + 1
		}
		if libc.Uint32FromInt32(libc.Int32FromUint8('A')) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) <= libc.Uint32FromInt32(libc.Int32FromUint8('Z')) {
			v2 = libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) - int32('A')
		} else {
			v2 = libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) - int32('a')
		}
		slot = v2
		tags = map6[slot]
		i = 0
		for {
			if !((*(*TAG)(__ccgo_ptr(tags + uintptr(i)*16))).Fname != libc.UintptrFromInt32(0)) {
				break
			}
			if off+libc.Uint32FromInt32(int32(uint32(*(*uint8)(unsafe.Pointer(tags + uintptr(i)*16 + 8))&0xff>>0))) <= (*MD_CTX)(unsafe.Pointer(ctx)).Fsize {
				if md_ascii_case_eq(tls, (*MD_CTX)(unsafe.Pointer(ctx)).Ftext+uintptr(off), (*(*TAG)(__ccgo_ptr(tags + uintptr(i)*16))).Fname, libc.Uint32FromInt32(int32(uint32(*(*uint8)(unsafe.Pointer(tags + uintptr(i)*16 + 8))&0xff>>0)))) != 0 {
					tmp = off + libc.Uint32FromInt32(int32(uint32(*(*uint8)(unsafe.Pointer(tags + uintptr(i)*16 + 8))&0xff>>0)))
					if tmp >= (*MD_CTX)(unsafe.Pointer(ctx)).Fsize {
						return int32(6)
					}
					if libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(tmp)))) == int32(' ') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(tmp)))) == int32('\t') || (libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(tmp)))) == int32('\r') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(tmp)))) == int32('\n')) || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(tmp)))) == int32('>') {
						return int32(6)
					}
					if tmp+uint32(1) < (*MD_CTX)(unsafe.Pointer(ctx)).Fsize && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(tmp)))) == int32('/') && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(tmp+libc.Uint32FromInt32(1))))) == int32('>') {
						return int32(6)
					}
					break
				}
			}
			goto _3
		_3:
			;
			i = i + 1
		}
	}
	/* Check for type 7: any COMPLETE other opening or closing tag. */
	if off+uint32(1) < (*MD_CTX)(unsafe.Pointer(ctx)).Fsize {
		if md_is_html_tag(tls, ctx, libc.UintptrFromInt32(0), uint32(0), beg, (*MD_CTX)(unsafe.Pointer(ctx)).Fsize, bp) != 0 {
			/* Only optional whitespace and new line may follow. */
			for *(*MD_OFFSET)(__ccgo_ptr(bp)) < (*MD_CTX)(unsafe.Pointer(ctx)).Fsize && (libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(*(*MD_OFFSET)(__ccgo_ptr(bp)))))) == int32(' ') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(*(*MD_OFFSET)(__ccgo_ptr(bp)))))) == int32('\t') || (libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(*(*MD_OFFSET)(__ccgo_ptr(bp)))))) == int32('\v') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(*(*MD_OFFSET)(__ccgo_ptr(bp)))))) == int32('\f'))) {
				*(*MD_OFFSET)(__ccgo_ptr(bp)) = *(*MD_OFFSET)(__ccgo_ptr(bp)) + 1
			}
			if *(*MD_OFFSET)(__ccgo_ptr(bp)) >= (*MD_CTX)(unsafe.Pointer(ctx)).Fsize || (libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(*(*MD_OFFSET)(__ccgo_ptr(bp)))))) == int32('\r') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(*(*MD_OFFSET)(__ccgo_ptr(bp)))))) == int32('\n')) {
				return int32(7)
			}
		}
	}
	return FALSE
}

/* Type 6 is started by a long list of allowed tags. We use two-level
 * tree to speed-up the search. */
var map6 = [26]uintptr{
	0:  uintptr(unsafe.Pointer(&a6)),
	1:  uintptr(unsafe.Pointer(&b6)),
	2:  uintptr(unsafe.Pointer(&c6)),
	3:  uintptr(unsafe.Pointer(&d6)),
	4:  uintptr(unsafe.Pointer(&xx)),
	5:  uintptr(unsafe.Pointer(&f6)),
	6:  uintptr(unsafe.Pointer(&xx)),
	7:  uintptr(unsafe.Pointer(&h6)),
	8:  uintptr(unsafe.Pointer(&i6)),
	9:  uintptr(unsafe.Pointer(&xx)),
	10: uintptr(unsafe.Pointer(&xx)),
	11: uintptr(unsafe.Pointer(&l6)),
	12: uintptr(unsafe.Pointer(&m6)),
	13: uintptr(unsafe.Pointer(&n6)),
	14: uintptr(unsafe.Pointer(&o6)),
	15: uintptr(unsafe.Pointer(&p6)),
	16: uintptr(unsafe.Pointer(&xx)),
	17: uintptr(unsafe.Pointer(&xx)),
	18: uintptr(unsafe.Pointer(&s6)),
	19: uintptr(unsafe.Pointer(&t6)),
	20: uintptr(unsafe.Pointer(&u6)),
	21: uintptr(unsafe.Pointer(&xx)),
	22: uintptr(unsafe.Pointer(&xx)),
	23: uintptr(unsafe.Pointer(&xx)),
	24: uintptr(unsafe.Pointer(&xx)),
	25: uintptr(unsafe.Pointer(&xx)),
}

// C documentation
//
//	/* Case sensitive check whether there is a substring 'what' between 'beg'
//	 * and end of line. */
func md_line_contains(tls *libc.TLS, ctx uintptr, beg MD_OFFSET, what uintptr, what_len MD_SIZE, p_end uintptr) (r int32) {
	var i MD_OFFSET
	_ = i
	i = beg
	for {
		if !(i+what_len < (*MD_CTX)(unsafe.Pointer(ctx)).Fsize) {
			break
		}
		if libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(i)))) == int32('\r') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(i)))) == int32('\n') {
			break
		}
		if libc.Xmemcmp(tls, (*MD_CTX)(unsafe.Pointer(ctx)).Ftext+uintptr(i), what, uint64(what_len)*uint64(1)) == 0 {
			*(*MD_OFFSET)(__ccgo_ptr(p_end)) = i + what_len
			return int32(TRUE)
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	*(*MD_OFFSET)(__ccgo_ptr(p_end)) = i
	return FALSE
}

// C documentation
//
//	/* Returns type of HTML block end condition or FALSE if not an end condition.
//	 *
//	 * Note it fills p_end even when it is not end condition as the caller
//	 * does not need to analyze contents of a raw HTML block.
//	 */
func md_is_html_block_end_condition(tls *libc.TLS, ctx uintptr, beg MD_OFFSET, p_end uintptr) (r int32) {
	var i, v11 int32
	var off MD_OFFSET
	_, _, _ = i, off, v11
	switch (*MD_CTX)(unsafe.Pointer(ctx)).Fhtml_block_type {
	case int32(1):
		goto _1
	case int32(2):
		goto _2
	case int32(3):
		goto _3
	case int32(4):
		goto _4
	case int32(5):
		goto _5
	case int32(7):
		goto _6
	case int32(6):
		goto _7
	default:
		goto _8
	}
	goto _9
_1:
	;
	off = beg
	for off+uint32(1) < (*MD_CTX)(unsafe.Pointer(ctx)).Fsize && !(libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('\r') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('\n')) {
		if libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('<') && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))) == int32('/') {
			i = 0
			for {
				if !(t1[i].Fname != libc.UintptrFromInt32(0)) {
					break
				}
				if off+uint32(2)+libc.Uint32FromInt32(int32(uint32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&t1)) + uintptr(i)*16 + 8))&0xff>>0))) < (*MD_CTX)(unsafe.Pointer(ctx)).Fsize {
					if md_ascii_case_eq(tls, (*MD_CTX)(unsafe.Pointer(ctx)).Ftext+uintptr(off+libc.Uint32FromInt32(2)), t1[i].Fname, libc.Uint32FromInt32(int32(uint32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&t1)) + uintptr(i)*16 + 8))&0xff>>0)))) != 0 && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(2)+libc.Uint32FromInt32(int32(uint32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&t1)) + uintptr(i)*16 + 8))&0xff>>0))))))) == int32('>') {
						*(*MD_OFFSET)(__ccgo_ptr(p_end)) = off + uint32(2) + libc.Uint32FromInt32(int32(uint32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&t1)) + uintptr(i)*16 + 8))&0xff>>0))) + uint32(1)
						return int32(TRUE)
					}
				}
				goto _10
			_10:
				;
				i = i + 1
			}
		}
		off = off + 1
	}
	*(*MD_OFFSET)(__ccgo_ptr(p_end)) = off
	return FALSE
_2:
	;
	if md_line_contains(tls, ctx, beg, __ccgo_ts+30, uint32(3), p_end) != 0 {
		v11 = int32(2)
	} else {
		v11 = FALSE
	}
	return v11
_3:
	;
	if md_line_contains(tls, ctx, beg, __ccgo_ts+34, uint32(2), p_end) != 0 {
		v11 = int32(3)
	} else {
		v11 = FALSE
	}
	return v11
_4:
	;
	if md_line_contains(tls, ctx, beg, __ccgo_ts+37, uint32(1), p_end) != 0 {
		v11 = int32(4)
	} else {
		v11 = FALSE
	}
	return v11
_5:
	;
	if md_line_contains(tls, ctx, beg, __ccgo_ts+39, uint32(3), p_end) != 0 {
		v11 = int32(5)
	} else {
		v11 = FALSE
	}
	return v11
_7:
	; /* Pass through */
_6:
	;
	if beg >= (*MD_CTX)(unsafe.Pointer(ctx)).Fsize || (libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(beg)))) == int32('\r') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(beg)))) == int32('\n')) {
		/* Blank line ends types 6 and 7. */
		*(*MD_OFFSET)(__ccgo_ptr(p_end)) = beg
		return (*MD_CTX)(unsafe.Pointer(ctx)).Fhtml_block_type
	}
	return FALSE
_8:
	;
_17:
	;
	libc.X__builtin_unreachable(tls)
	goto _16
_16:
	;
	if 0 != 0 {
		goto _17
	}
	goto _15
_15:
	;
_9:
	;
	return FALSE
}

func md_is_container_compatible(tls *libc.TLS, pivot uintptr, container uintptr) (r int32) {
	/* Block quote has no "items" like lists. */
	if libc.Int32FromUint8((*MD_CONTAINER)(unsafe.Pointer(container)).Fch) == int32('>') {
		return FALSE
	}
	if libc.Int32FromUint8((*MD_CONTAINER)(unsafe.Pointer(container)).Fch) != libc.Int32FromUint8((*MD_CONTAINER)(unsafe.Pointer(pivot)).Fch) {
		return FALSE
	}
	if (*MD_CONTAINER)(unsafe.Pointer(container)).Fmark_indent > (*MD_CONTAINER)(unsafe.Pointer(pivot)).Fcontents_indent {
		return FALSE
	}
	return int32(TRUE)
}

func md_push_container(tls *libc.TLS, ctx uintptr, container uintptr) (r int32) {
	var new_containers, v3 uintptr
	var v1 int32
	_, _, _ = new_containers, v1, v3
	if (*MD_CTX)(unsafe.Pointer(ctx)).Fn_containers >= (*MD_CTX)(unsafe.Pointer(ctx)).Falloc_containers {
		if (*MD_CTX)(unsafe.Pointer(ctx)).Falloc_containers > 0 {
			v1 = (*MD_CTX)(unsafe.Pointer(ctx)).Falloc_containers + (*MD_CTX)(unsafe.Pointer(ctx)).Falloc_containers/int32(2)
		} else {
			v1 = int32(16)
		}
		(*MD_CTX)(unsafe.Pointer(ctx)).Falloc_containers = v1
		new_containers = libc.Xrealloc(tls, (*MD_CTX)(unsafe.Pointer(ctx)).Fcontainers, libc.Uint64FromInt32((*MD_CTX)(unsafe.Pointer(ctx)).Falloc_containers)*uint64(28))
		if new_containers == libc.UintptrFromInt32(0) {
			if (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log != libc.UintptrFromInt32(0) {
				(*(*func(*libc.TLS, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log})))(tls, __ccgo_ts+43, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
			}
			return -int32(1)
		}
		(*MD_CTX)(unsafe.Pointer(ctx)).Fcontainers = new_containers
	}
	v3 = ctx + 552
	v1 = *(*int32)(unsafe.Pointer(v3))
	*(*int32)(unsafe.Pointer(v3)) = *(*int32)(unsafe.Pointer(v3)) + 1
	libc.Xmemcpy(tls, (*MD_CTX)(unsafe.Pointer(ctx)).Fcontainers+uintptr(v1)*28, container, uint64(28))
	return 0
}

func md_enter_child_containers(tls *libc.TLS, ctx uintptr, n_children int32) (r int32) {
	var c uintptr
	var i, is_ordered_list, ret, v10 int32
	_, _, _, _, _ = c, i, is_ordered_list, ret, v10
	ret = 0
	i = (*MD_CTX)(unsafe.Pointer(ctx)).Fn_containers - n_children
	for {
		if !(i < (*MD_CTX)(unsafe.Pointer(ctx)).Fn_containers) {
			break
		}
		c = (*MD_CTX)(unsafe.Pointer(ctx)).Fcontainers + uintptr(i)*28
		is_ordered_list = FALSE
		switch libc.Int32FromUint8((*MD_CONTAINER)(unsafe.Pointer(c)).Fch) {
		case int32('.'):
			goto _2
		case int32(')'):
			goto _3
		case int32('*'):
			goto _4
		case int32('+'):
			goto _5
		case int32('-'):
			goto _6
		case int32('>'):
			goto _7
		default:
			goto _8
		}
		goto _9
	_3:
		;
	_2:
		;
		is_ordered_list = int32(TRUE)
	_6:
		;
	_5:
		;
	_4:
		;
		/* Remember offset in ctx->block_bytes so we can revisit the
		 * block if we detect it is a loose list. */
		md_end_current_block(tls, ctx)
		(*MD_CONTAINER)(unsafe.Pointer(c)).Fblock_byte_off = libc.Uint32FromInt32((*MD_CTX)(unsafe.Pointer(ctx)).Fn_block_bytes)
		if is_ordered_list != 0 {
			v10 = int32(MD_BLOCK_OL)
		} else {
			v10 = int32(MD_BLOCK_UL)
		}
		ret = md_push_container_bytes(tls, ctx, v10, (*MD_CONTAINER)(unsafe.Pointer(c)).Fstart, uint32((*MD_CONTAINER)(unsafe.Pointer(c)).Fch), uint32(MD_BLOCK_CONTAINER_OPENER))
		if ret < 0 {
			goto abort
		}
		if int32(uint32(*(*uint16)(unsafe.Pointer(c + 4))&0xff00>>8)) != 0 {
			v10 = libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr((*MD_CONTAINER)(unsafe.Pointer(c)).Ftask_mark_off))))
		} else {
			v10 = 0
		}
		ret = md_push_container_bytes(tls, ctx, int32(MD_BLOCK_LI), (*MD_CONTAINER)(unsafe.Pointer(c)).Ftask_mark_off, libc.Uint32FromInt32(v10), uint32(MD_BLOCK_CONTAINER_OPENER))
		if ret < 0 {
			goto abort
		}
		goto _9
	_7:
		;
	_14:
		;
		ret = md_push_container_bytes(tls, ctx, int32(MD_BLOCK_QUOTE), uint32(0), uint32(0), uint32(MD_BLOCK_CONTAINER_OPENER))
		if ret < 0 {
			goto abort
		}
		goto _13
	_13:
		;
		if 0 != 0 {
			goto _14
		}
		goto _12
	_12:
		;
		goto _9
	_8:
		;
		libc.X__builtin_unreachable(tls)
		goto _9
	_9:
		;
		goto _1
	_1:
		;
		i = i + 1
	}
	goto abort
abort:
	;
	return ret
	return r
}

func md_leave_child_containers(tls *libc.TLS, ctx uintptr, n_keep int32) (r int32) {
	var c uintptr
	var is_ordered_list, ret, v12 int32
	_, _, _, _ = c, is_ordered_list, ret, v12
	ret = 0
	for (*MD_CTX)(unsafe.Pointer(ctx)).Fn_containers > n_keep {
		c = (*MD_CTX)(unsafe.Pointer(ctx)).Fcontainers + uintptr((*MD_CTX)(unsafe.Pointer(ctx)).Fn_containers-int32(1))*28
		is_ordered_list = FALSE
		switch libc.Int32FromUint8((*MD_CONTAINER)(unsafe.Pointer(c)).Fch) {
		case int32('.'):
			goto _1
		case int32(')'):
			goto _2
		case int32('*'):
			goto _3
		case int32('+'):
			goto _4
		case int32('-'):
			goto _5
		case int32('>'):
			goto _6
		default:
			goto _7
		}
		goto _8
	_2:
		;
	_1:
		;
		is_ordered_list = int32(TRUE)
	_5:
		;
	_4:
		;
	_3:
		;
	_11:
		;
		if int32(uint32(*(*uint16)(unsafe.Pointer(c + 4))&0xff00>>8)) != 0 {
			v12 = libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr((*MD_CONTAINER)(unsafe.Pointer(c)).Ftask_mark_off))))
		} else {
			v12 = 0
		}
		ret = md_push_container_bytes(tls, ctx, int32(MD_BLOCK_LI), (*MD_CONTAINER)(unsafe.Pointer(c)).Ftask_mark_off, libc.Uint32FromInt32(v12), uint32(MD_BLOCK_CONTAINER_CLOSER))
		if ret < 0 {
			goto abort
		}
		goto _10
	_10:
		;
		if 0 != 0 {
			goto _11
		}
		goto _9
	_9:
		;
		if is_ordered_list != 0 {
			v12 = int32(MD_BLOCK_OL)
		} else {
			v12 = int32(MD_BLOCK_UL)
		}
		ret = md_push_container_bytes(tls, ctx, v12, uint32(0), uint32((*MD_CONTAINER)(unsafe.Pointer(c)).Fch), uint32(MD_BLOCK_CONTAINER_CLOSER))
		if ret < 0 {
			goto abort
		}
		goto _8
	_6:
		;
		ret = md_push_container_bytes(tls, ctx, int32(MD_BLOCK_QUOTE), uint32(0), uint32(0), uint32(MD_BLOCK_CONTAINER_CLOSER))
		if ret < 0 {
			goto abort
		}
		goto _8
	_7:
		;
		libc.X__builtin_unreachable(tls)
		goto _8
	_8:
		;
		(*MD_CTX)(unsafe.Pointer(ctx)).Fn_containers = (*MD_CTX)(unsafe.Pointer(ctx)).Fn_containers - 1
	}
	goto abort
abort:
	;
	return ret
	return r
}

func md_is_container_mark(tls *libc.TLS, ctx uintptr, indent uint32, beg MD_OFFSET, p_end uintptr, p_container uintptr) (r int32) {
	var max_end, off MD_OFFSET
	_, _ = max_end, off
	off = beg
	if off >= (*MD_CTX)(unsafe.Pointer(ctx)).Fsize || indent >= (*MD_CTX)(unsafe.Pointer(ctx)).Fcode_indent_offset {
		return FALSE
	}
	/* Check for block quote mark. */
	if libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('>') {
		off = off + 1
		(*MD_CONTAINER)(unsafe.Pointer(p_container)).Fch = uint8('>')
		libc.SetBitFieldPtr16Uint32(p_container+4, libc.Uint32FromInt32(FALSE), 0, 0xff)
		libc.SetBitFieldPtr16Uint32(p_container+4, libc.Uint32FromInt32(FALSE), 8, 0xff00)
		(*MD_CONTAINER)(unsafe.Pointer(p_container)).Fmark_indent = indent
		(*MD_CONTAINER)(unsafe.Pointer(p_container)).Fcontents_indent = indent + uint32(1)
		*(*MD_OFFSET)(__ccgo_ptr(p_end)) = off
		return int32(TRUE)
	}
	/* Check for list item bullet mark. */
	if libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) != int32('\000') && libc.Xstrchr(tls, __ccgo_ts+791, libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off))))) != libc.UintptrFromInt32(0) && (off+uint32(1) >= (*MD_CTX)(unsafe.Pointer(ctx)).Fsize || (libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))) == int32(' ') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))) == int32('\t')) || (libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))) == int32('\r') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))) == int32('\n'))) {
		(*MD_CONTAINER)(unsafe.Pointer(p_container)).Fch = *(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))
		libc.SetBitFieldPtr16Uint32(p_container+4, libc.Uint32FromInt32(FALSE), 0, 0xff)
		libc.SetBitFieldPtr16Uint32(p_container+4, libc.Uint32FromInt32(FALSE), 8, 0xff00)
		(*MD_CONTAINER)(unsafe.Pointer(p_container)).Fmark_indent = indent
		(*MD_CONTAINER)(unsafe.Pointer(p_container)).Fcontents_indent = indent + uint32(1)
		*(*MD_OFFSET)(__ccgo_ptr(p_end)) = off + uint32(1)
		return int32(TRUE)
	}
	/* Check for ordered list item marks. */
	max_end = off + uint32(9)
	if max_end > (*MD_CTX)(unsafe.Pointer(ctx)).Fsize {
		max_end = (*MD_CTX)(unsafe.Pointer(ctx)).Fsize
	}
	(*MD_CONTAINER)(unsafe.Pointer(p_container)).Fstart = uint32(0)
	for off < max_end && (libc.Uint32FromInt32(libc.Int32FromUint8('0')) <= uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) && uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) <= libc.Uint32FromInt32(libc.Int32FromUint8('9'))) {
		(*MD_CONTAINER)(unsafe.Pointer(p_container)).Fstart = (*MD_CONTAINER)(unsafe.Pointer(p_container)).Fstart*uint32(10) + uint32(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) - uint32('0')
		off = off + 1
	}
	if off > beg && off < (*MD_CTX)(unsafe.Pointer(ctx)).Fsize && (libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('.') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32(')')) && (off+uint32(1) >= (*MD_CTX)(unsafe.Pointer(ctx)).Fsize || (libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))) == int32(' ') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))) == int32('\t')) || (libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))) == int32('\r') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off+libc.Uint32FromInt32(1))))) == int32('\n'))) {
		(*MD_CONTAINER)(unsafe.Pointer(p_container)).Fch = *(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))
		libc.SetBitFieldPtr16Uint32(p_container+4, libc.Uint32FromInt32(FALSE), 0, 0xff)
		libc.SetBitFieldPtr16Uint32(p_container+4, libc.Uint32FromInt32(FALSE), 8, 0xff00)
		(*MD_CONTAINER)(unsafe.Pointer(p_container)).Fmark_indent = indent
		(*MD_CONTAINER)(unsafe.Pointer(p_container)).Fcontents_indent = indent + off - beg + uint32(1)
		*(*MD_OFFSET)(__ccgo_ptr(p_end)) = off + uint32(1)
		return int32(TRUE)
	}
	return FALSE
}

func md_line_indentation(tls *libc.TLS, ctx uintptr, total_indent uint32, beg MD_OFFSET, p_end uintptr) (r uint32) {
	var indent uint32
	var off MD_OFFSET
	_, _ = indent, off
	off = beg
	indent = total_indent
	for off < (*MD_CTX)(unsafe.Pointer(ctx)).Fsize && (libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32(' ') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('\t')) {
		if libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(off)))) == int32('\t') {
			indent = (indent + uint32(4)) & libc.Uint32FromInt32(^libc.Int32FromInt32(3))
		} else {
			indent = indent + 1
		}
		off = off + 1
	}
	*(*MD_OFFSET)(__ccgo_ptr(p_end)) = off
	return indent - total_indent
}

var md_dummy_blank_line = MD_LINE_ANALYSIS{}

// C documentation
//
//	/* Analyze type of the line and find some its properties. This serves as a
//	 * main input for determining type and boundaries of a block. */
func md_analyze_line(tls *libc.TLS, ctx uintptr, beg MD_OFFSET, p_end uintptr, pivot_line uintptr, line uintptr) (r int32) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var block, c, c1, task_container, top_block, top_block1, v2 uintptr
	var html_block_type, n_brothers, n_children, n_parents, prev_line_has_list_loosening_effect, ret, v3 int32
	var tmp1, tmp2 MD_OFFSET
	var total_indent, v1 uint32
	var _ /* col_count at bp+48 */ uint32
	var _ /* container at bp+0 */ MD_CONTAINER
	var _ /* hr_killer at bp+32 */ MD_OFFSET
	var _ /* level at bp+36 */ uint32
	var _ /* level at bp+44 */ uint32
	var _ /* off at bp+28 */ MD_OFFSET
	var _ /* tmp at bp+40 */ MD_OFFSET
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = block, c, c1, html_block_type, n_brothers, n_children, n_parents, prev_line_has_list_loosening_effect, ret, task_container, tmp1, tmp2, top_block, top_block1, total_indent, v1, v2, v3
	total_indent = uint32(0)
	n_parents = 0
	n_brothers = 0
	n_children = 0
	*(*MD_CONTAINER)(__ccgo_ptr(bp)) = MD_CONTAINER{}
	prev_line_has_list_loosening_effect = (*MD_CTX)(unsafe.Pointer(ctx)).Flast_line_has_list_loosening_effect
	*(*MD_OFFSET)(__ccgo_ptr(bp + 28)) = beg
	*(*MD_OFFSET)(__ccgo_ptr(bp + 32)) = uint32(0)
	ret = 0
	(*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Findent = md_line_indentation(tls, ctx, total_indent, *(*MD_OFFSET)(__ccgo_ptr(bp + 28)), bp+28)
	total_indent = total_indent + (*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Findent
	(*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Fbeg = *(*MD_OFFSET)(__ccgo_ptr(bp + 28))
	(*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Fenforce_new_block = FALSE
	/* Given the indentation and block quote marks '>', determine how many of
	 * the current containers are our parents. */
	for n_parents < (*MD_CTX)(unsafe.Pointer(ctx)).Fn_containers {
		c = (*MD_CTX)(unsafe.Pointer(ctx)).Fcontainers + uintptr(n_parents)*28
		if libc.Int32FromUint8((*MD_CONTAINER)(unsafe.Pointer(c)).Fch) == int32('>') && (*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Findent < (*MD_CTX)(unsafe.Pointer(ctx)).Fcode_indent_offset && *(*MD_OFFSET)(__ccgo_ptr(bp + 28)) < (*MD_CTX)(unsafe.Pointer(ctx)).Fsize && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(*(*MD_OFFSET)(__ccgo_ptr(bp + 28)))))) == int32('>') {
			/* Block quote mark. */
			*(*MD_OFFSET)(__ccgo_ptr(bp + 28)) = *(*MD_OFFSET)(__ccgo_ptr(bp + 28)) + 1
			total_indent = total_indent + 1
			(*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Findent = md_line_indentation(tls, ctx, total_indent, *(*MD_OFFSET)(__ccgo_ptr(bp + 28)), bp+28)
			total_indent = total_indent + (*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Findent
			/* The optional 1st space after '>' is part of the block quote mark. */
			if (*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Findent > uint32(0) {
				(*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Findent = (*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Findent - 1
			}
			(*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Fbeg = *(*MD_OFFSET)(__ccgo_ptr(bp + 28))
		} else {
			if libc.Int32FromUint8((*MD_CONTAINER)(unsafe.Pointer(c)).Fch) != int32('>') && (*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Findent >= (*MD_CONTAINER)(unsafe.Pointer(c)).Fcontents_indent {
				/* List. */
				*(*uint32)(__ccgo_ptr(line + 20)) -= (*MD_CONTAINER)(unsafe.Pointer(c)).Fcontents_indent
			} else {
				break
			}
		}
		n_parents = n_parents + 1
	}
	if *(*MD_OFFSET)(__ccgo_ptr(bp + 28)) >= (*MD_CTX)(unsafe.Pointer(ctx)).Fsize || (libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(*(*MD_OFFSET)(__ccgo_ptr(bp + 28)))))) == int32('\r') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(*(*MD_OFFSET)(__ccgo_ptr(bp + 28)))))) == int32('\n')) {
		/* Blank line does not need any real indentation to be nested inside
		 * a list. */
		if n_brothers+n_children == 0 {
			for n_parents < (*MD_CTX)(unsafe.Pointer(ctx)).Fn_containers && libc.Int32FromUint8((*(*MD_CONTAINER)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Fcontainers + uintptr(n_parents)*28))).Fch) != int32('>') {
				n_parents = n_parents + 1
			}
		}
	}
	for int32(TRUE) != 0 {
		/* Check whether we are fenced code continuation. */
		if (*MD_LINE_ANALYSIS)(unsafe.Pointer(pivot_line)).Ftype1 == int32(MD_LINE_FENCEDCODE) {
			(*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Fbeg = *(*MD_OFFSET)(__ccgo_ptr(bp + 28))
			/* We are another MD_LINE_FENCEDCODE unless we are closing fence
			 * which we transform into MD_LINE_BLANK. */
			if (*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Findent < (*MD_CTX)(unsafe.Pointer(ctx)).Fcode_indent_offset {
				if md_is_closing_code_fence(tls, ctx, *(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr((*MD_LINE_ANALYSIS)(unsafe.Pointer(pivot_line)).Fbeg))), *(*MD_OFFSET)(__ccgo_ptr(bp + 28)), bp+28) != 0 {
					(*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Ftype1 = int32(MD_LINE_BLANK)
					(*MD_CTX)(unsafe.Pointer(ctx)).Flast_line_has_list_loosening_effect = FALSE
					break
				}
			}
			/* Change indentation accordingly to the initial code fence. */
			if n_parents == (*MD_CTX)(unsafe.Pointer(ctx)).Fn_containers {
				if (*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Findent > (*MD_LINE_ANALYSIS)(unsafe.Pointer(pivot_line)).Findent {
					*(*uint32)(__ccgo_ptr(line + 20)) -= (*MD_LINE_ANALYSIS)(unsafe.Pointer(pivot_line)).Findent
				} else {
					(*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Findent = uint32(0)
				}
				(*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Ftype1 = int32(MD_LINE_FENCEDCODE)
				break
			}
		}
		/* Check whether we are HTML block continuation. */
		if (*MD_LINE_ANALYSIS)(unsafe.Pointer(pivot_line)).Ftype1 == int32(MD_LINE_HTML) && (*MD_CTX)(unsafe.Pointer(ctx)).Fhtml_block_type > 0 {
			if n_parents < (*MD_CTX)(unsafe.Pointer(ctx)).Fn_containers {
				/* HTML block is implicitly ended if the enclosing container
				 * block ends. */
				(*MD_CTX)(unsafe.Pointer(ctx)).Fhtml_block_type = 0
			} else {
				html_block_type = md_is_html_block_end_condition(tls, ctx, *(*MD_OFFSET)(__ccgo_ptr(bp + 28)), bp+28)
				if html_block_type > 0 {
					if !(html_block_type == (*MD_CTX)(unsafe.Pointer(ctx)).Fhtml_block_type) {
						libc.X__builtin_unreachable(tls)
					}
					/* Make sure this is the last line of the block. */
					(*MD_CTX)(unsafe.Pointer(ctx)).Fhtml_block_type = 0
					/* Some end conditions serve as blank lines at the same time. */
					if html_block_type == int32(6) || html_block_type == int32(7) {
						(*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Ftype1 = int32(MD_LINE_BLANK)
						(*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Findent = uint32(0)
						break
					}
				}
				(*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Ftype1 = int32(MD_LINE_HTML)
				n_parents = (*MD_CTX)(unsafe.Pointer(ctx)).Fn_containers
				break
			}
		}
		/* Check for blank line. */
		if *(*MD_OFFSET)(__ccgo_ptr(bp + 28)) >= (*MD_CTX)(unsafe.Pointer(ctx)).Fsize || (libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(*(*MD_OFFSET)(__ccgo_ptr(bp + 28)))))) == int32('\r') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(*(*MD_OFFSET)(__ccgo_ptr(bp + 28)))))) == int32('\n')) {
			if (*MD_LINE_ANALYSIS)(unsafe.Pointer(pivot_line)).Ftype1 == int32(MD_LINE_INDENTEDCODE) && n_parents == (*MD_CTX)(unsafe.Pointer(ctx)).Fn_containers {
				(*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Ftype1 = int32(MD_LINE_INDENTEDCODE)
				if (*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Findent > (*MD_CTX)(unsafe.Pointer(ctx)).Fcode_indent_offset {
					*(*uint32)(__ccgo_ptr(line + 20)) -= (*MD_CTX)(unsafe.Pointer(ctx)).Fcode_indent_offset
				} else {
					(*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Findent = uint32(0)
				}
				(*MD_CTX)(unsafe.Pointer(ctx)).Flast_line_has_list_loosening_effect = FALSE
			} else {
				(*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Ftype1 = int32(MD_LINE_BLANK)
				(*MD_CTX)(unsafe.Pointer(ctx)).Flast_line_has_list_loosening_effect = libc.BoolInt32(n_parents > 0 && n_brothers+n_children == 0 && libc.Int32FromUint8((*(*MD_CONTAINER)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Fcontainers + uintptr(n_parents-int32(1))*28))).Fch) != int32('>'))
				/* See https://github.com/mity/md4c/issues/6
				 *
				 * This ugly checking tests we are in (yet empty) list item but
				 * not its very first line (i.e. not the line with the list
				 * item mark).
				 *
				 * If we are such a blank line, then any following non-blank
				 * line which would be part of the list item actually has to
				 * end the list because according to the specification, "a list
				 * item can begin with at most one blank line."
				 */
				if n_parents > 0 && libc.Int32FromUint8((*(*MD_CONTAINER)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Fcontainers + uintptr(n_parents-int32(1))*28))).Fch) != int32('>') && n_brothers+n_children == 0 && (*MD_CTX)(unsafe.Pointer(ctx)).Fcurrent_block == libc.UintptrFromInt32(0) && (*MD_CTX)(unsafe.Pointer(ctx)).Fn_block_bytes > libc.Int32FromInt64(8) {
					top_block = (*MD_CTX)(unsafe.Pointer(ctx)).Fblock_bytes + uintptr((*MD_CTX)(unsafe.Pointer(ctx)).Fn_block_bytes) - libc.UintptrFromInt64(8)
					if MD_BLOCKTYPE1(*(*uint32)(unsafe.Pointer(top_block + 0))&0xff>>0) == int32(MD_BLOCK_LI) {
						(*MD_CTX)(unsafe.Pointer(ctx)).Flast_list_item_starts_with_two_blank_lines = int32(TRUE)
					}
				}
			}
			break
		} else {
			/* This is the 2nd half of the hack. If the flag is set (i.e. there
			 * was a 2nd blank line at the beginning of the list item) and if
			 * we would otherwise still belong to the list item, we enforce
			 * the end of the list. */
			if (*MD_CTX)(unsafe.Pointer(ctx)).Flast_list_item_starts_with_two_blank_lines != 0 {
				if n_parents > 0 && n_parents == (*MD_CTX)(unsafe.Pointer(ctx)).Fn_containers && libc.Int32FromUint8((*(*MD_CONTAINER)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Fcontainers + uintptr(n_parents-int32(1))*28))).Fch) != int32('>') && n_brothers+n_children == 0 && (*MD_CTX)(unsafe.Pointer(ctx)).Fcurrent_block == libc.UintptrFromInt32(0) && (*MD_CTX)(unsafe.Pointer(ctx)).Fn_block_bytes > libc.Int32FromInt64(8) {
					top_block1 = (*MD_CTX)(unsafe.Pointer(ctx)).Fblock_bytes + uintptr((*MD_CTX)(unsafe.Pointer(ctx)).Fn_block_bytes) - libc.UintptrFromInt64(8)
					if MD_BLOCKTYPE1(*(*uint32)(unsafe.Pointer(top_block1 + 0))&0xff>>0) == int32(MD_BLOCK_LI) {
						n_parents = n_parents - 1
						(*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Findent = total_indent
						if n_parents > 0 {
							if (*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Findent < (*(*MD_CONTAINER)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Fcontainers + uintptr(n_parents-int32(1))*28))).Fcontents_indent {
								v1 = (*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Findent
							} else {
								v1 = (*(*MD_CONTAINER)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Fcontainers + uintptr(n_parents-int32(1))*28))).Fcontents_indent
							}
							*(*uint32)(__ccgo_ptr(line + 20)) -= v1
						}
					}
				}
				(*MD_CTX)(unsafe.Pointer(ctx)).Flast_list_item_starts_with_two_blank_lines = FALSE
			}
			(*MD_CTX)(unsafe.Pointer(ctx)).Flast_line_has_list_loosening_effect = FALSE
		}
		/* Check whether we are Setext underline. */
		if (*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Findent < (*MD_CTX)(unsafe.Pointer(ctx)).Fcode_indent_offset && (*MD_LINE_ANALYSIS)(unsafe.Pointer(pivot_line)).Ftype1 == int32(MD_LINE_TEXT) && *(*MD_OFFSET)(__ccgo_ptr(bp + 28)) < (*MD_CTX)(unsafe.Pointer(ctx)).Fsize && (libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(*(*MD_OFFSET)(__ccgo_ptr(bp + 28)))))) == int32('=') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(*(*MD_OFFSET)(__ccgo_ptr(bp + 28)))))) == int32('-')) && n_parents == (*MD_CTX)(unsafe.Pointer(ctx)).Fn_containers {
			if md_is_setext_underline(tls, ctx, *(*MD_OFFSET)(__ccgo_ptr(bp + 28)), bp+28, bp+36) != 0 {
				(*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Ftype1 = int32(MD_LINE_SETEXTUNDERLINE)
				(*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Fdata = *(*uint32)(__ccgo_ptr(bp + 36))
				break
			}
		}
		/* Check for thematic break line. */
		if (*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Findent < (*MD_CTX)(unsafe.Pointer(ctx)).Fcode_indent_offset && *(*MD_OFFSET)(__ccgo_ptr(bp + 28)) < (*MD_CTX)(unsafe.Pointer(ctx)).Fsize && *(*MD_OFFSET)(__ccgo_ptr(bp + 28)) >= *(*MD_OFFSET)(__ccgo_ptr(bp + 32)) && (libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(*(*MD_OFFSET)(__ccgo_ptr(bp + 28)))))) != int32('\000') && libc.Xstrchr(tls, __ccgo_ts+795, libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(*(*MD_OFFSET)(__ccgo_ptr(bp + 28))))))) != libc.UintptrFromInt32(0)) {
			if md_is_hr_line(tls, ctx, *(*MD_OFFSET)(__ccgo_ptr(bp + 28)), bp+28, bp+32) != 0 {
				(*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Ftype1 = int32(MD_LINE_HR)
				break
			}
		}
		/* Check for "brother" container. I.e. whether we are another list item
		 * in already started list. */
		if n_parents < (*MD_CTX)(unsafe.Pointer(ctx)).Fn_containers && n_brothers+n_children == 0 {
			if md_is_container_mark(tls, ctx, (*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Findent, *(*MD_OFFSET)(__ccgo_ptr(bp + 28)), bp+40, bp) != 0 && md_is_container_compatible(tls, (*MD_CTX)(unsafe.Pointer(ctx)).Fcontainers+uintptr(n_parents)*28, bp) != 0 {
				pivot_line = uintptr(unsafe.Pointer(&md_dummy_blank_line))
				*(*MD_OFFSET)(__ccgo_ptr(bp + 28)) = *(*MD_OFFSET)(__ccgo_ptr(bp + 40))
				total_indent = total_indent + ((*(*MD_CONTAINER)(__ccgo_ptr(bp))).Fcontents_indent - (*(*MD_CONTAINER)(__ccgo_ptr(bp))).Fmark_indent)
				(*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Findent = md_line_indentation(tls, ctx, total_indent, *(*MD_OFFSET)(__ccgo_ptr(bp + 28)), bp+28)
				total_indent = total_indent + (*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Findent
				(*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Fbeg = *(*MD_OFFSET)(__ccgo_ptr(bp + 28))
				/* Some of the following whitespace actually still belongs to the mark. */
				if *(*MD_OFFSET)(__ccgo_ptr(bp + 28)) >= (*MD_CTX)(unsafe.Pointer(ctx)).Fsize || (libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(*(*MD_OFFSET)(__ccgo_ptr(bp + 28)))))) == int32('\r') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(*(*MD_OFFSET)(__ccgo_ptr(bp + 28)))))) == int32('\n')) {
					(*(*MD_CONTAINER)(__ccgo_ptr(bp))).Fcontents_indent = (*(*MD_CONTAINER)(__ccgo_ptr(bp))).Fcontents_indent + 1
				} else {
					if (*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Findent <= (*MD_CTX)(unsafe.Pointer(ctx)).Fcode_indent_offset {
						(*(*MD_CONTAINER)(__ccgo_ptr(bp))).Fcontents_indent += (*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Findent
						(*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Findent = uint32(0)
					} else {
						(*(*MD_CONTAINER)(__ccgo_ptr(bp))).Fcontents_indent += uint32(1)
						(*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Findent = (*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Findent - 1
					}
				}
				(*(*MD_CONTAINER)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Fcontainers + uintptr(n_parents)*28))).Fmark_indent = (*(*MD_CONTAINER)(__ccgo_ptr(bp))).Fmark_indent
				(*(*MD_CONTAINER)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Fcontainers + uintptr(n_parents)*28))).Fcontents_indent = (*(*MD_CONTAINER)(__ccgo_ptr(bp))).Fcontents_indent
				n_brothers = n_brothers + 1
				continue
			}
		}
		/* Check for indented code.
		 * Note indented code block cannot interrupt a paragraph. */
		if (*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Findent >= (*MD_CTX)(unsafe.Pointer(ctx)).Fcode_indent_offset && (*MD_LINE_ANALYSIS)(unsafe.Pointer(pivot_line)).Ftype1 != int32(MD_LINE_TEXT) {
			(*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Ftype1 = int32(MD_LINE_INDENTEDCODE)
			*(*uint32)(__ccgo_ptr(line + 20)) -= (*MD_CTX)(unsafe.Pointer(ctx)).Fcode_indent_offset
			(*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Fdata = uint32(0)
			break
		}
		/* Check for start of a new container block. */
		if (*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Findent < (*MD_CTX)(unsafe.Pointer(ctx)).Fcode_indent_offset && md_is_container_mark(tls, ctx, (*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Findent, *(*MD_OFFSET)(__ccgo_ptr(bp + 28)), bp+28, bp) != 0 {
			if (*MD_LINE_ANALYSIS)(unsafe.Pointer(pivot_line)).Ftype1 == int32(MD_LINE_TEXT) && n_parents == (*MD_CTX)(unsafe.Pointer(ctx)).Fn_containers && (*(*MD_OFFSET)(__ccgo_ptr(bp + 28)) >= (*MD_CTX)(unsafe.Pointer(ctx)).Fsize || (libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(*(*MD_OFFSET)(__ccgo_ptr(bp + 28)))))) == int32('\r') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(*(*MD_OFFSET)(__ccgo_ptr(bp + 28)))))) == int32('\n'))) && libc.Int32FromUint8((*(*MD_CONTAINER)(__ccgo_ptr(bp))).Fch) != int32('>') {
				/* Noop. List mark followed by a blank line cannot interrupt a paragraph. */
			} else {
				if (*MD_LINE_ANALYSIS)(unsafe.Pointer(pivot_line)).Ftype1 == int32(MD_LINE_TEXT) && n_parents == (*MD_CTX)(unsafe.Pointer(ctx)).Fn_containers && (libc.Int32FromUint8((*(*MD_CONTAINER)(__ccgo_ptr(bp))).Fch) == int32('.') || libc.Int32FromUint8((*(*MD_CONTAINER)(__ccgo_ptr(bp))).Fch) == int32(')')) && (*(*MD_CONTAINER)(__ccgo_ptr(bp))).Fstart != uint32(1) {
					/* Noop. Ordered list cannot interrupt a paragraph unless the start index is 1. */
				} else {
					total_indent = total_indent + ((*(*MD_CONTAINER)(__ccgo_ptr(bp))).Fcontents_indent - (*(*MD_CONTAINER)(__ccgo_ptr(bp))).Fmark_indent)
					(*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Findent = md_line_indentation(tls, ctx, total_indent, *(*MD_OFFSET)(__ccgo_ptr(bp + 28)), bp+28)
					total_indent = total_indent + (*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Findent
					(*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Fbeg = *(*MD_OFFSET)(__ccgo_ptr(bp + 28))
					(*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Fdata = uint32((*(*MD_CONTAINER)(__ccgo_ptr(bp))).Fch)
					/* Some of the following whitespace actually still belongs to the mark. */
					if *(*MD_OFFSET)(__ccgo_ptr(bp + 28)) >= (*MD_CTX)(unsafe.Pointer(ctx)).Fsize || (libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(*(*MD_OFFSET)(__ccgo_ptr(bp + 28)))))) == int32('\r') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(*(*MD_OFFSET)(__ccgo_ptr(bp + 28)))))) == int32('\n')) {
						(*(*MD_CONTAINER)(__ccgo_ptr(bp))).Fcontents_indent = (*(*MD_CONTAINER)(__ccgo_ptr(bp))).Fcontents_indent + 1
					} else {
						if (*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Findent <= (*MD_CTX)(unsafe.Pointer(ctx)).Fcode_indent_offset {
							(*(*MD_CONTAINER)(__ccgo_ptr(bp))).Fcontents_indent += (*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Findent
							(*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Findent = uint32(0)
						} else {
							(*(*MD_CONTAINER)(__ccgo_ptr(bp))).Fcontents_indent += uint32(1)
							(*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Findent = (*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Findent - 1
						}
					}
					if n_brothers+n_children == 0 {
						pivot_line = uintptr(unsafe.Pointer(&md_dummy_blank_line))
					}
					if n_children == 0 {
						ret = md_leave_child_containers(tls, ctx, n_parents+n_brothers)
						if ret < 0 {
							goto abort
						}
					}
					n_children = n_children + 1
					ret = md_push_container(tls, ctx, bp)
					if ret < 0 {
						goto abort
					}
					continue
				}
			}
		}
		/* Check whether we are table continuation. */
		if (*MD_LINE_ANALYSIS)(unsafe.Pointer(pivot_line)).Ftype1 == int32(MD_LINE_TABLE) && n_parents == (*MD_CTX)(unsafe.Pointer(ctx)).Fn_containers {
			(*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Ftype1 = int32(MD_LINE_TABLE)
			break
		}
		/* Check for ATX header. */
		if (*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Findent < (*MD_CTX)(unsafe.Pointer(ctx)).Fcode_indent_offset && *(*MD_OFFSET)(__ccgo_ptr(bp + 28)) < (*MD_CTX)(unsafe.Pointer(ctx)).Fsize && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(*(*MD_OFFSET)(__ccgo_ptr(bp + 28)))))) == int32('#') {
			if md_is_atxheader_line(tls, ctx, *(*MD_OFFSET)(__ccgo_ptr(bp + 28)), line+12, bp+28, bp+44) != 0 {
				(*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Ftype1 = int32(MD_LINE_ATXHEADER)
				(*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Fdata = *(*uint32)(__ccgo_ptr(bp + 44))
				break
			}
		}
		/* Check whether we are starting code fence. */
		if (*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Findent < (*MD_CTX)(unsafe.Pointer(ctx)).Fcode_indent_offset && *(*MD_OFFSET)(__ccgo_ptr(bp + 28)) < (*MD_CTX)(unsafe.Pointer(ctx)).Fsize && (libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(*(*MD_OFFSET)(__ccgo_ptr(bp + 28)))))) == int32('`') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(*(*MD_OFFSET)(__ccgo_ptr(bp + 28)))))) == int32('~')) {
			if md_is_opening_code_fence(tls, ctx, *(*MD_OFFSET)(__ccgo_ptr(bp + 28)), bp+28) != 0 {
				(*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Ftype1 = int32(MD_LINE_FENCEDCODE)
				(*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Fdata = uint32(1)
				(*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Fenforce_new_block = int32(TRUE)
				break
			}
		}
		/* Check for start of raw HTML block. */
		if *(*MD_OFFSET)(__ccgo_ptr(bp + 28)) < (*MD_CTX)(unsafe.Pointer(ctx)).Fsize && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(*(*MD_OFFSET)(__ccgo_ptr(bp + 28)))))) == int32('<') && !((*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fflags&libc.Uint32FromInt32(MD_FLAG_NOHTMLBLOCKS) != 0) {
			(*MD_CTX)(unsafe.Pointer(ctx)).Fhtml_block_type = md_is_html_block_start_condition(tls, ctx, *(*MD_OFFSET)(__ccgo_ptr(bp + 28)))
			/* HTML block type 7 cannot interrupt paragraph. */
			if (*MD_CTX)(unsafe.Pointer(ctx)).Fhtml_block_type == int32(7) && (*MD_LINE_ANALYSIS)(unsafe.Pointer(pivot_line)).Ftype1 == int32(MD_LINE_TEXT) {
				(*MD_CTX)(unsafe.Pointer(ctx)).Fhtml_block_type = 0
			}
			if (*MD_CTX)(unsafe.Pointer(ctx)).Fhtml_block_type > 0 {
				/* The line itself also may immediately close the block. */
				if md_is_html_block_end_condition(tls, ctx, *(*MD_OFFSET)(__ccgo_ptr(bp + 28)), bp+28) == (*MD_CTX)(unsafe.Pointer(ctx)).Fhtml_block_type {
					/* Make sure this is the last line of the block. */
					(*MD_CTX)(unsafe.Pointer(ctx)).Fhtml_block_type = 0
				}
				(*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Fenforce_new_block = int32(TRUE)
				(*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Ftype1 = int32(MD_LINE_HTML)
				break
			}
		}
		/* Check for table underline. */
		if (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fflags&uint32(MD_FLAG_TABLES) != 0 && (*MD_LINE_ANALYSIS)(unsafe.Pointer(pivot_line)).Ftype1 == int32(MD_LINE_TEXT) && *(*MD_OFFSET)(__ccgo_ptr(bp + 28)) < (*MD_CTX)(unsafe.Pointer(ctx)).Fsize && (libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(*(*MD_OFFSET)(__ccgo_ptr(bp + 28)))))) == int32('|') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(*(*MD_OFFSET)(__ccgo_ptr(bp + 28)))))) == int32('-') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(*(*MD_OFFSET)(__ccgo_ptr(bp + 28)))))) == int32(':')) && n_parents == (*MD_CTX)(unsafe.Pointer(ctx)).Fn_containers {
			if (*MD_CTX)(unsafe.Pointer(ctx)).Fcurrent_block != libc.UintptrFromInt32(0) && (*MD_BLOCK)(unsafe.Pointer((*MD_CTX)(unsafe.Pointer(ctx)).Fcurrent_block)).Fn_lines == uint32(1) && md_is_table_underline(tls, ctx, *(*MD_OFFSET)(__ccgo_ptr(bp + 28)), bp+28, bp+48) != 0 {
				(*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Fdata = *(*uint32)(__ccgo_ptr(bp + 48))
				(*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Ftype1 = int32(MD_LINE_TABLEUNDERLINE)
				break
			}
		}
		/* By default, we are normal text line. */
		(*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Ftype1 = int32(MD_LINE_TEXT)
		if (*MD_LINE_ANALYSIS)(unsafe.Pointer(pivot_line)).Ftype1 == int32(MD_LINE_TEXT) && n_brothers+n_children == 0 {
			/* Lazy continuation. */
			n_parents = (*MD_CTX)(unsafe.Pointer(ctx)).Fn_containers
		}
		/* Check for task mark. */
		if (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fflags&uint32(MD_FLAG_TASKLISTS) != 0 && n_brothers+n_children > 0 && (libc.Int32FromUint8((*(*MD_CONTAINER)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Fcontainers + uintptr((*MD_CTX)(unsafe.Pointer(ctx)).Fn_containers-int32(1))*28))).Fch) != int32('\000') && libc.Xstrchr(tls, __ccgo_ts+799, libc.Int32FromUint8((*(*MD_CONTAINER)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Fcontainers + uintptr((*MD_CTX)(unsafe.Pointer(ctx)).Fn_containers-int32(1))*28))).Fch)) != libc.UintptrFromInt32(0)) {
			tmp1 = *(*MD_OFFSET)(__ccgo_ptr(bp + 28))
			for tmp1 < (*MD_CTX)(unsafe.Pointer(ctx)).Fsize && tmp1 < *(*MD_OFFSET)(__ccgo_ptr(bp + 28))+uint32(3) && (libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(tmp1)))) == int32(' ') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(tmp1)))) == int32('\t')) {
				tmp1 = tmp1 + 1
			}
			if tmp1+uint32(2) < (*MD_CTX)(unsafe.Pointer(ctx)).Fsize && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(tmp1)))) == int32('[') && (libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(tmp1+libc.Uint32FromInt32(1))))) != int32('\000') && libc.Xstrchr(tls, __ccgo_ts+805, libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(tmp1+libc.Uint32FromInt32(1)))))) != libc.UintptrFromInt32(0)) && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(tmp1+libc.Uint32FromInt32(2))))) == int32(']') && (tmp1+uint32(3) == (*MD_CTX)(unsafe.Pointer(ctx)).Fsize || (libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(tmp1+libc.Uint32FromInt32(3))))) == int32(' ') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(tmp1+libc.Uint32FromInt32(3))))) == int32('\t')) || (libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(tmp1+libc.Uint32FromInt32(3))))) == int32('\r') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(tmp1+libc.Uint32FromInt32(3))))) == int32('\n'))) {
				if n_children > 0 {
					v2 = (*MD_CTX)(unsafe.Pointer(ctx)).Fcontainers + uintptr((*MD_CTX)(unsafe.Pointer(ctx)).Fn_containers-int32(1))*28
				} else {
					v2 = bp
				}
				task_container = v2
				libc.SetBitFieldPtr16Uint32(task_container+4, libc.Uint32FromInt32(TRUE), 8, 0xff00)
				(*MD_CONTAINER)(unsafe.Pointer(task_container)).Ftask_mark_off = tmp1 + uint32(1)
				*(*MD_OFFSET)(__ccgo_ptr(bp + 28)) = tmp1 + uint32(3)
				for *(*MD_OFFSET)(__ccgo_ptr(bp + 28)) < (*MD_CTX)(unsafe.Pointer(ctx)).Fsize && (libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(*(*MD_OFFSET)(__ccgo_ptr(bp + 28)))))) == int32(' ') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(*(*MD_OFFSET)(__ccgo_ptr(bp + 28)))))) == int32('\t') || (libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(*(*MD_OFFSET)(__ccgo_ptr(bp + 28)))))) == int32('\v') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(*(*MD_OFFSET)(__ccgo_ptr(bp + 28)))))) == int32('\f'))) {
					*(*MD_OFFSET)(__ccgo_ptr(bp + 28)) = *(*MD_OFFSET)(__ccgo_ptr(bp + 28)) + 1
				}
				(*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Fbeg = *(*MD_OFFSET)(__ccgo_ptr(bp + 28))
			}
		}
		break
	}
	/* Scan for end of the line.
	 *
	 * Note this is quite a bottleneck of the parsing as we here iterate almost
	 * over compete document.
	 */
	/* Recent glibc versions have superbly optimized strcspn(), even using
	 * vectorization if available. */
	if (*MD_CTX)(unsafe.Pointer(ctx)).Fdoc_ends_with_newline != 0 && *(*MD_OFFSET)(__ccgo_ptr(bp + 28)) < (*MD_CTX)(unsafe.Pointer(ctx)).Fsize {
		for int32(TRUE) != 0 {
			*(*MD_OFFSET)(__ccgo_ptr(bp + 28)) = *(*MD_OFFSET)(__ccgo_ptr(bp + 28)) + uint32(libc.Xstrcspn(tls, (*MD_CTX)(unsafe.Pointer(ctx)).Ftext+uintptr(*(*MD_OFFSET)(__ccgo_ptr(bp + 28))), __ccgo_ts+809))
			/* strcspn() can stop on zero terminator; but that can appear
			 * anywhere in the Markfown input... */
			if libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(*(*MD_OFFSET)(__ccgo_ptr(bp + 28)))))) == int32('\000') {
				*(*MD_OFFSET)(__ccgo_ptr(bp + 28)) = *(*MD_OFFSET)(__ccgo_ptr(bp + 28)) + 1
			} else {
				break
			}
		}
	} else {
		/* Optimization: Use some loop unrolling. */
		for *(*MD_OFFSET)(__ccgo_ptr(bp + 28))+uint32(3) < (*MD_CTX)(unsafe.Pointer(ctx)).Fsize && !(libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(*(*MD_OFFSET)(__ccgo_ptr(bp + 28))+libc.Uint32FromInt32(0))))) == int32('\r') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(*(*MD_OFFSET)(__ccgo_ptr(bp + 28))+libc.Uint32FromInt32(0))))) == int32('\n')) && !(libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(*(*MD_OFFSET)(__ccgo_ptr(bp + 28))+libc.Uint32FromInt32(1))))) == int32('\r') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(*(*MD_OFFSET)(__ccgo_ptr(bp + 28))+libc.Uint32FromInt32(1))))) == int32('\n')) && !(libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(*(*MD_OFFSET)(__ccgo_ptr(bp + 28))+libc.Uint32FromInt32(2))))) == int32('\r') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(*(*MD_OFFSET)(__ccgo_ptr(bp + 28))+libc.Uint32FromInt32(2))))) == int32('\n')) && !(libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(*(*MD_OFFSET)(__ccgo_ptr(bp + 28))+libc.Uint32FromInt32(3))))) == int32('\r') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(*(*MD_OFFSET)(__ccgo_ptr(bp + 28))+libc.Uint32FromInt32(3))))) == int32('\n')) {
			*(*MD_OFFSET)(__ccgo_ptr(bp + 28)) = *(*MD_OFFSET)(__ccgo_ptr(bp + 28)) + uint32(4)
		}
		for *(*MD_OFFSET)(__ccgo_ptr(bp + 28)) < (*MD_CTX)(unsafe.Pointer(ctx)).Fsize && !(libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(*(*MD_OFFSET)(__ccgo_ptr(bp + 28)))))) == int32('\r') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(*(*MD_OFFSET)(__ccgo_ptr(bp + 28)))))) == int32('\n')) {
			*(*MD_OFFSET)(__ccgo_ptr(bp + 28)) = *(*MD_OFFSET)(__ccgo_ptr(bp + 28)) + 1
		}
	}
	/* Set end of the line. */
	(*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Fend = *(*MD_OFFSET)(__ccgo_ptr(bp + 28))
	/* But for ATX header, we should exclude the optional trailing mark. */
	if (*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Ftype1 == int32(MD_LINE_ATXHEADER) {
		tmp2 = (*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Fend
		for tmp2 > (*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Fbeg && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(tmp2-libc.Uint32FromInt32(1))))) == int32(' ') {
			tmp2 = tmp2 - 1
		}
		for tmp2 > (*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Fbeg && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(tmp2-libc.Uint32FromInt32(1))))) == int32('#') {
			tmp2 = tmp2 - 1
		}
		if tmp2 == (*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Fbeg || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(tmp2-libc.Uint32FromInt32(1))))) == int32(' ') || (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fflags&uint32(MD_FLAG_PERMISSIVEATXHEADERS) != 0 {
			(*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Fend = tmp2
		}
	}
	/* Trim trailing spaces. */
	if (*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Ftype1 != int32(MD_LINE_INDENTEDCODE) && (*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Ftype1 != int32(MD_LINE_FENCEDCODE) && (*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Ftype1 != int32(MD_LINE_HTML) {
		for (*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Fend > (*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Fbeg && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr((*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Fend-libc.Uint32FromInt32(1))))) == int32(' ') {
			(*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Fend = (*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Fend - 1
		}
	}
	/* Eat also the new line. */
	if *(*MD_OFFSET)(__ccgo_ptr(bp + 28)) < (*MD_CTX)(unsafe.Pointer(ctx)).Fsize && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(*(*MD_OFFSET)(__ccgo_ptr(bp + 28)))))) == int32('\r') {
		*(*MD_OFFSET)(__ccgo_ptr(bp + 28)) = *(*MD_OFFSET)(__ccgo_ptr(bp + 28)) + 1
	}
	if *(*MD_OFFSET)(__ccgo_ptr(bp + 28)) < (*MD_CTX)(unsafe.Pointer(ctx)).Fsize && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr(*(*MD_OFFSET)(__ccgo_ptr(bp + 28)))))) == int32('\n') {
		*(*MD_OFFSET)(__ccgo_ptr(bp + 28)) = *(*MD_OFFSET)(__ccgo_ptr(bp + 28)) + 1
	}
	*(*MD_OFFSET)(__ccgo_ptr(p_end)) = *(*MD_OFFSET)(__ccgo_ptr(bp + 28))
	/* If we belong to a list after seeing a blank line, the list is loose. */
	if prev_line_has_list_loosening_effect != 0 && (*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Ftype1 != int32(MD_LINE_BLANK) && n_parents+n_brothers > 0 {
		c1 = (*MD_CTX)(unsafe.Pointer(ctx)).Fcontainers + uintptr(n_parents+n_brothers-int32(1))*28
		if libc.Int32FromUint8((*MD_CONTAINER)(unsafe.Pointer(c1)).Fch) != int32('>') {
			block = (*MD_CTX)(unsafe.Pointer(ctx)).Fblock_bytes + uintptr((*MD_CONTAINER)(unsafe.Pointer(c1)).Fblock_byte_off)
			libc.AssignBitFieldPtr32Int32(block+0, int32(*(*uint32)(unsafe.Pointer(block + 0))&0xff00>>8)|int32(libc.Int32FromInt32(MD_BLOCK_LOOSE_LIST)), 8, 8, 0xff00)
		}
	}
	/* Leave any containers we are not part of anymore. */
	if n_children == 0 && n_parents+n_brothers < (*MD_CTX)(unsafe.Pointer(ctx)).Fn_containers {
		ret = md_leave_child_containers(tls, ctx, n_parents+n_brothers)
		if ret < 0 {
			goto abort
		}
	}
	/* Enter any container we found a mark for. */
	if n_brothers > 0 {
		if !(n_brothers == libc.Int32FromInt32(1)) {
			libc.X__builtin_unreachable(tls)
		}
		if int32(uint32(*(*uint16)(unsafe.Pointer((*MD_CTX)(unsafe.Pointer(ctx)).Fcontainers + uintptr(n_parents)*28 + 4))&0xff00>>8)) != 0 {
			v3 = libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr((*(*MD_CONTAINER)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Fcontainers + uintptr(n_parents)*28))).Ftask_mark_off))))
		} else {
			v3 = 0
		}
		ret = md_push_container_bytes(tls, ctx, int32(MD_BLOCK_LI), (*(*MD_CONTAINER)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Fcontainers + uintptr(n_parents)*28))).Ftask_mark_off, libc.Uint32FromInt32(v3), uint32(MD_BLOCK_CONTAINER_CLOSER))
		if ret < 0 {
			goto abort
		}
		if int32(uint32(*(*uint16)(unsafe.Pointer(bp + 4))&0xff00>>8)) != 0 {
			v3 = libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Ftext + uintptr((*(*MD_CONTAINER)(__ccgo_ptr(bp))).Ftask_mark_off))))
		} else {
			v3 = 0
		}
		ret = md_push_container_bytes(tls, ctx, int32(MD_BLOCK_LI), (*(*MD_CONTAINER)(__ccgo_ptr(bp))).Ftask_mark_off, libc.Uint32FromInt32(v3), uint32(MD_BLOCK_CONTAINER_OPENER))
		if ret < 0 {
			goto abort
		}
		libc.SetBitFieldPtr16Uint32((*MD_CTX)(unsafe.Pointer(ctx)).Fcontainers+uintptr(n_parents)*28+4, libc.Uint32FromInt32(int32(uint32(*(*uint16)(unsafe.Pointer(bp + 4))&0xff00>>8))), 8, 0xff00)
		(*(*MD_CONTAINER)(__ccgo_ptr((*MD_CTX)(unsafe.Pointer(ctx)).Fcontainers + uintptr(n_parents)*28))).Ftask_mark_off = (*(*MD_CONTAINER)(__ccgo_ptr(bp))).Ftask_mark_off
	}
	if n_children > 0 {
		ret = md_enter_child_containers(tls, ctx, n_children)
		if ret < 0 {
			goto abort
		}
	}
	goto abort
abort:
	;
	return ret
	return r
}

func md_process_line(tls *libc.TLS, ctx uintptr, p_pivot_line uintptr, line uintptr) (r int32) {
	var pivot_line uintptr
	var ret int32
	_, _ = pivot_line, ret
	pivot_line = *(*uintptr)(__ccgo_ptr(p_pivot_line))
	ret = 0
	/* Blank line ends current leaf block. */
	if (*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Ftype1 == int32(MD_LINE_BLANK) {
		ret = md_end_current_block(tls, ctx)
		if ret < 0 {
			goto abort
		}
		*(*uintptr)(__ccgo_ptr(p_pivot_line)) = uintptr(unsafe.Pointer(&md_dummy_blank_line))
		return 0
	}
	if (*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Fenforce_new_block != 0 {
		ret = md_end_current_block(tls, ctx)
		if ret < 0 {
			goto abort
		}
	}
	/* Some line types form block on their own. */
	if (*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Ftype1 == int32(MD_LINE_HR) || (*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Ftype1 == int32(MD_LINE_ATXHEADER) {
		ret = md_end_current_block(tls, ctx)
		if ret < 0 {
			goto abort
		}
		/* Add our single-line block. */
		ret = md_start_new_block(tls, ctx, line)
		if ret < 0 {
			goto abort
		}
		ret = md_add_line_into_current_block(tls, ctx, line)
		if ret < 0 {
			goto abort
		}
		ret = md_end_current_block(tls, ctx)
		if ret < 0 {
			goto abort
		}
		*(*uintptr)(__ccgo_ptr(p_pivot_line)) = uintptr(unsafe.Pointer(&md_dummy_blank_line))
		return 0
	}
	/* MD_LINE_SETEXTUNDERLINE changes meaning of the current block and ends it. */
	if (*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Ftype1 == int32(MD_LINE_SETEXTUNDERLINE) {
		if !((*MD_CTX)(unsafe.Pointer(ctx)).Fcurrent_block != libc.UintptrFromInt32(0)) {
			libc.X__builtin_unreachable(tls)
		}
		libc.SetBitFieldPtr32Int32((*MD_CTX)(unsafe.Pointer(ctx)).Fcurrent_block+0, int32(MD_BLOCK_H), 0, 0xff)
		libc.SetBitFieldPtr32Uint32((*MD_CTX)(unsafe.Pointer(ctx)).Fcurrent_block+0, (*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Fdata, 16, 0xffff0000)
		libc.AssignBitFieldPtr32Int32((*MD_CTX)(unsafe.Pointer(ctx)).Fcurrent_block+0, int32(*(*uint32)(unsafe.Pointer((*MD_CTX)(unsafe.Pointer(ctx)).Fcurrent_block + 0))&0xff00>>8)|int32(libc.Int32FromInt32(MD_BLOCK_SETEXT_HEADER)), 8, 8, 0xff00)
		ret = md_add_line_into_current_block(tls, ctx, line)
		if ret < 0 {
			goto abort
		}
		ret = md_end_current_block(tls, ctx)
		if ret < 0 {
			goto abort
		}
		if (*MD_CTX)(unsafe.Pointer(ctx)).Fcurrent_block == libc.UintptrFromInt32(0) {
			*(*uintptr)(__ccgo_ptr(p_pivot_line)) = uintptr(unsafe.Pointer(&md_dummy_blank_line))
		} else {
			/* This happens if we have consumed all the body as link ref. defs.
			 * and downgraded the underline into start of a new paragraph block. */
			(*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Ftype1 = int32(MD_LINE_TEXT)
			*(*uintptr)(__ccgo_ptr(p_pivot_line)) = line
		}
		return 0
	}
	/* MD_LINE_TABLEUNDERLINE changes meaning of the current block. */
	if (*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Ftype1 == int32(MD_LINE_TABLEUNDERLINE) {
		if !((*MD_CTX)(unsafe.Pointer(ctx)).Fcurrent_block != libc.UintptrFromInt32(0)) {
			libc.X__builtin_unreachable(tls)
		}
		if !((*MD_BLOCK)(unsafe.Pointer((*MD_CTX)(unsafe.Pointer(ctx)).Fcurrent_block)).Fn_lines == libc.Uint32FromInt32(1)) {
			libc.X__builtin_unreachable(tls)
		}
		libc.SetBitFieldPtr32Int32((*MD_CTX)(unsafe.Pointer(ctx)).Fcurrent_block+0, int32(MD_BLOCK_TABLE), 0, 0xff)
		libc.SetBitFieldPtr32Uint32((*MD_CTX)(unsafe.Pointer(ctx)).Fcurrent_block+0, (*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Fdata, 16, 0xffff0000)
		if !(pivot_line != uintptr(unsafe.Pointer(&md_dummy_blank_line))) {
			libc.X__builtin_unreachable(tls)
		}
		(*MD_LINE_ANALYSIS)(unsafe.Pointer(pivot_line)).Ftype1 = int32(MD_LINE_TABLE)
		ret = md_add_line_into_current_block(tls, ctx, line)
		if ret < 0 {
			goto abort
		}
		return 0
	}
	/* The current block also ends if the line has different type. */
	if (*MD_LINE_ANALYSIS)(unsafe.Pointer(line)).Ftype1 != (*MD_LINE_ANALYSIS)(unsafe.Pointer(pivot_line)).Ftype1 {
		ret = md_end_current_block(tls, ctx)
		if ret < 0 {
			goto abort
		}
	}
	/* The current line may start a new block. */
	if (*MD_CTX)(unsafe.Pointer(ctx)).Fcurrent_block == libc.UintptrFromInt32(0) {
		ret = md_start_new_block(tls, ctx, line)
		if ret < 0 {
			goto abort
		}
		*(*uintptr)(__ccgo_ptr(p_pivot_line)) = line
	}
	/* In all other cases the line is just a continuation of the current block. */
	ret = md_add_line_into_current_block(tls, ctx, line)
	if ret < 0 {
		goto abort
	}
	goto abort
abort:
	;
	return ret
	return r
}

func md_process_doc(tls *libc.TLS, ctx uintptr) (r int32) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var line, v1 uintptr
	var ret int32
	var _ /* line_buf at bp+8 */ [2]MD_LINE_ANALYSIS
	var _ /* off at bp+56 */ MD_OFFSET
	var _ /* pivot_line at bp+0 */ uintptr
	_, _, _ = line, ret, v1
	*(*uintptr)(__ccgo_ptr(bp)) = uintptr(unsafe.Pointer(&md_dummy_blank_line))
	line = bp + 8
	*(*MD_OFFSET)(__ccgo_ptr(bp + 56)) = uint32(0)
	ret = 0
	ret = (*(*func(*libc.TLS, MD_BLOCKTYPE1, uintptr, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fenter_block})))(tls, int32(MD_BLOCK_DOC), libc.UintptrFromInt32(0), (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
	if ret != 0 {
		if (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log != libc.UintptrFromInt32(0) {
			(*(*func(*libc.TLS, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log})))(tls, __ccgo_ts+288, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
		}
		goto abort
	}
	for *(*MD_OFFSET)(__ccgo_ptr(bp + 56)) < (*MD_CTX)(unsafe.Pointer(ctx)).Fsize {
		if line == *(*uintptr)(__ccgo_ptr(bp)) {
			if line == bp+8 {
				v1 = bp + 8 + 1*24
			} else {
				v1 = bp + 8
			}
			line = v1
		}
		ret = md_analyze_line(tls, ctx, *(*MD_OFFSET)(__ccgo_ptr(bp + 56)), bp+56, *(*uintptr)(__ccgo_ptr(bp)), line)
		if ret < 0 {
			goto abort
		}
		ret = md_process_line(tls, ctx, bp, line)
		if ret < 0 {
			goto abort
		}
	}
	md_end_current_block(tls, ctx)
	ret = md_build_ref_def_hashtable(tls, ctx)
	if ret < 0 {
		goto abort
	}
	/* Process all blocks. */
	ret = md_leave_child_containers(tls, ctx, 0)
	if ret < 0 {
		goto abort
	}
	ret = md_process_all_blocks(tls, ctx)
	if ret < 0 {
		goto abort
	}
	ret = (*(*func(*libc.TLS, MD_BLOCKTYPE1, uintptr, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fleave_block})))(tls, int32(MD_BLOCK_DOC), libc.UintptrFromInt32(0), (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
	if ret != 0 {
		if (*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log != libc.UintptrFromInt32(0) {
			(*(*func(*libc.TLS, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*MD_CTX)(unsafe.Pointer(ctx)).Fparser.Fdebug_log})))(tls, __ccgo_ts+325, (*MD_CTX)(unsafe.Pointer(ctx)).Fuserdata)
		}
		goto abort
	}
	goto abort
abort:
	;
	return ret
	return r
}

/********************
 ***  Public API  ***
 ********************/

func md_parse(tls *libc.TLS, text uintptr, size MD_SIZE, parser uintptr, userdata uintptr) (r int32) {
	bp := tls.Alloc(592)
	defer tls.Free(592)
	var i, ret int32
	var v1 uint32
	var _ /* ctx at bp+0 */ MD_CTX
	_, _, _ = i, ret, v1
	if (*MD_PARSER)(unsafe.Pointer(parser)).Fabi_version != uint32(0) {
		if (*MD_PARSER)(unsafe.Pointer(parser)).Fdebug_log != libc.UintptrFromInt32(0) {
			(*(*func(*libc.TLS, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*MD_PARSER)(unsafe.Pointer(parser)).Fdebug_log})))(tls, __ccgo_ts+812, userdata)
		}
		return -int32(1)
	}
	/* Setup context structure. */
	libc.Xmemset(tls, bp, 0, uint64(584))
	(*(*MD_CTX)(__ccgo_ptr(bp))).Ftext = text
	(*(*MD_CTX)(__ccgo_ptr(bp))).Fsize = size
	libc.Xmemcpy(tls, bp+16, parser, uint64(64))
	(*(*MD_CTX)(__ccgo_ptr(bp))).Fuserdata = userdata
	if (*(*MD_CTX)(__ccgo_ptr(bp))).Fparser.Fflags&uint32(MD_FLAG_NOINDENTEDCODEBLOCKS) != 0 {
		v1 = libc.Uint32FromInt32(-libc.Int32FromInt32(1))
	} else {
		v1 = uint32(4)
	}
	(*(*MD_CTX)(__ccgo_ptr(bp))).Fcode_indent_offset = v1
	md_build_mark_char_map(tls, bp)
	(*(*MD_CTX)(__ccgo_ptr(bp))).Fdoc_ends_with_newline = libc.BoolInt32(size > uint32(0) && (libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr(text + uintptr(size-uint32(1))))) == int32('\r') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr(text + uintptr(size-uint32(1))))) == int32('\n')))
	/* Reset all mark stacks and lists. */
	i = 0
	for {
		if !(i < libc.Int32FromUint64(libc.Uint64FromInt64(64)/libc.Uint64FromInt64(4))) {
			break
		}
		(*(*MD_MARKSTACK)(__ccgo_ptr(bp + 416 + uintptr(i)*4))).Ftop = -int32(1)
		goto _2
	_2:
		;
		i = i + 1
	}
	(*(*MD_CTX)(__ccgo_ptr(bp))).Fptr_stack.Ftop = -int32(1)
	(*(*MD_CTX)(__ccgo_ptr(bp))).Funresolved_link_head = -int32(1)
	(*(*MD_CTX)(__ccgo_ptr(bp))).Funresolved_link_tail = -int32(1)
	(*(*MD_CTX)(__ccgo_ptr(bp))).Ftable_cell_boundaries_head = -int32(1)
	(*(*MD_CTX)(__ccgo_ptr(bp))).Ftable_cell_boundaries_tail = -int32(1)
	/* All the work. */
	ret = md_process_doc(tls, bp)
	/* Clean-up. */
	md_free_ref_defs(tls, bp)
	md_free_ref_def_hashtable(tls, bp)
	libc.Xfree(tls, (*(*MD_CTX)(__ccgo_ptr(bp))).Fbuffer)
	libc.Xfree(tls, (*(*MD_CTX)(__ccgo_ptr(bp))).Fmarks)
	libc.Xfree(tls, (*(*MD_CTX)(__ccgo_ptr(bp))).Fblock_bytes)
	libc.Xfree(tls, (*(*MD_CTX)(__ccgo_ptr(bp))).Fcontainers)
	return ret
}

const MD_FLAG_COLLAPSEWHITESPACE1 = 0x0001
const MD_FLAG_HARD_SOFT_BREAKS1 = 0x8000
const MD_FLAG_LATEXMATHSPANS1 = 0x1000
const MD_FLAG_NOHTMLBLOCKS1 = 0x0020
const MD_FLAG_NOHTMLSPANS1 = 0x0040
const MD_FLAG_NOINDENTEDCODEBLOCKS1 = 0x0010
const MD_FLAG_PERMISSIVEATXHEADERS1 = 0x0002
const MD_FLAG_PERMISSIVEEMAILAUTOLINKS1 = 0x0008
const MD_FLAG_PERMISSIVEURLAUTOLINKS1 = 0x0004
const MD_FLAG_PERMISSIVEWWWAUTOLINKS1 = 0x0400
const MD_FLAG_STRIKETHROUGH1 = 0x0200
const MD_FLAG_TABLES1 = 0x0100
const MD_FLAG_TASKLISTS1 = 0x0800
const MD_FLAG_UNDERLINE1 = 0x4000
const MD_FLAG_WIKILINKS1 = 0x2000
const MD_HTML_FLAG_DEBUG = 1
const MD_HTML_FLAG_SKIP_UTF8_BOM = 4
const MD_HTML_FLAG_VERBATIM_ENTITIES = 2
const MD_HTML_FLAG_XHTML = 8
const NEED_HTML_ESC_FLAG = 1
const NEED_URL_ESC_FLAG = 2

type ENTITY = struct {
	Fname       uintptr
	Fcodepoints [2]uint32
}

type ENTITY_tag = ENTITY

type MD_HTML = struct {
	Fprocess_output      uintptr
	Fuserdata            uintptr
	Fflags               uint32
	Fimage_nesting_level int32
	Fescape_map          [256]uint8
}

type MD_HTML_tag = MD_HTML

/*****************************************
 ***  HTML rendering helper functions  ***
 *****************************************/

func render_verbatim(tls *libc.TLS, r uintptr, text uintptr, size MD_SIZE) {
	(*(*func(*libc.TLS, uintptr, MD_SIZE, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*MD_HTML)(unsafe.Pointer(r)).Fprocess_output})))(tls, text, size, (*MD_HTML)(unsafe.Pointer(r)).Fuserdata)
}

/* Keep this as a macro. Most compiler should then be smart enough to replace
 * the strlen() call with a compile-time constant if the string is a C literal. */

func render_html_escaped(tls *libc.TLS, r uintptr, data uintptr, size MD_SIZE) {
	var beg, off MD_OFFSET
	_, _ = beg, off
	beg = uint32(0)
	off = uint32(0)
	/* Some characters need to be escaped in normal HTML text. */
	for int32(1) != 0 {
		/* Optimization: Use some loop unrolling. */
		for off+uint32(3) < size && !(libc.Int32FromUint8(*(*uint8)(__ccgo_ptr(r + 24 + uintptr(*(*MD_CHAR)(__ccgo_ptr(data + uintptr(off+uint32(0))))))))&libc.Int32FromInt32(NEED_HTML_ESC_FLAG) != 0) && !(libc.Int32FromUint8(*(*uint8)(__ccgo_ptr(r + 24 + uintptr(*(*MD_CHAR)(__ccgo_ptr(data + uintptr(off+uint32(1))))))))&libc.Int32FromInt32(NEED_HTML_ESC_FLAG) != 0) && !(libc.Int32FromUint8(*(*uint8)(__ccgo_ptr(r + 24 + uintptr(*(*MD_CHAR)(__ccgo_ptr(data + uintptr(off+uint32(2))))))))&libc.Int32FromInt32(NEED_HTML_ESC_FLAG) != 0) && !(libc.Int32FromUint8(*(*uint8)(__ccgo_ptr(r + 24 + uintptr(*(*MD_CHAR)(__ccgo_ptr(data + uintptr(off+uint32(3))))))))&libc.Int32FromInt32(NEED_HTML_ESC_FLAG) != 0) {
			off = off + uint32(4)
		}
		for off < size && !(libc.Int32FromUint8(*(*uint8)(__ccgo_ptr(r + 24 + uintptr(*(*MD_CHAR)(__ccgo_ptr(data + uintptr(off)))))))&libc.Int32FromInt32(NEED_HTML_ESC_FLAG) != 0) {
			off = off + 1
		}
		if off > beg {
			render_verbatim(tls, r, data+uintptr(beg), off-beg)
		}
		if off < size {
			switch libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr(data + uintptr(off)))) {
			case int32('&'):
				render_verbatim(tls, r, __ccgo_ts+837, uint32(libc.Xstrlen(tls, __ccgo_ts+837)))
			case int32('<'):
				render_verbatim(tls, r, __ccgo_ts+843, uint32(libc.Xstrlen(tls, __ccgo_ts+843)))
			case int32('>'):
				render_verbatim(tls, r, __ccgo_ts+848, uint32(libc.Xstrlen(tls, __ccgo_ts+848)))
			case int32('"'):
				render_verbatim(tls, r, __ccgo_ts+853, uint32(libc.Xstrlen(tls, __ccgo_ts+853)))
				break
			}
			off = off + 1
		} else {
			break
		}
		beg = off
	}
}

func render_url_escaped(tls *libc.TLS, r uintptr, data uintptr, size MD_SIZE) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var beg, off MD_OFFSET
	var _ /* hex at bp+0 */ [3]uint8
	_, _ = beg, off
	beg = uint32(0)
	off = uint32(0)
	/* Some characters need to be escaped in URL attributes. */
	for int32(1) != 0 {
		for off < size && !(libc.Int32FromUint8(*(*uint8)(__ccgo_ptr(r + 24 + uintptr(*(*MD_CHAR)(__ccgo_ptr(data + uintptr(off)))))))&libc.Int32FromInt32(NEED_URL_ESC_FLAG) != 0) {
			off = off + 1
		}
		if off > beg {
			render_verbatim(tls, r, data+uintptr(beg), off-beg)
		}
		if off < size {
			switch libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr(data + uintptr(off)))) {
			case int32('&'):
				render_verbatim(tls, r, __ccgo_ts+837, uint32(libc.Xstrlen(tls, __ccgo_ts+837)))
			default:
				(*(*[3]uint8)(__ccgo_ptr(bp)))[0] = uint8('%')
				(*(*[3]uint8)(__ccgo_ptr(bp)))[int32(1)] = hex_chars[uint32(*(*MD_CHAR)(__ccgo_ptr(data + uintptr(off))))>>int32(4)&uint32(0xf)]
				(*(*[3]uint8)(__ccgo_ptr(bp)))[int32(2)] = hex_chars[uint32(*(*MD_CHAR)(__ccgo_ptr(data + uintptr(off))))>>0&uint32(0xf)]
				render_verbatim(tls, r, bp, uint32(3))
				break
			}
			off = off + 1
		} else {
			break
		}
		beg = off
	}
}

var hex_chars = [17]MD_CHAR{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F'}

func hex_val(tls *libc.TLS, ch uint8) (r uint32) {
	if int32('0') <= libc.Int32FromUint8(ch) && libc.Int32FromUint8(ch) <= int32('9') {
		return libc.Uint32FromInt32(libc.Int32FromUint8(ch) - int32('0'))
	}
	if int32('A') <= libc.Int32FromUint8(ch) && libc.Int32FromUint8(ch) <= int32('Z') {
		return libc.Uint32FromInt32(libc.Int32FromUint8(ch) - int32('A') + int32(10))
	} else {
		return libc.Uint32FromInt32(libc.Int32FromUint8(ch) - int32('a') + int32(10))
	}
	return r
}

func render_utf8_codepoint(tls *libc.TLS, r uintptr, codepoint uint32, __ccgo_fp_fn_append uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var n size_t
	var _ /* utf8 at bp+0 */ [4]uint8
	_ = n
	if codepoint <= uint32(0x7f) {
		n = uint64(1)
		(*(*[4]uint8)(__ccgo_ptr(bp)))[0] = uint8(codepoint)
	} else {
		if codepoint <= uint32(0x7ff) {
			n = uint64(2)
			(*(*[4]uint8)(__ccgo_ptr(bp)))[0] = uint8(uint32(0xc0) | codepoint>>int32(6)&uint32(0x1f))
			(*(*[4]uint8)(__ccgo_ptr(bp)))[int32(1)] = uint8(uint32(0x80) + codepoint>>0&uint32(0x3f))
		} else {
			if codepoint <= uint32(0xffff) {
				n = uint64(3)
				(*(*[4]uint8)(__ccgo_ptr(bp)))[0] = uint8(uint32(0xe0) | codepoint>>int32(12)&uint32(0xf))
				(*(*[4]uint8)(__ccgo_ptr(bp)))[int32(1)] = uint8(uint32(0x80) + codepoint>>int32(6)&uint32(0x3f))
				(*(*[4]uint8)(__ccgo_ptr(bp)))[int32(2)] = uint8(uint32(0x80) + codepoint>>0&uint32(0x3f))
			} else {
				n = uint64(4)
				(*(*[4]uint8)(__ccgo_ptr(bp)))[0] = uint8(uint32(0xf0) | codepoint>>int32(18)&uint32(0x7))
				(*(*[4]uint8)(__ccgo_ptr(bp)))[int32(1)] = uint8(uint32(0x80) + codepoint>>int32(12)&uint32(0x3f))
				(*(*[4]uint8)(__ccgo_ptr(bp)))[int32(2)] = uint8(uint32(0x80) + codepoint>>int32(6)&uint32(0x3f))
				(*(*[4]uint8)(__ccgo_ptr(bp)))[int32(3)] = uint8(uint32(0x80) + codepoint>>0&uint32(0x3f))
			}
		}
	}
	if uint32(0) < codepoint && codepoint <= uint32(0x10ffff) {
		(*(*func(*libc.TLS, uintptr, uintptr, MD_SIZE))(unsafe.Pointer(&struct{ uintptr }{__ccgo_fp_fn_append})))(tls, r, bp, uint32(n))
	} else {
		(*(*func(*libc.TLS, uintptr, uintptr, MD_SIZE))(unsafe.Pointer(&struct{ uintptr }{__ccgo_fp_fn_append})))(tls, r, uintptr(unsafe.Pointer(&utf8_replacement_char)), uint32(3))
	}
}

var utf8_replacement_char = [3]MD_CHAR{
	0: libc.Uint8FromInt32(0xef),
	1: libc.Uint8FromInt32(0xbf),
	2: libc.Uint8FromInt32(0xbd),
}

// C documentation
//
//	/* Translate entity to its UTF-8 equivalent, or output the verbatim one
//	 * if such entity is unknown (or if the translation is disabled). */
func render_entity(tls *libc.TLS, r uintptr, text uintptr, size MD_SIZE, __ccgo_fp_fn_append uintptr) {
	var codepoint uint32
	var ent uintptr
	var i, i1 MD_SIZE
	_, _, _, _ = codepoint, ent, i, i1
	if (*MD_HTML)(unsafe.Pointer(r)).Fflags&uint32(MD_HTML_FLAG_VERBATIM_ENTITIES) != 0 {
		render_verbatim(tls, r, text, size)
		return
	}
	/* We assume UTF-8 output is what is desired. */
	if size > uint32(3) && libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr(text + 1))) == int32('#') {
		codepoint = uint32(0)
		if libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr(text + 2))) == int32('x') || libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr(text + 2))) == int32('X') {
			i = uint32(3)
			for {
				if !(i < size-uint32(1)) {
					break
				}
				codepoint = uint32(16)*codepoint + hex_val(tls, *(*MD_CHAR)(__ccgo_ptr(text + uintptr(i))))
				goto _1
			_1:
				;
				i = i + 1
			}
		} else {
			i1 = uint32(2)
			for {
				if !(i1 < size-uint32(1)) {
					break
				}
				codepoint = uint32(10)*codepoint + libc.Uint32FromInt32(libc.Int32FromUint8(*(*MD_CHAR)(__ccgo_ptr(text + uintptr(i1))))-libc.Int32FromUint8('0'))
				goto _2
			_2:
				;
				i1 = i1 + 1
			}
		}
		render_utf8_codepoint(tls, r, codepoint, __ccgo_fp_fn_append)
		return
	} else {
		ent = entity_lookup(tls, text, uint64(size))
		if ent != libc.UintptrFromInt32(0) {
			render_utf8_codepoint(tls, r, *(*uint32)(__ccgo_ptr(ent + 8)), __ccgo_fp_fn_append)
			if *(*uint32)(__ccgo_ptr(ent + 8 + 1*4)) != 0 {
				render_utf8_codepoint(tls, r, *(*uint32)(__ccgo_ptr(ent + 8 + 1*4)), __ccgo_fp_fn_append)
			}
			return
		}
	}
	(*(*func(*libc.TLS, uintptr, uintptr, MD_SIZE))(unsafe.Pointer(&struct{ uintptr }{__ccgo_fp_fn_append})))(tls, r, text, size)
}

func render_attribute(tls *libc.TLS, r uintptr, attr uintptr, __ccgo_fp_fn_append uintptr) {
	var i int32
	var off MD_OFFSET
	var size MD_SIZE
	var text uintptr
	var type1 MD_TEXTTYPE1
	_, _, _, _, _ = i, off, size, text, type1
	i = 0
	for {
		if !(*(*MD_OFFSET)(__ccgo_ptr((*MD_ATTRIBUTE)(unsafe.Pointer(attr)).Fsubstr_offsets + uintptr(i)*4)) < (*MD_ATTRIBUTE)(unsafe.Pointer(attr)).Fsize) {
			break
		}
		type1 = *(*MD_TEXTTYPE1)(__ccgo_ptr((*MD_ATTRIBUTE)(unsafe.Pointer(attr)).Fsubstr_types + uintptr(i)*4))
		off = *(*MD_OFFSET)(__ccgo_ptr((*MD_ATTRIBUTE)(unsafe.Pointer(attr)).Fsubstr_offsets + uintptr(i)*4))
		size = *(*MD_OFFSET)(__ccgo_ptr((*MD_ATTRIBUTE)(unsafe.Pointer(attr)).Fsubstr_offsets + uintptr(i+int32(1))*4)) - off
		text = (*MD_ATTRIBUTE)(unsafe.Pointer(attr)).Ftext + uintptr(off)
		switch type1 {
		case int32(MD_TEXT_NULLCHAR):
			render_utf8_codepoint(tls, r, uint32(0x0000), __ccgo_fp(render_verbatim))
		case int32(MD_TEXT_ENTITY):
			render_entity(tls, r, text, size, __ccgo_fp_fn_append)
		default:
			(*(*func(*libc.TLS, uintptr, uintptr, MD_SIZE))(unsafe.Pointer(&struct{ uintptr }{__ccgo_fp_fn_append})))(tls, r, text, size)
			break
		}
		goto _1
	_1:
		;
		i = i + 1
	}
}

func render_open_ol_block(tls *libc.TLS, r uintptr, det uintptr) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	var _ /* buf at bp+0 */ [64]uint8
	if (*MD_BLOCK_OL_DETAIL)(unsafe.Pointer(det)).Fstart == uint32(1) {
		render_verbatim(tls, r, __ccgo_ts+860, uint32(libc.Xstrlen(tls, __ccgo_ts+860)))
		return
	}
	libc.X__builtin_snprintf(tls, bp, uint64(64), __ccgo_ts+866, libc.VaList(bp+72, (*MD_BLOCK_OL_DETAIL)(unsafe.Pointer(det)).Fstart))
	render_verbatim(tls, r, bp, uint32(libc.Xstrlen(tls, bp)))
}

func render_open_li_block(tls *libc.TLS, r uintptr, det uintptr) {
	if (*MD_BLOCK_LI_DETAIL)(unsafe.Pointer(det)).Fis_task != 0 {
		render_verbatim(tls, r, __ccgo_ts+883, uint32(libc.Xstrlen(tls, __ccgo_ts+883)))
		if libc.Int32FromUint8((*MD_BLOCK_LI_DETAIL)(unsafe.Pointer(det)).Ftask_mark) == int32('x') || libc.Int32FromUint8((*MD_BLOCK_LI_DETAIL)(unsafe.Pointer(det)).Ftask_mark) == int32('X') {
			render_verbatim(tls, r, __ccgo_ts+974, uint32(libc.Xstrlen(tls, __ccgo_ts+974)))
		}
		render_verbatim(tls, r, __ccgo_ts+37, uint32(libc.Xstrlen(tls, __ccgo_ts+37)))
	} else {
		render_verbatim(tls, r, __ccgo_ts+983, uint32(libc.Xstrlen(tls, __ccgo_ts+983)))
	}
}

func render_open_code_block(tls *libc.TLS, r uintptr, det uintptr) {
	render_verbatim(tls, r, __ccgo_ts+988, uint32(libc.Xstrlen(tls, __ccgo_ts+988)))
	/* If known, output the HTML 5 attribute class="language-LANGNAME". */
	if (*MD_BLOCK_CODE_DETAIL)(unsafe.Pointer(det)).Flang.Ftext != libc.UintptrFromInt32(0) {
		render_verbatim(tls, r, __ccgo_ts+999, uint32(libc.Xstrlen(tls, __ccgo_ts+999)))
		render_attribute(tls, r, det+32, __ccgo_fp(render_html_escaped))
		render_verbatim(tls, r, __ccgo_ts+1017, uint32(libc.Xstrlen(tls, __ccgo_ts+1017)))
	}
	render_verbatim(tls, r, __ccgo_ts+37, uint32(libc.Xstrlen(tls, __ccgo_ts+37)))
}

func render_open_td_block(tls *libc.TLS, r uintptr, cell_type uintptr, det uintptr) {
	render_verbatim(tls, r, __ccgo_ts+1019, uint32(libc.Xstrlen(tls, __ccgo_ts+1019)))
	render_verbatim(tls, r, cell_type, uint32(libc.Xstrlen(tls, cell_type)))
	switch (*MD_BLOCK_TD_DETAIL)(unsafe.Pointer(det)).Falign {
	case int32(MD_ALIGN_LEFT):
		render_verbatim(tls, r, __ccgo_ts+1021, uint32(libc.Xstrlen(tls, __ccgo_ts+1021)))
	case int32(MD_ALIGN_CENTER):
		render_verbatim(tls, r, __ccgo_ts+1036, uint32(libc.Xstrlen(tls, __ccgo_ts+1036)))
	case int32(MD_ALIGN_RIGHT):
		render_verbatim(tls, r, __ccgo_ts+1053, uint32(libc.Xstrlen(tls, __ccgo_ts+1053)))
	default:
		render_verbatim(tls, r, __ccgo_ts+37, uint32(libc.Xstrlen(tls, __ccgo_ts+37)))
		break
	}
}

func render_open_a_span(tls *libc.TLS, r uintptr, det uintptr) {
	render_verbatim(tls, r, __ccgo_ts+1069, uint32(libc.Xstrlen(tls, __ccgo_ts+1069)))
	render_attribute(tls, r, det, __ccgo_fp(render_url_escaped))
	if (*MD_SPAN_A_DETAIL)(unsafe.Pointer(det)).Ftitle.Ftext != libc.UintptrFromInt32(0) {
		render_verbatim(tls, r, __ccgo_ts+1079, uint32(libc.Xstrlen(tls, __ccgo_ts+1079)))
		render_attribute(tls, r, det+32, __ccgo_fp(render_html_escaped))
	}
	render_verbatim(tls, r, __ccgo_ts+1089, uint32(libc.Xstrlen(tls, __ccgo_ts+1089)))
}

func render_open_img_span(tls *libc.TLS, r uintptr, det uintptr) {
	render_verbatim(tls, r, __ccgo_ts+1092, uint32(libc.Xstrlen(tls, __ccgo_ts+1092)))
	render_attribute(tls, r, det, __ccgo_fp(render_url_escaped))
	render_verbatim(tls, r, __ccgo_ts+1103, uint32(libc.Xstrlen(tls, __ccgo_ts+1103)))
}

func render_close_img_span(tls *libc.TLS, r uintptr, det uintptr) {
	var v1, v2 uintptr
	_, _ = v1, v2
	if (*MD_SPAN_IMG_DETAIL)(unsafe.Pointer(det)).Ftitle.Ftext != libc.UintptrFromInt32(0) {
		render_verbatim(tls, r, __ccgo_ts+1079, uint32(libc.Xstrlen(tls, __ccgo_ts+1079)))
		render_attribute(tls, r, det+32, __ccgo_fp(render_html_escaped))
	}
	if (*MD_HTML)(unsafe.Pointer(r)).Fflags&uint32(MD_HTML_FLAG_XHTML) != 0 {
		v1 = __ccgo_ts + 1111
	} else {
		v1 = __ccgo_ts + 1089
	}
	if (*MD_HTML)(unsafe.Pointer(r)).Fflags&uint32(MD_HTML_FLAG_XHTML) != 0 {
		v2 = __ccgo_ts + 1111
	} else {
		v2 = __ccgo_ts + 1089
	}
	render_verbatim(tls, r, v1, uint32(libc.Xstrlen(tls, v2)))
}

func render_open_wikilink_span(tls *libc.TLS, r uintptr, det uintptr) {
	render_verbatim(tls, r, __ccgo_ts+1116, uint32(libc.Xstrlen(tls, __ccgo_ts+1116)))
	render_attribute(tls, r, det, __ccgo_fp(render_html_escaped))
	render_verbatim(tls, r, __ccgo_ts+1089, uint32(libc.Xstrlen(tls, __ccgo_ts+1089)))
}

/**************************************
 ***  HTML renderer implementation  ***
 **************************************/

func enter_block_callback(tls *libc.TLS, type1 MD_BLOCKTYPE1, detail uintptr, userdata uintptr) (r1 int32) {
	var r, v1, v2 uintptr
	_, _, _ = r, v1, v2
	r = userdata
	switch type1 {
	case int32(MD_BLOCK_DOC): /* noop */
	case int32(MD_BLOCK_QUOTE):
		render_verbatim(tls, r, __ccgo_ts+1172, uint32(libc.Xstrlen(tls, __ccgo_ts+1172)))
	case int32(MD_BLOCK_UL):
		render_verbatim(tls, r, __ccgo_ts+1186, uint32(libc.Xstrlen(tls, __ccgo_ts+1186)))
	case int32(MD_BLOCK_OL):
		render_open_ol_block(tls, r, detail)
	case int32(MD_BLOCK_LI):
		render_open_li_block(tls, r, detail)
	case int32(MD_BLOCK_HR):
		if (*MD_HTML)(unsafe.Pointer(r)).Fflags&uint32(MD_HTML_FLAG_XHTML) != 0 {
			v1 = __ccgo_ts + 1192
		} else {
			v1 = __ccgo_ts + 1200
		}
		if (*MD_HTML)(unsafe.Pointer(r)).Fflags&uint32(MD_HTML_FLAG_XHTML) != 0 {
			v2 = __ccgo_ts + 1192
		} else {
			v2 = __ccgo_ts + 1200
		}
		render_verbatim(tls, r, v1, uint32(libc.Xstrlen(tls, v2)))
	case int32(MD_BLOCK_H):
		render_verbatim(tls, r, head[(*MD_BLOCK_H_DETAIL)(unsafe.Pointer(detail)).Flevel-uint32(1)], uint32(libc.Xstrlen(tls, head[(*MD_BLOCK_H_DETAIL)(unsafe.Pointer(detail)).Flevel-uint32(1)])))
	case int32(MD_BLOCK_CODE):
		render_open_code_block(tls, r, detail)
	case int32(MD_BLOCK_HTML): /* noop */
	case int32(MD_BLOCK_P):
		render_verbatim(tls, r, __ccgo_ts+1206, uint32(libc.Xstrlen(tls, __ccgo_ts+1206)))
	case int32(MD_BLOCK_TABLE):
		render_verbatim(tls, r, __ccgo_ts+1210, uint32(libc.Xstrlen(tls, __ccgo_ts+1210)))
	case int32(MD_BLOCK_THEAD):
		render_verbatim(tls, r, __ccgo_ts+1219, uint32(libc.Xstrlen(tls, __ccgo_ts+1219)))
	case int32(MD_BLOCK_TBODY):
		render_verbatim(tls, r, __ccgo_ts+1228, uint32(libc.Xstrlen(tls, __ccgo_ts+1228)))
	case int32(MD_BLOCK_TR):
		render_verbatim(tls, r, __ccgo_ts+1237, uint32(libc.Xstrlen(tls, __ccgo_ts+1237)))
	case int32(MD_BLOCK_TH):
		render_open_td_block(tls, r, __ccgo_ts+755, detail)
	case int32(MD_BLOCK_TD):
		render_open_td_block(tls, r, __ccgo_ts+746, detail)
		break
	}
	return 0
}

var head = [6]uintptr{
	0: __ccgo_ts + 1142,
	1: __ccgo_ts + 1147,
	2: __ccgo_ts + 1152,
	3: __ccgo_ts + 1157,
	4: __ccgo_ts + 1162,
	5: __ccgo_ts + 1167,
}

func leave_block_callback(tls *libc.TLS, type1 MD_BLOCKTYPE1, detail uintptr, userdata uintptr) (r1 int32) {
	var r uintptr
	_ = r
	r = userdata
	switch type1 {
	case int32(MD_BLOCK_DOC): /*noop*/
	case int32(MD_BLOCK_QUOTE):
		render_verbatim(tls, r, __ccgo_ts+1285, uint32(libc.Xstrlen(tls, __ccgo_ts+1285)))
	case int32(MD_BLOCK_UL):
		render_verbatim(tls, r, __ccgo_ts+1300, uint32(libc.Xstrlen(tls, __ccgo_ts+1300)))
	case int32(MD_BLOCK_OL):
		render_verbatim(tls, r, __ccgo_ts+1307, uint32(libc.Xstrlen(tls, __ccgo_ts+1307)))
	case int32(MD_BLOCK_LI):
		render_verbatim(tls, r, __ccgo_ts+1314, uint32(libc.Xstrlen(tls, __ccgo_ts+1314)))
	case int32(MD_BLOCK_HR): /*noop*/
	case int32(MD_BLOCK_H):
		render_verbatim(tls, r, head1[(*MD_BLOCK_H_DETAIL)(unsafe.Pointer(detail)).Flevel-uint32(1)], uint32(libc.Xstrlen(tls, head1[(*MD_BLOCK_H_DETAIL)(unsafe.Pointer(detail)).Flevel-uint32(1)])))
	case int32(MD_BLOCK_CODE):
		render_verbatim(tls, r, __ccgo_ts+1321, uint32(libc.Xstrlen(tls, __ccgo_ts+1321)))
	case int32(MD_BLOCK_HTML): /* noop */
	case int32(MD_BLOCK_P):
		render_verbatim(tls, r, __ccgo_ts+1336, uint32(libc.Xstrlen(tls, __ccgo_ts+1336)))
	case int32(MD_BLOCK_TABLE):
		render_verbatim(tls, r, __ccgo_ts+1342, uint32(libc.Xstrlen(tls, __ccgo_ts+1342)))
	case int32(MD_BLOCK_THEAD):
		render_verbatim(tls, r, __ccgo_ts+1352, uint32(libc.Xstrlen(tls, __ccgo_ts+1352)))
	case int32(MD_BLOCK_TBODY):
		render_verbatim(tls, r, __ccgo_ts+1362, uint32(libc.Xstrlen(tls, __ccgo_ts+1362)))
	case int32(MD_BLOCK_TR):
		render_verbatim(tls, r, __ccgo_ts+1372, uint32(libc.Xstrlen(tls, __ccgo_ts+1372)))
	case int32(MD_BLOCK_TH):
		render_verbatim(tls, r, __ccgo_ts+1379, uint32(libc.Xstrlen(tls, __ccgo_ts+1379)))
	case int32(MD_BLOCK_TD):
		render_verbatim(tls, r, __ccgo_ts+1386, uint32(libc.Xstrlen(tls, __ccgo_ts+1386)))
		break
	}
	return 0
}

var head1 = [6]uintptr{
	0: __ccgo_ts + 1243,
	1: __ccgo_ts + 1250,
	2: __ccgo_ts + 1257,
	3: __ccgo_ts + 1264,
	4: __ccgo_ts + 1271,
	5: __ccgo_ts + 1278,
}

func enter_span_callback(tls *libc.TLS, type1 MD_SPANTYPE1, detail uintptr, userdata uintptr) (r1 int32) {
	var inside_img int32
	var r uintptr
	_, _ = inside_img, r
	r = userdata
	inside_img = libc.BoolInt32((*MD_HTML)(unsafe.Pointer(r)).Fimage_nesting_level > 0)
	/* We are inside a Markdown image label. Markdown allows to use any emphasis
	 * and other rich contents in that context similarly as in any link label.
	 *
	 * However, unlike in the case of links (where that contents becomescontents
	 * of the <a>...</a> tag), in the case of images the contents is supposed to
	 * fall into the attribute alt: <img alt="...">.
	 *
	 * In that context we naturally cannot output nested HTML tags. So lets
	 * suppress them and only output the plain text (i.e. what falls into text()
	 * callback).
	 *
	 * CommonMark specification declares this a recommended practice for HTML
	 * output.
	 */
	if type1 == int32(MD_SPAN_IMG) {
		(*MD_HTML)(unsafe.Pointer(r)).Fimage_nesting_level = (*MD_HTML)(unsafe.Pointer(r)).Fimage_nesting_level + 1
	}
	if inside_img != 0 {
		return 0
	}
	switch type1 {
	case int32(MD_SPAN_EM):
		render_verbatim(tls, r, __ccgo_ts+1393, uint32(libc.Xstrlen(tls, __ccgo_ts+1393)))
	case int32(MD_SPAN_STRONG):
		render_verbatim(tls, r, __ccgo_ts+1398, uint32(libc.Xstrlen(tls, __ccgo_ts+1398)))
	case int32(MD_SPAN_U):
		render_verbatim(tls, r, __ccgo_ts+1407, uint32(libc.Xstrlen(tls, __ccgo_ts+1407)))
	case int32(MD_SPAN_A):
		render_open_a_span(tls, r, detail)
	case int32(MD_SPAN_IMG):
		render_open_img_span(tls, r, detail)
	case int32(MD_SPAN_CODE):
		render_verbatim(tls, r, __ccgo_ts+1411, uint32(libc.Xstrlen(tls, __ccgo_ts+1411)))
	case int32(MD_SPAN_DEL):
		render_verbatim(tls, r, __ccgo_ts+1418, uint32(libc.Xstrlen(tls, __ccgo_ts+1418)))
	case int32(MD_SPAN_LATEXMATH):
		render_verbatim(tls, r, __ccgo_ts+1424, uint32(libc.Xstrlen(tls, __ccgo_ts+1424)))
	case int32(MD_SPAN_LATEXMATH_DISPLAY):
		render_verbatim(tls, r, __ccgo_ts+1437, uint32(libc.Xstrlen(tls, __ccgo_ts+1437)))
	case int32(MD_SPAN_WIKILINK1):
		render_open_wikilink_span(tls, r, detail)
		break
	}
	return 0
}

func leave_span_callback(tls *libc.TLS, type1 MD_SPANTYPE1, detail uintptr, userdata uintptr) (r1 int32) {
	var r uintptr
	_ = r
	r = userdata
	if type1 == int32(MD_SPAN_IMG) {
		(*MD_HTML)(unsafe.Pointer(r)).Fimage_nesting_level = (*MD_HTML)(unsafe.Pointer(r)).Fimage_nesting_level - 1
	}
	if (*MD_HTML)(unsafe.Pointer(r)).Fimage_nesting_level > 0 {
		return 0
	}
	switch type1 {
	case int32(MD_SPAN_EM):
		render_verbatim(tls, r, __ccgo_ts+1465, uint32(libc.Xstrlen(tls, __ccgo_ts+1465)))
	case int32(MD_SPAN_STRONG):
		render_verbatim(tls, r, __ccgo_ts+1471, uint32(libc.Xstrlen(tls, __ccgo_ts+1471)))
	case int32(MD_SPAN_U):
		render_verbatim(tls, r, __ccgo_ts+1481, uint32(libc.Xstrlen(tls, __ccgo_ts+1481)))
	case int32(MD_SPAN_A):
		render_verbatim(tls, r, __ccgo_ts+1486, uint32(libc.Xstrlen(tls, __ccgo_ts+1486)))
	case int32(MD_SPAN_IMG):
		render_close_img_span(tls, r, detail)
	case int32(MD_SPAN_CODE):
		render_verbatim(tls, r, __ccgo_ts+1491, uint32(libc.Xstrlen(tls, __ccgo_ts+1491)))
	case int32(MD_SPAN_DEL):
		render_verbatim(tls, r, __ccgo_ts+1499, uint32(libc.Xstrlen(tls, __ccgo_ts+1499)))
	case int32(MD_SPAN_LATEXMATH): /*fall through*/
		fallthrough
	case int32(MD_SPAN_LATEXMATH_DISPLAY):
		render_verbatim(tls, r, __ccgo_ts+1506, uint32(libc.Xstrlen(tls, __ccgo_ts+1506)))
	case int32(MD_SPAN_WIKILINK1):
		render_verbatim(tls, r, __ccgo_ts+1520, uint32(libc.Xstrlen(tls, __ccgo_ts+1520)))
		break
	}
	return 0
}

func text_callback(tls *libc.TLS, type1 MD_TEXTTYPE1, text uintptr, size MD_SIZE, userdata uintptr) (r1 int32) {
	var r, v1, v2, v3, v4 uintptr
	_, _, _, _, _ = r, v1, v2, v3, v4
	r = userdata
	switch type1 {
	case int32(MD_TEXT_NULLCHAR):
		render_utf8_codepoint(tls, r, uint32(0x0000), __ccgo_fp(render_verbatim))
	case int32(MD_TEXT_BR):
		if (*MD_HTML)(unsafe.Pointer(r)).Fimage_nesting_level == 0 {
			if (*MD_HTML)(unsafe.Pointer(r)).Fflags&uint32(MD_HTML_FLAG_XHTML) != 0 {
				v2 = __ccgo_ts + 1534
			} else {
				v2 = __ccgo_ts + 1542
			}
			v1 = v2
		} else {
			v1 = __ccgo_ts + 268
		}
		if (*MD_HTML)(unsafe.Pointer(r)).Fimage_nesting_level == 0 {
			if (*MD_HTML)(unsafe.Pointer(r)).Fflags&uint32(MD_HTML_FLAG_XHTML) != 0 {
				v4 = __ccgo_ts + 1534
			} else {
				v4 = __ccgo_ts + 1542
			}
			v3 = v4
		} else {
			v3 = __ccgo_ts + 268
		}
		render_verbatim(tls, r, v1, uint32(libc.Xstrlen(tls, v3)))
	case int32(MD_TEXT_SOFTBR):
		if (*MD_HTML)(unsafe.Pointer(r)).Fimage_nesting_level == 0 {
			v1 = __ccgo_ts + 286
		} else {
			v1 = __ccgo_ts + 268
		}
		if (*MD_HTML)(unsafe.Pointer(r)).Fimage_nesting_level == 0 {
			v2 = __ccgo_ts + 286
		} else {
			v2 = __ccgo_ts + 268
		}
		render_verbatim(tls, r, v1, uint32(libc.Xstrlen(tls, v2)))
	case int32(MD_TEXT_HTML):
		render_verbatim(tls, r, text, size)
	case int32(MD_TEXT_ENTITY):
		render_entity(tls, r, text, size, __ccgo_fp(render_html_escaped))
	default:
		render_html_escaped(tls, r, text, size)
		break
	}
	return 0
}

func debug_log_callback(tls *libc.TLS, msg uintptr, userdata uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var r uintptr
	_ = r
	r = userdata
	if (*MD_HTML)(unsafe.Pointer(r)).Fflags&uint32(MD_HTML_FLAG_DEBUG) != 0 {
		libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts+1548, libc.VaList(bp+8, msg))
	}
}

type __ccgo_fp__Xmd_html_2 = func(*libc.TLS, uintptr, uint32, uintptr)

func md_html(tls *libc.TLS, input uintptr, input_size MD_SIZE, __ccgo_fp_process_output uintptr, userdata uintptr, parser_flags uint32, renderer_flags uint32) (r int32) {
	bp := tls.Alloc(352)
	defer tls.Free(352)
	var ch uint8
	var i int32
	var v2 uintptr
	var _ /* parser at bp+280 */ MD_PARSER
	var _ /* render at bp+0 */ MD_HTML
	_, _, _ = ch, i, v2
	*(*MD_HTML)(__ccgo_ptr(bp)) = MD_HTML{
		Fprocess_output: __ccgo_fp_process_output,
		Fuserdata:       userdata,
		Fflags:          renderer_flags,
	}
	*(*MD_PARSER)(__ccgo_ptr(bp + 280)) = MD_PARSER{
		Fflags:       parser_flags,
		Fenter_block: __ccgo_fp(enter_block_callback),
		Fleave_block: __ccgo_fp(leave_block_callback),
		Fenter_span:  __ccgo_fp(enter_span_callback),
		Fleave_span:  __ccgo_fp(leave_span_callback),
		Ftext:        __ccgo_fp(text_callback),
		Fdebug_log:   __ccgo_fp(debug_log_callback),
	}
	/* Build map of characters which need escaping. */
	i = 0
	for {
		if !(i < int32(256)) {
			break
		}
		ch = libc.Uint8FromInt32(i)
		if libc.Xstrchr(tls, __ccgo_ts+1558, libc.Int32FromUint8(ch)) != libc.UintptrFromInt32(0) {
			v2 = bp + 24 + uintptr(i)
			*(*uint8)(unsafe.Pointer(v2)) = uint8(int32(*(*uint8)(unsafe.Pointer(v2))) | libc.Int32FromInt32(NEED_HTML_ESC_FLAG))
		}
		if !(int32('a') <= libc.Int32FromUint8(ch) && libc.Int32FromUint8(ch) <= int32('z') || int32('A') <= libc.Int32FromUint8(ch) && libc.Int32FromUint8(ch) <= int32('Z') || int32('0') <= libc.Int32FromUint8(ch) && libc.Int32FromUint8(ch) <= int32('9')) && libc.Xstrchr(tls, __ccgo_ts+1563, libc.Int32FromUint8(ch)) == libc.UintptrFromInt32(0) {
			v2 = bp + 24 + uintptr(i)
			*(*uint8)(unsafe.Pointer(v2)) = uint8(int32(*(*uint8)(unsafe.Pointer(v2))) | libc.Int32FromInt32(NEED_URL_ESC_FLAG))
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	/* Consider skipping UTF-8 byte order mark (BOM). */
	if renderer_flags&uint32(MD_HTML_FLAG_SKIP_UTF8_BOM) != 0 && libc.Bool(uint64(1) == uint64(1)) {
		if uint64(input_size) >= uint64(3) && libc.Xmemcmp(tls, input, uintptr(unsafe.Pointer(&bom)), uint64(3)) == 0 {
			input = input + uintptr(3)
			input_size = uint32(uint64(input_size) - libc.Uint64FromInt64(3))
		}
	}
	return md_parse(tls, input, input_size, bp+280, bp)
}

var bom = [3]MD_CHAR{
	0: libc.Uint8FromInt32(0xef),
	1: libc.Uint8FromInt32(0xbb),
	2: libc.Uint8FromInt32(0xbf),
}

// C documentation
//
//	/* Generated by scripts/build_enity_map.py. */
var ENTITY_MAP = [2125]ENTITY{
	0: {
		Fname: __ccgo_ts + 1585,
		Fcodepoints: [2]uint32{
			0: uint32(198),
		},
	},
	1: {
		Fname: __ccgo_ts + 1593,
		Fcodepoints: [2]uint32{
			0: uint32(38),
		},
	},
	2: {
		Fname: __ccgo_ts + 1599,
		Fcodepoints: [2]uint32{
			0: uint32(193),
		},
	},
	3: {
		Fname: __ccgo_ts + 1608,
		Fcodepoints: [2]uint32{
			0: uint32(258),
		},
	},
	4: {
		Fname: __ccgo_ts + 1617,
		Fcodepoints: [2]uint32{
			0: uint32(194),
		},
	},
	5: {
		Fname: __ccgo_ts + 1625,
		Fcodepoints: [2]uint32{
			0: uint32(1040),
		},
	},
	6: {
		Fname: __ccgo_ts + 1631,
		Fcodepoints: [2]uint32{
			0: uint32(120068),
		},
	},
	7: {
		Fname: __ccgo_ts + 1637,
		Fcodepoints: [2]uint32{
			0: uint32(192),
		},
	},
	8: {
		Fname: __ccgo_ts + 1646,
		Fcodepoints: [2]uint32{
			0: uint32(913),
		},
	},
	9: {
		Fname: __ccgo_ts + 1654,
		Fcodepoints: [2]uint32{
			0: uint32(256),
		},
	},
	10: {
		Fname: __ccgo_ts + 1662,
		Fcodepoints: [2]uint32{
			0: uint32(10835),
		},
	},
	11: {
		Fname: __ccgo_ts + 1668,
		Fcodepoints: [2]uint32{
			0: uint32(260),
		},
	},
	12: {
		Fname: __ccgo_ts + 1676,
		Fcodepoints: [2]uint32{
			0: uint32(120120),
		},
	},
	13: {
		Fname: __ccgo_ts + 1683,
		Fcodepoints: [2]uint32{
			0: uint32(8289),
		},
	},
	14: {
		Fname: __ccgo_ts + 1699,
		Fcodepoints: [2]uint32{
			0: uint32(197),
		},
	},
	15: {
		Fname: __ccgo_ts + 1707,
		Fcodepoints: [2]uint32{
			0: uint32(119964),
		},
	},
	16: {
		Fname: __ccgo_ts + 1714,
		Fcodepoints: [2]uint32{
			0: uint32(8788),
		},
	},
	17: {
		Fname: __ccgo_ts + 1723,
		Fcodepoints: [2]uint32{
			0: uint32(195),
		},
	},
	18: {
		Fname: __ccgo_ts + 1732,
		Fcodepoints: [2]uint32{
			0: uint32(196),
		},
	},
	19: {
		Fname: __ccgo_ts + 1739,
		Fcodepoints: [2]uint32{
			0: uint32(8726),
		},
	},
	20: {
		Fname: __ccgo_ts + 1751,
		Fcodepoints: [2]uint32{
			0: uint32(10983),
		},
	},
	21: {
		Fname: __ccgo_ts + 1758,
		Fcodepoints: [2]uint32{
			0: uint32(8966),
		},
	},
	22: {
		Fname: __ccgo_ts + 1767,
		Fcodepoints: [2]uint32{
			0: uint32(1041),
		},
	},
	23: {
		Fname: __ccgo_ts + 1773,
		Fcodepoints: [2]uint32{
			0: uint32(8757),
		},
	},
	24: {
		Fname: __ccgo_ts + 1783,
		Fcodepoints: [2]uint32{
			0: uint32(8492),
		},
	},
	25: {
		Fname: __ccgo_ts + 1796,
		Fcodepoints: [2]uint32{
			0: uint32(914),
		},
	},
	26: {
		Fname: __ccgo_ts + 1803,
		Fcodepoints: [2]uint32{
			0: uint32(120069),
		},
	},
	27: {
		Fname: __ccgo_ts + 1809,
		Fcodepoints: [2]uint32{
			0: uint32(120121),
		},
	},
	28: {
		Fname: __ccgo_ts + 1816,
		Fcodepoints: [2]uint32{
			0: uint32(728),
		},
	},
	29: {
		Fname: __ccgo_ts + 1824,
		Fcodepoints: [2]uint32{
			0: uint32(8492),
		},
	},
	30: {
		Fname: __ccgo_ts + 1831,
		Fcodepoints: [2]uint32{
			0: uint32(8782),
		},
	},
	31: {
		Fname: __ccgo_ts + 1840,
		Fcodepoints: [2]uint32{
			0: uint32(1063),
		},
	},
	32: {
		Fname: __ccgo_ts + 1847,
		Fcodepoints: [2]uint32{
			0: uint32(169),
		},
	},
	33: {
		Fname: __ccgo_ts + 1854,
		Fcodepoints: [2]uint32{
			0: uint32(262),
		},
	},
	34: {
		Fname: __ccgo_ts + 1863,
		Fcodepoints: [2]uint32{
			0: uint32(8914),
		},
	},
	35: {
		Fname: __ccgo_ts + 1869,
		Fcodepoints: [2]uint32{
			0: uint32(8517),
		},
	},
	36: {
		Fname: __ccgo_ts + 1892,
		Fcodepoints: [2]uint32{
			0: uint32(8493),
		},
	},
	37: {
		Fname: __ccgo_ts + 1902,
		Fcodepoints: [2]uint32{
			0: uint32(268),
		},
	},
	38: {
		Fname: __ccgo_ts + 1911,
		Fcodepoints: [2]uint32{
			0: uint32(199),
		},
	},
	39: {
		Fname: __ccgo_ts + 1920,
		Fcodepoints: [2]uint32{
			0: uint32(264),
		},
	},
	40: {
		Fname: __ccgo_ts + 1928,
		Fcodepoints: [2]uint32{
			0: uint32(8752),
		},
	},
	41: {
		Fname: __ccgo_ts + 1938,
		Fcodepoints: [2]uint32{
			0: uint32(266),
		},
	},
	42: {
		Fname: __ccgo_ts + 1945,
		Fcodepoints: [2]uint32{
			0: uint32(184),
		},
	},
	43: {
		Fname: __ccgo_ts + 1955,
		Fcodepoints: [2]uint32{
			0: uint32(183),
		},
	},
	44: {
		Fname: __ccgo_ts + 1967,
		Fcodepoints: [2]uint32{
			0: uint32(8493),
		},
	},
	45: {
		Fname: __ccgo_ts + 1973,
		Fcodepoints: [2]uint32{
			0: uint32(935),
		},
	},
	46: {
		Fname: __ccgo_ts + 1979,
		Fcodepoints: [2]uint32{
			0: uint32(8857),
		},
	},
	47: {
		Fname: __ccgo_ts + 1991,
		Fcodepoints: [2]uint32{
			0: uint32(8854),
		},
	},
	48: {
		Fname: __ccgo_ts + 2005,
		Fcodepoints: [2]uint32{
			0: uint32(8853),
		},
	},
	49: {
		Fname: __ccgo_ts + 2018,
		Fcodepoints: [2]uint32{
			0: uint32(8855),
		},
	},
	50: {
		Fname: __ccgo_ts + 2032,
		Fcodepoints: [2]uint32{
			0: uint32(8754),
		},
	},
	51: {
		Fname: __ccgo_ts + 2059,
		Fcodepoints: [2]uint32{
			0: uint32(8221),
		},
	},
	52: {
		Fname: __ccgo_ts + 2083,
		Fcodepoints: [2]uint32{
			0: uint32(8217),
		},
	},
	53: {
		Fname: __ccgo_ts + 2101,
		Fcodepoints: [2]uint32{
			0: uint32(8759),
		},
	},
	54: {
		Fname: __ccgo_ts + 2109,
		Fcodepoints: [2]uint32{
			0: uint32(10868),
		},
	},
	55: {
		Fname: __ccgo_ts + 2118,
		Fcodepoints: [2]uint32{
			0: uint32(8801),
		},
	},
	56: {
		Fname: __ccgo_ts + 2130,
		Fcodepoints: [2]uint32{
			0: uint32(8751),
		},
	},
	57: {
		Fname: __ccgo_ts + 2139,
		Fcodepoints: [2]uint32{
			0: uint32(8750),
		},
	},
	58: {
		Fname: __ccgo_ts + 2157,
		Fcodepoints: [2]uint32{
			0: uint32(8450),
		},
	},
	59: {
		Fname: __ccgo_ts + 2164,
		Fcodepoints: [2]uint32{
			0: uint32(8720),
		},
	},
	60: {
		Fname: __ccgo_ts + 2176,
		Fcodepoints: [2]uint32{
			0: uint32(8755),
		},
	},
	61: {
		Fname: __ccgo_ts + 2210,
		Fcodepoints: [2]uint32{
			0: uint32(10799),
		},
	},
	62: {
		Fname: __ccgo_ts + 2218,
		Fcodepoints: [2]uint32{
			0: uint32(119966),
		},
	},
	63: {
		Fname: __ccgo_ts + 2225,
		Fcodepoints: [2]uint32{
			0: uint32(8915),
		},
	},
	64: {
		Fname: __ccgo_ts + 2231,
		Fcodepoints: [2]uint32{
			0: uint32(8781),
		},
	},
	65: {
		Fname: __ccgo_ts + 2240,
		Fcodepoints: [2]uint32{
			0: uint32(8517),
		},
	},
	66: {
		Fname: __ccgo_ts + 2245,
		Fcodepoints: [2]uint32{
			0: uint32(10513),
		},
	},
	67: {
		Fname: __ccgo_ts + 2256,
		Fcodepoints: [2]uint32{
			0: uint32(1026),
		},
	},
	68: {
		Fname: __ccgo_ts + 2263,
		Fcodepoints: [2]uint32{
			0: uint32(1029),
		},
	},
	69: {
		Fname: __ccgo_ts + 2270,
		Fcodepoints: [2]uint32{
			0: uint32(1039),
		},
	},
	70: {
		Fname: __ccgo_ts + 2277,
		Fcodepoints: [2]uint32{
			0: uint32(8225),
		},
	},
	71: {
		Fname: __ccgo_ts + 2286,
		Fcodepoints: [2]uint32{
			0: uint32(8609),
		},
	},
	72: {
		Fname: __ccgo_ts + 2293,
		Fcodepoints: [2]uint32{
			0: uint32(10980),
		},
	},
	73: {
		Fname: __ccgo_ts + 2301,
		Fcodepoints: [2]uint32{
			0: uint32(270),
		},
	},
	74: {
		Fname: __ccgo_ts + 2310,
		Fcodepoints: [2]uint32{
			0: uint32(1044),
		},
	},
	75: {
		Fname: __ccgo_ts + 2316,
		Fcodepoints: [2]uint32{
			0: uint32(8711),
		},
	},
	76: {
		Fname: __ccgo_ts + 2322,
		Fcodepoints: [2]uint32{
			0: uint32(916),
		},
	},
	77: {
		Fname: __ccgo_ts + 2330,
		Fcodepoints: [2]uint32{
			0: uint32(120071),
		},
	},
	78: {
		Fname: __ccgo_ts + 2336,
		Fcodepoints: [2]uint32{
			0: uint32(180),
		},
	},
	79: {
		Fname: __ccgo_ts + 2355,
		Fcodepoints: [2]uint32{
			0: uint32(729),
		},
	},
	80: {
		Fname: __ccgo_ts + 2372,
		Fcodepoints: [2]uint32{
			0: uint32(733),
		},
	},
	81: {
		Fname: __ccgo_ts + 2397,
		Fcodepoints: [2]uint32{
			0: uint32(96),
		},
	},
	82: {
		Fname: __ccgo_ts + 2416,
		Fcodepoints: [2]uint32{
			0: uint32(732),
		},
	},
	83: {
		Fname: __ccgo_ts + 2435,
		Fcodepoints: [2]uint32{
			0: uint32(8900),
		},
	},
	84: {
		Fname: __ccgo_ts + 2445,
		Fcodepoints: [2]uint32{
			0: uint32(8518),
		},
	},
	85: {
		Fname: __ccgo_ts + 2461,
		Fcodepoints: [2]uint32{
			0: uint32(120123),
		},
	},
	86: {
		Fname: __ccgo_ts + 2468,
		Fcodepoints: [2]uint32{
			0: uint32(168),
		},
	},
	87: {
		Fname: __ccgo_ts + 2474,
		Fcodepoints: [2]uint32{
			0: uint32(8412),
		},
	},
	88: {
		Fname: __ccgo_ts + 2483,
		Fcodepoints: [2]uint32{
			0: uint32(8784),
		},
	},
	89: {
		Fname: __ccgo_ts + 2494,
		Fcodepoints: [2]uint32{
			0: uint32(8751),
		},
	},
	90: {
		Fname: __ccgo_ts + 2518,
		Fcodepoints: [2]uint32{
			0: uint32(168),
		},
	},
	91: {
		Fname: __ccgo_ts + 2530,
		Fcodepoints: [2]uint32{
			0: uint32(8659),
		},
	},
	92: {
		Fname: __ccgo_ts + 2548,
		Fcodepoints: [2]uint32{
			0: uint32(8656),
		},
	},
	93: {
		Fname: __ccgo_ts + 2566,
		Fcodepoints: [2]uint32{
			0: uint32(8660),
		},
	},
	94: {
		Fname: __ccgo_ts + 2589,
		Fcodepoints: [2]uint32{
			0: uint32(10980),
		},
	},
	95: {
		Fname: __ccgo_ts + 2605,
		Fcodepoints: [2]uint32{
			0: uint32(10232),
		},
	},
	96: {
		Fname: __ccgo_ts + 2627,
		Fcodepoints: [2]uint32{
			0: uint32(10234),
		},
	},
	97: {
		Fname: __ccgo_ts + 2654,
		Fcodepoints: [2]uint32{
			0: uint32(10233),
		},
	},
	98: {
		Fname: __ccgo_ts + 2677,
		Fcodepoints: [2]uint32{
			0: uint32(8658),
		},
	},
	99: {
		Fname: __ccgo_ts + 2696,
		Fcodepoints: [2]uint32{
			0: uint32(8872),
		},
	},
	100: {
		Fname: __ccgo_ts + 2713,
		Fcodepoints: [2]uint32{
			0: uint32(8657),
		},
	},
	101: {
		Fname: __ccgo_ts + 2729,
		Fcodepoints: [2]uint32{
			0: uint32(8661),
		},
	},
	102: {
		Fname: __ccgo_ts + 2749,
		Fcodepoints: [2]uint32{
			0: uint32(8741),
		},
	},
	103: {
		Fname: __ccgo_ts + 2769,
		Fcodepoints: [2]uint32{
			0: uint32(8595),
		},
	},
	104: {
		Fname: __ccgo_ts + 2781,
		Fcodepoints: [2]uint32{
			0: uint32(10515),
		},
	},
	105: {
		Fname: __ccgo_ts + 2796,
		Fcodepoints: [2]uint32{
			0: uint32(8693),
		},
	},
	106: {
		Fname: __ccgo_ts + 2815,
		Fcodepoints: [2]uint32{
			0: uint32(785),
		},
	},
	107: {
		Fname: __ccgo_ts + 2827,
		Fcodepoints: [2]uint32{
			0: uint32(10576),
		},
	},
	108: {
		Fname: __ccgo_ts + 2849,
		Fcodepoints: [2]uint32{
			0: uint32(10590),
		},
	},
	109: {
		Fname: __ccgo_ts + 2869,
		Fcodepoints: [2]uint32{
			0: uint32(8637),
		},
	},
	110: {
		Fname: __ccgo_ts + 2886,
		Fcodepoints: [2]uint32{
			0: uint32(10582),
		},
	},
	111: {
		Fname: __ccgo_ts + 2906,
		Fcodepoints: [2]uint32{
			0: uint32(10591),
		},
	},
	112: {
		Fname: __ccgo_ts + 2927,
		Fcodepoints: [2]uint32{
			0: uint32(8641),
		},
	},
	113: {
		Fname: __ccgo_ts + 2945,
		Fcodepoints: [2]uint32{
			0: uint32(10583),
		},
	},
	114: {
		Fname: __ccgo_ts + 2966,
		Fcodepoints: [2]uint32{
			0: uint32(8868),
		},
	},
	115: {
		Fname: __ccgo_ts + 2976,
		Fcodepoints: [2]uint32{
			0: uint32(8615),
		},
	},
	116: {
		Fname: __ccgo_ts + 2991,
		Fcodepoints: [2]uint32{
			0: uint32(8659),
		},
	},
	117: {
		Fname: __ccgo_ts + 3003,
		Fcodepoints: [2]uint32{
			0: uint32(119967),
		},
	},
	118: {
		Fname: __ccgo_ts + 3010,
		Fcodepoints: [2]uint32{
			0: uint32(272),
		},
	},
	119: {
		Fname: __ccgo_ts + 3019,
		Fcodepoints: [2]uint32{
			0: uint32(330),
		},
	},
	120: {
		Fname: __ccgo_ts + 3025,
		Fcodepoints: [2]uint32{
			0: uint32(208),
		},
	},
	121: {
		Fname: __ccgo_ts + 3031,
		Fcodepoints: [2]uint32{
			0: uint32(201),
		},
	},
	122: {
		Fname: __ccgo_ts + 3040,
		Fcodepoints: [2]uint32{
			0: uint32(282),
		},
	},
	123: {
		Fname: __ccgo_ts + 3049,
		Fcodepoints: [2]uint32{
			0: uint32(202),
		},
	},
	124: {
		Fname: __ccgo_ts + 3057,
		Fcodepoints: [2]uint32{
			0: uint32(1069),
		},
	},
	125: {
		Fname: __ccgo_ts + 3063,
		Fcodepoints: [2]uint32{
			0: uint32(278),
		},
	},
	126: {
		Fname: __ccgo_ts + 3070,
		Fcodepoints: [2]uint32{
			0: uint32(120072),
		},
	},
	127: {
		Fname: __ccgo_ts + 3076,
		Fcodepoints: [2]uint32{
			0: uint32(200),
		},
	},
	128: {
		Fname: __ccgo_ts + 3085,
		Fcodepoints: [2]uint32{
			0: uint32(8712),
		},
	},
	129: {
		Fname: __ccgo_ts + 3095,
		Fcodepoints: [2]uint32{
			0: uint32(274),
		},
	},
	130: {
		Fname: __ccgo_ts + 3103,
		Fcodepoints: [2]uint32{
			0: uint32(9723),
		},
	},
	131: {
		Fname: __ccgo_ts + 3122,
		Fcodepoints: [2]uint32{
			0: uint32(9643),
		},
	},
	132: {
		Fname: __ccgo_ts + 3145,
		Fcodepoints: [2]uint32{
			0: uint32(280),
		},
	},
	133: {
		Fname: __ccgo_ts + 3153,
		Fcodepoints: [2]uint32{
			0: uint32(120124),
		},
	},
	134: {
		Fname: __ccgo_ts + 3160,
		Fcodepoints: [2]uint32{
			0: uint32(917),
		},
	},
	135: {
		Fname: __ccgo_ts + 3170,
		Fcodepoints: [2]uint32{
			0: uint32(10869),
		},
	},
	136: {
		Fname: __ccgo_ts + 3178,
		Fcodepoints: [2]uint32{
			0: uint32(8770),
		},
	},
	137: {
		Fname: __ccgo_ts + 3191,
		Fcodepoints: [2]uint32{
			0: uint32(8652),
		},
	},
	138: {
		Fname: __ccgo_ts + 3205,
		Fcodepoints: [2]uint32{
			0: uint32(8496),
		},
	},
	139: {
		Fname: __ccgo_ts + 3212,
		Fcodepoints: [2]uint32{
			0: uint32(10867),
		},
	},
	140: {
		Fname: __ccgo_ts + 3219,
		Fcodepoints: [2]uint32{
			0: uint32(919),
		},
	},
	141: {
		Fname: __ccgo_ts + 3225,
		Fcodepoints: [2]uint32{
			0: uint32(203),
		},
	},
	142: {
		Fname: __ccgo_ts + 3232,
		Fcodepoints: [2]uint32{
			0: uint32(8707),
		},
	},
	143: {
		Fname: __ccgo_ts + 3241,
		Fcodepoints: [2]uint32{
			0: uint32(8519),
		},
	},
	144: {
		Fname: __ccgo_ts + 3256,
		Fcodepoints: [2]uint32{
			0: uint32(1060),
		},
	},
	145: {
		Fname: __ccgo_ts + 3262,
		Fcodepoints: [2]uint32{
			0: uint32(120073),
		},
	},
	146: {
		Fname: __ccgo_ts + 3268,
		Fcodepoints: [2]uint32{
			0: uint32(9724),
		},
	},
	147: {
		Fname: __ccgo_ts + 3288,
		Fcodepoints: [2]uint32{
			0: uint32(9642),
		},
	},
	148: {
		Fname: __ccgo_ts + 3312,
		Fcodepoints: [2]uint32{
			0: uint32(120125),
		},
	},
	149: {
		Fname: __ccgo_ts + 3319,
		Fcodepoints: [2]uint32{
			0: uint32(8704),
		},
	},
	150: {
		Fname: __ccgo_ts + 3328,
		Fcodepoints: [2]uint32{
			0: uint32(8497),
		},
	},
	151: {
		Fname: __ccgo_ts + 3341,
		Fcodepoints: [2]uint32{
			0: uint32(8497),
		},
	},
	152: {
		Fname: __ccgo_ts + 3348,
		Fcodepoints: [2]uint32{
			0: uint32(1027),
		},
	},
	153: {
		Fname: __ccgo_ts + 3355,
		Fcodepoints: [2]uint32{
			0: uint32(62),
		},
	},
	154: {
		Fname: __ccgo_ts + 3360,
		Fcodepoints: [2]uint32{
			0: uint32(915),
		},
	},
	155: {
		Fname: __ccgo_ts + 3368,
		Fcodepoints: [2]uint32{
			0: uint32(988),
		},
	},
	156: {
		Fname: __ccgo_ts + 3377,
		Fcodepoints: [2]uint32{
			0: uint32(286),
		},
	},
	157: {
		Fname: __ccgo_ts + 3386,
		Fcodepoints: [2]uint32{
			0: uint32(290),
		},
	},
	158: {
		Fname: __ccgo_ts + 3395,
		Fcodepoints: [2]uint32{
			0: uint32(284),
		},
	},
	159: {
		Fname: __ccgo_ts + 3403,
		Fcodepoints: [2]uint32{
			0: uint32(1043),
		},
	},
	160: {
		Fname: __ccgo_ts + 3409,
		Fcodepoints: [2]uint32{
			0: uint32(288),
		},
	},
	161: {
		Fname: __ccgo_ts + 3416,
		Fcodepoints: [2]uint32{
			0: uint32(120074),
		},
	},
	162: {
		Fname: __ccgo_ts + 3422,
		Fcodepoints: [2]uint32{
			0: uint32(8921),
		},
	},
	163: {
		Fname: __ccgo_ts + 3427,
		Fcodepoints: [2]uint32{
			0: uint32(120126),
		},
	},
	164: {
		Fname: __ccgo_ts + 3434,
		Fcodepoints: [2]uint32{
			0: uint32(8805),
		},
	},
	165: {
		Fname: __ccgo_ts + 3449,
		Fcodepoints: [2]uint32{
			0: uint32(8923),
		},
	},
	166: {
		Fname: __ccgo_ts + 3468,
		Fcodepoints: [2]uint32{
			0: uint32(8807),
		},
	},
	167: {
		Fname: __ccgo_ts + 3487,
		Fcodepoints: [2]uint32{
			0: uint32(10914),
		},
	},
	168: {
		Fname: __ccgo_ts + 3504,
		Fcodepoints: [2]uint32{
			0: uint32(8823),
		},
	},
	169: {
		Fname: __ccgo_ts + 3518,
		Fcodepoints: [2]uint32{
			0: uint32(10878),
		},
	},
	170: {
		Fname: __ccgo_ts + 3538,
		Fcodepoints: [2]uint32{
			0: uint32(8819),
		},
	},
	171: {
		Fname: __ccgo_ts + 3553,
		Fcodepoints: [2]uint32{
			0: uint32(119970),
		},
	},
	172: {
		Fname: __ccgo_ts + 3560,
		Fcodepoints: [2]uint32{
			0: uint32(8811),
		},
	},
	173: {
		Fname: __ccgo_ts + 3565,
		Fcodepoints: [2]uint32{
			0: uint32(1066),
		},
	},
	174: {
		Fname: __ccgo_ts + 3574,
		Fcodepoints: [2]uint32{
			0: uint32(711),
		},
	},
	175: {
		Fname: __ccgo_ts + 3582,
		Fcodepoints: [2]uint32{
			0: uint32(94),
		},
	},
	176: {
		Fname: __ccgo_ts + 3588,
		Fcodepoints: [2]uint32{
			0: uint32(292),
		},
	},
	177: {
		Fname: __ccgo_ts + 3596,
		Fcodepoints: [2]uint32{
			0: uint32(8460),
		},
	},
	178: {
		Fname: __ccgo_ts + 3602,
		Fcodepoints: [2]uint32{
			0: uint32(8459),
		},
	},
	179: {
		Fname: __ccgo_ts + 3617,
		Fcodepoints: [2]uint32{
			0: uint32(8461),
		},
	},
	180: {
		Fname: __ccgo_ts + 3624,
		Fcodepoints: [2]uint32{
			0: uint32(9472),
		},
	},
	181: {
		Fname: __ccgo_ts + 3641,
		Fcodepoints: [2]uint32{
			0: uint32(8459),
		},
	},
	182: {
		Fname: __ccgo_ts + 3648,
		Fcodepoints: [2]uint32{
			0: uint32(294),
		},
	},
	183: {
		Fname: __ccgo_ts + 3657,
		Fcodepoints: [2]uint32{
			0: uint32(8782),
		},
	},
	184: {
		Fname: __ccgo_ts + 3672,
		Fcodepoints: [2]uint32{
			0: uint32(8783),
		},
	},
	185: {
		Fname: __ccgo_ts + 3684,
		Fcodepoints: [2]uint32{
			0: uint32(1045),
		},
	},
	186: {
		Fname: __ccgo_ts + 3691,
		Fcodepoints: [2]uint32{
			0: uint32(306),
		},
	},
	187: {
		Fname: __ccgo_ts + 3699,
		Fcodepoints: [2]uint32{
			0: uint32(1025),
		},
	},
	188: {
		Fname: __ccgo_ts + 3706,
		Fcodepoints: [2]uint32{
			0: uint32(205),
		},
	},
	189: {
		Fname: __ccgo_ts + 3715,
		Fcodepoints: [2]uint32{
			0: uint32(206),
		},
	},
	190: {
		Fname: __ccgo_ts + 3723,
		Fcodepoints: [2]uint32{
			0: uint32(1048),
		},
	},
	191: {
		Fname: __ccgo_ts + 3729,
		Fcodepoints: [2]uint32{
			0: uint32(304),
		},
	},
	192: {
		Fname: __ccgo_ts + 3736,
		Fcodepoints: [2]uint32{
			0: uint32(8465),
		},
	},
	193: {
		Fname: __ccgo_ts + 3742,
		Fcodepoints: [2]uint32{
			0: uint32(204),
		},
	},
	194: {
		Fname: __ccgo_ts + 3751,
		Fcodepoints: [2]uint32{
			0: uint32(8465),
		},
	},
	195: {
		Fname: __ccgo_ts + 3756,
		Fcodepoints: [2]uint32{
			0: uint32(298),
		},
	},
	196: {
		Fname: __ccgo_ts + 3764,
		Fcodepoints: [2]uint32{
			0: uint32(8520),
		},
	},
	197: {
		Fname: __ccgo_ts + 3777,
		Fcodepoints: [2]uint32{
			0: uint32(8658),
		},
	},
	198: {
		Fname: __ccgo_ts + 3787,
		Fcodepoints: [2]uint32{
			0: uint32(8748),
		},
	},
	199: {
		Fname: __ccgo_ts + 3793,
		Fcodepoints: [2]uint32{
			0: uint32(8747),
		},
	},
	200: {
		Fname: __ccgo_ts + 3804,
		Fcodepoints: [2]uint32{
			0: uint32(8898),
		},
	},
	201: {
		Fname: __ccgo_ts + 3819,
		Fcodepoints: [2]uint32{
			0: uint32(8291),
		},
	},
	202: {
		Fname: __ccgo_ts + 3836,
		Fcodepoints: [2]uint32{
			0: uint32(8290),
		},
	},
	203: {
		Fname: __ccgo_ts + 3853,
		Fcodepoints: [2]uint32{
			0: uint32(302),
		},
	},
	204: {
		Fname: __ccgo_ts + 3861,
		Fcodepoints: [2]uint32{
			0: uint32(120128),
		},
	},
	205: {
		Fname: __ccgo_ts + 3868,
		Fcodepoints: [2]uint32{
			0: uint32(921),
		},
	},
	206: {
		Fname: __ccgo_ts + 3875,
		Fcodepoints: [2]uint32{
			0: uint32(8464),
		},
	},
	207: {
		Fname: __ccgo_ts + 3882,
		Fcodepoints: [2]uint32{
			0: uint32(296),
		},
	},
	208: {
		Fname: __ccgo_ts + 3891,
		Fcodepoints: [2]uint32{
			0: uint32(1030),
		},
	},
	209: {
		Fname: __ccgo_ts + 3899,
		Fcodepoints: [2]uint32{
			0: uint32(207),
		},
	},
	210: {
		Fname: __ccgo_ts + 3906,
		Fcodepoints: [2]uint32{
			0: uint32(308),
		},
	},
	211: {
		Fname: __ccgo_ts + 3914,
		Fcodepoints: [2]uint32{
			0: uint32(1049),
		},
	},
	212: {
		Fname: __ccgo_ts + 3920,
		Fcodepoints: [2]uint32{
			0: uint32(120077),
		},
	},
	213: {
		Fname: __ccgo_ts + 3926,
		Fcodepoints: [2]uint32{
			0: uint32(120129),
		},
	},
	214: {
		Fname: __ccgo_ts + 3933,
		Fcodepoints: [2]uint32{
			0: uint32(119973),
		},
	},
	215: {
		Fname: __ccgo_ts + 3940,
		Fcodepoints: [2]uint32{
			0: uint32(1032),
		},
	},
	216: {
		Fname: __ccgo_ts + 3949,
		Fcodepoints: [2]uint32{
			0: uint32(1028),
		},
	},
	217: {
		Fname: __ccgo_ts + 3957,
		Fcodepoints: [2]uint32{
			0: uint32(1061),
		},
	},
	218: {
		Fname: __ccgo_ts + 3964,
		Fcodepoints: [2]uint32{
			0: uint32(1036),
		},
	},
	219: {
		Fname: __ccgo_ts + 3971,
		Fcodepoints: [2]uint32{
			0: uint32(922),
		},
	},
	220: {
		Fname: __ccgo_ts + 3979,
		Fcodepoints: [2]uint32{
			0: uint32(310),
		},
	},
	221: {
		Fname: __ccgo_ts + 3988,
		Fcodepoints: [2]uint32{
			0: uint32(1050),
		},
	},
	222: {
		Fname: __ccgo_ts + 3994,
		Fcodepoints: [2]uint32{
			0: uint32(120078),
		},
	},
	223: {
		Fname: __ccgo_ts + 4000,
		Fcodepoints: [2]uint32{
			0: uint32(120130),
		},
	},
	224: {
		Fname: __ccgo_ts + 4007,
		Fcodepoints: [2]uint32{
			0: uint32(119974),
		},
	},
	225: {
		Fname: __ccgo_ts + 4014,
		Fcodepoints: [2]uint32{
			0: uint32(1033),
		},
	},
	226: {
		Fname: __ccgo_ts + 4021,
		Fcodepoints: [2]uint32{
			0: uint32(60),
		},
	},
	227: {
		Fname: __ccgo_ts + 4026,
		Fcodepoints: [2]uint32{
			0: uint32(313),
		},
	},
	228: {
		Fname: __ccgo_ts + 4035,
		Fcodepoints: [2]uint32{
			0: uint32(923),
		},
	},
	229: {
		Fname: __ccgo_ts + 4044,
		Fcodepoints: [2]uint32{
			0: uint32(10218),
		},
	},
	230: {
		Fname: __ccgo_ts + 4051,
		Fcodepoints: [2]uint32{
			0: uint32(8466),
		},
	},
	231: {
		Fname: __ccgo_ts + 4064,
		Fcodepoints: [2]uint32{
			0: uint32(8606),
		},
	},
	232: {
		Fname: __ccgo_ts + 4071,
		Fcodepoints: [2]uint32{
			0: uint32(317),
		},
	},
	233: {
		Fname: __ccgo_ts + 4080,
		Fcodepoints: [2]uint32{
			0: uint32(315),
		},
	},
	234: {
		Fname: __ccgo_ts + 4089,
		Fcodepoints: [2]uint32{
			0: uint32(1051),
		},
	},
	235: {
		Fname: __ccgo_ts + 4095,
		Fcodepoints: [2]uint32{
			0: uint32(10216),
		},
	},
	236: {
		Fname: __ccgo_ts + 4114,
		Fcodepoints: [2]uint32{
			0: uint32(8592),
		},
	},
	237: {
		Fname: __ccgo_ts + 4126,
		Fcodepoints: [2]uint32{
			0: uint32(8676),
		},
	},
	238: {
		Fname: __ccgo_ts + 4141,
		Fcodepoints: [2]uint32{
			0: uint32(8646),
		},
	},
	239: {
		Fname: __ccgo_ts + 4163,
		Fcodepoints: [2]uint32{
			0: uint32(8968),
		},
	},
	240: {
		Fname: __ccgo_ts + 4177,
		Fcodepoints: [2]uint32{
			0: uint32(10214),
		},
	},
	241: {
		Fname: __ccgo_ts + 4197,
		Fcodepoints: [2]uint32{
			0: uint32(10593),
		},
	},
	242: {
		Fname: __ccgo_ts + 4217,
		Fcodepoints: [2]uint32{
			0: uint32(8643),
		},
	},
	243: {
		Fname: __ccgo_ts + 4234,
		Fcodepoints: [2]uint32{
			0: uint32(10585),
		},
	},
	244: {
		Fname: __ccgo_ts + 4254,
		Fcodepoints: [2]uint32{
			0: uint32(8970),
		},
	},
	245: {
		Fname: __ccgo_ts + 4266,
		Fcodepoints: [2]uint32{
			0: uint32(8596),
		},
	},
	246: {
		Fname: __ccgo_ts + 4283,
		Fcodepoints: [2]uint32{
			0: uint32(10574),
		},
	},
	247: {
		Fname: __ccgo_ts + 4301,
		Fcodepoints: [2]uint32{
			0: uint32(8867),
		},
	},
	248: {
		Fname: __ccgo_ts + 4311,
		Fcodepoints: [2]uint32{
			0: uint32(8612),
		},
	},
	249: {
		Fname: __ccgo_ts + 4326,
		Fcodepoints: [2]uint32{
			0: uint32(10586),
		},
	},
	250: {
		Fname: __ccgo_ts + 4342,
		Fcodepoints: [2]uint32{
			0: uint32(8882),
		},
	},
	251: {
		Fname: __ccgo_ts + 4357,
		Fcodepoints: [2]uint32{
			0: uint32(10703),
		},
	},
	252: {
		Fname: __ccgo_ts + 4375,
		Fcodepoints: [2]uint32{
			0: uint32(8884),
		},
	},
	253: {
		Fname: __ccgo_ts + 4395,
		Fcodepoints: [2]uint32{
			0: uint32(10577),
		},
	},
	254: {
		Fname: __ccgo_ts + 4414,
		Fcodepoints: [2]uint32{
			0: uint32(10592),
		},
	},
	255: {
		Fname: __ccgo_ts + 4432,
		Fcodepoints: [2]uint32{
			0: uint32(8639),
		},
	},
	256: {
		Fname: __ccgo_ts + 4447,
		Fcodepoints: [2]uint32{
			0: uint32(10584),
		},
	},
	257: {
		Fname: __ccgo_ts + 4465,
		Fcodepoints: [2]uint32{
			0: uint32(8636),
		},
	},
	258: {
		Fname: __ccgo_ts + 4478,
		Fcodepoints: [2]uint32{
			0: uint32(10578),
		},
	},
	259: {
		Fname: __ccgo_ts + 4494,
		Fcodepoints: [2]uint32{
			0: uint32(8656),
		},
	},
	260: {
		Fname: __ccgo_ts + 4506,
		Fcodepoints: [2]uint32{
			0: uint32(8660),
		},
	},
	261: {
		Fname: __ccgo_ts + 4523,
		Fcodepoints: [2]uint32{
			0: uint32(8922),
		},
	},
	262: {
		Fname: __ccgo_ts + 4542,
		Fcodepoints: [2]uint32{
			0: uint32(8806),
		},
	},
	263: {
		Fname: __ccgo_ts + 4558,
		Fcodepoints: [2]uint32{
			0: uint32(8822),
		},
	},
	264: {
		Fname: __ccgo_ts + 4572,
		Fcodepoints: [2]uint32{
			0: uint32(10913),
		},
	},
	265: {
		Fname: __ccgo_ts + 4583,
		Fcodepoints: [2]uint32{
			0: uint32(10877),
		},
	},
	266: {
		Fname: __ccgo_ts + 4600,
		Fcodepoints: [2]uint32{
			0: uint32(8818),
		},
	},
	267: {
		Fname: __ccgo_ts + 4612,
		Fcodepoints: [2]uint32{
			0: uint32(120079),
		},
	},
	268: {
		Fname: __ccgo_ts + 4618,
		Fcodepoints: [2]uint32{
			0: uint32(8920),
		},
	},
	269: {
		Fname: __ccgo_ts + 4623,
		Fcodepoints: [2]uint32{
			0: uint32(8666),
		},
	},
	270: {
		Fname: __ccgo_ts + 4636,
		Fcodepoints: [2]uint32{
			0: uint32(319),
		},
	},
	271: {
		Fname: __ccgo_ts + 4645,
		Fcodepoints: [2]uint32{
			0: uint32(10229),
		},
	},
	272: {
		Fname: __ccgo_ts + 4661,
		Fcodepoints: [2]uint32{
			0: uint32(10231),
		},
	},
	273: {
		Fname: __ccgo_ts + 4682,
		Fcodepoints: [2]uint32{
			0: uint32(10230),
		},
	},
	274: {
		Fname: __ccgo_ts + 4699,
		Fcodepoints: [2]uint32{
			0: uint32(10232),
		},
	},
	275: {
		Fname: __ccgo_ts + 4715,
		Fcodepoints: [2]uint32{
			0: uint32(10234),
		},
	},
	276: {
		Fname: __ccgo_ts + 4736,
		Fcodepoints: [2]uint32{
			0: uint32(10233),
		},
	},
	277: {
		Fname: __ccgo_ts + 4753,
		Fcodepoints: [2]uint32{
			0: uint32(120131),
		},
	},
	278: {
		Fname: __ccgo_ts + 4760,
		Fcodepoints: [2]uint32{
			0: uint32(8601),
		},
	},
	279: {
		Fname: __ccgo_ts + 4777,
		Fcodepoints: [2]uint32{
			0: uint32(8600),
		},
	},
	280: {
		Fname: __ccgo_ts + 4795,
		Fcodepoints: [2]uint32{
			0: uint32(8466),
		},
	},
	281: {
		Fname: __ccgo_ts + 4802,
		Fcodepoints: [2]uint32{
			0: uint32(8624),
		},
	},
	282: {
		Fname: __ccgo_ts + 4808,
		Fcodepoints: [2]uint32{
			0: uint32(321),
		},
	},
	283: {
		Fname: __ccgo_ts + 4817,
		Fcodepoints: [2]uint32{
			0: uint32(8810),
		},
	},
	284: {
		Fname: __ccgo_ts + 4822,
		Fcodepoints: [2]uint32{
			0: uint32(10501),
		},
	},
	285: {
		Fname: __ccgo_ts + 4828,
		Fcodepoints: [2]uint32{
			0: uint32(1052),
		},
	},
	286: {
		Fname: __ccgo_ts + 4834,
		Fcodepoints: [2]uint32{
			0: uint32(8287),
		},
	},
	287: {
		Fname: __ccgo_ts + 4848,
		Fcodepoints: [2]uint32{
			0: uint32(8499),
		},
	},
	288: {
		Fname: __ccgo_ts + 4860,
		Fcodepoints: [2]uint32{
			0: uint32(120080),
		},
	},
	289: {
		Fname: __ccgo_ts + 4866,
		Fcodepoints: [2]uint32{
			0: uint32(8723),
		},
	},
	290: {
		Fname: __ccgo_ts + 4878,
		Fcodepoints: [2]uint32{
			0: uint32(120132),
		},
	},
	291: {
		Fname: __ccgo_ts + 4885,
		Fcodepoints: [2]uint32{
			0: uint32(8499),
		},
	},
	292: {
		Fname: __ccgo_ts + 4892,
		Fcodepoints: [2]uint32{
			0: uint32(924),
		},
	},
	293: {
		Fname: __ccgo_ts + 4897,
		Fcodepoints: [2]uint32{
			0: uint32(1034),
		},
	},
	294: {
		Fname: __ccgo_ts + 4904,
		Fcodepoints: [2]uint32{
			0: uint32(323),
		},
	},
	295: {
		Fname: __ccgo_ts + 4913,
		Fcodepoints: [2]uint32{
			0: uint32(327),
		},
	},
	296: {
		Fname: __ccgo_ts + 4922,
		Fcodepoints: [2]uint32{
			0: uint32(325),
		},
	},
	297: {
		Fname: __ccgo_ts + 4931,
		Fcodepoints: [2]uint32{
			0: uint32(1053),
		},
	},
	298: {
		Fname: __ccgo_ts + 4937,
		Fcodepoints: [2]uint32{
			0: uint32(8203),
		},
	},
	299: {
		Fname: __ccgo_ts + 4959,
		Fcodepoints: [2]uint32{
			0: uint32(8203),
		},
	},
	300: {
		Fname: __ccgo_ts + 4980,
		Fcodepoints: [2]uint32{
			0: uint32(8203),
		},
	},
	301: {
		Fname: __ccgo_ts + 5000,
		Fcodepoints: [2]uint32{
			0: uint32(8203),
		},
	},
	302: {
		Fname: __ccgo_ts + 5024,
		Fcodepoints: [2]uint32{
			0: uint32(8811),
		},
	},
	303: {
		Fname: __ccgo_ts + 5047,
		Fcodepoints: [2]uint32{
			0: uint32(8810),
		},
	},
	304: {
		Fname: __ccgo_ts + 5064,
		Fcodepoints: [2]uint32{
			0: uint32(10),
		},
	},
	305: {
		Fname: __ccgo_ts + 5074,
		Fcodepoints: [2]uint32{
			0: uint32(120081),
		},
	},
	306: {
		Fname: __ccgo_ts + 5080,
		Fcodepoints: [2]uint32{
			0: uint32(8288),
		},
	},
	307: {
		Fname: __ccgo_ts + 5090,
		Fcodepoints: [2]uint32{
			0: uint32(160),
		},
	},
	308: {
		Fname: __ccgo_ts + 5109,
		Fcodepoints: [2]uint32{
			0: uint32(8469),
		},
	},
	309: {
		Fname: __ccgo_ts + 5116,
		Fcodepoints: [2]uint32{
			0: uint32(10988),
		},
	},
	310: {
		Fname: __ccgo_ts + 5122,
		Fcodepoints: [2]uint32{
			0: uint32(8802),
		},
	},
	311: {
		Fname: __ccgo_ts + 5137,
		Fcodepoints: [2]uint32{
			0: uint32(8813),
		},
	},
	312: {
		Fname: __ccgo_ts + 5149,
		Fcodepoints: [2]uint32{
			0: uint32(8742),
		},
	},
	313: {
		Fname: __ccgo_ts + 5172,
		Fcodepoints: [2]uint32{
			0: uint32(8713),
		},
	},
	314: {
		Fname: __ccgo_ts + 5185,
		Fcodepoints: [2]uint32{
			0: uint32(8800),
		},
	},
	315: {
		Fname: __ccgo_ts + 5196,
		Fcodepoints: [2]uint32{
			0: uint32(8770),
			1: uint32(824),
		},
	},
	316: {
		Fname: __ccgo_ts + 5212,
		Fcodepoints: [2]uint32{
			0: uint32(8708),
		},
	},
	317: {
		Fname: __ccgo_ts + 5224,
		Fcodepoints: [2]uint32{
			0: uint32(8815),
		},
	},
	318: {
		Fname: __ccgo_ts + 5237,
		Fcodepoints: [2]uint32{
			0: uint32(8817),
		},
	},
	319: {
		Fname: __ccgo_ts + 5255,
		Fcodepoints: [2]uint32{
			0: uint32(8807),
			1: uint32(824),
		},
	},
	320: {
		Fname: __ccgo_ts + 5277,
		Fcodepoints: [2]uint32{
			0: uint32(8811),
			1: uint32(824),
		},
	},
	321: {
		Fname: __ccgo_ts + 5297,
		Fcodepoints: [2]uint32{
			0: uint32(8825),
		},
	},
	322: {
		Fname: __ccgo_ts + 5314,
		Fcodepoints: [2]uint32{
			0: uint32(10878),
			1: uint32(824),
		},
	},
	323: {
		Fname: __ccgo_ts + 5337,
		Fcodepoints: [2]uint32{
			0: uint32(8821),
		},
	},
	324: {
		Fname: __ccgo_ts + 5355,
		Fcodepoints: [2]uint32{
			0: uint32(8782),
			1: uint32(824),
		},
	},
	325: {
		Fname: __ccgo_ts + 5373,
		Fcodepoints: [2]uint32{
			0: uint32(8783),
			1: uint32(824),
		},
	},
	326: {
		Fname: __ccgo_ts + 5388,
		Fcodepoints: [2]uint32{
			0: uint32(8938),
		},
	},
	327: {
		Fname: __ccgo_ts + 5406,
		Fcodepoints: [2]uint32{
			0: uint32(10703),
			1: uint32(824),
		},
	},
	328: {
		Fname: __ccgo_ts + 5427,
		Fcodepoints: [2]uint32{
			0: uint32(8940),
		},
	},
	329: {
		Fname: __ccgo_ts + 5450,
		Fcodepoints: [2]uint32{
			0: uint32(8814),
		},
	},
	330: {
		Fname: __ccgo_ts + 5460,
		Fcodepoints: [2]uint32{
			0: uint32(8816),
		},
	},
	331: {
		Fname: __ccgo_ts + 5475,
		Fcodepoints: [2]uint32{
			0: uint32(8824),
		},
	},
	332: {
		Fname: __ccgo_ts + 5492,
		Fcodepoints: [2]uint32{
			0: uint32(8810),
			1: uint32(824),
		},
	},
	333: {
		Fname: __ccgo_ts + 5506,
		Fcodepoints: [2]uint32{
			0: uint32(10877),
			1: uint32(824),
		},
	},
	334: {
		Fname: __ccgo_ts + 5526,
		Fcodepoints: [2]uint32{
			0: uint32(8820),
		},
	},
	335: {
		Fname: __ccgo_ts + 5541,
		Fcodepoints: [2]uint32{
			0: uint32(10914),
			1: uint32(824),
		},
	},
	336: {
		Fname: __ccgo_ts + 5567,
		Fcodepoints: [2]uint32{
			0: uint32(10913),
			1: uint32(824),
		},
	},
	337: {
		Fname: __ccgo_ts + 5587,
		Fcodepoints: [2]uint32{
			0: uint32(8832),
		},
	},
	338: {
		Fname: __ccgo_ts + 5601,
		Fcodepoints: [2]uint32{
			0: uint32(10927),
			1: uint32(824),
		},
	},
	339: {
		Fname: __ccgo_ts + 5620,
		Fcodepoints: [2]uint32{
			0: uint32(8928),
		},
	},
	340: {
		Fname: __ccgo_ts + 5644,
		Fcodepoints: [2]uint32{
			0: uint32(8716),
		},
	},
	341: {
		Fname: __ccgo_ts + 5664,
		Fcodepoints: [2]uint32{
			0: uint32(8939),
		},
	},
	342: {
		Fname: __ccgo_ts + 5683,
		Fcodepoints: [2]uint32{
			0: uint32(10704),
			1: uint32(824),
		},
	},
	343: {
		Fname: __ccgo_ts + 5705,
		Fcodepoints: [2]uint32{
			0: uint32(8941),
		},
	},
	344: {
		Fname: __ccgo_ts + 5729,
		Fcodepoints: [2]uint32{
			0: uint32(8847),
			1: uint32(824),
		},
	},
	345: {
		Fname: __ccgo_ts + 5747,
		Fcodepoints: [2]uint32{
			0: uint32(8930),
		},
	},
	346: {
		Fname: __ccgo_ts + 5770,
		Fcodepoints: [2]uint32{
			0: uint32(8848),
			1: uint32(824),
		},
	},
	347: {
		Fname: __ccgo_ts + 5790,
		Fcodepoints: [2]uint32{
			0: uint32(8931),
		},
	},
	348: {
		Fname: __ccgo_ts + 5815,
		Fcodepoints: [2]uint32{
			0: uint32(8834),
			1: uint32(8402),
		},
	},
	349: {
		Fname: __ccgo_ts + 5827,
		Fcodepoints: [2]uint32{
			0: uint32(8840),
		},
	},
	350: {
		Fname: __ccgo_ts + 5844,
		Fcodepoints: [2]uint32{
			0: uint32(8833),
		},
	},
	351: {
		Fname: __ccgo_ts + 5858,
		Fcodepoints: [2]uint32{
			0: uint32(10928),
			1: uint32(824),
		},
	},
	352: {
		Fname: __ccgo_ts + 5877,
		Fcodepoints: [2]uint32{
			0: uint32(8929),
		},
	},
	353: {
		Fname: __ccgo_ts + 5901,
		Fcodepoints: [2]uint32{
			0: uint32(8831),
			1: uint32(824),
		},
	},
	354: {
		Fname: __ccgo_ts + 5920,
		Fcodepoints: [2]uint32{
			0: uint32(8835),
			1: uint32(8402),
		},
	},
	355: {
		Fname: __ccgo_ts + 5934,
		Fcodepoints: [2]uint32{
			0: uint32(8841),
		},
	},
	356: {
		Fname: __ccgo_ts + 5953,
		Fcodepoints: [2]uint32{
			0: uint32(8769),
		},
	},
	357: {
		Fname: __ccgo_ts + 5964,
		Fcodepoints: [2]uint32{
			0: uint32(8772),
		},
	},
	358: {
		Fname: __ccgo_ts + 5980,
		Fcodepoints: [2]uint32{
			0: uint32(8775),
		},
	},
	359: {
		Fname: __ccgo_ts + 6000,
		Fcodepoints: [2]uint32{
			0: uint32(8777),
		},
	},
	360: {
		Fname: __ccgo_ts + 6016,
		Fcodepoints: [2]uint32{
			0: uint32(8740),
		},
	},
	361: {
		Fname: __ccgo_ts + 6033,
		Fcodepoints: [2]uint32{
			0: uint32(119977),
		},
	},
	362: {
		Fname: __ccgo_ts + 6040,
		Fcodepoints: [2]uint32{
			0: uint32(209),
		},
	},
	363: {
		Fname: __ccgo_ts + 6049,
		Fcodepoints: [2]uint32{
			0: uint32(925),
		},
	},
	364: {
		Fname: __ccgo_ts + 6054,
		Fcodepoints: [2]uint32{
			0: uint32(338),
		},
	},
	365: {
		Fname: __ccgo_ts + 6062,
		Fcodepoints: [2]uint32{
			0: uint32(211),
		},
	},
	366: {
		Fname: __ccgo_ts + 6071,
		Fcodepoints: [2]uint32{
			0: uint32(212),
		},
	},
	367: {
		Fname: __ccgo_ts + 6079,
		Fcodepoints: [2]uint32{
			0: uint32(1054),
		},
	},
	368: {
		Fname: __ccgo_ts + 6085,
		Fcodepoints: [2]uint32{
			0: uint32(336),
		},
	},
	369: {
		Fname: __ccgo_ts + 6094,
		Fcodepoints: [2]uint32{
			0: uint32(120082),
		},
	},
	370: {
		Fname: __ccgo_ts + 6100,
		Fcodepoints: [2]uint32{
			0: uint32(210),
		},
	},
	371: {
		Fname: __ccgo_ts + 6109,
		Fcodepoints: [2]uint32{
			0: uint32(332),
		},
	},
	372: {
		Fname: __ccgo_ts + 6117,
		Fcodepoints: [2]uint32{
			0: uint32(937),
		},
	},
	373: {
		Fname: __ccgo_ts + 6125,
		Fcodepoints: [2]uint32{
			0: uint32(927),
		},
	},
	374: {
		Fname: __ccgo_ts + 6135,
		Fcodepoints: [2]uint32{
			0: uint32(120134),
		},
	},
	375: {
		Fname: __ccgo_ts + 6142,
		Fcodepoints: [2]uint32{
			0: uint32(8220),
		},
	},
	376: {
		Fname: __ccgo_ts + 6165,
		Fcodepoints: [2]uint32{
			0: uint32(8216),
		},
	},
	377: {
		Fname: __ccgo_ts + 6182,
		Fcodepoints: [2]uint32{
			0: uint32(10836),
		},
	},
	378: {
		Fname: __ccgo_ts + 6187,
		Fcodepoints: [2]uint32{
			0: uint32(119978),
		},
	},
	379: {
		Fname: __ccgo_ts + 6194,
		Fcodepoints: [2]uint32{
			0: uint32(216),
		},
	},
	380: {
		Fname: __ccgo_ts + 6203,
		Fcodepoints: [2]uint32{
			0: uint32(213),
		},
	},
	381: {
		Fname: __ccgo_ts + 6212,
		Fcodepoints: [2]uint32{
			0: uint32(10807),
		},
	},
	382: {
		Fname: __ccgo_ts + 6221,
		Fcodepoints: [2]uint32{
			0: uint32(214),
		},
	},
	383: {
		Fname: __ccgo_ts + 6228,
		Fcodepoints: [2]uint32{
			0: uint32(8254),
		},
	},
	384: {
		Fname: __ccgo_ts + 6238,
		Fcodepoints: [2]uint32{
			0: uint32(9182),
		},
	},
	385: {
		Fname: __ccgo_ts + 6250,
		Fcodepoints: [2]uint32{
			0: uint32(9140),
		},
	},
	386: {
		Fname: __ccgo_ts + 6264,
		Fcodepoints: [2]uint32{
			0: uint32(9180),
		},
	},
	387: {
		Fname: __ccgo_ts + 6282,
		Fcodepoints: [2]uint32{
			0: uint32(8706),
		},
	},
	388: {
		Fname: __ccgo_ts + 6293,
		Fcodepoints: [2]uint32{
			0: uint32(1055),
		},
	},
	389: {
		Fname: __ccgo_ts + 6299,
		Fcodepoints: [2]uint32{
			0: uint32(120083),
		},
	},
	390: {
		Fname: __ccgo_ts + 6305,
		Fcodepoints: [2]uint32{
			0: uint32(934),
		},
	},
	391: {
		Fname: __ccgo_ts + 6311,
		Fcodepoints: [2]uint32{
			0: uint32(928),
		},
	},
	392: {
		Fname: __ccgo_ts + 6316,
		Fcodepoints: [2]uint32{
			0: uint32(177),
		},
	},
	393: {
		Fname: __ccgo_ts + 6328,
		Fcodepoints: [2]uint32{
			0: uint32(8460),
		},
	},
	394: {
		Fname: __ccgo_ts + 6344,
		Fcodepoints: [2]uint32{
			0: uint32(8473),
		},
	},
	395: {
		Fname: __ccgo_ts + 6351,
		Fcodepoints: [2]uint32{
			0: uint32(10939),
		},
	},
	396: {
		Fname: __ccgo_ts + 6356,
		Fcodepoints: [2]uint32{
			0: uint32(8826),
		},
	},
	397: {
		Fname: __ccgo_ts + 6367,
		Fcodepoints: [2]uint32{
			0: uint32(10927),
		},
	},
	398: {
		Fname: __ccgo_ts + 6383,
		Fcodepoints: [2]uint32{
			0: uint32(8828),
		},
	},
	399: {
		Fname: __ccgo_ts + 6404,
		Fcodepoints: [2]uint32{
			0: uint32(8830),
		},
	},
	400: {
		Fname: __ccgo_ts + 6420,
		Fcodepoints: [2]uint32{
			0: uint32(8243),
		},
	},
	401: {
		Fname: __ccgo_ts + 6428,
		Fcodepoints: [2]uint32{
			0: uint32(8719),
		},
	},
	402: {
		Fname: __ccgo_ts + 6438,
		Fcodepoints: [2]uint32{
			0: uint32(8759),
		},
	},
	403: {
		Fname: __ccgo_ts + 6451,
		Fcodepoints: [2]uint32{
			0: uint32(8733),
		},
	},
	404: {
		Fname: __ccgo_ts + 6466,
		Fcodepoints: [2]uint32{
			0: uint32(119979),
		},
	},
	405: {
		Fname: __ccgo_ts + 6473,
		Fcodepoints: [2]uint32{
			0: uint32(936),
		},
	},
	406: {
		Fname: __ccgo_ts + 6479,
		Fcodepoints: [2]uint32{
			0: uint32(34),
		},
	},
	407: {
		Fname: __ccgo_ts + 6486,
		Fcodepoints: [2]uint32{
			0: uint32(120084),
		},
	},
	408: {
		Fname: __ccgo_ts + 6492,
		Fcodepoints: [2]uint32{
			0: uint32(8474),
		},
	},
	409: {
		Fname: __ccgo_ts + 6499,
		Fcodepoints: [2]uint32{
			0: uint32(119980),
		},
	},
	410: {
		Fname: __ccgo_ts + 6506,
		Fcodepoints: [2]uint32{
			0: uint32(10512),
		},
	},
	411: {
		Fname: __ccgo_ts + 6514,
		Fcodepoints: [2]uint32{
			0: uint32(174),
		},
	},
	412: {
		Fname: __ccgo_ts + 6520,
		Fcodepoints: [2]uint32{
			0: uint32(340),
		},
	},
	413: {
		Fname: __ccgo_ts + 6529,
		Fcodepoints: [2]uint32{
			0: uint32(10219),
		},
	},
	414: {
		Fname: __ccgo_ts + 6536,
		Fcodepoints: [2]uint32{
			0: uint32(8608),
		},
	},
	415: {
		Fname: __ccgo_ts + 6543,
		Fcodepoints: [2]uint32{
			0: uint32(10518),
		},
	},
	416: {
		Fname: __ccgo_ts + 6552,
		Fcodepoints: [2]uint32{
			0: uint32(344),
		},
	},
	417: {
		Fname: __ccgo_ts + 6561,
		Fcodepoints: [2]uint32{
			0: uint32(342),
		},
	},
	418: {
		Fname: __ccgo_ts + 6570,
		Fcodepoints: [2]uint32{
			0: uint32(1056),
		},
	},
	419: {
		Fname: __ccgo_ts + 6576,
		Fcodepoints: [2]uint32{
			0: uint32(8476),
		},
	},
	420: {
		Fname: __ccgo_ts + 6581,
		Fcodepoints: [2]uint32{
			0: uint32(8715),
		},
	},
	421: {
		Fname: __ccgo_ts + 6598,
		Fcodepoints: [2]uint32{
			0: uint32(8651),
		},
	},
	422: {
		Fname: __ccgo_ts + 6619,
		Fcodepoints: [2]uint32{
			0: uint32(10607),
		},
	},
	423: {
		Fname: __ccgo_ts + 6642,
		Fcodepoints: [2]uint32{
			0: uint32(8476),
		},
	},
	424: {
		Fname: __ccgo_ts + 6648,
		Fcodepoints: [2]uint32{
			0: uint32(929),
		},
	},
	425: {
		Fname: __ccgo_ts + 6654,
		Fcodepoints: [2]uint32{
			0: uint32(10217),
		},
	},
	426: {
		Fname: __ccgo_ts + 6674,
		Fcodepoints: [2]uint32{
			0: uint32(8594),
		},
	},
	427: {
		Fname: __ccgo_ts + 6687,
		Fcodepoints: [2]uint32{
			0: uint32(8677),
		},
	},
	428: {
		Fname: __ccgo_ts + 6703,
		Fcodepoints: [2]uint32{
			0: uint32(8644),
		},
	},
	429: {
		Fname: __ccgo_ts + 6725,
		Fcodepoints: [2]uint32{
			0: uint32(8969),
		},
	},
	430: {
		Fname: __ccgo_ts + 6740,
		Fcodepoints: [2]uint32{
			0: uint32(10215),
		},
	},
	431: {
		Fname: __ccgo_ts + 6761,
		Fcodepoints: [2]uint32{
			0: uint32(10589),
		},
	},
	432: {
		Fname: __ccgo_ts + 6782,
		Fcodepoints: [2]uint32{
			0: uint32(8642),
		},
	},
	433: {
		Fname: __ccgo_ts + 6800,
		Fcodepoints: [2]uint32{
			0: uint32(10581),
		},
	},
	434: {
		Fname: __ccgo_ts + 6821,
		Fcodepoints: [2]uint32{
			0: uint32(8971),
		},
	},
	435: {
		Fname: __ccgo_ts + 6834,
		Fcodepoints: [2]uint32{
			0: uint32(8866),
		},
	},
	436: {
		Fname: __ccgo_ts + 6845,
		Fcodepoints: [2]uint32{
			0: uint32(8614),
		},
	},
	437: {
		Fname: __ccgo_ts + 6861,
		Fcodepoints: [2]uint32{
			0: uint32(10587),
		},
	},
	438: {
		Fname: __ccgo_ts + 6878,
		Fcodepoints: [2]uint32{
			0: uint32(8883),
		},
	},
	439: {
		Fname: __ccgo_ts + 6894,
		Fcodepoints: [2]uint32{
			0: uint32(10704),
		},
	},
	440: {
		Fname: __ccgo_ts + 6913,
		Fcodepoints: [2]uint32{
			0: uint32(8885),
		},
	},
	441: {
		Fname: __ccgo_ts + 6934,
		Fcodepoints: [2]uint32{
			0: uint32(10575),
		},
	},
	442: {
		Fname: __ccgo_ts + 6954,
		Fcodepoints: [2]uint32{
			0: uint32(10588),
		},
	},
	443: {
		Fname: __ccgo_ts + 6973,
		Fcodepoints: [2]uint32{
			0: uint32(8638),
		},
	},
	444: {
		Fname: __ccgo_ts + 6989,
		Fcodepoints: [2]uint32{
			0: uint32(10580),
		},
	},
	445: {
		Fname: __ccgo_ts + 7008,
		Fcodepoints: [2]uint32{
			0: uint32(8640),
		},
	},
	446: {
		Fname: __ccgo_ts + 7022,
		Fcodepoints: [2]uint32{
			0: uint32(10579),
		},
	},
	447: {
		Fname: __ccgo_ts + 7039,
		Fcodepoints: [2]uint32{
			0: uint32(8658),
		},
	},
	448: {
		Fname: __ccgo_ts + 7052,
		Fcodepoints: [2]uint32{
			0: uint32(8477),
		},
	},
	449: {
		Fname: __ccgo_ts + 7059,
		Fcodepoints: [2]uint32{
			0: uint32(10608),
		},
	},
	450: {
		Fname: __ccgo_ts + 7074,
		Fcodepoints: [2]uint32{
			0: uint32(8667),
		},
	},
	451: {
		Fname: __ccgo_ts + 7088,
		Fcodepoints: [2]uint32{
			0: uint32(8475),
		},
	},
	452: {
		Fname: __ccgo_ts + 7095,
		Fcodepoints: [2]uint32{
			0: uint32(8625),
		},
	},
	453: {
		Fname: __ccgo_ts + 7101,
		Fcodepoints: [2]uint32{
			0: uint32(10740),
		},
	},
	454: {
		Fname: __ccgo_ts + 7115,
		Fcodepoints: [2]uint32{
			0: uint32(1065),
		},
	},
	455: {
		Fname: __ccgo_ts + 7124,
		Fcodepoints: [2]uint32{
			0: uint32(1064),
		},
	},
	456: {
		Fname: __ccgo_ts + 7131,
		Fcodepoints: [2]uint32{
			0: uint32(1068),
		},
	},
	457: {
		Fname: __ccgo_ts + 7140,
		Fcodepoints: [2]uint32{
			0: uint32(346),
		},
	},
	458: {
		Fname: __ccgo_ts + 7149,
		Fcodepoints: [2]uint32{
			0: uint32(10940),
		},
	},
	459: {
		Fname: __ccgo_ts + 7154,
		Fcodepoints: [2]uint32{
			0: uint32(352),
		},
	},
	460: {
		Fname: __ccgo_ts + 7163,
		Fcodepoints: [2]uint32{
			0: uint32(350),
		},
	},
	461: {
		Fname: __ccgo_ts + 7172,
		Fcodepoints: [2]uint32{
			0: uint32(348),
		},
	},
	462: {
		Fname: __ccgo_ts + 7180,
		Fcodepoints: [2]uint32{
			0: uint32(1057),
		},
	},
	463: {
		Fname: __ccgo_ts + 7186,
		Fcodepoints: [2]uint32{
			0: uint32(120086),
		},
	},
	464: {
		Fname: __ccgo_ts + 7192,
		Fcodepoints: [2]uint32{
			0: uint32(8595),
		},
	},
	465: {
		Fname: __ccgo_ts + 7209,
		Fcodepoints: [2]uint32{
			0: uint32(8592),
		},
	},
	466: {
		Fname: __ccgo_ts + 7226,
		Fcodepoints: [2]uint32{
			0: uint32(8594),
		},
	},
	467: {
		Fname: __ccgo_ts + 7244,
		Fcodepoints: [2]uint32{
			0: uint32(8593),
		},
	},
	468: {
		Fname: __ccgo_ts + 7259,
		Fcodepoints: [2]uint32{
			0: uint32(931),
		},
	},
	469: {
		Fname: __ccgo_ts + 7267,
		Fcodepoints: [2]uint32{
			0: uint32(8728),
		},
	},
	470: {
		Fname: __ccgo_ts + 7281,
		Fcodepoints: [2]uint32{
			0: uint32(120138),
		},
	},
	471: {
		Fname: __ccgo_ts + 7288,
		Fcodepoints: [2]uint32{
			0: uint32(8730),
		},
	},
	472: {
		Fname: __ccgo_ts + 7295,
		Fcodepoints: [2]uint32{
			0: uint32(9633),
		},
	},
	473: {
		Fname: __ccgo_ts + 7304,
		Fcodepoints: [2]uint32{
			0: uint32(8851),
		},
	},
	474: {
		Fname: __ccgo_ts + 7325,
		Fcodepoints: [2]uint32{
			0: uint32(8847),
		},
	},
	475: {
		Fname: __ccgo_ts + 7340,
		Fcodepoints: [2]uint32{
			0: uint32(8849),
		},
	},
	476: {
		Fname: __ccgo_ts + 7360,
		Fcodepoints: [2]uint32{
			0: uint32(8848),
		},
	},
	477: {
		Fname: __ccgo_ts + 7377,
		Fcodepoints: [2]uint32{
			0: uint32(8850),
		},
	},
	478: {
		Fname: __ccgo_ts + 7399,
		Fcodepoints: [2]uint32{
			0: uint32(8852),
		},
	},
	479: {
		Fname: __ccgo_ts + 7413,
		Fcodepoints: [2]uint32{
			0: uint32(119982),
		},
	},
	480: {
		Fname: __ccgo_ts + 7420,
		Fcodepoints: [2]uint32{
			0: uint32(8902),
		},
	},
	481: {
		Fname: __ccgo_ts + 7427,
		Fcodepoints: [2]uint32{
			0: uint32(8912),
		},
	},
	482: {
		Fname: __ccgo_ts + 7433,
		Fcodepoints: [2]uint32{
			0: uint32(8912),
		},
	},
	483: {
		Fname: __ccgo_ts + 7442,
		Fcodepoints: [2]uint32{
			0: uint32(8838),
		},
	},
	484: {
		Fname: __ccgo_ts + 7456,
		Fcodepoints: [2]uint32{
			0: uint32(8827),
		},
	},
	485: {
		Fname: __ccgo_ts + 7467,
		Fcodepoints: [2]uint32{
			0: uint32(10928),
		},
	},
	486: {
		Fname: __ccgo_ts + 7483,
		Fcodepoints: [2]uint32{
			0: uint32(8829),
		},
	},
	487: {
		Fname: __ccgo_ts + 7504,
		Fcodepoints: [2]uint32{
			0: uint32(8831),
		},
	},
	488: {
		Fname: __ccgo_ts + 7520,
		Fcodepoints: [2]uint32{
			0: uint32(8715),
		},
	},
	489: {
		Fname: __ccgo_ts + 7531,
		Fcodepoints: [2]uint32{
			0: uint32(8721),
		},
	},
	490: {
		Fname: __ccgo_ts + 7537,
		Fcodepoints: [2]uint32{
			0: uint32(8913),
		},
	},
	491: {
		Fname: __ccgo_ts + 7543,
		Fcodepoints: [2]uint32{
			0: uint32(8835),
		},
	},
	492: {
		Fname: __ccgo_ts + 7554,
		Fcodepoints: [2]uint32{
			0: uint32(8839),
		},
	},
	493: {
		Fname: __ccgo_ts + 7570,
		Fcodepoints: [2]uint32{
			0: uint32(8913),
		},
	},
	494: {
		Fname: __ccgo_ts + 7579,
		Fcodepoints: [2]uint32{
			0: uint32(222),
		},
	},
	495: {
		Fname: __ccgo_ts + 7587,
		Fcodepoints: [2]uint32{
			0: uint32(8482),
		},
	},
	496: {
		Fname: __ccgo_ts + 7595,
		Fcodepoints: [2]uint32{
			0: uint32(1035),
		},
	},
	497: {
		Fname: __ccgo_ts + 7603,
		Fcodepoints: [2]uint32{
			0: uint32(1062),
		},
	},
	498: {
		Fname: __ccgo_ts + 7610,
		Fcodepoints: [2]uint32{
			0: uint32(9),
		},
	},
	499: {
		Fname: __ccgo_ts + 7616,
		Fcodepoints: [2]uint32{
			0: uint32(932),
		},
	},
	500: {
		Fname: __ccgo_ts + 7622,
		Fcodepoints: [2]uint32{
			0: uint32(356),
		},
	},
	501: {
		Fname: __ccgo_ts + 7631,
		Fcodepoints: [2]uint32{
			0: uint32(354),
		},
	},
	502: {
		Fname: __ccgo_ts + 7640,
		Fcodepoints: [2]uint32{
			0: uint32(1058),
		},
	},
	503: {
		Fname: __ccgo_ts + 7646,
		Fcodepoints: [2]uint32{
			0: uint32(120087),
		},
	},
	504: {
		Fname: __ccgo_ts + 7652,
		Fcodepoints: [2]uint32{
			0: uint32(8756),
		},
	},
	505: {
		Fname: __ccgo_ts + 7664,
		Fcodepoints: [2]uint32{
			0: uint32(920),
		},
	},
	506: {
		Fname: __ccgo_ts + 7672,
		Fcodepoints: [2]uint32{
			0: uint32(8287),
			1: uint32(8202),
		},
	},
	507: {
		Fname: __ccgo_ts + 7685,
		Fcodepoints: [2]uint32{
			0: uint32(8201),
		},
	},
	508: {
		Fname: __ccgo_ts + 7697,
		Fcodepoints: [2]uint32{
			0: uint32(8764),
		},
	},
	509: {
		Fname: __ccgo_ts + 7705,
		Fcodepoints: [2]uint32{
			0: uint32(8771),
		},
	},
	510: {
		Fname: __ccgo_ts + 7718,
		Fcodepoints: [2]uint32{
			0: uint32(8773),
		},
	},
	511: {
		Fname: __ccgo_ts + 7735,
		Fcodepoints: [2]uint32{
			0: uint32(8776),
		},
	},
	512: {
		Fname: __ccgo_ts + 7748,
		Fcodepoints: [2]uint32{
			0: uint32(120139),
		},
	},
	513: {
		Fname: __ccgo_ts + 7755,
		Fcodepoints: [2]uint32{
			0: uint32(8411),
		},
	},
	514: {
		Fname: __ccgo_ts + 7767,
		Fcodepoints: [2]uint32{
			0: uint32(119983),
		},
	},
	515: {
		Fname: __ccgo_ts + 7774,
		Fcodepoints: [2]uint32{
			0: uint32(358),
		},
	},
	516: {
		Fname: __ccgo_ts + 7783,
		Fcodepoints: [2]uint32{
			0: uint32(218),
		},
	},
	517: {
		Fname: __ccgo_ts + 7792,
		Fcodepoints: [2]uint32{
			0: uint32(8607),
		},
	},
	518: {
		Fname: __ccgo_ts + 7799,
		Fcodepoints: [2]uint32{
			0: uint32(10569),
		},
	},
	519: {
		Fname: __ccgo_ts + 7810,
		Fcodepoints: [2]uint32{
			0: uint32(1038),
		},
	},
	520: {
		Fname: __ccgo_ts + 7818,
		Fcodepoints: [2]uint32{
			0: uint32(364),
		},
	},
	521: {
		Fname: __ccgo_ts + 7827,
		Fcodepoints: [2]uint32{
			0: uint32(219),
		},
	},
	522: {
		Fname: __ccgo_ts + 7835,
		Fcodepoints: [2]uint32{
			0: uint32(1059),
		},
	},
	523: {
		Fname: __ccgo_ts + 7841,
		Fcodepoints: [2]uint32{
			0: uint32(368),
		},
	},
	524: {
		Fname: __ccgo_ts + 7850,
		Fcodepoints: [2]uint32{
			0: uint32(120088),
		},
	},
	525: {
		Fname: __ccgo_ts + 7856,
		Fcodepoints: [2]uint32{
			0: uint32(217),
		},
	},
	526: {
		Fname: __ccgo_ts + 7865,
		Fcodepoints: [2]uint32{
			0: uint32(362),
		},
	},
	527: {
		Fname: __ccgo_ts + 7873,
		Fcodepoints: [2]uint32{
			0: uint32(95),
		},
	},
	528: {
		Fname: __ccgo_ts + 7884,
		Fcodepoints: [2]uint32{
			0: uint32(9183),
		},
	},
	529: {
		Fname: __ccgo_ts + 7897,
		Fcodepoints: [2]uint32{
			0: uint32(9141),
		},
	},
	530: {
		Fname: __ccgo_ts + 7912,
		Fcodepoints: [2]uint32{
			0: uint32(9181),
		},
	},
	531: {
		Fname: __ccgo_ts + 7931,
		Fcodepoints: [2]uint32{
			0: uint32(8899),
		},
	},
	532: {
		Fname: __ccgo_ts + 7939,
		Fcodepoints: [2]uint32{
			0: uint32(8846),
		},
	},
	533: {
		Fname: __ccgo_ts + 7951,
		Fcodepoints: [2]uint32{
			0: uint32(370),
		},
	},
	534: {
		Fname: __ccgo_ts + 7959,
		Fcodepoints: [2]uint32{
			0: uint32(120140),
		},
	},
	535: {
		Fname: __ccgo_ts + 7966,
		Fcodepoints: [2]uint32{
			0: uint32(8593),
		},
	},
	536: {
		Fname: __ccgo_ts + 7976,
		Fcodepoints: [2]uint32{
			0: uint32(10514),
		},
	},
	537: {
		Fname: __ccgo_ts + 7989,
		Fcodepoints: [2]uint32{
			0: uint32(8645),
		},
	},
	538: {
		Fname: __ccgo_ts + 8008,
		Fcodepoints: [2]uint32{
			0: uint32(8597),
		},
	},
	539: {
		Fname: __ccgo_ts + 8022,
		Fcodepoints: [2]uint32{
			0: uint32(10606),
		},
	},
	540: {
		Fname: __ccgo_ts + 8038,
		Fcodepoints: [2]uint32{
			0: uint32(8869),
		},
	},
	541: {
		Fname: __ccgo_ts + 8046,
		Fcodepoints: [2]uint32{
			0: uint32(8613),
		},
	},
	542: {
		Fname: __ccgo_ts + 8059,
		Fcodepoints: [2]uint32{
			0: uint32(8657),
		},
	},
	543: {
		Fname: __ccgo_ts + 8069,
		Fcodepoints: [2]uint32{
			0: uint32(8661),
		},
	},
	544: {
		Fname: __ccgo_ts + 8083,
		Fcodepoints: [2]uint32{
			0: uint32(8598),
		},
	},
	545: {
		Fname: __ccgo_ts + 8100,
		Fcodepoints: [2]uint32{
			0: uint32(8599),
		},
	},
	546: {
		Fname: __ccgo_ts + 8118,
		Fcodepoints: [2]uint32{
			0: uint32(978),
		},
	},
	547: {
		Fname: __ccgo_ts + 8125,
		Fcodepoints: [2]uint32{
			0: uint32(933),
		},
	},
	548: {
		Fname: __ccgo_ts + 8135,
		Fcodepoints: [2]uint32{
			0: uint32(366),
		},
	},
	549: {
		Fname: __ccgo_ts + 8143,
		Fcodepoints: [2]uint32{
			0: uint32(119984),
		},
	},
	550: {
		Fname: __ccgo_ts + 8150,
		Fcodepoints: [2]uint32{
			0: uint32(360),
		},
	},
	551: {
		Fname: __ccgo_ts + 8159,
		Fcodepoints: [2]uint32{
			0: uint32(220),
		},
	},
	552: {
		Fname: __ccgo_ts + 8166,
		Fcodepoints: [2]uint32{
			0: uint32(8875),
		},
	},
	553: {
		Fname: __ccgo_ts + 8174,
		Fcodepoints: [2]uint32{
			0: uint32(10987),
		},
	},
	554: {
		Fname: __ccgo_ts + 8181,
		Fcodepoints: [2]uint32{
			0: uint32(1042),
		},
	},
	555: {
		Fname: __ccgo_ts + 8187,
		Fcodepoints: [2]uint32{
			0: uint32(8873),
		},
	},
	556: {
		Fname: __ccgo_ts + 8195,
		Fcodepoints: [2]uint32{
			0: uint32(10982),
		},
	},
	557: {
		Fname: __ccgo_ts + 8204,
		Fcodepoints: [2]uint32{
			0: uint32(8897),
		},
	},
	558: {
		Fname: __ccgo_ts + 8210,
		Fcodepoints: [2]uint32{
			0: uint32(8214),
		},
	},
	559: {
		Fname: __ccgo_ts + 8219,
		Fcodepoints: [2]uint32{
			0: uint32(8214),
		},
	},
	560: {
		Fname: __ccgo_ts + 8226,
		Fcodepoints: [2]uint32{
			0: uint32(8739),
		},
	},
	561: {
		Fname: __ccgo_ts + 8240,
		Fcodepoints: [2]uint32{
			0: uint32(124),
		},
	},
	562: {
		Fname: __ccgo_ts + 8255,
		Fcodepoints: [2]uint32{
			0: uint32(10072),
		},
	},
	563: {
		Fname: __ccgo_ts + 8275,
		Fcodepoints: [2]uint32{
			0: uint32(8768),
		},
	},
	564: {
		Fname: __ccgo_ts + 8291,
		Fcodepoints: [2]uint32{
			0: uint32(8202),
		},
	},
	565: {
		Fname: __ccgo_ts + 8307,
		Fcodepoints: [2]uint32{
			0: uint32(120089),
		},
	},
	566: {
		Fname: __ccgo_ts + 8313,
		Fcodepoints: [2]uint32{
			0: uint32(120141),
		},
	},
	567: {
		Fname: __ccgo_ts + 8320,
		Fcodepoints: [2]uint32{
			0: uint32(119985),
		},
	},
	568: {
		Fname: __ccgo_ts + 8327,
		Fcodepoints: [2]uint32{
			0: uint32(8874),
		},
	},
	569: {
		Fname: __ccgo_ts + 8336,
		Fcodepoints: [2]uint32{
			0: uint32(372),
		},
	},
	570: {
		Fname: __ccgo_ts + 8344,
		Fcodepoints: [2]uint32{
			0: uint32(8896),
		},
	},
	571: {
		Fname: __ccgo_ts + 8352,
		Fcodepoints: [2]uint32{
			0: uint32(120090),
		},
	},
	572: {
		Fname: __ccgo_ts + 8358,
		Fcodepoints: [2]uint32{
			0: uint32(120142),
		},
	},
	573: {
		Fname: __ccgo_ts + 8365,
		Fcodepoints: [2]uint32{
			0: uint32(119986),
		},
	},
	574: {
		Fname: __ccgo_ts + 8372,
		Fcodepoints: [2]uint32{
			0: uint32(120091),
		},
	},
	575: {
		Fname: __ccgo_ts + 8378,
		Fcodepoints: [2]uint32{
			0: uint32(926),
		},
	},
	576: {
		Fname: __ccgo_ts + 8383,
		Fcodepoints: [2]uint32{
			0: uint32(120143),
		},
	},
	577: {
		Fname: __ccgo_ts + 8390,
		Fcodepoints: [2]uint32{
			0: uint32(119987),
		},
	},
	578: {
		Fname: __ccgo_ts + 8397,
		Fcodepoints: [2]uint32{
			0: uint32(1071),
		},
	},
	579: {
		Fname: __ccgo_ts + 8404,
		Fcodepoints: [2]uint32{
			0: uint32(1031),
		},
	},
	580: {
		Fname: __ccgo_ts + 8411,
		Fcodepoints: [2]uint32{
			0: uint32(1070),
		},
	},
	581: {
		Fname: __ccgo_ts + 8418,
		Fcodepoints: [2]uint32{
			0: uint32(221),
		},
	},
	582: {
		Fname: __ccgo_ts + 8427,
		Fcodepoints: [2]uint32{
			0: uint32(374),
		},
	},
	583: {
		Fname: __ccgo_ts + 8435,
		Fcodepoints: [2]uint32{
			0: uint32(1067),
		},
	},
	584: {
		Fname: __ccgo_ts + 8441,
		Fcodepoints: [2]uint32{
			0: uint32(120092),
		},
	},
	585: {
		Fname: __ccgo_ts + 8447,
		Fcodepoints: [2]uint32{
			0: uint32(120144),
		},
	},
	586: {
		Fname: __ccgo_ts + 8454,
		Fcodepoints: [2]uint32{
			0: uint32(119988),
		},
	},
	587: {
		Fname: __ccgo_ts + 8461,
		Fcodepoints: [2]uint32{
			0: uint32(376),
		},
	},
	588: {
		Fname: __ccgo_ts + 8468,
		Fcodepoints: [2]uint32{
			0: uint32(1046),
		},
	},
	589: {
		Fname: __ccgo_ts + 8475,
		Fcodepoints: [2]uint32{
			0: uint32(377),
		},
	},
	590: {
		Fname: __ccgo_ts + 8484,
		Fcodepoints: [2]uint32{
			0: uint32(381),
		},
	},
	591: {
		Fname: __ccgo_ts + 8493,
		Fcodepoints: [2]uint32{
			0: uint32(1047),
		},
	},
	592: {
		Fname: __ccgo_ts + 8499,
		Fcodepoints: [2]uint32{
			0: uint32(379),
		},
	},
	593: {
		Fname: __ccgo_ts + 8506,
		Fcodepoints: [2]uint32{
			0: uint32(8203),
		},
	},
	594: {
		Fname: __ccgo_ts + 8523,
		Fcodepoints: [2]uint32{
			0: uint32(918),
		},
	},
	595: {
		Fname: __ccgo_ts + 8530,
		Fcodepoints: [2]uint32{
			0: uint32(8488),
		},
	},
	596: {
		Fname: __ccgo_ts + 8536,
		Fcodepoints: [2]uint32{
			0: uint32(8484),
		},
	},
	597: {
		Fname: __ccgo_ts + 8543,
		Fcodepoints: [2]uint32{
			0: uint32(119989),
		},
	},
	598: {
		Fname: __ccgo_ts + 8550,
		Fcodepoints: [2]uint32{
			0: uint32(225),
		},
	},
	599: {
		Fname: __ccgo_ts + 8559,
		Fcodepoints: [2]uint32{
			0: uint32(259),
		},
	},
	600: {
		Fname: __ccgo_ts + 8568,
		Fcodepoints: [2]uint32{
			0: uint32(8766),
		},
	},
	601: {
		Fname: __ccgo_ts + 8573,
		Fcodepoints: [2]uint32{
			0: uint32(8766),
			1: uint32(819),
		},
	},
	602: {
		Fname: __ccgo_ts + 8579,
		Fcodepoints: [2]uint32{
			0: uint32(8767),
		},
	},
	603: {
		Fname: __ccgo_ts + 8585,
		Fcodepoints: [2]uint32{
			0: uint32(226),
		},
	},
	604: {
		Fname: __ccgo_ts + 8593,
		Fcodepoints: [2]uint32{
			0: uint32(180),
		},
	},
	605: {
		Fname: __ccgo_ts + 8601,
		Fcodepoints: [2]uint32{
			0: uint32(1072),
		},
	},
	606: {
		Fname: __ccgo_ts + 8607,
		Fcodepoints: [2]uint32{
			0: uint32(230),
		},
	},
	607: {
		Fname: __ccgo_ts + 8615,
		Fcodepoints: [2]uint32{
			0: uint32(8289),
		},
	},
	608: {
		Fname: __ccgo_ts + 8620,
		Fcodepoints: [2]uint32{
			0: uint32(120094),
		},
	},
	609: {
		Fname: __ccgo_ts + 8626,
		Fcodepoints: [2]uint32{
			0: uint32(224),
		},
	},
	610: {
		Fname: __ccgo_ts + 8635,
		Fcodepoints: [2]uint32{
			0: uint32(8501),
		},
	},
	611: {
		Fname: __ccgo_ts + 8645,
		Fcodepoints: [2]uint32{
			0: uint32(8501),
		},
	},
	612: {
		Fname: __ccgo_ts + 8653,
		Fcodepoints: [2]uint32{
			0: uint32(945),
		},
	},
	613: {
		Fname: __ccgo_ts + 8661,
		Fcodepoints: [2]uint32{
			0: uint32(257),
		},
	},
	614: {
		Fname: __ccgo_ts + 8669,
		Fcodepoints: [2]uint32{
			0: uint32(10815),
		},
	},
	615: {
		Fname: __ccgo_ts + 837,
		Fcodepoints: [2]uint32{
			0: uint32(38),
		},
	},
	616: {
		Fname: __ccgo_ts + 8677,
		Fcodepoints: [2]uint32{
			0: uint32(8743),
		},
	},
	617: {
		Fname: __ccgo_ts + 8683,
		Fcodepoints: [2]uint32{
			0: uint32(10837),
		},
	},
	618: {
		Fname: __ccgo_ts + 8692,
		Fcodepoints: [2]uint32{
			0: uint32(10844),
		},
	},
	619: {
		Fname: __ccgo_ts + 8699,
		Fcodepoints: [2]uint32{
			0: uint32(10840),
		},
	},
	620: {
		Fname: __ccgo_ts + 8710,
		Fcodepoints: [2]uint32{
			0: uint32(10842),
		},
	},
	621: {
		Fname: __ccgo_ts + 8717,
		Fcodepoints: [2]uint32{
			0: uint32(8736),
		},
	},
	622: {
		Fname: __ccgo_ts + 8723,
		Fcodepoints: [2]uint32{
			0: uint32(10660),
		},
	},
	623: {
		Fname: __ccgo_ts + 8730,
		Fcodepoints: [2]uint32{
			0: uint32(8736),
		},
	},
	624: {
		Fname: __ccgo_ts + 8738,
		Fcodepoints: [2]uint32{
			0: uint32(8737),
		},
	},
	625: {
		Fname: __ccgo_ts + 8747,
		Fcodepoints: [2]uint32{
			0: uint32(10664),
		},
	},
	626: {
		Fname: __ccgo_ts + 8758,
		Fcodepoints: [2]uint32{
			0: uint32(10665),
		},
	},
	627: {
		Fname: __ccgo_ts + 8769,
		Fcodepoints: [2]uint32{
			0: uint32(10666),
		},
	},
	628: {
		Fname: __ccgo_ts + 8780,
		Fcodepoints: [2]uint32{
			0: uint32(10667),
		},
	},
	629: {
		Fname: __ccgo_ts + 8791,
		Fcodepoints: [2]uint32{
			0: uint32(10668),
		},
	},
	630: {
		Fname: __ccgo_ts + 8802,
		Fcodepoints: [2]uint32{
			0: uint32(10669),
		},
	},
	631: {
		Fname: __ccgo_ts + 8813,
		Fcodepoints: [2]uint32{
			0: uint32(10670),
		},
	},
	632: {
		Fname: __ccgo_ts + 8824,
		Fcodepoints: [2]uint32{
			0: uint32(10671),
		},
	},
	633: {
		Fname: __ccgo_ts + 8835,
		Fcodepoints: [2]uint32{
			0: uint32(8735),
		},
	},
	634: {
		Fname: __ccgo_ts + 8843,
		Fcodepoints: [2]uint32{
			0: uint32(8894),
		},
	},
	635: {
		Fname: __ccgo_ts + 8853,
		Fcodepoints: [2]uint32{
			0: uint32(10653),
		},
	},
	636: {
		Fname: __ccgo_ts + 8864,
		Fcodepoints: [2]uint32{
			0: uint32(8738),
		},
	},
	637: {
		Fname: __ccgo_ts + 8873,
		Fcodepoints: [2]uint32{
			0: uint32(197),
		},
	},
	638: {
		Fname: __ccgo_ts + 8881,
		Fcodepoints: [2]uint32{
			0: uint32(9084),
		},
	},
	639: {
		Fname: __ccgo_ts + 8891,
		Fcodepoints: [2]uint32{
			0: uint32(261),
		},
	},
	640: {
		Fname: __ccgo_ts + 8899,
		Fcodepoints: [2]uint32{
			0: uint32(120146),
		},
	},
	641: {
		Fname: __ccgo_ts + 8906,
		Fcodepoints: [2]uint32{
			0: uint32(8776),
		},
	},
	642: {
		Fname: __ccgo_ts + 8911,
		Fcodepoints: [2]uint32{
			0: uint32(10864),
		},
	},
	643: {
		Fname: __ccgo_ts + 8917,
		Fcodepoints: [2]uint32{
			0: uint32(10863),
		},
	},
	644: {
		Fname: __ccgo_ts + 8926,
		Fcodepoints: [2]uint32{
			0: uint32(8778),
		},
	},
	645: {
		Fname: __ccgo_ts + 8932,
		Fcodepoints: [2]uint32{
			0: uint32(8779),
		},
	},
	646: {
		Fname: __ccgo_ts + 8939,
		Fcodepoints: [2]uint32{
			0: uint32(39),
		},
	},
	647: {
		Fname: __ccgo_ts + 8946,
		Fcodepoints: [2]uint32{
			0: uint32(8776),
		},
	},
	648: {
		Fname: __ccgo_ts + 8955,
		Fcodepoints: [2]uint32{
			0: uint32(8778),
		},
	},
	649: {
		Fname: __ccgo_ts + 8966,
		Fcodepoints: [2]uint32{
			0: uint32(229),
		},
	},
	650: {
		Fname: __ccgo_ts + 8974,
		Fcodepoints: [2]uint32{
			0: uint32(119990),
		},
	},
	651: {
		Fname: __ccgo_ts + 8981,
		Fcodepoints: [2]uint32{
			0: uint32(42),
		},
	},
	652: {
		Fname: __ccgo_ts + 8987,
		Fcodepoints: [2]uint32{
			0: uint32(8776),
		},
	},
	653: {
		Fname: __ccgo_ts + 8995,
		Fcodepoints: [2]uint32{
			0: uint32(8781),
		},
	},
	654: {
		Fname: __ccgo_ts + 9005,
		Fcodepoints: [2]uint32{
			0: uint32(227),
		},
	},
	655: {
		Fname: __ccgo_ts + 9014,
		Fcodepoints: [2]uint32{
			0: uint32(228),
		},
	},
	656: {
		Fname: __ccgo_ts + 9021,
		Fcodepoints: [2]uint32{
			0: uint32(8755),
		},
	},
	657: {
		Fname: __ccgo_ts + 9032,
		Fcodepoints: [2]uint32{
			0: uint32(10769),
		},
	},
	658: {
		Fname: __ccgo_ts + 9040,
		Fcodepoints: [2]uint32{
			0: uint32(10989),
		},
	},
	659: {
		Fname: __ccgo_ts + 9047,
		Fcodepoints: [2]uint32{
			0: uint32(8780),
		},
	},
	660: {
		Fname: __ccgo_ts + 9058,
		Fcodepoints: [2]uint32{
			0: uint32(1014),
		},
	},
	661: {
		Fname: __ccgo_ts + 9072,
		Fcodepoints: [2]uint32{
			0: uint32(8245),
		},
	},
	662: {
		Fname: __ccgo_ts + 9084,
		Fcodepoints: [2]uint32{
			0: uint32(8765),
		},
	},
	663: {
		Fname: __ccgo_ts + 9094,
		Fcodepoints: [2]uint32{
			0: uint32(8909),
		},
	},
	664: {
		Fname: __ccgo_ts + 9106,
		Fcodepoints: [2]uint32{
			0: uint32(8893),
		},
	},
	665: {
		Fname: __ccgo_ts + 9115,
		Fcodepoints: [2]uint32{
			0: uint32(8965),
		},
	},
	666: {
		Fname: __ccgo_ts + 9124,
		Fcodepoints: [2]uint32{
			0: uint32(8965),
		},
	},
	667: {
		Fname: __ccgo_ts + 9135,
		Fcodepoints: [2]uint32{
			0: uint32(9141),
		},
	},
	668: {
		Fname: __ccgo_ts + 9142,
		Fcodepoints: [2]uint32{
			0: uint32(9142),
		},
	},
	669: {
		Fname: __ccgo_ts + 9153,
		Fcodepoints: [2]uint32{
			0: uint32(8780),
		},
	},
	670: {
		Fname: __ccgo_ts + 9161,
		Fcodepoints: [2]uint32{
			0: uint32(1073),
		},
	},
	671: {
		Fname: __ccgo_ts + 9167,
		Fcodepoints: [2]uint32{
			0: uint32(8222),
		},
	},
	672: {
		Fname: __ccgo_ts + 9175,
		Fcodepoints: [2]uint32{
			0: uint32(8757),
		},
	},
	673: {
		Fname: __ccgo_ts + 9184,
		Fcodepoints: [2]uint32{
			0: uint32(8757),
		},
	},
	674: {
		Fname: __ccgo_ts + 9194,
		Fcodepoints: [2]uint32{
			0: uint32(10672),
		},
	},
	675: {
		Fname: __ccgo_ts + 9204,
		Fcodepoints: [2]uint32{
			0: uint32(1014),
		},
	},
	676: {
		Fname: __ccgo_ts + 9212,
		Fcodepoints: [2]uint32{
			0: uint32(8492),
		},
	},
	677: {
		Fname: __ccgo_ts + 9221,
		Fcodepoints: [2]uint32{
			0: uint32(946),
		},
	},
	678: {
		Fname: __ccgo_ts + 9228,
		Fcodepoints: [2]uint32{
			0: uint32(8502),
		},
	},
	679: {
		Fname: __ccgo_ts + 9235,
		Fcodepoints: [2]uint32{
			0: uint32(8812),
		},
	},
	680: {
		Fname: __ccgo_ts + 9245,
		Fcodepoints: [2]uint32{
			0: uint32(120095),
		},
	},
	681: {
		Fname: __ccgo_ts + 9251,
		Fcodepoints: [2]uint32{
			0: uint32(8898),
		},
	},
	682: {
		Fname: __ccgo_ts + 9260,
		Fcodepoints: [2]uint32{
			0: uint32(9711),
		},
	},
	683: {
		Fname: __ccgo_ts + 9270,
		Fcodepoints: [2]uint32{
			0: uint32(8899),
		},
	},
	684: {
		Fname: __ccgo_ts + 9279,
		Fcodepoints: [2]uint32{
			0: uint32(10752),
		},
	},
	685: {
		Fname: __ccgo_ts + 9289,
		Fcodepoints: [2]uint32{
			0: uint32(10753),
		},
	},
	686: {
		Fname: __ccgo_ts + 9300,
		Fcodepoints: [2]uint32{
			0: uint32(10754),
		},
	},
	687: {
		Fname: __ccgo_ts + 9312,
		Fcodepoints: [2]uint32{
			0: uint32(10758),
		},
	},
	688: {
		Fname: __ccgo_ts + 9323,
		Fcodepoints: [2]uint32{
			0: uint32(9733),
		},
	},
	689: {
		Fname: __ccgo_ts + 9333,
		Fcodepoints: [2]uint32{
			0: uint32(9661),
		},
	},
	690: {
		Fname: __ccgo_ts + 9351,
		Fcodepoints: [2]uint32{
			0: uint32(9651),
		},
	},
	691: {
		Fname: __ccgo_ts + 9367,
		Fcodepoints: [2]uint32{
			0: uint32(10756),
		},
	},
	692: {
		Fname: __ccgo_ts + 9378,
		Fcodepoints: [2]uint32{
			0: uint32(8897),
		},
	},
	693: {
		Fname: __ccgo_ts + 9387,
		Fcodepoints: [2]uint32{
			0: uint32(8896),
		},
	},
	694: {
		Fname: __ccgo_ts + 9398,
		Fcodepoints: [2]uint32{
			0: uint32(10509),
		},
	},
	695: {
		Fname: __ccgo_ts + 9407,
		Fcodepoints: [2]uint32{
			0: uint32(10731),
		},
	},
	696: {
		Fname: __ccgo_ts + 9422,
		Fcodepoints: [2]uint32{
			0: uint32(9642),
		},
	},
	697: {
		Fname: __ccgo_ts + 9436,
		Fcodepoints: [2]uint32{
			0: uint32(9652),
		},
	},
	698: {
		Fname: __ccgo_ts + 9452,
		Fcodepoints: [2]uint32{
			0: uint32(9662),
		},
	},
	699: {
		Fname: __ccgo_ts + 9472,
		Fcodepoints: [2]uint32{
			0: uint32(9666),
		},
	},
	700: {
		Fname: __ccgo_ts + 9492,
		Fcodepoints: [2]uint32{
			0: uint32(9656),
		},
	},
	701: {
		Fname: __ccgo_ts + 9513,
		Fcodepoints: [2]uint32{
			0: uint32(9251),
		},
	},
	702: {
		Fname: __ccgo_ts + 9521,
		Fcodepoints: [2]uint32{
			0: uint32(9618),
		},
	},
	703: {
		Fname: __ccgo_ts + 9529,
		Fcodepoints: [2]uint32{
			0: uint32(9617),
		},
	},
	704: {
		Fname: __ccgo_ts + 9537,
		Fcodepoints: [2]uint32{
			0: uint32(9619),
		},
	},
	705: {
		Fname: __ccgo_ts + 9545,
		Fcodepoints: [2]uint32{
			0: uint32(9608),
		},
	},
	706: {
		Fname: __ccgo_ts + 9553,
		Fcodepoints: [2]uint32{
			0: uint32(61),
			1: uint32(8421),
		},
	},
	707: {
		Fname: __ccgo_ts + 9559,
		Fcodepoints: [2]uint32{
			0: uint32(8801),
			1: uint32(8421),
		},
	},
	708: {
		Fname: __ccgo_ts + 9569,
		Fcodepoints: [2]uint32{
			0: uint32(8976),
		},
	},
	709: {
		Fname: __ccgo_ts + 9576,
		Fcodepoints: [2]uint32{
			0: uint32(120147),
		},
	},
	710: {
		Fname: __ccgo_ts + 9583,
		Fcodepoints: [2]uint32{
			0: uint32(8869),
		},
	},
	711: {
		Fname: __ccgo_ts + 9589,
		Fcodepoints: [2]uint32{
			0: uint32(8869),
		},
	},
	712: {
		Fname: __ccgo_ts + 9598,
		Fcodepoints: [2]uint32{
			0: uint32(8904),
		},
	},
	713: {
		Fname: __ccgo_ts + 9607,
		Fcodepoints: [2]uint32{
			0: uint32(9559),
		},
	},
	714: {
		Fname: __ccgo_ts + 9615,
		Fcodepoints: [2]uint32{
			0: uint32(9556),
		},
	},
	715: {
		Fname: __ccgo_ts + 9623,
		Fcodepoints: [2]uint32{
			0: uint32(9558),
		},
	},
	716: {
		Fname: __ccgo_ts + 9631,
		Fcodepoints: [2]uint32{
			0: uint32(9555),
		},
	},
	717: {
		Fname: __ccgo_ts + 9639,
		Fcodepoints: [2]uint32{
			0: uint32(9552),
		},
	},
	718: {
		Fname: __ccgo_ts + 9646,
		Fcodepoints: [2]uint32{
			0: uint32(9574),
		},
	},
	719: {
		Fname: __ccgo_ts + 9654,
		Fcodepoints: [2]uint32{
			0: uint32(9577),
		},
	},
	720: {
		Fname: __ccgo_ts + 9662,
		Fcodepoints: [2]uint32{
			0: uint32(9572),
		},
	},
	721: {
		Fname: __ccgo_ts + 9670,
		Fcodepoints: [2]uint32{
			0: uint32(9575),
		},
	},
	722: {
		Fname: __ccgo_ts + 9678,
		Fcodepoints: [2]uint32{
			0: uint32(9565),
		},
	},
	723: {
		Fname: __ccgo_ts + 9686,
		Fcodepoints: [2]uint32{
			0: uint32(9562),
		},
	},
	724: {
		Fname: __ccgo_ts + 9694,
		Fcodepoints: [2]uint32{
			0: uint32(9564),
		},
	},
	725: {
		Fname: __ccgo_ts + 9702,
		Fcodepoints: [2]uint32{
			0: uint32(9561),
		},
	},
	726: {
		Fname: __ccgo_ts + 9710,
		Fcodepoints: [2]uint32{
			0: uint32(9553),
		},
	},
	727: {
		Fname: __ccgo_ts + 9717,
		Fcodepoints: [2]uint32{
			0: uint32(9580),
		},
	},
	728: {
		Fname: __ccgo_ts + 9725,
		Fcodepoints: [2]uint32{
			0: uint32(9571),
		},
	},
	729: {
		Fname: __ccgo_ts + 9733,
		Fcodepoints: [2]uint32{
			0: uint32(9568),
		},
	},
	730: {
		Fname: __ccgo_ts + 9741,
		Fcodepoints: [2]uint32{
			0: uint32(9579),
		},
	},
	731: {
		Fname: __ccgo_ts + 9749,
		Fcodepoints: [2]uint32{
			0: uint32(9570),
		},
	},
	732: {
		Fname: __ccgo_ts + 9757,
		Fcodepoints: [2]uint32{
			0: uint32(9567),
		},
	},
	733: {
		Fname: __ccgo_ts + 9765,
		Fcodepoints: [2]uint32{
			0: uint32(10697),
		},
	},
	734: {
		Fname: __ccgo_ts + 9774,
		Fcodepoints: [2]uint32{
			0: uint32(9557),
		},
	},
	735: {
		Fname: __ccgo_ts + 9782,
		Fcodepoints: [2]uint32{
			0: uint32(9554),
		},
	},
	736: {
		Fname: __ccgo_ts + 9790,
		Fcodepoints: [2]uint32{
			0: uint32(9488),
		},
	},
	737: {
		Fname: __ccgo_ts + 9798,
		Fcodepoints: [2]uint32{
			0: uint32(9484),
		},
	},
	738: {
		Fname: __ccgo_ts + 9806,
		Fcodepoints: [2]uint32{
			0: uint32(9472),
		},
	},
	739: {
		Fname: __ccgo_ts + 9813,
		Fcodepoints: [2]uint32{
			0: uint32(9573),
		},
	},
	740: {
		Fname: __ccgo_ts + 9821,
		Fcodepoints: [2]uint32{
			0: uint32(9576),
		},
	},
	741: {
		Fname: __ccgo_ts + 9829,
		Fcodepoints: [2]uint32{
			0: uint32(9516),
		},
	},
	742: {
		Fname: __ccgo_ts + 9837,
		Fcodepoints: [2]uint32{
			0: uint32(9524),
		},
	},
	743: {
		Fname: __ccgo_ts + 9845,
		Fcodepoints: [2]uint32{
			0: uint32(8863),
		},
	},
	744: {
		Fname: __ccgo_ts + 9856,
		Fcodepoints: [2]uint32{
			0: uint32(8862),
		},
	},
	745: {
		Fname: __ccgo_ts + 9866,
		Fcodepoints: [2]uint32{
			0: uint32(8864),
		},
	},
	746: {
		Fname: __ccgo_ts + 9877,
		Fcodepoints: [2]uint32{
			0: uint32(9563),
		},
	},
	747: {
		Fname: __ccgo_ts + 9885,
		Fcodepoints: [2]uint32{
			0: uint32(9560),
		},
	},
	748: {
		Fname: __ccgo_ts + 9893,
		Fcodepoints: [2]uint32{
			0: uint32(9496),
		},
	},
	749: {
		Fname: __ccgo_ts + 9901,
		Fcodepoints: [2]uint32{
			0: uint32(9492),
		},
	},
	750: {
		Fname: __ccgo_ts + 9909,
		Fcodepoints: [2]uint32{
			0: uint32(9474),
		},
	},
	751: {
		Fname: __ccgo_ts + 9916,
		Fcodepoints: [2]uint32{
			0: uint32(9578),
		},
	},
	752: {
		Fname: __ccgo_ts + 9924,
		Fcodepoints: [2]uint32{
			0: uint32(9569),
		},
	},
	753: {
		Fname: __ccgo_ts + 9932,
		Fcodepoints: [2]uint32{
			0: uint32(9566),
		},
	},
	754: {
		Fname: __ccgo_ts + 9940,
		Fcodepoints: [2]uint32{
			0: uint32(9532),
		},
	},
	755: {
		Fname: __ccgo_ts + 9948,
		Fcodepoints: [2]uint32{
			0: uint32(9508),
		},
	},
	756: {
		Fname: __ccgo_ts + 9956,
		Fcodepoints: [2]uint32{
			0: uint32(9500),
		},
	},
	757: {
		Fname: __ccgo_ts + 9964,
		Fcodepoints: [2]uint32{
			0: uint32(8245),
		},
	},
	758: {
		Fname: __ccgo_ts + 9973,
		Fcodepoints: [2]uint32{
			0: uint32(728),
		},
	},
	759: {
		Fname: __ccgo_ts + 9981,
		Fcodepoints: [2]uint32{
			0: uint32(166),
		},
	},
	760: {
		Fname: __ccgo_ts + 9990,
		Fcodepoints: [2]uint32{
			0: uint32(119991),
		},
	},
	761: {
		Fname: __ccgo_ts + 9997,
		Fcodepoints: [2]uint32{
			0: uint32(8271),
		},
	},
	762: {
		Fname: __ccgo_ts + 10005,
		Fcodepoints: [2]uint32{
			0: uint32(8765),
		},
	},
	763: {
		Fname: __ccgo_ts + 10012,
		Fcodepoints: [2]uint32{
			0: uint32(8909),
		},
	},
	764: {
		Fname: __ccgo_ts + 10020,
		Fcodepoints: [2]uint32{
			0: uint32(92),
		},
	},
	765: {
		Fname: __ccgo_ts + 10027,
		Fcodepoints: [2]uint32{
			0: uint32(10693),
		},
	},
	766: {
		Fname: __ccgo_ts + 10035,
		Fcodepoints: [2]uint32{
			0: uint32(10184),
		},
	},
	767: {
		Fname: __ccgo_ts + 10046,
		Fcodepoints: [2]uint32{
			0: uint32(8226),
		},
	},
	768: {
		Fname: __ccgo_ts + 10053,
		Fcodepoints: [2]uint32{
			0: uint32(8226),
		},
	},
	769: {
		Fname: __ccgo_ts + 10062,
		Fcodepoints: [2]uint32{
			0: uint32(8782),
		},
	},
	770: {
		Fname: __ccgo_ts + 10069,
		Fcodepoints: [2]uint32{
			0: uint32(10926),
		},
	},
	771: {
		Fname: __ccgo_ts + 10077,
		Fcodepoints: [2]uint32{
			0: uint32(8783),
		},
	},
	772: {
		Fname: __ccgo_ts + 10085,
		Fcodepoints: [2]uint32{
			0: uint32(8783),
		},
	},
	773: {
		Fname: __ccgo_ts + 10094,
		Fcodepoints: [2]uint32{
			0: uint32(263),
		},
	},
	774: {
		Fname: __ccgo_ts + 10103,
		Fcodepoints: [2]uint32{
			0: uint32(8745),
		},
	},
	775: {
		Fname: __ccgo_ts + 10109,
		Fcodepoints: [2]uint32{
			0: uint32(10820),
		},
	},
	776: {
		Fname: __ccgo_ts + 10118,
		Fcodepoints: [2]uint32{
			0: uint32(10825),
		},
	},
	777: {
		Fname: __ccgo_ts + 10129,
		Fcodepoints: [2]uint32{
			0: uint32(10827),
		},
	},
	778: {
		Fname: __ccgo_ts + 10138,
		Fcodepoints: [2]uint32{
			0: uint32(10823),
		},
	},
	779: {
		Fname: __ccgo_ts + 10147,
		Fcodepoints: [2]uint32{
			0: uint32(10816),
		},
	},
	780: {
		Fname: __ccgo_ts + 10156,
		Fcodepoints: [2]uint32{
			0: uint32(8745),
			1: uint32(65024),
		},
	},
	781: {
		Fname: __ccgo_ts + 10163,
		Fcodepoints: [2]uint32{
			0: uint32(8257),
		},
	},
	782: {
		Fname: __ccgo_ts + 10171,
		Fcodepoints: [2]uint32{
			0: uint32(711),
		},
	},
	783: {
		Fname: __ccgo_ts + 10179,
		Fcodepoints: [2]uint32{
			0: uint32(10829),
		},
	},
	784: {
		Fname: __ccgo_ts + 10187,
		Fcodepoints: [2]uint32{
			0: uint32(269),
		},
	},
	785: {
		Fname: __ccgo_ts + 10196,
		Fcodepoints: [2]uint32{
			0: uint32(231),
		},
	},
	786: {
		Fname: __ccgo_ts + 10205,
		Fcodepoints: [2]uint32{
			0: uint32(265),
		},
	},
	787: {
		Fname: __ccgo_ts + 10213,
		Fcodepoints: [2]uint32{
			0: uint32(10828),
		},
	},
	788: {
		Fname: __ccgo_ts + 10221,
		Fcodepoints: [2]uint32{
			0: uint32(10832),
		},
	},
	789: {
		Fname: __ccgo_ts + 10231,
		Fcodepoints: [2]uint32{
			0: uint32(267),
		},
	},
	790: {
		Fname: __ccgo_ts + 10238,
		Fcodepoints: [2]uint32{
			0: uint32(184),
		},
	},
	791: {
		Fname: __ccgo_ts + 10246,
		Fcodepoints: [2]uint32{
			0: uint32(10674),
		},
	},
	792: {
		Fname: __ccgo_ts + 10256,
		Fcodepoints: [2]uint32{
			0: uint32(162),
		},
	},
	793: {
		Fname: __ccgo_ts + 10263,
		Fcodepoints: [2]uint32{
			0: uint32(183),
		},
	},
	794: {
		Fname: __ccgo_ts + 10275,
		Fcodepoints: [2]uint32{
			0: uint32(120096),
		},
	},
	795: {
		Fname: __ccgo_ts + 10281,
		Fcodepoints: [2]uint32{
			0: uint32(1095),
		},
	},
	796: {
		Fname: __ccgo_ts + 10288,
		Fcodepoints: [2]uint32{
			0: uint32(10003),
		},
	},
	797: {
		Fname: __ccgo_ts + 10296,
		Fcodepoints: [2]uint32{
			0: uint32(10003),
		},
	},
	798: {
		Fname: __ccgo_ts + 10308,
		Fcodepoints: [2]uint32{
			0: uint32(967),
		},
	},
	799: {
		Fname: __ccgo_ts + 10314,
		Fcodepoints: [2]uint32{
			0: uint32(9675),
		},
	},
	800: {
		Fname: __ccgo_ts + 10320,
		Fcodepoints: [2]uint32{
			0: uint32(10691),
		},
	},
	801: {
		Fname: __ccgo_ts + 10327,
		Fcodepoints: [2]uint32{
			0: uint32(710),
		},
	},
	802: {
		Fname: __ccgo_ts + 10334,
		Fcodepoints: [2]uint32{
			0: uint32(8791),
		},
	},
	803: {
		Fname: __ccgo_ts + 10343,
		Fcodepoints: [2]uint32{
			0: uint32(8634),
		},
	},
	804: {
		Fname: __ccgo_ts + 10361,
		Fcodepoints: [2]uint32{
			0: uint32(8635),
		},
	},
	805: {
		Fname: __ccgo_ts + 10380,
		Fcodepoints: [2]uint32{
			0: uint32(174),
		},
	},
	806: {
		Fname: __ccgo_ts + 10391,
		Fcodepoints: [2]uint32{
			0: uint32(9416),
		},
	},
	807: {
		Fname: __ccgo_ts + 10402,
		Fcodepoints: [2]uint32{
			0: uint32(8859),
		},
	},
	808: {
		Fname: __ccgo_ts + 10415,
		Fcodepoints: [2]uint32{
			0: uint32(8858),
		},
	},
	809: {
		Fname: __ccgo_ts + 10429,
		Fcodepoints: [2]uint32{
			0: uint32(8861),
		},
	},
	810: {
		Fname: __ccgo_ts + 10443,
		Fcodepoints: [2]uint32{
			0: uint32(8791),
		},
	},
	811: {
		Fname: __ccgo_ts + 10450,
		Fcodepoints: [2]uint32{
			0: uint32(10768),
		},
	},
	812: {
		Fname: __ccgo_ts + 10461,
		Fcodepoints: [2]uint32{
			0: uint32(10991),
		},
	},
	813: {
		Fname: __ccgo_ts + 10470,
		Fcodepoints: [2]uint32{
			0: uint32(10690),
		},
	},
	814: {
		Fname: __ccgo_ts + 10480,
		Fcodepoints: [2]uint32{
			0: uint32(9827),
		},
	},
	815: {
		Fname: __ccgo_ts + 10488,
		Fcodepoints: [2]uint32{
			0: uint32(9827),
		},
	},
	816: {
		Fname: __ccgo_ts + 10499,
		Fcodepoints: [2]uint32{
			0: uint32(58),
		},
	},
	817: {
		Fname: __ccgo_ts + 10507,
		Fcodepoints: [2]uint32{
			0: uint32(8788),
		},
	},
	818: {
		Fname: __ccgo_ts + 10516,
		Fcodepoints: [2]uint32{
			0: uint32(8788),
		},
	},
	819: {
		Fname: __ccgo_ts + 10526,
		Fcodepoints: [2]uint32{
			0: uint32(44),
		},
	},
	820: {
		Fname: __ccgo_ts + 10534,
		Fcodepoints: [2]uint32{
			0: uint32(64),
		},
	},
	821: {
		Fname: __ccgo_ts + 10543,
		Fcodepoints: [2]uint32{
			0: uint32(8705),
		},
	},
	822: {
		Fname: __ccgo_ts + 10550,
		Fcodepoints: [2]uint32{
			0: uint32(8728),
		},
	},
	823: {
		Fname: __ccgo_ts + 10559,
		Fcodepoints: [2]uint32{
			0: uint32(8705),
		},
	},
	824: {
		Fname: __ccgo_ts + 10572,
		Fcodepoints: [2]uint32{
			0: uint32(8450),
		},
	},
	825: {
		Fname: __ccgo_ts + 10584,
		Fcodepoints: [2]uint32{
			0: uint32(8773),
		},
	},
	826: {
		Fname: __ccgo_ts + 10591,
		Fcodepoints: [2]uint32{
			0: uint32(10861),
		},
	},
	827: {
		Fname: __ccgo_ts + 10601,
		Fcodepoints: [2]uint32{
			0: uint32(8750),
		},
	},
	828: {
		Fname: __ccgo_ts + 10610,
		Fcodepoints: [2]uint32{
			0: uint32(120148),
		},
	},
	829: {
		Fname: __ccgo_ts + 10617,
		Fcodepoints: [2]uint32{
			0: uint32(8720),
		},
	},
	830: {
		Fname: __ccgo_ts + 10626,
		Fcodepoints: [2]uint32{
			0: uint32(169),
		},
	},
	831: {
		Fname: __ccgo_ts + 10633,
		Fcodepoints: [2]uint32{
			0: uint32(8471),
		},
	},
	832: {
		Fname: __ccgo_ts + 10642,
		Fcodepoints: [2]uint32{
			0: uint32(8629),
		},
	},
	833: {
		Fname: __ccgo_ts + 10650,
		Fcodepoints: [2]uint32{
			0: uint32(10007),
		},
	},
	834: {
		Fname: __ccgo_ts + 10658,
		Fcodepoints: [2]uint32{
			0: uint32(119992),
		},
	},
	835: {
		Fname: __ccgo_ts + 10665,
		Fcodepoints: [2]uint32{
			0: uint32(10959),
		},
	},
	836: {
		Fname: __ccgo_ts + 10672,
		Fcodepoints: [2]uint32{
			0: uint32(10961),
		},
	},
	837: {
		Fname: __ccgo_ts + 10680,
		Fcodepoints: [2]uint32{
			0: uint32(10960),
		},
	},
	838: {
		Fname: __ccgo_ts + 10687,
		Fcodepoints: [2]uint32{
			0: uint32(10962),
		},
	},
	839: {
		Fname: __ccgo_ts + 10695,
		Fcodepoints: [2]uint32{
			0: uint32(8943),
		},
	},
	840: {
		Fname: __ccgo_ts + 10703,
		Fcodepoints: [2]uint32{
			0: uint32(10552),
		},
	},
	841: {
		Fname: __ccgo_ts + 10713,
		Fcodepoints: [2]uint32{
			0: uint32(10549),
		},
	},
	842: {
		Fname: __ccgo_ts + 10723,
		Fcodepoints: [2]uint32{
			0: uint32(8926),
		},
	},
	843: {
		Fname: __ccgo_ts + 10731,
		Fcodepoints: [2]uint32{
			0: uint32(8927),
		},
	},
	844: {
		Fname: __ccgo_ts + 10739,
		Fcodepoints: [2]uint32{
			0: uint32(8630),
		},
	},
	845: {
		Fname: __ccgo_ts + 10748,
		Fcodepoints: [2]uint32{
			0: uint32(10557),
		},
	},
	846: {
		Fname: __ccgo_ts + 10758,
		Fcodepoints: [2]uint32{
			0: uint32(8746),
		},
	},
	847: {
		Fname: __ccgo_ts + 10764,
		Fcodepoints: [2]uint32{
			0: uint32(10824),
		},
	},
	848: {
		Fname: __ccgo_ts + 10775,
		Fcodepoints: [2]uint32{
			0: uint32(10822),
		},
	},
	849: {
		Fname: __ccgo_ts + 10784,
		Fcodepoints: [2]uint32{
			0: uint32(10826),
		},
	},
	850: {
		Fname: __ccgo_ts + 10793,
		Fcodepoints: [2]uint32{
			0: uint32(8845),
		},
	},
	851: {
		Fname: __ccgo_ts + 10802,
		Fcodepoints: [2]uint32{
			0: uint32(10821),
		},
	},
	852: {
		Fname: __ccgo_ts + 10810,
		Fcodepoints: [2]uint32{
			0: uint32(8746),
			1: uint32(65024),
		},
	},
	853: {
		Fname: __ccgo_ts + 10817,
		Fcodepoints: [2]uint32{
			0: uint32(8631),
		},
	},
	854: {
		Fname: __ccgo_ts + 10826,
		Fcodepoints: [2]uint32{
			0: uint32(10556),
		},
	},
	855: {
		Fname: __ccgo_ts + 10836,
		Fcodepoints: [2]uint32{
			0: uint32(8926),
		},
	},
	856: {
		Fname: __ccgo_ts + 10850,
		Fcodepoints: [2]uint32{
			0: uint32(8927),
		},
	},
	857: {
		Fname: __ccgo_ts + 10864,
		Fcodepoints: [2]uint32{
			0: uint32(8910),
		},
	},
	858: {
		Fname: __ccgo_ts + 10875,
		Fcodepoints: [2]uint32{
			0: uint32(8911),
		},
	},
	859: {
		Fname: __ccgo_ts + 10888,
		Fcodepoints: [2]uint32{
			0: uint32(164),
		},
	},
	860: {
		Fname: __ccgo_ts + 10897,
		Fcodepoints: [2]uint32{
			0: uint32(8630),
		},
	},
	861: {
		Fname: __ccgo_ts + 10914,
		Fcodepoints: [2]uint32{
			0: uint32(8631),
		},
	},
	862: {
		Fname: __ccgo_ts + 10932,
		Fcodepoints: [2]uint32{
			0: uint32(8910),
		},
	},
	863: {
		Fname: __ccgo_ts + 10940,
		Fcodepoints: [2]uint32{
			0: uint32(8911),
		},
	},
	864: {
		Fname: __ccgo_ts + 10948,
		Fcodepoints: [2]uint32{
			0: uint32(8754),
		},
	},
	865: {
		Fname: __ccgo_ts + 10959,
		Fcodepoints: [2]uint32{
			0: uint32(8753),
		},
	},
	866: {
		Fname: __ccgo_ts + 10967,
		Fcodepoints: [2]uint32{
			0: uint32(9005),
		},
	},
	867: {
		Fname: __ccgo_ts + 10976,
		Fcodepoints: [2]uint32{
			0: uint32(8659),
		},
	},
	868: {
		Fname: __ccgo_ts + 10983,
		Fcodepoints: [2]uint32{
			0: uint32(10597),
		},
	},
	869: {
		Fname: __ccgo_ts + 10990,
		Fcodepoints: [2]uint32{
			0: uint32(8224),
		},
	},
	870: {
		Fname: __ccgo_ts + 10999,
		Fcodepoints: [2]uint32{
			0: uint32(8504),
		},
	},
	871: {
		Fname: __ccgo_ts + 11008,
		Fcodepoints: [2]uint32{
			0: uint32(8595),
		},
	},
	872: {
		Fname: __ccgo_ts + 11015,
		Fcodepoints: [2]uint32{
			0: uint32(8208),
		},
	},
	873: {
		Fname: __ccgo_ts + 11022,
		Fcodepoints: [2]uint32{
			0: uint32(8867),
		},
	},
	874: {
		Fname: __ccgo_ts + 11030,
		Fcodepoints: [2]uint32{
			0: uint32(10511),
		},
	},
	875: {
		Fname: __ccgo_ts + 11040,
		Fcodepoints: [2]uint32{
			0: uint32(733),
		},
	},
	876: {
		Fname: __ccgo_ts + 11048,
		Fcodepoints: [2]uint32{
			0: uint32(271),
		},
	},
	877: {
		Fname: __ccgo_ts + 11057,
		Fcodepoints: [2]uint32{
			0: uint32(1076),
		},
	},
	878: {
		Fname: __ccgo_ts + 11063,
		Fcodepoints: [2]uint32{
			0: uint32(8518),
		},
	},
	879: {
		Fname: __ccgo_ts + 11068,
		Fcodepoints: [2]uint32{
			0: uint32(8225),
		},
	},
	880: {
		Fname: __ccgo_ts + 11078,
		Fcodepoints: [2]uint32{
			0: uint32(8650),
		},
	},
	881: {
		Fname: __ccgo_ts + 11086,
		Fcodepoints: [2]uint32{
			0: uint32(10871),
		},
	},
	882: {
		Fname: __ccgo_ts + 11096,
		Fcodepoints: [2]uint32{
			0: uint32(176),
		},
	},
	883: {
		Fname: __ccgo_ts + 11102,
		Fcodepoints: [2]uint32{
			0: uint32(948),
		},
	},
	884: {
		Fname: __ccgo_ts + 11110,
		Fcodepoints: [2]uint32{
			0: uint32(10673),
		},
	},
	885: {
		Fname: __ccgo_ts + 11120,
		Fcodepoints: [2]uint32{
			0: uint32(10623),
		},
	},
	886: {
		Fname: __ccgo_ts + 11129,
		Fcodepoints: [2]uint32{
			0: uint32(120097),
		},
	},
	887: {
		Fname: __ccgo_ts + 11135,
		Fcodepoints: [2]uint32{
			0: uint32(8643),
		},
	},
	888: {
		Fname: __ccgo_ts + 11143,
		Fcodepoints: [2]uint32{
			0: uint32(8642),
		},
	},
	889: {
		Fname: __ccgo_ts + 11151,
		Fcodepoints: [2]uint32{
			0: uint32(8900),
		},
	},
	890: {
		Fname: __ccgo_ts + 11158,
		Fcodepoints: [2]uint32{
			0: uint32(8900),
		},
	},
	891: {
		Fname: __ccgo_ts + 11168,
		Fcodepoints: [2]uint32{
			0: uint32(9830),
		},
	},
	892: {
		Fname: __ccgo_ts + 11182,
		Fcodepoints: [2]uint32{
			0: uint32(9830),
		},
	},
	893: {
		Fname: __ccgo_ts + 11190,
		Fcodepoints: [2]uint32{
			0: uint32(168),
		},
	},
	894: {
		Fname: __ccgo_ts + 11196,
		Fcodepoints: [2]uint32{
			0: uint32(989),
		},
	},
	895: {
		Fname: __ccgo_ts + 11206,
		Fcodepoints: [2]uint32{
			0: uint32(8946),
		},
	},
	896: {
		Fname: __ccgo_ts + 11214,
		Fcodepoints: [2]uint32{
			0: uint32(247),
		},
	},
	897: {
		Fname: __ccgo_ts + 11220,
		Fcodepoints: [2]uint32{
			0: uint32(247),
		},
	},
	898: {
		Fname: __ccgo_ts + 11229,
		Fcodepoints: [2]uint32{
			0: uint32(8903),
		},
	},
	899: {
		Fname: __ccgo_ts + 11245,
		Fcodepoints: [2]uint32{
			0: uint32(8903),
		},
	},
	900: {
		Fname: __ccgo_ts + 11254,
		Fcodepoints: [2]uint32{
			0: uint32(1106),
		},
	},
	901: {
		Fname: __ccgo_ts + 11261,
		Fcodepoints: [2]uint32{
			0: uint32(8990),
		},
	},
	902: {
		Fname: __ccgo_ts + 11270,
		Fcodepoints: [2]uint32{
			0: uint32(8973),
		},
	},
	903: {
		Fname: __ccgo_ts + 11279,
		Fcodepoints: [2]uint32{
			0: uint32(36),
		},
	},
	904: {
		Fname: __ccgo_ts + 11288,
		Fcodepoints: [2]uint32{
			0: uint32(120149),
		},
	},
	905: {
		Fname: __ccgo_ts + 11295,
		Fcodepoints: [2]uint32{
			0: uint32(729),
		},
	},
	906: {
		Fname: __ccgo_ts + 11301,
		Fcodepoints: [2]uint32{
			0: uint32(8784),
		},
	},
	907: {
		Fname: __ccgo_ts + 11309,
		Fcodepoints: [2]uint32{
			0: uint32(8785),
		},
	},
	908: {
		Fname: __ccgo_ts + 11320,
		Fcodepoints: [2]uint32{
			0: uint32(8760),
		},
	},
	909: {
		Fname: __ccgo_ts + 11331,
		Fcodepoints: [2]uint32{
			0: uint32(8724),
		},
	},
	910: {
		Fname: __ccgo_ts + 11341,
		Fcodepoints: [2]uint32{
			0: uint32(8865),
		},
	},
	911: {
		Fname: __ccgo_ts + 11353,
		Fcodepoints: [2]uint32{
			0: uint32(8966),
		},
	},
	912: {
		Fname: __ccgo_ts + 11370,
		Fcodepoints: [2]uint32{
			0: uint32(8595),
		},
	},
	913: {
		Fname: __ccgo_ts + 11382,
		Fcodepoints: [2]uint32{
			0: uint32(8650),
		},
	},
	914: {
		Fname: __ccgo_ts + 11399,
		Fcodepoints: [2]uint32{
			0: uint32(8643),
		},
	},
	915: {
		Fname: __ccgo_ts + 11417,
		Fcodepoints: [2]uint32{
			0: uint32(8642),
		},
	},
	916: {
		Fname: __ccgo_ts + 11436,
		Fcodepoints: [2]uint32{
			0: uint32(10512),
		},
	},
	917: {
		Fname: __ccgo_ts + 11447,
		Fcodepoints: [2]uint32{
			0: uint32(8991),
		},
	},
	918: {
		Fname: __ccgo_ts + 11456,
		Fcodepoints: [2]uint32{
			0: uint32(8972),
		},
	},
	919: {
		Fname: __ccgo_ts + 11465,
		Fcodepoints: [2]uint32{
			0: uint32(119993),
		},
	},
	920: {
		Fname: __ccgo_ts + 11472,
		Fcodepoints: [2]uint32{
			0: uint32(1109),
		},
	},
	921: {
		Fname: __ccgo_ts + 11479,
		Fcodepoints: [2]uint32{
			0: uint32(10742),
		},
	},
	922: {
		Fname: __ccgo_ts + 11486,
		Fcodepoints: [2]uint32{
			0: uint32(273),
		},
	},
	923: {
		Fname: __ccgo_ts + 11495,
		Fcodepoints: [2]uint32{
			0: uint32(8945),
		},
	},
	924: {
		Fname: __ccgo_ts + 11503,
		Fcodepoints: [2]uint32{
			0: uint32(9663),
		},
	},
	925: {
		Fname: __ccgo_ts + 11510,
		Fcodepoints: [2]uint32{
			0: uint32(9662),
		},
	},
	926: {
		Fname: __ccgo_ts + 11518,
		Fcodepoints: [2]uint32{
			0: uint32(8693),
		},
	},
	927: {
		Fname: __ccgo_ts + 11526,
		Fcodepoints: [2]uint32{
			0: uint32(10607),
		},
	},
	928: {
		Fname: __ccgo_ts + 11534,
		Fcodepoints: [2]uint32{
			0: uint32(10662),
		},
	},
	929: {
		Fname: __ccgo_ts + 11544,
		Fcodepoints: [2]uint32{
			0: uint32(1119),
		},
	},
	930: {
		Fname: __ccgo_ts + 11551,
		Fcodepoints: [2]uint32{
			0: uint32(10239),
		},
	},
	931: {
		Fname: __ccgo_ts + 11562,
		Fcodepoints: [2]uint32{
			0: uint32(10871),
		},
	},
	932: {
		Fname: __ccgo_ts + 11570,
		Fcodepoints: [2]uint32{
			0: uint32(8785),
		},
	},
	933: {
		Fname: __ccgo_ts + 11577,
		Fcodepoints: [2]uint32{
			0: uint32(233),
		},
	},
	934: {
		Fname: __ccgo_ts + 11586,
		Fcodepoints: [2]uint32{
			0: uint32(10862),
		},
	},
	935: {
		Fname: __ccgo_ts + 11595,
		Fcodepoints: [2]uint32{
			0: uint32(283),
		},
	},
	936: {
		Fname: __ccgo_ts + 11604,
		Fcodepoints: [2]uint32{
			0: uint32(8790),
		},
	},
	937: {
		Fname: __ccgo_ts + 11611,
		Fcodepoints: [2]uint32{
			0: uint32(234),
		},
	},
	938: {
		Fname: __ccgo_ts + 11619,
		Fcodepoints: [2]uint32{
			0: uint32(8789),
		},
	},
	939: {
		Fname: __ccgo_ts + 11628,
		Fcodepoints: [2]uint32{
			0: uint32(1101),
		},
	},
	940: {
		Fname: __ccgo_ts + 11634,
		Fcodepoints: [2]uint32{
			0: uint32(279),
		},
	},
	941: {
		Fname: __ccgo_ts + 11641,
		Fcodepoints: [2]uint32{
			0: uint32(8519),
		},
	},
	942: {
		Fname: __ccgo_ts + 11646,
		Fcodepoints: [2]uint32{
			0: uint32(8786),
		},
	},
	943: {
		Fname: __ccgo_ts + 11654,
		Fcodepoints: [2]uint32{
			0: uint32(120098),
		},
	},
	944: {
		Fname: __ccgo_ts + 11660,
		Fcodepoints: [2]uint32{
			0: uint32(10906),
		},
	},
	945: {
		Fname: __ccgo_ts + 11665,
		Fcodepoints: [2]uint32{
			0: uint32(232),
		},
	},
	946: {
		Fname: __ccgo_ts + 11674,
		Fcodepoints: [2]uint32{
			0: uint32(10902),
		},
	},
	947: {
		Fname: __ccgo_ts + 11680,
		Fcodepoints: [2]uint32{
			0: uint32(10904),
		},
	},
	948: {
		Fname: __ccgo_ts + 11689,
		Fcodepoints: [2]uint32{
			0: uint32(10905),
		},
	},
	949: {
		Fname: __ccgo_ts + 11694,
		Fcodepoints: [2]uint32{
			0: uint32(9191),
		},
	},
	950: {
		Fname: __ccgo_ts + 11705,
		Fcodepoints: [2]uint32{
			0: uint32(8467),
		},
	},
	951: {
		Fname: __ccgo_ts + 11711,
		Fcodepoints: [2]uint32{
			0: uint32(10901),
		},
	},
	952: {
		Fname: __ccgo_ts + 11717,
		Fcodepoints: [2]uint32{
			0: uint32(10903),
		},
	},
	953: {
		Fname: __ccgo_ts + 11726,
		Fcodepoints: [2]uint32{
			0: uint32(275),
		},
	},
	954: {
		Fname: __ccgo_ts + 11734,
		Fcodepoints: [2]uint32{
			0: uint32(8709),
		},
	},
	955: {
		Fname: __ccgo_ts + 11742,
		Fcodepoints: [2]uint32{
			0: uint32(8709),
		},
	},
	956: {
		Fname: __ccgo_ts + 11753,
		Fcodepoints: [2]uint32{
			0: uint32(8709),
		},
	},
	957: {
		Fname: __ccgo_ts + 11762,
		Fcodepoints: [2]uint32{
			0: uint32(8196),
		},
	},
	958: {
		Fname: __ccgo_ts + 11771,
		Fcodepoints: [2]uint32{
			0: uint32(8197),
		},
	},
	959: {
		Fname: __ccgo_ts + 11780,
		Fcodepoints: [2]uint32{
			0: uint32(8195),
		},
	},
	960: {
		Fname: __ccgo_ts + 11787,
		Fcodepoints: [2]uint32{
			0: uint32(331),
		},
	},
	961: {
		Fname: __ccgo_ts + 11793,
		Fcodepoints: [2]uint32{
			0: uint32(8194),
		},
	},
	962: {
		Fname: __ccgo_ts + 11800,
		Fcodepoints: [2]uint32{
			0: uint32(281),
		},
	},
	963: {
		Fname: __ccgo_ts + 11808,
		Fcodepoints: [2]uint32{
			0: uint32(120150),
		},
	},
	964: {
		Fname: __ccgo_ts + 11815,
		Fcodepoints: [2]uint32{
			0: uint32(8917),
		},
	},
	965: {
		Fname: __ccgo_ts + 11822,
		Fcodepoints: [2]uint32{
			0: uint32(10723),
		},
	},
	966: {
		Fname: __ccgo_ts + 11831,
		Fcodepoints: [2]uint32{
			0: uint32(10865),
		},
	},
	967: {
		Fname: __ccgo_ts + 11839,
		Fcodepoints: [2]uint32{
			0: uint32(949),
		},
	},
	968: {
		Fname: __ccgo_ts + 11846,
		Fcodepoints: [2]uint32{
			0: uint32(949),
		},
	},
	969: {
		Fname: __ccgo_ts + 11856,
		Fcodepoints: [2]uint32{
			0: uint32(1013),
		},
	},
	970: {
		Fname: __ccgo_ts + 11864,
		Fcodepoints: [2]uint32{
			0: uint32(8790),
		},
	},
	971: {
		Fname: __ccgo_ts + 11873,
		Fcodepoints: [2]uint32{
			0: uint32(8789),
		},
	},
	972: {
		Fname: __ccgo_ts + 11883,
		Fcodepoints: [2]uint32{
			0: uint32(8770),
		},
	},
	973: {
		Fname: __ccgo_ts + 11891,
		Fcodepoints: [2]uint32{
			0: uint32(10902),
		},
	},
	974: {
		Fname: __ccgo_ts + 11904,
		Fcodepoints: [2]uint32{
			0: uint32(10901),
		},
	},
	975: {
		Fname: __ccgo_ts + 11918,
		Fcodepoints: [2]uint32{
			0: uint32(61),
		},
	},
	976: {
		Fname: __ccgo_ts + 11927,
		Fcodepoints: [2]uint32{
			0: uint32(8799),
		},
	},
	977: {
		Fname: __ccgo_ts + 11936,
		Fcodepoints: [2]uint32{
			0: uint32(8801),
		},
	},
	978: {
		Fname: __ccgo_ts + 11944,
		Fcodepoints: [2]uint32{
			0: uint32(10872),
		},
	},
	979: {
		Fname: __ccgo_ts + 11954,
		Fcodepoints: [2]uint32{
			0: uint32(10725),
		},
	},
	980: {
		Fname: __ccgo_ts + 11965,
		Fcodepoints: [2]uint32{
			0: uint32(8787),
		},
	},
	981: {
		Fname: __ccgo_ts + 11973,
		Fcodepoints: [2]uint32{
			0: uint32(10609),
		},
	},
	982: {
		Fname: __ccgo_ts + 11981,
		Fcodepoints: [2]uint32{
			0: uint32(8495),
		},
	},
	983: {
		Fname: __ccgo_ts + 11988,
		Fcodepoints: [2]uint32{
			0: uint32(8784),
		},
	},
	984: {
		Fname: __ccgo_ts + 11996,
		Fcodepoints: [2]uint32{
			0: uint32(8770),
		},
	},
	985: {
		Fname: __ccgo_ts + 12003,
		Fcodepoints: [2]uint32{
			0: uint32(951),
		},
	},
	986: {
		Fname: __ccgo_ts + 12009,
		Fcodepoints: [2]uint32{
			0: uint32(240),
		},
	},
	987: {
		Fname: __ccgo_ts + 12015,
		Fcodepoints: [2]uint32{
			0: uint32(235),
		},
	},
	988: {
		Fname: __ccgo_ts + 12022,
		Fcodepoints: [2]uint32{
			0: uint32(8364),
		},
	},
	989: {
		Fname: __ccgo_ts + 12029,
		Fcodepoints: [2]uint32{
			0: uint32(33),
		},
	},
	990: {
		Fname: __ccgo_ts + 12036,
		Fcodepoints: [2]uint32{
			0: uint32(8707),
		},
	},
	991: {
		Fname: __ccgo_ts + 12044,
		Fcodepoints: [2]uint32{
			0: uint32(8496),
		},
	},
	992: {
		Fname: __ccgo_ts + 12058,
		Fcodepoints: [2]uint32{
			0: uint32(8519),
		},
	},
	993: {
		Fname: __ccgo_ts + 12073,
		Fcodepoints: [2]uint32{
			0: uint32(8786),
		},
	},
	994: {
		Fname: __ccgo_ts + 12089,
		Fcodepoints: [2]uint32{
			0: uint32(1092),
		},
	},
	995: {
		Fname: __ccgo_ts + 12095,
		Fcodepoints: [2]uint32{
			0: uint32(9792),
		},
	},
	996: {
		Fname: __ccgo_ts + 12104,
		Fcodepoints: [2]uint32{
			0: uint32(64259),
		},
	},
	997: {
		Fname: __ccgo_ts + 12113,
		Fcodepoints: [2]uint32{
			0: uint32(64256),
		},
	},
	998: {
		Fname: __ccgo_ts + 12121,
		Fcodepoints: [2]uint32{
			0: uint32(64260),
		},
	},
	999: {
		Fname: __ccgo_ts + 12130,
		Fcodepoints: [2]uint32{
			0: uint32(120099),
		},
	},
	1000: {
		Fname: __ccgo_ts + 12136,
		Fcodepoints: [2]uint32{
			0: uint32(64257),
		},
	},
	1001: {
		Fname: __ccgo_ts + 12144,
		Fcodepoints: [2]uint32{
			0: uint32(102),
			1: uint32(106),
		},
	},
	1002: {
		Fname: __ccgo_ts + 12152,
		Fcodepoints: [2]uint32{
			0: uint32(9837),
		},
	},
	1003: {
		Fname: __ccgo_ts + 12159,
		Fcodepoints: [2]uint32{
			0: uint32(64258),
		},
	},
	1004: {
		Fname: __ccgo_ts + 12167,
		Fcodepoints: [2]uint32{
			0: uint32(9649),
		},
	},
	1005: {
		Fname: __ccgo_ts + 12175,
		Fcodepoints: [2]uint32{
			0: uint32(402),
		},
	},
	1006: {
		Fname: __ccgo_ts + 12182,
		Fcodepoints: [2]uint32{
			0: uint32(120151),
		},
	},
	1007: {
		Fname: __ccgo_ts + 12189,
		Fcodepoints: [2]uint32{
			0: uint32(8704),
		},
	},
	1008: {
		Fname: __ccgo_ts + 12198,
		Fcodepoints: [2]uint32{
			0: uint32(8916),
		},
	},
	1009: {
		Fname: __ccgo_ts + 12205,
		Fcodepoints: [2]uint32{
			0: uint32(10969),
		},
	},
	1010: {
		Fname: __ccgo_ts + 12213,
		Fcodepoints: [2]uint32{
			0: uint32(10765),
		},
	},
	1011: {
		Fname: __ccgo_ts + 12224,
		Fcodepoints: [2]uint32{
			0: uint32(189),
		},
	},
	1012: {
		Fname: __ccgo_ts + 12233,
		Fcodepoints: [2]uint32{
			0: uint32(8531),
		},
	},
	1013: {
		Fname: __ccgo_ts + 12242,
		Fcodepoints: [2]uint32{
			0: uint32(188),
		},
	},
	1014: {
		Fname: __ccgo_ts + 12251,
		Fcodepoints: [2]uint32{
			0: uint32(8533),
		},
	},
	1015: {
		Fname: __ccgo_ts + 12260,
		Fcodepoints: [2]uint32{
			0: uint32(8537),
		},
	},
	1016: {
		Fname: __ccgo_ts + 12269,
		Fcodepoints: [2]uint32{
			0: uint32(8539),
		},
	},
	1017: {
		Fname: __ccgo_ts + 12278,
		Fcodepoints: [2]uint32{
			0: uint32(8532),
		},
	},
	1018: {
		Fname: __ccgo_ts + 12287,
		Fcodepoints: [2]uint32{
			0: uint32(8534),
		},
	},
	1019: {
		Fname: __ccgo_ts + 12296,
		Fcodepoints: [2]uint32{
			0: uint32(190),
		},
	},
	1020: {
		Fname: __ccgo_ts + 12305,
		Fcodepoints: [2]uint32{
			0: uint32(8535),
		},
	},
	1021: {
		Fname: __ccgo_ts + 12314,
		Fcodepoints: [2]uint32{
			0: uint32(8540),
		},
	},
	1022: {
		Fname: __ccgo_ts + 12323,
		Fcodepoints: [2]uint32{
			0: uint32(8536),
		},
	},
	1023: {
		Fname: __ccgo_ts + 12332,
		Fcodepoints: [2]uint32{
			0: uint32(8538),
		},
	},
	1024: {
		Fname: __ccgo_ts + 12341,
		Fcodepoints: [2]uint32{
			0: uint32(8541),
		},
	},
	1025: {
		Fname: __ccgo_ts + 12350,
		Fcodepoints: [2]uint32{
			0: uint32(8542),
		},
	},
	1026: {
		Fname: __ccgo_ts + 12359,
		Fcodepoints: [2]uint32{
			0: uint32(8260),
		},
	},
	1027: {
		Fname: __ccgo_ts + 12367,
		Fcodepoints: [2]uint32{
			0: uint32(8994),
		},
	},
	1028: {
		Fname: __ccgo_ts + 12375,
		Fcodepoints: [2]uint32{
			0: uint32(119995),
		},
	},
	1029: {
		Fname: __ccgo_ts + 12382,
		Fcodepoints: [2]uint32{
			0: uint32(8807),
		},
	},
	1030: {
		Fname: __ccgo_ts + 12387,
		Fcodepoints: [2]uint32{
			0: uint32(10892),
		},
	},
	1031: {
		Fname: __ccgo_ts + 12393,
		Fcodepoints: [2]uint32{
			0: uint32(501),
		},
	},
	1032: {
		Fname: __ccgo_ts + 12402,
		Fcodepoints: [2]uint32{
			0: uint32(947),
		},
	},
	1033: {
		Fname: __ccgo_ts + 12410,
		Fcodepoints: [2]uint32{
			0: uint32(989),
		},
	},
	1034: {
		Fname: __ccgo_ts + 12419,
		Fcodepoints: [2]uint32{
			0: uint32(10886),
		},
	},
	1035: {
		Fname: __ccgo_ts + 12425,
		Fcodepoints: [2]uint32{
			0: uint32(287),
		},
	},
	1036: {
		Fname: __ccgo_ts + 12434,
		Fcodepoints: [2]uint32{
			0: uint32(285),
		},
	},
	1037: {
		Fname: __ccgo_ts + 12442,
		Fcodepoints: [2]uint32{
			0: uint32(1075),
		},
	},
	1038: {
		Fname: __ccgo_ts + 12448,
		Fcodepoints: [2]uint32{
			0: uint32(289),
		},
	},
	1039: {
		Fname: __ccgo_ts + 12455,
		Fcodepoints: [2]uint32{
			0: uint32(8805),
		},
	},
	1040: {
		Fname: __ccgo_ts + 12460,
		Fcodepoints: [2]uint32{
			0: uint32(8923),
		},
	},
	1041: {
		Fname: __ccgo_ts + 12466,
		Fcodepoints: [2]uint32{
			0: uint32(8805),
		},
	},
	1042: {
		Fname: __ccgo_ts + 12472,
		Fcodepoints: [2]uint32{
			0: uint32(8807),
		},
	},
	1043: {
		Fname: __ccgo_ts + 12479,
		Fcodepoints: [2]uint32{
			0: uint32(10878),
		},
	},
	1044: {
		Fname: __ccgo_ts + 12490,
		Fcodepoints: [2]uint32{
			0: uint32(10878),
		},
	},
	1045: {
		Fname: __ccgo_ts + 12496,
		Fcodepoints: [2]uint32{
			0: uint32(10921),
		},
	},
	1046: {
		Fname: __ccgo_ts + 12504,
		Fcodepoints: [2]uint32{
			0: uint32(10880),
		},
	},
	1047: {
		Fname: __ccgo_ts + 12513,
		Fcodepoints: [2]uint32{
			0: uint32(10882),
		},
	},
	1048: {
		Fname: __ccgo_ts + 12523,
		Fcodepoints: [2]uint32{
			0: uint32(10884),
		},
	},
	1049: {
		Fname: __ccgo_ts + 12534,
		Fcodepoints: [2]uint32{
			0: uint32(8923),
			1: uint32(65024),
		},
	},
	1050: {
		Fname: __ccgo_ts + 12541,
		Fcodepoints: [2]uint32{
			0: uint32(10900),
		},
	},
	1051: {
		Fname: __ccgo_ts + 12550,
		Fcodepoints: [2]uint32{
			0: uint32(120100),
		},
	},
	1052: {
		Fname: __ccgo_ts + 12556,
		Fcodepoints: [2]uint32{
			0: uint32(8811),
		},
	},
	1053: {
		Fname: __ccgo_ts + 12561,
		Fcodepoints: [2]uint32{
			0: uint32(8921),
		},
	},
	1054: {
		Fname: __ccgo_ts + 12567,
		Fcodepoints: [2]uint32{
			0: uint32(8503),
		},
	},
	1055: {
		Fname: __ccgo_ts + 12575,
		Fcodepoints: [2]uint32{
			0: uint32(1107),
		},
	},
	1056: {
		Fname: __ccgo_ts + 12582,
		Fcodepoints: [2]uint32{
			0: uint32(8823),
		},
	},
	1057: {
		Fname: __ccgo_ts + 12587,
		Fcodepoints: [2]uint32{
			0: uint32(10898),
		},
	},
	1058: {
		Fname: __ccgo_ts + 12593,
		Fcodepoints: [2]uint32{
			0: uint32(10917),
		},
	},
	1059: {
		Fname: __ccgo_ts + 12599,
		Fcodepoints: [2]uint32{
			0: uint32(10916),
		},
	},
	1060: {
		Fname: __ccgo_ts + 12605,
		Fcodepoints: [2]uint32{
			0: uint32(8809),
		},
	},
	1061: {
		Fname: __ccgo_ts + 12611,
		Fcodepoints: [2]uint32{
			0: uint32(10890),
		},
	},
	1062: {
		Fname: __ccgo_ts + 12618,
		Fcodepoints: [2]uint32{
			0: uint32(10890),
		},
	},
	1063: {
		Fname: __ccgo_ts + 12629,
		Fcodepoints: [2]uint32{
			0: uint32(10888),
		},
	},
	1064: {
		Fname: __ccgo_ts + 12635,
		Fcodepoints: [2]uint32{
			0: uint32(10888),
		},
	},
	1065: {
		Fname: __ccgo_ts + 12642,
		Fcodepoints: [2]uint32{
			0: uint32(8809),
		},
	},
	1066: {
		Fname: __ccgo_ts + 12650,
		Fcodepoints: [2]uint32{
			0: uint32(8935),
		},
	},
	1067: {
		Fname: __ccgo_ts + 12658,
		Fcodepoints: [2]uint32{
			0: uint32(120152),
		},
	},
	1068: {
		Fname: __ccgo_ts + 12665,
		Fcodepoints: [2]uint32{
			0: uint32(96),
		},
	},
	1069: {
		Fname: __ccgo_ts + 12673,
		Fcodepoints: [2]uint32{
			0: uint32(8458),
		},
	},
	1070: {
		Fname: __ccgo_ts + 12680,
		Fcodepoints: [2]uint32{
			0: uint32(8819),
		},
	},
	1071: {
		Fname: __ccgo_ts + 12687,
		Fcodepoints: [2]uint32{
			0: uint32(10894),
		},
	},
	1072: {
		Fname: __ccgo_ts + 12695,
		Fcodepoints: [2]uint32{
			0: uint32(10896),
		},
	},
	1073: {
		Fname: __ccgo_ts + 848,
		Fcodepoints: [2]uint32{
			0: uint32(62),
		},
	},
	1074: {
		Fname: __ccgo_ts + 12703,
		Fcodepoints: [2]uint32{
			0: uint32(10919),
		},
	},
	1075: {
		Fname: __ccgo_ts + 12710,
		Fcodepoints: [2]uint32{
			0: uint32(10874),
		},
	},
	1076: {
		Fname: __ccgo_ts + 12718,
		Fcodepoints: [2]uint32{
			0: uint32(8919),
		},
	},
	1077: {
		Fname: __ccgo_ts + 12726,
		Fcodepoints: [2]uint32{
			0: uint32(10645),
		},
	},
	1078: {
		Fname: __ccgo_ts + 12735,
		Fcodepoints: [2]uint32{
			0: uint32(10876),
		},
	},
	1079: {
		Fname: __ccgo_ts + 12745,
		Fcodepoints: [2]uint32{
			0: uint32(10886),
		},
	},
	1080: {
		Fname: __ccgo_ts + 12757,
		Fcodepoints: [2]uint32{
			0: uint32(10616),
		},
	},
	1081: {
		Fname: __ccgo_ts + 12766,
		Fcodepoints: [2]uint32{
			0: uint32(8919),
		},
	},
	1082: {
		Fname: __ccgo_ts + 12775,
		Fcodepoints: [2]uint32{
			0: uint32(8923),
		},
	},
	1083: {
		Fname: __ccgo_ts + 12787,
		Fcodepoints: [2]uint32{
			0: uint32(10892),
		},
	},
	1084: {
		Fname: __ccgo_ts + 12800,
		Fcodepoints: [2]uint32{
			0: uint32(8823),
		},
	},
	1085: {
		Fname: __ccgo_ts + 12810,
		Fcodepoints: [2]uint32{
			0: uint32(8819),
		},
	},
	1086: {
		Fname: __ccgo_ts + 12819,
		Fcodepoints: [2]uint32{
			0: uint32(8809),
			1: uint32(65024),
		},
	},
	1087: {
		Fname: __ccgo_ts + 12831,
		Fcodepoints: [2]uint32{
			0: uint32(8809),
			1: uint32(65024),
		},
	},
	1088: {
		Fname: __ccgo_ts + 12838,
		Fcodepoints: [2]uint32{
			0: uint32(8660),
		},
	},
	1089: {
		Fname: __ccgo_ts + 12845,
		Fcodepoints: [2]uint32{
			0: uint32(8202),
		},
	},
	1090: {
		Fname: __ccgo_ts + 12854,
		Fcodepoints: [2]uint32{
			0: uint32(189),
		},
	},
	1091: {
		Fname: __ccgo_ts + 12861,
		Fcodepoints: [2]uint32{
			0: uint32(8459),
		},
	},
	1092: {
		Fname: __ccgo_ts + 12870,
		Fcodepoints: [2]uint32{
			0: uint32(1098),
		},
	},
	1093: {
		Fname: __ccgo_ts + 12879,
		Fcodepoints: [2]uint32{
			0: uint32(8596),
		},
	},
	1094: {
		Fname: __ccgo_ts + 12886,
		Fcodepoints: [2]uint32{
			0: uint32(10568),
		},
	},
	1095: {
		Fname: __ccgo_ts + 12896,
		Fcodepoints: [2]uint32{
			0: uint32(8621),
		},
	},
	1096: {
		Fname: __ccgo_ts + 12904,
		Fcodepoints: [2]uint32{
			0: uint32(8463),
		},
	},
	1097: {
		Fname: __ccgo_ts + 12911,
		Fcodepoints: [2]uint32{
			0: uint32(293),
		},
	},
	1098: {
		Fname: __ccgo_ts + 12919,
		Fcodepoints: [2]uint32{
			0: uint32(9829),
		},
	},
	1099: {
		Fname: __ccgo_ts + 12928,
		Fcodepoints: [2]uint32{
			0: uint32(9829),
		},
	},
	1100: {
		Fname: __ccgo_ts + 12940,
		Fcodepoints: [2]uint32{
			0: uint32(8230),
		},
	},
	1101: {
		Fname: __ccgo_ts + 12949,
		Fcodepoints: [2]uint32{
			0: uint32(8889),
		},
	},
	1102: {
		Fname: __ccgo_ts + 12958,
		Fcodepoints: [2]uint32{
			0: uint32(120101),
		},
	},
	1103: {
		Fname: __ccgo_ts + 12964,
		Fcodepoints: [2]uint32{
			0: uint32(10533),
		},
	},
	1104: {
		Fname: __ccgo_ts + 12975,
		Fcodepoints: [2]uint32{
			0: uint32(10534),
		},
	},
	1105: {
		Fname: __ccgo_ts + 12986,
		Fcodepoints: [2]uint32{
			0: uint32(8703),
		},
	},
	1106: {
		Fname: __ccgo_ts + 12994,
		Fcodepoints: [2]uint32{
			0: uint32(8763),
		},
	},
	1107: {
		Fname: __ccgo_ts + 13003,
		Fcodepoints: [2]uint32{
			0: uint32(8617),
		},
	},
	1108: {
		Fname: __ccgo_ts + 13019,
		Fcodepoints: [2]uint32{
			0: uint32(8618),
		},
	},
	1109: {
		Fname: __ccgo_ts + 13036,
		Fcodepoints: [2]uint32{
			0: uint32(120153),
		},
	},
	1110: {
		Fname: __ccgo_ts + 13043,
		Fcodepoints: [2]uint32{
			0: uint32(8213),
		},
	},
	1111: {
		Fname: __ccgo_ts + 13052,
		Fcodepoints: [2]uint32{
			0: uint32(119997),
		},
	},
	1112: {
		Fname: __ccgo_ts + 13059,
		Fcodepoints: [2]uint32{
			0: uint32(8463),
		},
	},
	1113: {
		Fname: __ccgo_ts + 13068,
		Fcodepoints: [2]uint32{
			0: uint32(295),
		},
	},
	1114: {
		Fname: __ccgo_ts + 13077,
		Fcodepoints: [2]uint32{
			0: uint32(8259),
		},
	},
	1115: {
		Fname: __ccgo_ts + 13086,
		Fcodepoints: [2]uint32{
			0: uint32(8208),
		},
	},
	1116: {
		Fname: __ccgo_ts + 13095,
		Fcodepoints: [2]uint32{
			0: uint32(237),
		},
	},
	1117: {
		Fname: __ccgo_ts + 13104,
		Fcodepoints: [2]uint32{
			0: uint32(8291),
		},
	},
	1118: {
		Fname: __ccgo_ts + 13109,
		Fcodepoints: [2]uint32{
			0: uint32(238),
		},
	},
	1119: {
		Fname: __ccgo_ts + 13117,
		Fcodepoints: [2]uint32{
			0: uint32(1080),
		},
	},
	1120: {
		Fname: __ccgo_ts + 13123,
		Fcodepoints: [2]uint32{
			0: uint32(1077),
		},
	},
	1121: {
		Fname: __ccgo_ts + 13130,
		Fcodepoints: [2]uint32{
			0: uint32(161),
		},
	},
	1122: {
		Fname: __ccgo_ts + 13138,
		Fcodepoints: [2]uint32{
			0: uint32(8660),
		},
	},
	1123: {
		Fname: __ccgo_ts + 13144,
		Fcodepoints: [2]uint32{
			0: uint32(120102),
		},
	},
	1124: {
		Fname: __ccgo_ts + 13150,
		Fcodepoints: [2]uint32{
			0: uint32(236),
		},
	},
	1125: {
		Fname: __ccgo_ts + 13159,
		Fcodepoints: [2]uint32{
			0: uint32(8520),
		},
	},
	1126: {
		Fname: __ccgo_ts + 13164,
		Fcodepoints: [2]uint32{
			0: uint32(10764),
		},
	},
	1127: {
		Fname: __ccgo_ts + 13173,
		Fcodepoints: [2]uint32{
			0: uint32(8749),
		},
	},
	1128: {
		Fname: __ccgo_ts + 13181,
		Fcodepoints: [2]uint32{
			0: uint32(10716),
		},
	},
	1129: {
		Fname: __ccgo_ts + 13190,
		Fcodepoints: [2]uint32{
			0: uint32(8489),
		},
	},
	1130: {
		Fname: __ccgo_ts + 13198,
		Fcodepoints: [2]uint32{
			0: uint32(307),
		},
	},
	1131: {
		Fname: __ccgo_ts + 13206,
		Fcodepoints: [2]uint32{
			0: uint32(299),
		},
	},
	1132: {
		Fname: __ccgo_ts + 13214,
		Fcodepoints: [2]uint32{
			0: uint32(8465),
		},
	},
	1133: {
		Fname: __ccgo_ts + 13222,
		Fcodepoints: [2]uint32{
			0: uint32(8464),
		},
	},
	1134: {
		Fname: __ccgo_ts + 13233,
		Fcodepoints: [2]uint32{
			0: uint32(8465),
		},
	},
	1135: {
		Fname: __ccgo_ts + 13244,
		Fcodepoints: [2]uint32{
			0: uint32(305),
		},
	},
	1136: {
		Fname: __ccgo_ts + 13252,
		Fcodepoints: [2]uint32{
			0: uint32(8887),
		},
	},
	1137: {
		Fname: __ccgo_ts + 13259,
		Fcodepoints: [2]uint32{
			0: uint32(437),
		},
	},
	1138: {
		Fname: __ccgo_ts + 13267,
		Fcodepoints: [2]uint32{
			0: uint32(8712),
		},
	},
	1139: {
		Fname: __ccgo_ts + 13272,
		Fcodepoints: [2]uint32{
			0: uint32(8453),
		},
	},
	1140: {
		Fname: __ccgo_ts + 13281,
		Fcodepoints: [2]uint32{
			0: uint32(8734),
		},
	},
	1141: {
		Fname: __ccgo_ts + 13289,
		Fcodepoints: [2]uint32{
			0: uint32(10717),
		},
	},
	1142: {
		Fname: __ccgo_ts + 13300,
		Fcodepoints: [2]uint32{
			0: uint32(305),
		},
	},
	1143: {
		Fname: __ccgo_ts + 13309,
		Fcodepoints: [2]uint32{
			0: uint32(8747),
		},
	},
	1144: {
		Fname: __ccgo_ts + 13315,
		Fcodepoints: [2]uint32{
			0: uint32(8890),
		},
	},
	1145: {
		Fname: __ccgo_ts + 13324,
		Fcodepoints: [2]uint32{
			0: uint32(8484),
		},
	},
	1146: {
		Fname: __ccgo_ts + 13335,
		Fcodepoints: [2]uint32{
			0: uint32(8890),
		},
	},
	1147: {
		Fname: __ccgo_ts + 13346,
		Fcodepoints: [2]uint32{
			0: uint32(10775),
		},
	},
	1148: {
		Fname: __ccgo_ts + 13357,
		Fcodepoints: [2]uint32{
			0: uint32(10812),
		},
	},
	1149: {
		Fname: __ccgo_ts + 13367,
		Fcodepoints: [2]uint32{
			0: uint32(1105),
		},
	},
	1150: {
		Fname: __ccgo_ts + 13374,
		Fcodepoints: [2]uint32{
			0: uint32(303),
		},
	},
	1151: {
		Fname: __ccgo_ts + 13382,
		Fcodepoints: [2]uint32{
			0: uint32(120154),
		},
	},
	1152: {
		Fname: __ccgo_ts + 13389,
		Fcodepoints: [2]uint32{
			0: uint32(953),
		},
	},
	1153: {
		Fname: __ccgo_ts + 13396,
		Fcodepoints: [2]uint32{
			0: uint32(10812),
		},
	},
	1154: {
		Fname: __ccgo_ts + 13404,
		Fcodepoints: [2]uint32{
			0: uint32(191),
		},
	},
	1155: {
		Fname: __ccgo_ts + 13413,
		Fcodepoints: [2]uint32{
			0: uint32(119998),
		},
	},
	1156: {
		Fname: __ccgo_ts + 13420,
		Fcodepoints: [2]uint32{
			0: uint32(8712),
		},
	},
	1157: {
		Fname: __ccgo_ts + 13427,
		Fcodepoints: [2]uint32{
			0: uint32(8953),
		},
	},
	1158: {
		Fname: __ccgo_ts + 13435,
		Fcodepoints: [2]uint32{
			0: uint32(8949),
		},
	},
	1159: {
		Fname: __ccgo_ts + 13445,
		Fcodepoints: [2]uint32{
			0: uint32(8948),
		},
	},
	1160: {
		Fname: __ccgo_ts + 13453,
		Fcodepoints: [2]uint32{
			0: uint32(8947),
		},
	},
	1161: {
		Fname: __ccgo_ts + 13462,
		Fcodepoints: [2]uint32{
			0: uint32(8712),
		},
	},
	1162: {
		Fname: __ccgo_ts + 13470,
		Fcodepoints: [2]uint32{
			0: uint32(8290),
		},
	},
	1163: {
		Fname: __ccgo_ts + 13475,
		Fcodepoints: [2]uint32{
			0: uint32(297),
		},
	},
	1164: {
		Fname: __ccgo_ts + 13484,
		Fcodepoints: [2]uint32{
			0: uint32(1110),
		},
	},
	1165: {
		Fname: __ccgo_ts + 13492,
		Fcodepoints: [2]uint32{
			0: uint32(239),
		},
	},
	1166: {
		Fname: __ccgo_ts + 13499,
		Fcodepoints: [2]uint32{
			0: uint32(309),
		},
	},
	1167: {
		Fname: __ccgo_ts + 13507,
		Fcodepoints: [2]uint32{
			0: uint32(1081),
		},
	},
	1168: {
		Fname: __ccgo_ts + 13513,
		Fcodepoints: [2]uint32{
			0: uint32(120103),
		},
	},
	1169: {
		Fname: __ccgo_ts + 13519,
		Fcodepoints: [2]uint32{
			0: uint32(567),
		},
	},
	1170: {
		Fname: __ccgo_ts + 13527,
		Fcodepoints: [2]uint32{
			0: uint32(120155),
		},
	},
	1171: {
		Fname: __ccgo_ts + 13534,
		Fcodepoints: [2]uint32{
			0: uint32(119999),
		},
	},
	1172: {
		Fname: __ccgo_ts + 13541,
		Fcodepoints: [2]uint32{
			0: uint32(1112),
		},
	},
	1173: {
		Fname: __ccgo_ts + 13550,
		Fcodepoints: [2]uint32{
			0: uint32(1108),
		},
	},
	1174: {
		Fname: __ccgo_ts + 13558,
		Fcodepoints: [2]uint32{
			0: uint32(954),
		},
	},
	1175: {
		Fname: __ccgo_ts + 13566,
		Fcodepoints: [2]uint32{
			0: uint32(1008),
		},
	},
	1176: {
		Fname: __ccgo_ts + 13575,
		Fcodepoints: [2]uint32{
			0: uint32(311),
		},
	},
	1177: {
		Fname: __ccgo_ts + 13584,
		Fcodepoints: [2]uint32{
			0: uint32(1082),
		},
	},
	1178: {
		Fname: __ccgo_ts + 13590,
		Fcodepoints: [2]uint32{
			0: uint32(120104),
		},
	},
	1179: {
		Fname: __ccgo_ts + 13596,
		Fcodepoints: [2]uint32{
			0: uint32(312),
		},
	},
	1180: {
		Fname: __ccgo_ts + 13605,
		Fcodepoints: [2]uint32{
			0: uint32(1093),
		},
	},
	1181: {
		Fname: __ccgo_ts + 13612,
		Fcodepoints: [2]uint32{
			0: uint32(1116),
		},
	},
	1182: {
		Fname: __ccgo_ts + 13619,
		Fcodepoints: [2]uint32{
			0: uint32(120156),
		},
	},
	1183: {
		Fname: __ccgo_ts + 13626,
		Fcodepoints: [2]uint32{
			0: uint32(120000),
		},
	},
	1184: {
		Fname: __ccgo_ts + 13633,
		Fcodepoints: [2]uint32{
			0: uint32(8666),
		},
	},
	1185: {
		Fname: __ccgo_ts + 13641,
		Fcodepoints: [2]uint32{
			0: uint32(8656),
		},
	},
	1186: {
		Fname: __ccgo_ts + 13648,
		Fcodepoints: [2]uint32{
			0: uint32(10523),
		},
	},
	1187: {
		Fname: __ccgo_ts + 13657,
		Fcodepoints: [2]uint32{
			0: uint32(10510),
		},
	},
	1188: {
		Fname: __ccgo_ts + 13665,
		Fcodepoints: [2]uint32{
			0: uint32(8806),
		},
	},
	1189: {
		Fname: __ccgo_ts + 13670,
		Fcodepoints: [2]uint32{
			0: uint32(10891),
		},
	},
	1190: {
		Fname: __ccgo_ts + 13676,
		Fcodepoints: [2]uint32{
			0: uint32(10594),
		},
	},
	1191: {
		Fname: __ccgo_ts + 13683,
		Fcodepoints: [2]uint32{
			0: uint32(314),
		},
	},
	1192: {
		Fname: __ccgo_ts + 13692,
		Fcodepoints: [2]uint32{
			0: uint32(10676),
		},
	},
	1193: {
		Fname: __ccgo_ts + 13703,
		Fcodepoints: [2]uint32{
			0: uint32(8466),
		},
	},
	1194: {
		Fname: __ccgo_ts + 13712,
		Fcodepoints: [2]uint32{
			0: uint32(955),
		},
	},
	1195: {
		Fname: __ccgo_ts + 13721,
		Fcodepoints: [2]uint32{
			0: uint32(10216),
		},
	},
	1196: {
		Fname: __ccgo_ts + 13728,
		Fcodepoints: [2]uint32{
			0: uint32(10641),
		},
	},
	1197: {
		Fname: __ccgo_ts + 13736,
		Fcodepoints: [2]uint32{
			0: uint32(10216),
		},
	},
	1198: {
		Fname: __ccgo_ts + 13745,
		Fcodepoints: [2]uint32{
			0: uint32(10885),
		},
	},
	1199: {
		Fname: __ccgo_ts + 13751,
		Fcodepoints: [2]uint32{
			0: uint32(171),
		},
	},
	1200: {
		Fname: __ccgo_ts + 13759,
		Fcodepoints: [2]uint32{
			0: uint32(8592),
		},
	},
	1201: {
		Fname: __ccgo_ts + 13766,
		Fcodepoints: [2]uint32{
			0: uint32(8676),
		},
	},
	1202: {
		Fname: __ccgo_ts + 13774,
		Fcodepoints: [2]uint32{
			0: uint32(10527),
		},
	},
	1203: {
		Fname: __ccgo_ts + 13784,
		Fcodepoints: [2]uint32{
			0: uint32(10525),
		},
	},
	1204: {
		Fname: __ccgo_ts + 13793,
		Fcodepoints: [2]uint32{
			0: uint32(8617),
		},
	},
	1205: {
		Fname: __ccgo_ts + 13802,
		Fcodepoints: [2]uint32{
			0: uint32(8619),
		},
	},
	1206: {
		Fname: __ccgo_ts + 13811,
		Fcodepoints: [2]uint32{
			0: uint32(10553),
		},
	},
	1207: {
		Fname: __ccgo_ts + 13820,
		Fcodepoints: [2]uint32{
			0: uint32(10611),
		},
	},
	1208: {
		Fname: __ccgo_ts + 13830,
		Fcodepoints: [2]uint32{
			0: uint32(8610),
		},
	},
	1209: {
		Fname: __ccgo_ts + 13839,
		Fcodepoints: [2]uint32{
			0: uint32(10923),
		},
	},
	1210: {
		Fname: __ccgo_ts + 13845,
		Fcodepoints: [2]uint32{
			0: uint32(10521),
		},
	},
	1211: {
		Fname: __ccgo_ts + 13854,
		Fcodepoints: [2]uint32{
			0: uint32(10925),
		},
	},
	1212: {
		Fname: __ccgo_ts + 13861,
		Fcodepoints: [2]uint32{
			0: uint32(10925),
			1: uint32(65024),
		},
	},
	1213: {
		Fname: __ccgo_ts + 13869,
		Fcodepoints: [2]uint32{
			0: uint32(10508),
		},
	},
	1214: {
		Fname: __ccgo_ts + 13877,
		Fcodepoints: [2]uint32{
			0: uint32(10098),
		},
	},
	1215: {
		Fname: __ccgo_ts + 13885,
		Fcodepoints: [2]uint32{
			0: uint32(123),
		},
	},
	1216: {
		Fname: __ccgo_ts + 13894,
		Fcodepoints: [2]uint32{
			0: uint32(91),
		},
	},
	1217: {
		Fname: __ccgo_ts + 13903,
		Fcodepoints: [2]uint32{
			0: uint32(10635),
		},
	},
	1218: {
		Fname: __ccgo_ts + 13911,
		Fcodepoints: [2]uint32{
			0: uint32(10639),
		},
	},
	1219: {
		Fname: __ccgo_ts + 13921,
		Fcodepoints: [2]uint32{
			0: uint32(10637),
		},
	},
	1220: {
		Fname: __ccgo_ts + 13931,
		Fcodepoints: [2]uint32{
			0: uint32(318),
		},
	},
	1221: {
		Fname: __ccgo_ts + 13940,
		Fcodepoints: [2]uint32{
			0: uint32(316),
		},
	},
	1222: {
		Fname: __ccgo_ts + 13949,
		Fcodepoints: [2]uint32{
			0: uint32(8968),
		},
	},
	1223: {
		Fname: __ccgo_ts + 13957,
		Fcodepoints: [2]uint32{
			0: uint32(123),
		},
	},
	1224: {
		Fname: __ccgo_ts + 13964,
		Fcodepoints: [2]uint32{
			0: uint32(1083),
		},
	},
	1225: {
		Fname: __ccgo_ts + 13970,
		Fcodepoints: [2]uint32{
			0: uint32(10550),
		},
	},
	1226: {
		Fname: __ccgo_ts + 13977,
		Fcodepoints: [2]uint32{
			0: uint32(8220),
		},
	},
	1227: {
		Fname: __ccgo_ts + 13985,
		Fcodepoints: [2]uint32{
			0: uint32(8222),
		},
	},
	1228: {
		Fname: __ccgo_ts + 13994,
		Fcodepoints: [2]uint32{
			0: uint32(10599),
		},
	},
	1229: {
		Fname: __ccgo_ts + 14004,
		Fcodepoints: [2]uint32{
			0: uint32(10571),
		},
	},
	1230: {
		Fname: __ccgo_ts + 14015,
		Fcodepoints: [2]uint32{
			0: uint32(8626),
		},
	},
	1231: {
		Fname: __ccgo_ts + 14022,
		Fcodepoints: [2]uint32{
			0: uint32(8804),
		},
	},
	1232: {
		Fname: __ccgo_ts + 14027,
		Fcodepoints: [2]uint32{
			0: uint32(8592),
		},
	},
	1233: {
		Fname: __ccgo_ts + 14039,
		Fcodepoints: [2]uint32{
			0: uint32(8610),
		},
	},
	1234: {
		Fname: __ccgo_ts + 14055,
		Fcodepoints: [2]uint32{
			0: uint32(8637),
		},
	},
	1235: {
		Fname: __ccgo_ts + 14073,
		Fcodepoints: [2]uint32{
			0: uint32(8636),
		},
	},
	1236: {
		Fname: __ccgo_ts + 14089,
		Fcodepoints: [2]uint32{
			0: uint32(8647),
		},
	},
	1237: {
		Fname: __ccgo_ts + 14106,
		Fcodepoints: [2]uint32{
			0: uint32(8596),
		},
	},
	1238: {
		Fname: __ccgo_ts + 14123,
		Fcodepoints: [2]uint32{
			0: uint32(8646),
		},
	},
	1239: {
		Fname: __ccgo_ts + 14141,
		Fcodepoints: [2]uint32{
			0: uint32(8651),
		},
	},
	1240: {
		Fname: __ccgo_ts + 14161,
		Fcodepoints: [2]uint32{
			0: uint32(8621),
		},
	},
	1241: {
		Fname: __ccgo_ts + 14183,
		Fcodepoints: [2]uint32{
			0: uint32(8907),
		},
	},
	1242: {
		Fname: __ccgo_ts + 14200,
		Fcodepoints: [2]uint32{
			0: uint32(8922),
		},
	},
	1243: {
		Fname: __ccgo_ts + 14206,
		Fcodepoints: [2]uint32{
			0: uint32(8804),
		},
	},
	1244: {
		Fname: __ccgo_ts + 14212,
		Fcodepoints: [2]uint32{
			0: uint32(8806),
		},
	},
	1245: {
		Fname: __ccgo_ts + 14219,
		Fcodepoints: [2]uint32{
			0: uint32(10877),
		},
	},
	1246: {
		Fname: __ccgo_ts + 14230,
		Fcodepoints: [2]uint32{
			0: uint32(10877),
		},
	},
	1247: {
		Fname: __ccgo_ts + 14236,
		Fcodepoints: [2]uint32{
			0: uint32(10920),
		},
	},
	1248: {
		Fname: __ccgo_ts + 14244,
		Fcodepoints: [2]uint32{
			0: uint32(10879),
		},
	},
	1249: {
		Fname: __ccgo_ts + 14253,
		Fcodepoints: [2]uint32{
			0: uint32(10881),
		},
	},
	1250: {
		Fname: __ccgo_ts + 14263,
		Fcodepoints: [2]uint32{
			0: uint32(10883),
		},
	},
	1251: {
		Fname: __ccgo_ts + 14274,
		Fcodepoints: [2]uint32{
			0: uint32(8922),
			1: uint32(65024),
		},
	},
	1252: {
		Fname: __ccgo_ts + 14281,
		Fcodepoints: [2]uint32{
			0: uint32(10899),
		},
	},
	1253: {
		Fname: __ccgo_ts + 14290,
		Fcodepoints: [2]uint32{
			0: uint32(10885),
		},
	},
	1254: {
		Fname: __ccgo_ts + 14303,
		Fcodepoints: [2]uint32{
			0: uint32(8918),
		},
	},
	1255: {
		Fname: __ccgo_ts + 14313,
		Fcodepoints: [2]uint32{
			0: uint32(8922),
		},
	},
	1256: {
		Fname: __ccgo_ts + 14325,
		Fcodepoints: [2]uint32{
			0: uint32(10891),
		},
	},
	1257: {
		Fname: __ccgo_ts + 14338,
		Fcodepoints: [2]uint32{
			0: uint32(8822),
		},
	},
	1258: {
		Fname: __ccgo_ts + 14348,
		Fcodepoints: [2]uint32{
			0: uint32(8818),
		},
	},
	1259: {
		Fname: __ccgo_ts + 14358,
		Fcodepoints: [2]uint32{
			0: uint32(10620),
		},
	},
	1260: {
		Fname: __ccgo_ts + 14367,
		Fcodepoints: [2]uint32{
			0: uint32(8970),
		},
	},
	1261: {
		Fname: __ccgo_ts + 14376,
		Fcodepoints: [2]uint32{
			0: uint32(120105),
		},
	},
	1262: {
		Fname: __ccgo_ts + 14382,
		Fcodepoints: [2]uint32{
			0: uint32(8822),
		},
	},
	1263: {
		Fname: __ccgo_ts + 14387,
		Fcodepoints: [2]uint32{
			0: uint32(10897),
		},
	},
	1264: {
		Fname: __ccgo_ts + 14393,
		Fcodepoints: [2]uint32{
			0: uint32(8637),
		},
	},
	1265: {
		Fname: __ccgo_ts + 14401,
		Fcodepoints: [2]uint32{
			0: uint32(8636),
		},
	},
	1266: {
		Fname: __ccgo_ts + 14409,
		Fcodepoints: [2]uint32{
			0: uint32(10602),
		},
	},
	1267: {
		Fname: __ccgo_ts + 14418,
		Fcodepoints: [2]uint32{
			0: uint32(9604),
		},
	},
	1268: {
		Fname: __ccgo_ts + 14426,
		Fcodepoints: [2]uint32{
			0: uint32(1113),
		},
	},
	1269: {
		Fname: __ccgo_ts + 14433,
		Fcodepoints: [2]uint32{
			0: uint32(8810),
		},
	},
	1270: {
		Fname: __ccgo_ts + 14438,
		Fcodepoints: [2]uint32{
			0: uint32(8647),
		},
	},
	1271: {
		Fname: __ccgo_ts + 14446,
		Fcodepoints: [2]uint32{
			0: uint32(8990),
		},
	},
	1272: {
		Fname: __ccgo_ts + 14457,
		Fcodepoints: [2]uint32{
			0: uint32(10603),
		},
	},
	1273: {
		Fname: __ccgo_ts + 14466,
		Fcodepoints: [2]uint32{
			0: uint32(9722),
		},
	},
	1274: {
		Fname: __ccgo_ts + 14474,
		Fcodepoints: [2]uint32{
			0: uint32(320),
		},
	},
	1275: {
		Fname: __ccgo_ts + 14483,
		Fcodepoints: [2]uint32{
			0: uint32(9136),
		},
	},
	1276: {
		Fname: __ccgo_ts + 14492,
		Fcodepoints: [2]uint32{
			0: uint32(9136),
		},
	},
	1277: {
		Fname: __ccgo_ts + 14505,
		Fcodepoints: [2]uint32{
			0: uint32(8808),
		},
	},
	1278: {
		Fname: __ccgo_ts + 14511,
		Fcodepoints: [2]uint32{
			0: uint32(10889),
		},
	},
	1279: {
		Fname: __ccgo_ts + 14518,
		Fcodepoints: [2]uint32{
			0: uint32(10889),
		},
	},
	1280: {
		Fname: __ccgo_ts + 14529,
		Fcodepoints: [2]uint32{
			0: uint32(10887),
		},
	},
	1281: {
		Fname: __ccgo_ts + 14535,
		Fcodepoints: [2]uint32{
			0: uint32(10887),
		},
	},
	1282: {
		Fname: __ccgo_ts + 14542,
		Fcodepoints: [2]uint32{
			0: uint32(8808),
		},
	},
	1283: {
		Fname: __ccgo_ts + 14550,
		Fcodepoints: [2]uint32{
			0: uint32(8934),
		},
	},
	1284: {
		Fname: __ccgo_ts + 14558,
		Fcodepoints: [2]uint32{
			0: uint32(10220),
		},
	},
	1285: {
		Fname: __ccgo_ts + 14566,
		Fcodepoints: [2]uint32{
			0: uint32(8701),
		},
	},
	1286: {
		Fname: __ccgo_ts + 14574,
		Fcodepoints: [2]uint32{
			0: uint32(10214),
		},
	},
	1287: {
		Fname: __ccgo_ts + 14582,
		Fcodepoints: [2]uint32{
			0: uint32(10229),
		},
	},
	1288: {
		Fname: __ccgo_ts + 14598,
		Fcodepoints: [2]uint32{
			0: uint32(10231),
		},
	},
	1289: {
		Fname: __ccgo_ts + 14619,
		Fcodepoints: [2]uint32{
			0: uint32(10236),
		},
	},
	1290: {
		Fname: __ccgo_ts + 14632,
		Fcodepoints: [2]uint32{
			0: uint32(10230),
		},
	},
	1291: {
		Fname: __ccgo_ts + 14649,
		Fcodepoints: [2]uint32{
			0: uint32(8619),
		},
	},
	1292: {
		Fname: __ccgo_ts + 14665,
		Fcodepoints: [2]uint32{
			0: uint32(8620),
		},
	},
	1293: {
		Fname: __ccgo_ts + 14682,
		Fcodepoints: [2]uint32{
			0: uint32(10629),
		},
	},
	1294: {
		Fname: __ccgo_ts + 14690,
		Fcodepoints: [2]uint32{
			0: uint32(120157),
		},
	},
	1295: {
		Fname: __ccgo_ts + 14697,
		Fcodepoints: [2]uint32{
			0: uint32(10797),
		},
	},
	1296: {
		Fname: __ccgo_ts + 14706,
		Fcodepoints: [2]uint32{
			0: uint32(10804),
		},
	},
	1297: {
		Fname: __ccgo_ts + 14716,
		Fcodepoints: [2]uint32{
			0: uint32(8727),
		},
	},
	1298: {
		Fname: __ccgo_ts + 14725,
		Fcodepoints: [2]uint32{
			0: uint32(95),
		},
	},
	1299: {
		Fname: __ccgo_ts + 14734,
		Fcodepoints: [2]uint32{
			0: uint32(9674),
		},
	},
	1300: {
		Fname: __ccgo_ts + 14740,
		Fcodepoints: [2]uint32{
			0: uint32(9674),
		},
	},
	1301: {
		Fname: __ccgo_ts + 14750,
		Fcodepoints: [2]uint32{
			0: uint32(10731),
		},
	},
	1302: {
		Fname: __ccgo_ts + 14757,
		Fcodepoints: [2]uint32{
			0: uint32(40),
		},
	},
	1303: {
		Fname: __ccgo_ts + 14764,
		Fcodepoints: [2]uint32{
			0: uint32(10643),
		},
	},
	1304: {
		Fname: __ccgo_ts + 14773,
		Fcodepoints: [2]uint32{
			0: uint32(8646),
		},
	},
	1305: {
		Fname: __ccgo_ts + 14781,
		Fcodepoints: [2]uint32{
			0: uint32(8991),
		},
	},
	1306: {
		Fname: __ccgo_ts + 14792,
		Fcodepoints: [2]uint32{
			0: uint32(8651),
		},
	},
	1307: {
		Fname: __ccgo_ts + 14800,
		Fcodepoints: [2]uint32{
			0: uint32(10605),
		},
	},
	1308: {
		Fname: __ccgo_ts + 14809,
		Fcodepoints: [2]uint32{
			0: uint32(8206),
		},
	},
	1309: {
		Fname: __ccgo_ts + 14815,
		Fcodepoints: [2]uint32{
			0: uint32(8895),
		},
	},
	1310: {
		Fname: __ccgo_ts + 14823,
		Fcodepoints: [2]uint32{
			0: uint32(8249),
		},
	},
	1311: {
		Fname: __ccgo_ts + 14832,
		Fcodepoints: [2]uint32{
			0: uint32(120001),
		},
	},
	1312: {
		Fname: __ccgo_ts + 14839,
		Fcodepoints: [2]uint32{
			0: uint32(8624),
		},
	},
	1313: {
		Fname: __ccgo_ts + 14845,
		Fcodepoints: [2]uint32{
			0: uint32(8818),
		},
	},
	1314: {
		Fname: __ccgo_ts + 14852,
		Fcodepoints: [2]uint32{
			0: uint32(10893),
		},
	},
	1315: {
		Fname: __ccgo_ts + 14860,
		Fcodepoints: [2]uint32{
			0: uint32(10895),
		},
	},
	1316: {
		Fname: __ccgo_ts + 14868,
		Fcodepoints: [2]uint32{
			0: uint32(91),
		},
	},
	1317: {
		Fname: __ccgo_ts + 14875,
		Fcodepoints: [2]uint32{
			0: uint32(8216),
		},
	},
	1318: {
		Fname: __ccgo_ts + 14883,
		Fcodepoints: [2]uint32{
			0: uint32(8218),
		},
	},
	1319: {
		Fname: __ccgo_ts + 14892,
		Fcodepoints: [2]uint32{
			0: uint32(322),
		},
	},
	1320: {
		Fname: __ccgo_ts + 843,
		Fcodepoints: [2]uint32{
			0: uint32(60),
		},
	},
	1321: {
		Fname: __ccgo_ts + 14901,
		Fcodepoints: [2]uint32{
			0: uint32(10918),
		},
	},
	1322: {
		Fname: __ccgo_ts + 14908,
		Fcodepoints: [2]uint32{
			0: uint32(10873),
		},
	},
	1323: {
		Fname: __ccgo_ts + 14916,
		Fcodepoints: [2]uint32{
			0: uint32(8918),
		},
	},
	1324: {
		Fname: __ccgo_ts + 14924,
		Fcodepoints: [2]uint32{
			0: uint32(8907),
		},
	},
	1325: {
		Fname: __ccgo_ts + 14933,
		Fcodepoints: [2]uint32{
			0: uint32(8905),
		},
	},
	1326: {
		Fname: __ccgo_ts + 14942,
		Fcodepoints: [2]uint32{
			0: uint32(10614),
		},
	},
	1327: {
		Fname: __ccgo_ts + 14951,
		Fcodepoints: [2]uint32{
			0: uint32(10875),
		},
	},
	1328: {
		Fname: __ccgo_ts + 14961,
		Fcodepoints: [2]uint32{
			0: uint32(10646),
		},
	},
	1329: {
		Fname: __ccgo_ts + 14970,
		Fcodepoints: [2]uint32{
			0: uint32(9667),
		},
	},
	1330: {
		Fname: __ccgo_ts + 14977,
		Fcodepoints: [2]uint32{
			0: uint32(8884),
		},
	},
	1331: {
		Fname: __ccgo_ts + 14985,
		Fcodepoints: [2]uint32{
			0: uint32(9666),
		},
	},
	1332: {
		Fname: __ccgo_ts + 14993,
		Fcodepoints: [2]uint32{
			0: uint32(10570),
		},
	},
	1333: {
		Fname: __ccgo_ts + 15004,
		Fcodepoints: [2]uint32{
			0: uint32(10598),
		},
	},
	1334: {
		Fname: __ccgo_ts + 15014,
		Fcodepoints: [2]uint32{
			0: uint32(8808),
			1: uint32(65024),
		},
	},
	1335: {
		Fname: __ccgo_ts + 15026,
		Fcodepoints: [2]uint32{
			0: uint32(8808),
			1: uint32(65024),
		},
	},
	1336: {
		Fname: __ccgo_ts + 15033,
		Fcodepoints: [2]uint32{
			0: uint32(8762),
		},
	},
	1337: {
		Fname: __ccgo_ts + 15041,
		Fcodepoints: [2]uint32{
			0: uint32(175),
		},
	},
	1338: {
		Fname: __ccgo_ts + 15048,
		Fcodepoints: [2]uint32{
			0: uint32(9794),
		},
	},
	1339: {
		Fname: __ccgo_ts + 15055,
		Fcodepoints: [2]uint32{
			0: uint32(10016),
		},
	},
	1340: {
		Fname: __ccgo_ts + 15062,
		Fcodepoints: [2]uint32{
			0: uint32(10016),
		},
	},
	1341: {
		Fname: __ccgo_ts + 15072,
		Fcodepoints: [2]uint32{
			0: uint32(8614),
		},
	},
	1342: {
		Fname: __ccgo_ts + 15078,
		Fcodepoints: [2]uint32{
			0: uint32(8614),
		},
	},
	1343: {
		Fname: __ccgo_ts + 15087,
		Fcodepoints: [2]uint32{
			0: uint32(8615),
		},
	},
	1344: {
		Fname: __ccgo_ts + 15100,
		Fcodepoints: [2]uint32{
			0: uint32(8612),
		},
	},
	1345: {
		Fname: __ccgo_ts + 15113,
		Fcodepoints: [2]uint32{
			0: uint32(8613),
		},
	},
	1346: {
		Fname: __ccgo_ts + 15124,
		Fcodepoints: [2]uint32{
			0: uint32(9646),
		},
	},
	1347: {
		Fname: __ccgo_ts + 15133,
		Fcodepoints: [2]uint32{
			0: uint32(10793),
		},
	},
	1348: {
		Fname: __ccgo_ts + 15142,
		Fcodepoints: [2]uint32{
			0: uint32(1084),
		},
	},
	1349: {
		Fname: __ccgo_ts + 15148,
		Fcodepoints: [2]uint32{
			0: uint32(8212),
		},
	},
	1350: {
		Fname: __ccgo_ts + 15156,
		Fcodepoints: [2]uint32{
			0: uint32(8737),
		},
	},
	1351: {
		Fname: __ccgo_ts + 15172,
		Fcodepoints: [2]uint32{
			0: uint32(120106),
		},
	},
	1352: {
		Fname: __ccgo_ts + 15178,
		Fcodepoints: [2]uint32{
			0: uint32(8487),
		},
	},
	1353: {
		Fname: __ccgo_ts + 15184,
		Fcodepoints: [2]uint32{
			0: uint32(181),
		},
	},
	1354: {
		Fname: __ccgo_ts + 15192,
		Fcodepoints: [2]uint32{
			0: uint32(8739),
		},
	},
	1355: {
		Fname: __ccgo_ts + 15198,
		Fcodepoints: [2]uint32{
			0: uint32(42),
		},
	},
	1356: {
		Fname: __ccgo_ts + 15207,
		Fcodepoints: [2]uint32{
			0: uint32(10992),
		},
	},
	1357: {
		Fname: __ccgo_ts + 15216,
		Fcodepoints: [2]uint32{
			0: uint32(183),
		},
	},
	1358: {
		Fname: __ccgo_ts + 15225,
		Fcodepoints: [2]uint32{
			0: uint32(8722),
		},
	},
	1359: {
		Fname: __ccgo_ts + 15233,
		Fcodepoints: [2]uint32{
			0: uint32(8863),
		},
	},
	1360: {
		Fname: __ccgo_ts + 15242,
		Fcodepoints: [2]uint32{
			0: uint32(8760),
		},
	},
	1361: {
		Fname: __ccgo_ts + 15251,
		Fcodepoints: [2]uint32{
			0: uint32(10794),
		},
	},
	1362: {
		Fname: __ccgo_ts + 15261,
		Fcodepoints: [2]uint32{
			0: uint32(10971),
		},
	},
	1363: {
		Fname: __ccgo_ts + 15268,
		Fcodepoints: [2]uint32{
			0: uint32(8230),
		},
	},
	1364: {
		Fname: __ccgo_ts + 15275,
		Fcodepoints: [2]uint32{
			0: uint32(8723),
		},
	},
	1365: {
		Fname: __ccgo_ts + 15284,
		Fcodepoints: [2]uint32{
			0: uint32(8871),
		},
	},
	1366: {
		Fname: __ccgo_ts + 15293,
		Fcodepoints: [2]uint32{
			0: uint32(120158),
		},
	},
	1367: {
		Fname: __ccgo_ts + 15300,
		Fcodepoints: [2]uint32{
			0: uint32(8723),
		},
	},
	1368: {
		Fname: __ccgo_ts + 15305,
		Fcodepoints: [2]uint32{
			0: uint32(120002),
		},
	},
	1369: {
		Fname: __ccgo_ts + 15312,
		Fcodepoints: [2]uint32{
			0: uint32(8766),
		},
	},
	1370: {
		Fname: __ccgo_ts + 15321,
		Fcodepoints: [2]uint32{
			0: uint32(956),
		},
	},
	1371: {
		Fname: __ccgo_ts + 15326,
		Fcodepoints: [2]uint32{
			0: uint32(8888),
		},
	},
	1372: {
		Fname: __ccgo_ts + 15337,
		Fcodepoints: [2]uint32{
			0: uint32(8888),
		},
	},
	1373: {
		Fname: __ccgo_ts + 15345,
		Fcodepoints: [2]uint32{
			0: uint32(8921),
			1: uint32(824),
		},
	},
	1374: {
		Fname: __ccgo_ts + 15351,
		Fcodepoints: [2]uint32{
			0: uint32(8811),
			1: uint32(8402),
		},
	},
	1375: {
		Fname: __ccgo_ts + 15357,
		Fcodepoints: [2]uint32{
			0: uint32(8811),
			1: uint32(824),
		},
	},
	1376: {
		Fname: __ccgo_ts + 15364,
		Fcodepoints: [2]uint32{
			0: uint32(8653),
		},
	},
	1377: {
		Fname: __ccgo_ts + 15377,
		Fcodepoints: [2]uint32{
			0: uint32(8654),
		},
	},
	1378: {
		Fname: __ccgo_ts + 15395,
		Fcodepoints: [2]uint32{
			0: uint32(8920),
			1: uint32(824),
		},
	},
	1379: {
		Fname: __ccgo_ts + 15401,
		Fcodepoints: [2]uint32{
			0: uint32(8810),
			1: uint32(8402),
		},
	},
	1380: {
		Fname: __ccgo_ts + 15407,
		Fcodepoints: [2]uint32{
			0: uint32(8810),
			1: uint32(824),
		},
	},
	1381: {
		Fname: __ccgo_ts + 15414,
		Fcodepoints: [2]uint32{
			0: uint32(8655),
		},
	},
	1382: {
		Fname: __ccgo_ts + 15428,
		Fcodepoints: [2]uint32{
			0: uint32(8879),
		},
	},
	1383: {
		Fname: __ccgo_ts + 15437,
		Fcodepoints: [2]uint32{
			0: uint32(8878),
		},
	},
	1384: {
		Fname: __ccgo_ts + 15446,
		Fcodepoints: [2]uint32{
			0: uint32(8711),
		},
	},
	1385: {
		Fname: __ccgo_ts + 15454,
		Fcodepoints: [2]uint32{
			0: uint32(324),
		},
	},
	1386: {
		Fname: __ccgo_ts + 15463,
		Fcodepoints: [2]uint32{
			0: uint32(8736),
			1: uint32(8402),
		},
	},
	1387: {
		Fname: __ccgo_ts + 15470,
		Fcodepoints: [2]uint32{
			0: uint32(8777),
		},
	},
	1388: {
		Fname: __ccgo_ts + 15476,
		Fcodepoints: [2]uint32{
			0: uint32(10864),
			1: uint32(824),
		},
	},
	1389: {
		Fname: __ccgo_ts + 15483,
		Fcodepoints: [2]uint32{
			0: uint32(8779),
			1: uint32(824),
		},
	},
	1390: {
		Fname: __ccgo_ts + 15491,
		Fcodepoints: [2]uint32{
			0: uint32(329),
		},
	},
	1391: {
		Fname: __ccgo_ts + 15499,
		Fcodepoints: [2]uint32{
			0: uint32(8777),
		},
	},
	1392: {
		Fname: __ccgo_ts + 15509,
		Fcodepoints: [2]uint32{
			0: uint32(9838),
		},
	},
	1393: {
		Fname: __ccgo_ts + 15517,
		Fcodepoints: [2]uint32{
			0: uint32(9838),
		},
	},
	1394: {
		Fname: __ccgo_ts + 15527,
		Fcodepoints: [2]uint32{
			0: uint32(8469),
		},
	},
	1395: {
		Fname: __ccgo_ts + 15538,
		Fcodepoints: [2]uint32{
			0: uint32(160),
		},
	},
	1396: {
		Fname: __ccgo_ts + 15545,
		Fcodepoints: [2]uint32{
			0: uint32(8782),
			1: uint32(824),
		},
	},
	1397: {
		Fname: __ccgo_ts + 15553,
		Fcodepoints: [2]uint32{
			0: uint32(8783),
			1: uint32(824),
		},
	},
	1398: {
		Fname: __ccgo_ts + 15562,
		Fcodepoints: [2]uint32{
			0: uint32(10819),
		},
	},
	1399: {
		Fname: __ccgo_ts + 15569,
		Fcodepoints: [2]uint32{
			0: uint32(328),
		},
	},
	1400: {
		Fname: __ccgo_ts + 15578,
		Fcodepoints: [2]uint32{
			0: uint32(326),
		},
	},
	1401: {
		Fname: __ccgo_ts + 15587,
		Fcodepoints: [2]uint32{
			0: uint32(8775),
		},
	},
	1402: {
		Fname: __ccgo_ts + 15595,
		Fcodepoints: [2]uint32{
			0: uint32(10861),
			1: uint32(824),
		},
	},
	1403: {
		Fname: __ccgo_ts + 15606,
		Fcodepoints: [2]uint32{
			0: uint32(10818),
		},
	},
	1404: {
		Fname: __ccgo_ts + 15613,
		Fcodepoints: [2]uint32{
			0: uint32(1085),
		},
	},
	1405: {
		Fname: __ccgo_ts + 15619,
		Fcodepoints: [2]uint32{
			0: uint32(8211),
		},
	},
	1406: {
		Fname: __ccgo_ts + 15627,
		Fcodepoints: [2]uint32{
			0: uint32(8800),
		},
	},
	1407: {
		Fname: __ccgo_ts + 15632,
		Fcodepoints: [2]uint32{
			0: uint32(8663),
		},
	},
	1408: {
		Fname: __ccgo_ts + 15640,
		Fcodepoints: [2]uint32{
			0: uint32(10532),
		},
	},
	1409: {
		Fname: __ccgo_ts + 15649,
		Fcodepoints: [2]uint32{
			0: uint32(8599),
		},
	},
	1410: {
		Fname: __ccgo_ts + 15657,
		Fcodepoints: [2]uint32{
			0: uint32(8599),
		},
	},
	1411: {
		Fname: __ccgo_ts + 15667,
		Fcodepoints: [2]uint32{
			0: uint32(8784),
			1: uint32(824),
		},
	},
	1412: {
		Fname: __ccgo_ts + 15675,
		Fcodepoints: [2]uint32{
			0: uint32(8802),
		},
	},
	1413: {
		Fname: __ccgo_ts + 15684,
		Fcodepoints: [2]uint32{
			0: uint32(10536),
		},
	},
	1414: {
		Fname: __ccgo_ts + 15693,
		Fcodepoints: [2]uint32{
			0: uint32(8770),
			1: uint32(824),
		},
	},
	1415: {
		Fname: __ccgo_ts + 15701,
		Fcodepoints: [2]uint32{
			0: uint32(8708),
		},
	},
	1416: {
		Fname: __ccgo_ts + 15710,
		Fcodepoints: [2]uint32{
			0: uint32(8708),
		},
	},
	1417: {
		Fname: __ccgo_ts + 15720,
		Fcodepoints: [2]uint32{
			0: uint32(120107),
		},
	},
	1418: {
		Fname: __ccgo_ts + 15726,
		Fcodepoints: [2]uint32{
			0: uint32(8807),
			1: uint32(824),
		},
	},
	1419: {
		Fname: __ccgo_ts + 15732,
		Fcodepoints: [2]uint32{
			0: uint32(8817),
		},
	},
	1420: {
		Fname: __ccgo_ts + 15738,
		Fcodepoints: [2]uint32{
			0: uint32(8817),
		},
	},
	1421: {
		Fname: __ccgo_ts + 15745,
		Fcodepoints: [2]uint32{
			0: uint32(8807),
			1: uint32(824),
		},
	},
	1422: {
		Fname: __ccgo_ts + 15753,
		Fcodepoints: [2]uint32{
			0: uint32(10878),
			1: uint32(824),
		},
	},
	1423: {
		Fname: __ccgo_ts + 15765,
		Fcodepoints: [2]uint32{
			0: uint32(10878),
			1: uint32(824),
		},
	},
	1424: {
		Fname: __ccgo_ts + 15772,
		Fcodepoints: [2]uint32{
			0: uint32(8821),
		},
	},
	1425: {
		Fname: __ccgo_ts + 15780,
		Fcodepoints: [2]uint32{
			0: uint32(8815),
		},
	},
	1426: {
		Fname: __ccgo_ts + 15786,
		Fcodepoints: [2]uint32{
			0: uint32(8815),
		},
	},
	1427: {
		Fname: __ccgo_ts + 15793,
		Fcodepoints: [2]uint32{
			0: uint32(8654),
		},
	},
	1428: {
		Fname: __ccgo_ts + 15801,
		Fcodepoints: [2]uint32{
			0: uint32(8622),
		},
	},
	1429: {
		Fname: __ccgo_ts + 15809,
		Fcodepoints: [2]uint32{
			0: uint32(10994),
		},
	},
	1430: {
		Fname: __ccgo_ts + 15817,
		Fcodepoints: [2]uint32{
			0: uint32(8715),
		},
	},
	1431: {
		Fname: __ccgo_ts + 15822,
		Fcodepoints: [2]uint32{
			0: uint32(8956),
		},
	},
	1432: {
		Fname: __ccgo_ts + 15828,
		Fcodepoints: [2]uint32{
			0: uint32(8954),
		},
	},
	1433: {
		Fname: __ccgo_ts + 15835,
		Fcodepoints: [2]uint32{
			0: uint32(8715),
		},
	},
	1434: {
		Fname: __ccgo_ts + 15841,
		Fcodepoints: [2]uint32{
			0: uint32(1114),
		},
	},
	1435: {
		Fname: __ccgo_ts + 15848,
		Fcodepoints: [2]uint32{
			0: uint32(8653),
		},
	},
	1436: {
		Fname: __ccgo_ts + 15856,
		Fcodepoints: [2]uint32{
			0: uint32(8806),
			1: uint32(824),
		},
	},
	1437: {
		Fname: __ccgo_ts + 15862,
		Fcodepoints: [2]uint32{
			0: uint32(8602),
		},
	},
	1438: {
		Fname: __ccgo_ts + 15870,
		Fcodepoints: [2]uint32{
			0: uint32(8229),
		},
	},
	1439: {
		Fname: __ccgo_ts + 15877,
		Fcodepoints: [2]uint32{
			0: uint32(8816),
		},
	},
	1440: {
		Fname: __ccgo_ts + 15883,
		Fcodepoints: [2]uint32{
			0: uint32(8602),
		},
	},
	1441: {
		Fname: __ccgo_ts + 15896,
		Fcodepoints: [2]uint32{
			0: uint32(8622),
		},
	},
	1442: {
		Fname: __ccgo_ts + 15914,
		Fcodepoints: [2]uint32{
			0: uint32(8816),
		},
	},
	1443: {
		Fname: __ccgo_ts + 15921,
		Fcodepoints: [2]uint32{
			0: uint32(8806),
			1: uint32(824),
		},
	},
	1444: {
		Fname: __ccgo_ts + 15929,
		Fcodepoints: [2]uint32{
			0: uint32(10877),
			1: uint32(824),
		},
	},
	1445: {
		Fname: __ccgo_ts + 15941,
		Fcodepoints: [2]uint32{
			0: uint32(10877),
			1: uint32(824),
		},
	},
	1446: {
		Fname: __ccgo_ts + 15948,
		Fcodepoints: [2]uint32{
			0: uint32(8814),
		},
	},
	1447: {
		Fname: __ccgo_ts + 15956,
		Fcodepoints: [2]uint32{
			0: uint32(8820),
		},
	},
	1448: {
		Fname: __ccgo_ts + 15964,
		Fcodepoints: [2]uint32{
			0: uint32(8814),
		},
	},
	1449: {
		Fname: __ccgo_ts + 15970,
		Fcodepoints: [2]uint32{
			0: uint32(8938),
		},
	},
	1450: {
		Fname: __ccgo_ts + 15978,
		Fcodepoints: [2]uint32{
			0: uint32(8940),
		},
	},
	1451: {
		Fname: __ccgo_ts + 15987,
		Fcodepoints: [2]uint32{
			0: uint32(8740),
		},
	},
	1452: {
		Fname: __ccgo_ts + 15994,
		Fcodepoints: [2]uint32{
			0: uint32(120159),
		},
	},
	1453: {
		Fname: __ccgo_ts + 16001,
		Fcodepoints: [2]uint32{
			0: uint32(172),
		},
	},
	1454: {
		Fname: __ccgo_ts + 16007,
		Fcodepoints: [2]uint32{
			0: uint32(8713),
		},
	},
	1455: {
		Fname: __ccgo_ts + 16015,
		Fcodepoints: [2]uint32{
			0: uint32(8953),
			1: uint32(824),
		},
	},
	1456: {
		Fname: __ccgo_ts + 16024,
		Fcodepoints: [2]uint32{
			0: uint32(8949),
			1: uint32(824),
		},
	},
	1457: {
		Fname: __ccgo_ts + 16035,
		Fcodepoints: [2]uint32{
			0: uint32(8713),
		},
	},
	1458: {
		Fname: __ccgo_ts + 16045,
		Fcodepoints: [2]uint32{
			0: uint32(8951),
		},
	},
	1459: {
		Fname: __ccgo_ts + 16055,
		Fcodepoints: [2]uint32{
			0: uint32(8950),
		},
	},
	1460: {
		Fname: __ccgo_ts + 16065,
		Fcodepoints: [2]uint32{
			0: uint32(8716),
		},
	},
	1461: {
		Fname: __ccgo_ts + 16073,
		Fcodepoints: [2]uint32{
			0: uint32(8716),
		},
	},
	1462: {
		Fname: __ccgo_ts + 16083,
		Fcodepoints: [2]uint32{
			0: uint32(8958),
		},
	},
	1463: {
		Fname: __ccgo_ts + 16093,
		Fcodepoints: [2]uint32{
			0: uint32(8957),
		},
	},
	1464: {
		Fname: __ccgo_ts + 16103,
		Fcodepoints: [2]uint32{
			0: uint32(8742),
		},
	},
	1465: {
		Fname: __ccgo_ts + 16110,
		Fcodepoints: [2]uint32{
			0: uint32(8742),
		},
	},
	1466: {
		Fname: __ccgo_ts + 16122,
		Fcodepoints: [2]uint32{
			0: uint32(11005),
			1: uint32(8421),
		},
	},
	1467: {
		Fname: __ccgo_ts + 16131,
		Fcodepoints: [2]uint32{
			0: uint32(8706),
			1: uint32(824),
		},
	},
	1468: {
		Fname: __ccgo_ts + 16139,
		Fcodepoints: [2]uint32{
			0: uint32(10772),
		},
	},
	1469: {
		Fname: __ccgo_ts + 16149,
		Fcodepoints: [2]uint32{
			0: uint32(8832),
		},
	},
	1470: {
		Fname: __ccgo_ts + 16155,
		Fcodepoints: [2]uint32{
			0: uint32(8928),
		},
	},
	1471: {
		Fname: __ccgo_ts + 16164,
		Fcodepoints: [2]uint32{
			0: uint32(10927),
			1: uint32(824),
		},
	},
	1472: {
		Fname: __ccgo_ts + 16171,
		Fcodepoints: [2]uint32{
			0: uint32(8832),
		},
	},
	1473: {
		Fname: __ccgo_ts + 16179,
		Fcodepoints: [2]uint32{
			0: uint32(10927),
			1: uint32(824),
		},
	},
	1474: {
		Fname: __ccgo_ts + 16189,
		Fcodepoints: [2]uint32{
			0: uint32(8655),
		},
	},
	1475: {
		Fname: __ccgo_ts + 16197,
		Fcodepoints: [2]uint32{
			0: uint32(8603),
		},
	},
	1476: {
		Fname: __ccgo_ts + 16205,
		Fcodepoints: [2]uint32{
			0: uint32(10547),
			1: uint32(824),
		},
	},
	1477: {
		Fname: __ccgo_ts + 16214,
		Fcodepoints: [2]uint32{
			0: uint32(8605),
			1: uint32(824),
		},
	},
	1478: {
		Fname: __ccgo_ts + 16223,
		Fcodepoints: [2]uint32{
			0: uint32(8603),
		},
	},
	1479: {
		Fname: __ccgo_ts + 16237,
		Fcodepoints: [2]uint32{
			0: uint32(8939),
		},
	},
	1480: {
		Fname: __ccgo_ts + 16245,
		Fcodepoints: [2]uint32{
			0: uint32(8941),
		},
	},
	1481: {
		Fname: __ccgo_ts + 16254,
		Fcodepoints: [2]uint32{
			0: uint32(8833),
		},
	},
	1482: {
		Fname: __ccgo_ts + 16260,
		Fcodepoints: [2]uint32{
			0: uint32(8929),
		},
	},
	1483: {
		Fname: __ccgo_ts + 16269,
		Fcodepoints: [2]uint32{
			0: uint32(10928),
			1: uint32(824),
		},
	},
	1484: {
		Fname: __ccgo_ts + 16276,
		Fcodepoints: [2]uint32{
			0: uint32(120003),
		},
	},
	1485: {
		Fname: __ccgo_ts + 16283,
		Fcodepoints: [2]uint32{
			0: uint32(8740),
		},
	},
	1486: {
		Fname: __ccgo_ts + 16295,
		Fcodepoints: [2]uint32{
			0: uint32(8742),
		},
	},
	1487: {
		Fname: __ccgo_ts + 16312,
		Fcodepoints: [2]uint32{
			0: uint32(8769),
		},
	},
	1488: {
		Fname: __ccgo_ts + 16319,
		Fcodepoints: [2]uint32{
			0: uint32(8772),
		},
	},
	1489: {
		Fname: __ccgo_ts + 16327,
		Fcodepoints: [2]uint32{
			0: uint32(8772),
		},
	},
	1490: {
		Fname: __ccgo_ts + 16336,
		Fcodepoints: [2]uint32{
			0: uint32(8740),
		},
	},
	1491: {
		Fname: __ccgo_ts + 16344,
		Fcodepoints: [2]uint32{
			0: uint32(8742),
		},
	},
	1492: {
		Fname: __ccgo_ts + 16352,
		Fcodepoints: [2]uint32{
			0: uint32(8930),
		},
	},
	1493: {
		Fname: __ccgo_ts + 16362,
		Fcodepoints: [2]uint32{
			0: uint32(8931),
		},
	},
	1494: {
		Fname: __ccgo_ts + 16372,
		Fcodepoints: [2]uint32{
			0: uint32(8836),
		},
	},
	1495: {
		Fname: __ccgo_ts + 16379,
		Fcodepoints: [2]uint32{
			0: uint32(10949),
			1: uint32(824),
		},
	},
	1496: {
		Fname: __ccgo_ts + 16387,
		Fcodepoints: [2]uint32{
			0: uint32(8840),
		},
	},
	1497: {
		Fname: __ccgo_ts + 16395,
		Fcodepoints: [2]uint32{
			0: uint32(8834),
			1: uint32(8402),
		},
	},
	1498: {
		Fname: __ccgo_ts + 16405,
		Fcodepoints: [2]uint32{
			0: uint32(8840),
		},
	},
	1499: {
		Fname: __ccgo_ts + 16417,
		Fcodepoints: [2]uint32{
			0: uint32(10949),
			1: uint32(824),
		},
	},
	1500: {
		Fname: __ccgo_ts + 16430,
		Fcodepoints: [2]uint32{
			0: uint32(8833),
		},
	},
	1501: {
		Fname: __ccgo_ts + 16438,
		Fcodepoints: [2]uint32{
			0: uint32(10928),
			1: uint32(824),
		},
	},
	1502: {
		Fname: __ccgo_ts + 16448,
		Fcodepoints: [2]uint32{
			0: uint32(8837),
		},
	},
	1503: {
		Fname: __ccgo_ts + 16455,
		Fcodepoints: [2]uint32{
			0: uint32(10950),
			1: uint32(824),
		},
	},
	1504: {
		Fname: __ccgo_ts + 16463,
		Fcodepoints: [2]uint32{
			0: uint32(8841),
		},
	},
	1505: {
		Fname: __ccgo_ts + 16471,
		Fcodepoints: [2]uint32{
			0: uint32(8835),
			1: uint32(8402),
		},
	},
	1506: {
		Fname: __ccgo_ts + 16481,
		Fcodepoints: [2]uint32{
			0: uint32(8841),
		},
	},
	1507: {
		Fname: __ccgo_ts + 16493,
		Fcodepoints: [2]uint32{
			0: uint32(10950),
			1: uint32(824),
		},
	},
	1508: {
		Fname: __ccgo_ts + 16506,
		Fcodepoints: [2]uint32{
			0: uint32(8825),
		},
	},
	1509: {
		Fname: __ccgo_ts + 16513,
		Fcodepoints: [2]uint32{
			0: uint32(241),
		},
	},
	1510: {
		Fname: __ccgo_ts + 16522,
		Fcodepoints: [2]uint32{
			0: uint32(8824),
		},
	},
	1511: {
		Fname: __ccgo_ts + 16529,
		Fcodepoints: [2]uint32{
			0: uint32(8938),
		},
	},
	1512: {
		Fname: __ccgo_ts + 16545,
		Fcodepoints: [2]uint32{
			0: uint32(8940),
		},
	},
	1513: {
		Fname: __ccgo_ts + 16563,
		Fcodepoints: [2]uint32{
			0: uint32(8939),
		},
	},
	1514: {
		Fname: __ccgo_ts + 16580,
		Fcodepoints: [2]uint32{
			0: uint32(8941),
		},
	},
	1515: {
		Fname: __ccgo_ts + 16599,
		Fcodepoints: [2]uint32{
			0: uint32(957),
		},
	},
	1516: {
		Fname: __ccgo_ts + 16604,
		Fcodepoints: [2]uint32{
			0: uint32(35),
		},
	},
	1517: {
		Fname: __ccgo_ts + 16610,
		Fcodepoints: [2]uint32{
			0: uint32(8470),
		},
	},
	1518: {
		Fname: __ccgo_ts + 16619,
		Fcodepoints: [2]uint32{
			0: uint32(8199),
		},
	},
	1519: {
		Fname: __ccgo_ts + 16627,
		Fcodepoints: [2]uint32{
			0: uint32(8877),
		},
	},
	1520: {
		Fname: __ccgo_ts + 16636,
		Fcodepoints: [2]uint32{
			0: uint32(10500),
		},
	},
	1521: {
		Fname: __ccgo_ts + 16645,
		Fcodepoints: [2]uint32{
			0: uint32(8781),
			1: uint32(8402),
		},
	},
	1522: {
		Fname: __ccgo_ts + 16652,
		Fcodepoints: [2]uint32{
			0: uint32(8876),
		},
	},
	1523: {
		Fname: __ccgo_ts + 16661,
		Fcodepoints: [2]uint32{
			0: uint32(8805),
			1: uint32(8402),
		},
	},
	1524: {
		Fname: __ccgo_ts + 16668,
		Fcodepoints: [2]uint32{
			0: uint32(62),
			1: uint32(8402),
		},
	},
	1525: {
		Fname: __ccgo_ts + 16675,
		Fcodepoints: [2]uint32{
			0: uint32(10718),
		},
	},
	1526: {
		Fname: __ccgo_ts + 16685,
		Fcodepoints: [2]uint32{
			0: uint32(10498),
		},
	},
	1527: {
		Fname: __ccgo_ts + 16694,
		Fcodepoints: [2]uint32{
			0: uint32(8804),
			1: uint32(8402),
		},
	},
	1528: {
		Fname: __ccgo_ts + 16701,
		Fcodepoints: [2]uint32{
			0: uint32(60),
			1: uint32(8402),
		},
	},
	1529: {
		Fname: __ccgo_ts + 16708,
		Fcodepoints: [2]uint32{
			0: uint32(8884),
			1: uint32(8402),
		},
	},
	1530: {
		Fname: __ccgo_ts + 16718,
		Fcodepoints: [2]uint32{
			0: uint32(10499),
		},
	},
	1531: {
		Fname: __ccgo_ts + 16727,
		Fcodepoints: [2]uint32{
			0: uint32(8885),
			1: uint32(8402),
		},
	},
	1532: {
		Fname: __ccgo_ts + 16737,
		Fcodepoints: [2]uint32{
			0: uint32(8764),
			1: uint32(8402),
		},
	},
	1533: {
		Fname: __ccgo_ts + 16745,
		Fcodepoints: [2]uint32{
			0: uint32(8662),
		},
	},
	1534: {
		Fname: __ccgo_ts + 16753,
		Fcodepoints: [2]uint32{
			0: uint32(10531),
		},
	},
	1535: {
		Fname: __ccgo_ts + 16762,
		Fcodepoints: [2]uint32{
			0: uint32(8598),
		},
	},
	1536: {
		Fname: __ccgo_ts + 16770,
		Fcodepoints: [2]uint32{
			0: uint32(8598),
		},
	},
	1537: {
		Fname: __ccgo_ts + 16780,
		Fcodepoints: [2]uint32{
			0: uint32(10535),
		},
	},
	1538: {
		Fname: __ccgo_ts + 16789,
		Fcodepoints: [2]uint32{
			0: uint32(9416),
		},
	},
	1539: {
		Fname: __ccgo_ts + 16794,
		Fcodepoints: [2]uint32{
			0: uint32(243),
		},
	},
	1540: {
		Fname: __ccgo_ts + 16803,
		Fcodepoints: [2]uint32{
			0: uint32(8859),
		},
	},
	1541: {
		Fname: __ccgo_ts + 16810,
		Fcodepoints: [2]uint32{
			0: uint32(8858),
		},
	},
	1542: {
		Fname: __ccgo_ts + 16817,
		Fcodepoints: [2]uint32{
			0: uint32(244),
		},
	},
	1543: {
		Fname: __ccgo_ts + 16825,
		Fcodepoints: [2]uint32{
			0: uint32(1086),
		},
	},
	1544: {
		Fname: __ccgo_ts + 16831,
		Fcodepoints: [2]uint32{
			0: uint32(8861),
		},
	},
	1545: {
		Fname: __ccgo_ts + 16839,
		Fcodepoints: [2]uint32{
			0: uint32(337),
		},
	},
	1546: {
		Fname: __ccgo_ts + 16848,
		Fcodepoints: [2]uint32{
			0: uint32(10808),
		},
	},
	1547: {
		Fname: __ccgo_ts + 16855,
		Fcodepoints: [2]uint32{
			0: uint32(8857),
		},
	},
	1548: {
		Fname: __ccgo_ts + 16862,
		Fcodepoints: [2]uint32{
			0: uint32(10684),
		},
	},
	1549: {
		Fname: __ccgo_ts + 16871,
		Fcodepoints: [2]uint32{
			0: uint32(339),
		},
	},
	1550: {
		Fname: __ccgo_ts + 16879,
		Fcodepoints: [2]uint32{
			0: uint32(10687),
		},
	},
	1551: {
		Fname: __ccgo_ts + 16887,
		Fcodepoints: [2]uint32{
			0: uint32(120108),
		},
	},
	1552: {
		Fname: __ccgo_ts + 16893,
		Fcodepoints: [2]uint32{
			0: uint32(731),
		},
	},
	1553: {
		Fname: __ccgo_ts + 16900,
		Fcodepoints: [2]uint32{
			0: uint32(242),
		},
	},
	1554: {
		Fname: __ccgo_ts + 16909,
		Fcodepoints: [2]uint32{
			0: uint32(10689),
		},
	},
	1555: {
		Fname: __ccgo_ts + 16915,
		Fcodepoints: [2]uint32{
			0: uint32(10677),
		},
	},
	1556: {
		Fname: __ccgo_ts + 16923,
		Fcodepoints: [2]uint32{
			0: uint32(937),
		},
	},
	1557: {
		Fname: __ccgo_ts + 16929,
		Fcodepoints: [2]uint32{
			0: uint32(8750),
		},
	},
	1558: {
		Fname: __ccgo_ts + 16936,
		Fcodepoints: [2]uint32{
			0: uint32(8634),
		},
	},
	1559: {
		Fname: __ccgo_ts + 16944,
		Fcodepoints: [2]uint32{
			0: uint32(10686),
		},
	},
	1560: {
		Fname: __ccgo_ts + 16952,
		Fcodepoints: [2]uint32{
			0: uint32(10683),
		},
	},
	1561: {
		Fname: __ccgo_ts + 16962,
		Fcodepoints: [2]uint32{
			0: uint32(8254),
		},
	},
	1562: {
		Fname: __ccgo_ts + 16970,
		Fcodepoints: [2]uint32{
			0: uint32(10688),
		},
	},
	1563: {
		Fname: __ccgo_ts + 16976,
		Fcodepoints: [2]uint32{
			0: uint32(333),
		},
	},
	1564: {
		Fname: __ccgo_ts + 16984,
		Fcodepoints: [2]uint32{
			0: uint32(969),
		},
	},
	1565: {
		Fname: __ccgo_ts + 16992,
		Fcodepoints: [2]uint32{
			0: uint32(959),
		},
	},
	1566: {
		Fname: __ccgo_ts + 17002,
		Fcodepoints: [2]uint32{
			0: uint32(10678),
		},
	},
	1567: {
		Fname: __ccgo_ts + 17009,
		Fcodepoints: [2]uint32{
			0: uint32(8854),
		},
	},
	1568: {
		Fname: __ccgo_ts + 17018,
		Fcodepoints: [2]uint32{
			0: uint32(120160),
		},
	},
	1569: {
		Fname: __ccgo_ts + 17025,
		Fcodepoints: [2]uint32{
			0: uint32(10679),
		},
	},
	1570: {
		Fname: __ccgo_ts + 17032,
		Fcodepoints: [2]uint32{
			0: uint32(10681),
		},
	},
	1571: {
		Fname: __ccgo_ts + 17040,
		Fcodepoints: [2]uint32{
			0: uint32(8853),
		},
	},
	1572: {
		Fname: __ccgo_ts + 17048,
		Fcodepoints: [2]uint32{
			0: uint32(8744),
		},
	},
	1573: {
		Fname: __ccgo_ts + 17053,
		Fcodepoints: [2]uint32{
			0: uint32(8635),
		},
	},
	1574: {
		Fname: __ccgo_ts + 17061,
		Fcodepoints: [2]uint32{
			0: uint32(10845),
		},
	},
	1575: {
		Fname: __ccgo_ts + 17067,
		Fcodepoints: [2]uint32{
			0: uint32(8500),
		},
	},
	1576: {
		Fname: __ccgo_ts + 17075,
		Fcodepoints: [2]uint32{
			0: uint32(8500),
		},
	},
	1577: {
		Fname: __ccgo_ts + 17085,
		Fcodepoints: [2]uint32{
			0: uint32(170),
		},
	},
	1578: {
		Fname: __ccgo_ts + 17092,
		Fcodepoints: [2]uint32{
			0: uint32(186),
		},
	},
	1579: {
		Fname: __ccgo_ts + 17099,
		Fcodepoints: [2]uint32{
			0: uint32(8886),
		},
	},
	1580: {
		Fname: __ccgo_ts + 17108,
		Fcodepoints: [2]uint32{
			0: uint32(10838),
		},
	},
	1581: {
		Fname: __ccgo_ts + 17115,
		Fcodepoints: [2]uint32{
			0: uint32(10839),
		},
	},
	1582: {
		Fname: __ccgo_ts + 17125,
		Fcodepoints: [2]uint32{
			0: uint32(10843),
		},
	},
	1583: {
		Fname: __ccgo_ts + 17131,
		Fcodepoints: [2]uint32{
			0: uint32(8500),
		},
	},
	1584: {
		Fname: __ccgo_ts + 17138,
		Fcodepoints: [2]uint32{
			0: uint32(248),
		},
	},
	1585: {
		Fname: __ccgo_ts + 17147,
		Fcodepoints: [2]uint32{
			0: uint32(8856),
		},
	},
	1586: {
		Fname: __ccgo_ts + 17154,
		Fcodepoints: [2]uint32{
			0: uint32(245),
		},
	},
	1587: {
		Fname: __ccgo_ts + 17163,
		Fcodepoints: [2]uint32{
			0: uint32(8855),
		},
	},
	1588: {
		Fname: __ccgo_ts + 17172,
		Fcodepoints: [2]uint32{
			0: uint32(10806),
		},
	},
	1589: {
		Fname: __ccgo_ts + 17183,
		Fcodepoints: [2]uint32{
			0: uint32(246),
		},
	},
	1590: {
		Fname: __ccgo_ts + 17190,
		Fcodepoints: [2]uint32{
			0: uint32(9021),
		},
	},
	1591: {
		Fname: __ccgo_ts + 17198,
		Fcodepoints: [2]uint32{
			0: uint32(8741),
		},
	},
	1592: {
		Fname: __ccgo_ts + 17204,
		Fcodepoints: [2]uint32{
			0: uint32(182),
		},
	},
	1593: {
		Fname: __ccgo_ts + 17211,
		Fcodepoints: [2]uint32{
			0: uint32(8741),
		},
	},
	1594: {
		Fname: __ccgo_ts + 17222,
		Fcodepoints: [2]uint32{
			0: uint32(10995),
		},
	},
	1595: {
		Fname: __ccgo_ts + 17231,
		Fcodepoints: [2]uint32{
			0: uint32(11005),
		},
	},
	1596: {
		Fname: __ccgo_ts + 17239,
		Fcodepoints: [2]uint32{
			0: uint32(8706),
		},
	},
	1597: {
		Fname: __ccgo_ts + 17246,
		Fcodepoints: [2]uint32{
			0: uint32(1087),
		},
	},
	1598: {
		Fname: __ccgo_ts + 17252,
		Fcodepoints: [2]uint32{
			0: uint32(37),
		},
	},
	1599: {
		Fname: __ccgo_ts + 17261,
		Fcodepoints: [2]uint32{
			0: uint32(46),
		},
	},
	1600: {
		Fname: __ccgo_ts + 17270,
		Fcodepoints: [2]uint32{
			0: uint32(8240),
		},
	},
	1601: {
		Fname: __ccgo_ts + 17279,
		Fcodepoints: [2]uint32{
			0: uint32(8869),
		},
	},
	1602: {
		Fname: __ccgo_ts + 17286,
		Fcodepoints: [2]uint32{
			0: uint32(8241),
		},
	},
	1603: {
		Fname: __ccgo_ts + 17296,
		Fcodepoints: [2]uint32{
			0: uint32(120109),
		},
	},
	1604: {
		Fname: __ccgo_ts + 17302,
		Fcodepoints: [2]uint32{
			0: uint32(966),
		},
	},
	1605: {
		Fname: __ccgo_ts + 17308,
		Fcodepoints: [2]uint32{
			0: uint32(981),
		},
	},
	1606: {
		Fname: __ccgo_ts + 17315,
		Fcodepoints: [2]uint32{
			0: uint32(8499),
		},
	},
	1607: {
		Fname: __ccgo_ts + 17324,
		Fcodepoints: [2]uint32{
			0: uint32(9742),
		},
	},
	1608: {
		Fname: __ccgo_ts + 17332,
		Fcodepoints: [2]uint32{
			0: uint32(960),
		},
	},
	1609: {
		Fname: __ccgo_ts + 17337,
		Fcodepoints: [2]uint32{
			0: uint32(8916),
		},
	},
	1610: {
		Fname: __ccgo_ts + 17349,
		Fcodepoints: [2]uint32{
			0: uint32(982),
		},
	},
	1611: {
		Fname: __ccgo_ts + 17355,
		Fcodepoints: [2]uint32{
			0: uint32(8463),
		},
	},
	1612: {
		Fname: __ccgo_ts + 17364,
		Fcodepoints: [2]uint32{
			0: uint32(8462),
		},
	},
	1613: {
		Fname: __ccgo_ts + 17374,
		Fcodepoints: [2]uint32{
			0: uint32(8463),
		},
	},
	1614: {
		Fname: __ccgo_ts + 17383,
		Fcodepoints: [2]uint32{
			0: uint32(43),
		},
	},
	1615: {
		Fname: __ccgo_ts + 17390,
		Fcodepoints: [2]uint32{
			0: uint32(10787),
		},
	},
	1616: {
		Fname: __ccgo_ts + 17401,
		Fcodepoints: [2]uint32{
			0: uint32(8862),
		},
	},
	1617: {
		Fname: __ccgo_ts + 17409,
		Fcodepoints: [2]uint32{
			0: uint32(10786),
		},
	},
	1618: {
		Fname: __ccgo_ts + 17419,
		Fcodepoints: [2]uint32{
			0: uint32(8724),
		},
	},
	1619: {
		Fname: __ccgo_ts + 17428,
		Fcodepoints: [2]uint32{
			0: uint32(10789),
		},
	},
	1620: {
		Fname: __ccgo_ts + 17437,
		Fcodepoints: [2]uint32{
			0: uint32(10866),
		},
	},
	1621: {
		Fname: __ccgo_ts + 17445,
		Fcodepoints: [2]uint32{
			0: uint32(177),
		},
	},
	1622: {
		Fname: __ccgo_ts + 17454,
		Fcodepoints: [2]uint32{
			0: uint32(10790),
		},
	},
	1623: {
		Fname: __ccgo_ts + 17464,
		Fcodepoints: [2]uint32{
			0: uint32(10791),
		},
	},
	1624: {
		Fname: __ccgo_ts + 17474,
		Fcodepoints: [2]uint32{
			0: uint32(177),
		},
	},
	1625: {
		Fname: __ccgo_ts + 17479,
		Fcodepoints: [2]uint32{
			0: uint32(10773),
		},
	},
	1626: {
		Fname: __ccgo_ts + 17490,
		Fcodepoints: [2]uint32{
			0: uint32(120161),
		},
	},
	1627: {
		Fname: __ccgo_ts + 17497,
		Fcodepoints: [2]uint32{
			0: uint32(163),
		},
	},
	1628: {
		Fname: __ccgo_ts + 17505,
		Fcodepoints: [2]uint32{
			0: uint32(8826),
		},
	},
	1629: {
		Fname: __ccgo_ts + 17510,
		Fcodepoints: [2]uint32{
			0: uint32(10931),
		},
	},
	1630: {
		Fname: __ccgo_ts + 17516,
		Fcodepoints: [2]uint32{
			0: uint32(10935),
		},
	},
	1631: {
		Fname: __ccgo_ts + 17523,
		Fcodepoints: [2]uint32{
			0: uint32(8828),
		},
	},
	1632: {
		Fname: __ccgo_ts + 17531,
		Fcodepoints: [2]uint32{
			0: uint32(10927),
		},
	},
	1633: {
		Fname: __ccgo_ts + 17537,
		Fcodepoints: [2]uint32{
			0: uint32(8826),
		},
	},
	1634: {
		Fname: __ccgo_ts + 17544,
		Fcodepoints: [2]uint32{
			0: uint32(10935),
		},
	},
	1635: {
		Fname: __ccgo_ts + 17557,
		Fcodepoints: [2]uint32{
			0: uint32(8828),
		},
	},
	1636: {
		Fname: __ccgo_ts + 17571,
		Fcodepoints: [2]uint32{
			0: uint32(10927),
		},
	},
	1637: {
		Fname: __ccgo_ts + 17580,
		Fcodepoints: [2]uint32{
			0: uint32(10937),
		},
	},
	1638: {
		Fname: __ccgo_ts + 17594,
		Fcodepoints: [2]uint32{
			0: uint32(10933),
		},
	},
	1639: {
		Fname: __ccgo_ts + 17605,
		Fcodepoints: [2]uint32{
			0: uint32(8936),
		},
	},
	1640: {
		Fname: __ccgo_ts + 17616,
		Fcodepoints: [2]uint32{
			0: uint32(8830),
		},
	},
	1641: {
		Fname: __ccgo_ts + 17626,
		Fcodepoints: [2]uint32{
			0: uint32(8242),
		},
	},
	1642: {
		Fname: __ccgo_ts + 17634,
		Fcodepoints: [2]uint32{
			0: uint32(8473),
		},
	},
	1643: {
		Fname: __ccgo_ts + 17643,
		Fcodepoints: [2]uint32{
			0: uint32(10933),
		},
	},
	1644: {
		Fname: __ccgo_ts + 17650,
		Fcodepoints: [2]uint32{
			0: uint32(10937),
		},
	},
	1645: {
		Fname: __ccgo_ts + 17658,
		Fcodepoints: [2]uint32{
			0: uint32(8936),
		},
	},
	1646: {
		Fname: __ccgo_ts + 17667,
		Fcodepoints: [2]uint32{
			0: uint32(8719),
		},
	},
	1647: {
		Fname: __ccgo_ts + 17674,
		Fcodepoints: [2]uint32{
			0: uint32(9006),
		},
	},
	1648: {
		Fname: __ccgo_ts + 17685,
		Fcodepoints: [2]uint32{
			0: uint32(8978),
		},
	},
	1649: {
		Fname: __ccgo_ts + 17696,
		Fcodepoints: [2]uint32{
			0: uint32(8979),
		},
	},
	1650: {
		Fname: __ccgo_ts + 17707,
		Fcodepoints: [2]uint32{
			0: uint32(8733),
		},
	},
	1651: {
		Fname: __ccgo_ts + 17714,
		Fcodepoints: [2]uint32{
			0: uint32(8733),
		},
	},
	1652: {
		Fname: __ccgo_ts + 17723,
		Fcodepoints: [2]uint32{
			0: uint32(8830),
		},
	},
	1653: {
		Fname: __ccgo_ts + 17731,
		Fcodepoints: [2]uint32{
			0: uint32(8880),
		},
	},
	1654: {
		Fname: __ccgo_ts + 17740,
		Fcodepoints: [2]uint32{
			0: uint32(120005),
		},
	},
	1655: {
		Fname: __ccgo_ts + 17747,
		Fcodepoints: [2]uint32{
			0: uint32(968),
		},
	},
	1656: {
		Fname: __ccgo_ts + 17753,
		Fcodepoints: [2]uint32{
			0: uint32(8200),
		},
	},
	1657: {
		Fname: __ccgo_ts + 17762,
		Fcodepoints: [2]uint32{
			0: uint32(120110),
		},
	},
	1658: {
		Fname: __ccgo_ts + 17768,
		Fcodepoints: [2]uint32{
			0: uint32(10764),
		},
	},
	1659: {
		Fname: __ccgo_ts + 17775,
		Fcodepoints: [2]uint32{
			0: uint32(120162),
		},
	},
	1660: {
		Fname: __ccgo_ts + 17782,
		Fcodepoints: [2]uint32{
			0: uint32(8279),
		},
	},
	1661: {
		Fname: __ccgo_ts + 17791,
		Fcodepoints: [2]uint32{
			0: uint32(120006),
		},
	},
	1662: {
		Fname: __ccgo_ts + 17798,
		Fcodepoints: [2]uint32{
			0: uint32(8461),
		},
	},
	1663: {
		Fname: __ccgo_ts + 17812,
		Fcodepoints: [2]uint32{
			0: uint32(10774),
		},
	},
	1664: {
		Fname: __ccgo_ts + 17822,
		Fcodepoints: [2]uint32{
			0: uint32(63),
		},
	},
	1665: {
		Fname: __ccgo_ts + 17830,
		Fcodepoints: [2]uint32{
			0: uint32(8799),
		},
	},
	1666: {
		Fname: __ccgo_ts + 853,
		Fcodepoints: [2]uint32{
			0: uint32(34),
		},
	},
	1667: {
		Fname: __ccgo_ts + 17840,
		Fcodepoints: [2]uint32{
			0: uint32(8667),
		},
	},
	1668: {
		Fname: __ccgo_ts + 17848,
		Fcodepoints: [2]uint32{
			0: uint32(8658),
		},
	},
	1669: {
		Fname: __ccgo_ts + 17855,
		Fcodepoints: [2]uint32{
			0: uint32(10524),
		},
	},
	1670: {
		Fname: __ccgo_ts + 17864,
		Fcodepoints: [2]uint32{
			0: uint32(10511),
		},
	},
	1671: {
		Fname: __ccgo_ts + 17872,
		Fcodepoints: [2]uint32{
			0: uint32(10596),
		},
	},
	1672: {
		Fname: __ccgo_ts + 17879,
		Fcodepoints: [2]uint32{
			0: uint32(8765),
			1: uint32(817),
		},
	},
	1673: {
		Fname: __ccgo_ts + 17886,
		Fcodepoints: [2]uint32{
			0: uint32(341),
		},
	},
	1674: {
		Fname: __ccgo_ts + 17895,
		Fcodepoints: [2]uint32{
			0: uint32(8730),
		},
	},
	1675: {
		Fname: __ccgo_ts + 17903,
		Fcodepoints: [2]uint32{
			0: uint32(10675),
		},
	},
	1676: {
		Fname: __ccgo_ts + 17914,
		Fcodepoints: [2]uint32{
			0: uint32(10217),
		},
	},
	1677: {
		Fname: __ccgo_ts + 17921,
		Fcodepoints: [2]uint32{
			0: uint32(10642),
		},
	},
	1678: {
		Fname: __ccgo_ts + 17929,
		Fcodepoints: [2]uint32{
			0: uint32(10661),
		},
	},
	1679: {
		Fname: __ccgo_ts + 17937,
		Fcodepoints: [2]uint32{
			0: uint32(10217),
		},
	},
	1680: {
		Fname: __ccgo_ts + 17946,
		Fcodepoints: [2]uint32{
			0: uint32(187),
		},
	},
	1681: {
		Fname: __ccgo_ts + 17954,
		Fcodepoints: [2]uint32{
			0: uint32(8594),
		},
	},
	1682: {
		Fname: __ccgo_ts + 17961,
		Fcodepoints: [2]uint32{
			0: uint32(10613),
		},
	},
	1683: {
		Fname: __ccgo_ts + 17970,
		Fcodepoints: [2]uint32{
			0: uint32(8677),
		},
	},
	1684: {
		Fname: __ccgo_ts + 17978,
		Fcodepoints: [2]uint32{
			0: uint32(10528),
		},
	},
	1685: {
		Fname: __ccgo_ts + 17988,
		Fcodepoints: [2]uint32{
			0: uint32(10547),
		},
	},
	1686: {
		Fname: __ccgo_ts + 17996,
		Fcodepoints: [2]uint32{
			0: uint32(10526),
		},
	},
	1687: {
		Fname: __ccgo_ts + 18005,
		Fcodepoints: [2]uint32{
			0: uint32(8618),
		},
	},
	1688: {
		Fname: __ccgo_ts + 18014,
		Fcodepoints: [2]uint32{
			0: uint32(8620),
		},
	},
	1689: {
		Fname: __ccgo_ts + 18023,
		Fcodepoints: [2]uint32{
			0: uint32(10565),
		},
	},
	1690: {
		Fname: __ccgo_ts + 18032,
		Fcodepoints: [2]uint32{
			0: uint32(10612),
		},
	},
	1691: {
		Fname: __ccgo_ts + 18042,
		Fcodepoints: [2]uint32{
			0: uint32(8611),
		},
	},
	1692: {
		Fname: __ccgo_ts + 18051,
		Fcodepoints: [2]uint32{
			0: uint32(8605),
		},
	},
	1693: {
		Fname: __ccgo_ts + 18059,
		Fcodepoints: [2]uint32{
			0: uint32(10522),
		},
	},
	1694: {
		Fname: __ccgo_ts + 18068,
		Fcodepoints: [2]uint32{
			0: uint32(8758),
		},
	},
	1695: {
		Fname: __ccgo_ts + 18076,
		Fcodepoints: [2]uint32{
			0: uint32(8474),
		},
	},
	1696: {
		Fname: __ccgo_ts + 18088,
		Fcodepoints: [2]uint32{
			0: uint32(10509),
		},
	},
	1697: {
		Fname: __ccgo_ts + 18096,
		Fcodepoints: [2]uint32{
			0: uint32(10099),
		},
	},
	1698: {
		Fname: __ccgo_ts + 18104,
		Fcodepoints: [2]uint32{
			0: uint32(125),
		},
	},
	1699: {
		Fname: __ccgo_ts + 18113,
		Fcodepoints: [2]uint32{
			0: uint32(93),
		},
	},
	1700: {
		Fname: __ccgo_ts + 18122,
		Fcodepoints: [2]uint32{
			0: uint32(10636),
		},
	},
	1701: {
		Fname: __ccgo_ts + 18130,
		Fcodepoints: [2]uint32{
			0: uint32(10638),
		},
	},
	1702: {
		Fname: __ccgo_ts + 18140,
		Fcodepoints: [2]uint32{
			0: uint32(10640),
		},
	},
	1703: {
		Fname: __ccgo_ts + 18150,
		Fcodepoints: [2]uint32{
			0: uint32(345),
		},
	},
	1704: {
		Fname: __ccgo_ts + 18159,
		Fcodepoints: [2]uint32{
			0: uint32(343),
		},
	},
	1705: {
		Fname: __ccgo_ts + 18168,
		Fcodepoints: [2]uint32{
			0: uint32(8969),
		},
	},
	1706: {
		Fname: __ccgo_ts + 18176,
		Fcodepoints: [2]uint32{
			0: uint32(125),
		},
	},
	1707: {
		Fname: __ccgo_ts + 18183,
		Fcodepoints: [2]uint32{
			0: uint32(1088),
		},
	},
	1708: {
		Fname: __ccgo_ts + 18189,
		Fcodepoints: [2]uint32{
			0: uint32(10551),
		},
	},
	1709: {
		Fname: __ccgo_ts + 18196,
		Fcodepoints: [2]uint32{
			0: uint32(10601),
		},
	},
	1710: {
		Fname: __ccgo_ts + 18206,
		Fcodepoints: [2]uint32{
			0: uint32(8221),
		},
	},
	1711: {
		Fname: __ccgo_ts + 18214,
		Fcodepoints: [2]uint32{
			0: uint32(8221),
		},
	},
	1712: {
		Fname: __ccgo_ts + 18223,
		Fcodepoints: [2]uint32{
			0: uint32(8627),
		},
	},
	1713: {
		Fname: __ccgo_ts + 18230,
		Fcodepoints: [2]uint32{
			0: uint32(8476),
		},
	},
	1714: {
		Fname: __ccgo_ts + 18237,
		Fcodepoints: [2]uint32{
			0: uint32(8475),
		},
	},
	1715: {
		Fname: __ccgo_ts + 18247,
		Fcodepoints: [2]uint32{
			0: uint32(8476),
		},
	},
	1716: {
		Fname: __ccgo_ts + 18258,
		Fcodepoints: [2]uint32{
			0: uint32(8477),
		},
	},
	1717: {
		Fname: __ccgo_ts + 18266,
		Fcodepoints: [2]uint32{
			0: uint32(9645),
		},
	},
	1718: {
		Fname: __ccgo_ts + 18273,
		Fcodepoints: [2]uint32{
			0: uint32(174),
		},
	},
	1719: {
		Fname: __ccgo_ts + 18279,
		Fcodepoints: [2]uint32{
			0: uint32(10621),
		},
	},
	1720: {
		Fname: __ccgo_ts + 18288,
		Fcodepoints: [2]uint32{
			0: uint32(8971),
		},
	},
	1721: {
		Fname: __ccgo_ts + 18297,
		Fcodepoints: [2]uint32{
			0: uint32(120111),
		},
	},
	1722: {
		Fname: __ccgo_ts + 18303,
		Fcodepoints: [2]uint32{
			0: uint32(8641),
		},
	},
	1723: {
		Fname: __ccgo_ts + 18311,
		Fcodepoints: [2]uint32{
			0: uint32(8640),
		},
	},
	1724: {
		Fname: __ccgo_ts + 18319,
		Fcodepoints: [2]uint32{
			0: uint32(10604),
		},
	},
	1725: {
		Fname: __ccgo_ts + 18328,
		Fcodepoints: [2]uint32{
			0: uint32(961),
		},
	},
	1726: {
		Fname: __ccgo_ts + 18334,
		Fcodepoints: [2]uint32{
			0: uint32(1009),
		},
	},
	1727: {
		Fname: __ccgo_ts + 18341,
		Fcodepoints: [2]uint32{
			0: uint32(8594),
		},
	},
	1728: {
		Fname: __ccgo_ts + 18354,
		Fcodepoints: [2]uint32{
			0: uint32(8611),
		},
	},
	1729: {
		Fname: __ccgo_ts + 18371,
		Fcodepoints: [2]uint32{
			0: uint32(8641),
		},
	},
	1730: {
		Fname: __ccgo_ts + 18390,
		Fcodepoints: [2]uint32{
			0: uint32(8640),
		},
	},
	1731: {
		Fname: __ccgo_ts + 18407,
		Fcodepoints: [2]uint32{
			0: uint32(8644),
		},
	},
	1732: {
		Fname: __ccgo_ts + 18425,
		Fcodepoints: [2]uint32{
			0: uint32(8652),
		},
	},
	1733: {
		Fname: __ccgo_ts + 18445,
		Fcodepoints: [2]uint32{
			0: uint32(8649),
		},
	},
	1734: {
		Fname: __ccgo_ts + 18464,
		Fcodepoints: [2]uint32{
			0: uint32(8605),
		},
	},
	1735: {
		Fname: __ccgo_ts + 18482,
		Fcodepoints: [2]uint32{
			0: uint32(8908),
		},
	},
	1736: {
		Fname: __ccgo_ts + 18500,
		Fcodepoints: [2]uint32{
			0: uint32(730),
		},
	},
	1737: {
		Fname: __ccgo_ts + 18507,
		Fcodepoints: [2]uint32{
			0: uint32(8787),
		},
	},
	1738: {
		Fname: __ccgo_ts + 18522,
		Fcodepoints: [2]uint32{
			0: uint32(8644),
		},
	},
	1739: {
		Fname: __ccgo_ts + 18530,
		Fcodepoints: [2]uint32{
			0: uint32(8652),
		},
	},
	1740: {
		Fname: __ccgo_ts + 18538,
		Fcodepoints: [2]uint32{
			0: uint32(8207),
		},
	},
	1741: {
		Fname: __ccgo_ts + 18544,
		Fcodepoints: [2]uint32{
			0: uint32(9137),
		},
	},
	1742: {
		Fname: __ccgo_ts + 18553,
		Fcodepoints: [2]uint32{
			0: uint32(9137),
		},
	},
	1743: {
		Fname: __ccgo_ts + 18566,
		Fcodepoints: [2]uint32{
			0: uint32(10990),
		},
	},
	1744: {
		Fname: __ccgo_ts + 18574,
		Fcodepoints: [2]uint32{
			0: uint32(10221),
		},
	},
	1745: {
		Fname: __ccgo_ts + 18582,
		Fcodepoints: [2]uint32{
			0: uint32(8702),
		},
	},
	1746: {
		Fname: __ccgo_ts + 18590,
		Fcodepoints: [2]uint32{
			0: uint32(10215),
		},
	},
	1747: {
		Fname: __ccgo_ts + 18598,
		Fcodepoints: [2]uint32{
			0: uint32(10630),
		},
	},
	1748: {
		Fname: __ccgo_ts + 18606,
		Fcodepoints: [2]uint32{
			0: uint32(120163),
		},
	},
	1749: {
		Fname: __ccgo_ts + 18613,
		Fcodepoints: [2]uint32{
			0: uint32(10798),
		},
	},
	1750: {
		Fname: __ccgo_ts + 18622,
		Fcodepoints: [2]uint32{
			0: uint32(10805),
		},
	},
	1751: {
		Fname: __ccgo_ts + 18632,
		Fcodepoints: [2]uint32{
			0: uint32(41),
		},
	},
	1752: {
		Fname: __ccgo_ts + 18639,
		Fcodepoints: [2]uint32{
			0: uint32(10644),
		},
	},
	1753: {
		Fname: __ccgo_ts + 18648,
		Fcodepoints: [2]uint32{
			0: uint32(10770),
		},
	},
	1754: {
		Fname: __ccgo_ts + 18659,
		Fcodepoints: [2]uint32{
			0: uint32(8649),
		},
	},
	1755: {
		Fname: __ccgo_ts + 18667,
		Fcodepoints: [2]uint32{
			0: uint32(8250),
		},
	},
	1756: {
		Fname: __ccgo_ts + 18676,
		Fcodepoints: [2]uint32{
			0: uint32(120007),
		},
	},
	1757: {
		Fname: __ccgo_ts + 18683,
		Fcodepoints: [2]uint32{
			0: uint32(8625),
		},
	},
	1758: {
		Fname: __ccgo_ts + 18689,
		Fcodepoints: [2]uint32{
			0: uint32(93),
		},
	},
	1759: {
		Fname: __ccgo_ts + 18696,
		Fcodepoints: [2]uint32{
			0: uint32(8217),
		},
	},
	1760: {
		Fname: __ccgo_ts + 18704,
		Fcodepoints: [2]uint32{
			0: uint32(8217),
		},
	},
	1761: {
		Fname: __ccgo_ts + 18713,
		Fcodepoints: [2]uint32{
			0: uint32(8908),
		},
	},
	1762: {
		Fname: __ccgo_ts + 18722,
		Fcodepoints: [2]uint32{
			0: uint32(8906),
		},
	},
	1763: {
		Fname: __ccgo_ts + 18731,
		Fcodepoints: [2]uint32{
			0: uint32(9657),
		},
	},
	1764: {
		Fname: __ccgo_ts + 18738,
		Fcodepoints: [2]uint32{
			0: uint32(8885),
		},
	},
	1765: {
		Fname: __ccgo_ts + 18746,
		Fcodepoints: [2]uint32{
			0: uint32(9656),
		},
	},
	1766: {
		Fname: __ccgo_ts + 18754,
		Fcodepoints: [2]uint32{
			0: uint32(10702),
		},
	},
	1767: {
		Fname: __ccgo_ts + 18765,
		Fcodepoints: [2]uint32{
			0: uint32(10600),
		},
	},
	1768: {
		Fname: __ccgo_ts + 18775,
		Fcodepoints: [2]uint32{
			0: uint32(8478),
		},
	},
	1769: {
		Fname: __ccgo_ts + 18780,
		Fcodepoints: [2]uint32{
			0: uint32(347),
		},
	},
	1770: {
		Fname: __ccgo_ts + 18789,
		Fcodepoints: [2]uint32{
			0: uint32(8218),
		},
	},
	1771: {
		Fname: __ccgo_ts + 18797,
		Fcodepoints: [2]uint32{
			0: uint32(8827),
		},
	},
	1772: {
		Fname: __ccgo_ts + 18802,
		Fcodepoints: [2]uint32{
			0: uint32(10932),
		},
	},
	1773: {
		Fname: __ccgo_ts + 18808,
		Fcodepoints: [2]uint32{
			0: uint32(10936),
		},
	},
	1774: {
		Fname: __ccgo_ts + 18815,
		Fcodepoints: [2]uint32{
			0: uint32(353),
		},
	},
	1775: {
		Fname: __ccgo_ts + 18824,
		Fcodepoints: [2]uint32{
			0: uint32(8829),
		},
	},
	1776: {
		Fname: __ccgo_ts + 18832,
		Fcodepoints: [2]uint32{
			0: uint32(10928),
		},
	},
	1777: {
		Fname: __ccgo_ts + 18838,
		Fcodepoints: [2]uint32{
			0: uint32(351),
		},
	},
	1778: {
		Fname: __ccgo_ts + 18847,
		Fcodepoints: [2]uint32{
			0: uint32(349),
		},
	},
	1779: {
		Fname: __ccgo_ts + 18855,
		Fcodepoints: [2]uint32{
			0: uint32(10934),
		},
	},
	1780: {
		Fname: __ccgo_ts + 18862,
		Fcodepoints: [2]uint32{
			0: uint32(10938),
		},
	},
	1781: {
		Fname: __ccgo_ts + 18870,
		Fcodepoints: [2]uint32{
			0: uint32(8937),
		},
	},
	1782: {
		Fname: __ccgo_ts + 18879,
		Fcodepoints: [2]uint32{
			0: uint32(10771),
		},
	},
	1783: {
		Fname: __ccgo_ts + 18890,
		Fcodepoints: [2]uint32{
			0: uint32(8831),
		},
	},
	1784: {
		Fname: __ccgo_ts + 18898,
		Fcodepoints: [2]uint32{
			0: uint32(1089),
		},
	},
	1785: {
		Fname: __ccgo_ts + 18904,
		Fcodepoints: [2]uint32{
			0: uint32(8901),
		},
	},
	1786: {
		Fname: __ccgo_ts + 18911,
		Fcodepoints: [2]uint32{
			0: uint32(8865),
		},
	},
	1787: {
		Fname: __ccgo_ts + 18919,
		Fcodepoints: [2]uint32{
			0: uint32(10854),
		},
	},
	1788: {
		Fname: __ccgo_ts + 18927,
		Fcodepoints: [2]uint32{
			0: uint32(8664),
		},
	},
	1789: {
		Fname: __ccgo_ts + 18935,
		Fcodepoints: [2]uint32{
			0: uint32(10533),
		},
	},
	1790: {
		Fname: __ccgo_ts + 18944,
		Fcodepoints: [2]uint32{
			0: uint32(8600),
		},
	},
	1791: {
		Fname: __ccgo_ts + 18952,
		Fcodepoints: [2]uint32{
			0: uint32(8600),
		},
	},
	1792: {
		Fname: __ccgo_ts + 18962,
		Fcodepoints: [2]uint32{
			0: uint32(167),
		},
	},
	1793: {
		Fname: __ccgo_ts + 18969,
		Fcodepoints: [2]uint32{
			0: uint32(59),
		},
	},
	1794: {
		Fname: __ccgo_ts + 18976,
		Fcodepoints: [2]uint32{
			0: uint32(10537),
		},
	},
	1795: {
		Fname: __ccgo_ts + 18985,
		Fcodepoints: [2]uint32{
			0: uint32(8726),
		},
	},
	1796: {
		Fname: __ccgo_ts + 18996,
		Fcodepoints: [2]uint32{
			0: uint32(8726),
		},
	},
	1797: {
		Fname: __ccgo_ts + 19004,
		Fcodepoints: [2]uint32{
			0: uint32(10038),
		},
	},
	1798: {
		Fname: __ccgo_ts + 19011,
		Fcodepoints: [2]uint32{
			0: uint32(120112),
		},
	},
	1799: {
		Fname: __ccgo_ts + 19017,
		Fcodepoints: [2]uint32{
			0: uint32(8994),
		},
	},
	1800: {
		Fname: __ccgo_ts + 19026,
		Fcodepoints: [2]uint32{
			0: uint32(9839),
		},
	},
	1801: {
		Fname: __ccgo_ts + 19034,
		Fcodepoints: [2]uint32{
			0: uint32(1097),
		},
	},
	1802: {
		Fname: __ccgo_ts + 19043,
		Fcodepoints: [2]uint32{
			0: uint32(1096),
		},
	},
	1803: {
		Fname: __ccgo_ts + 19050,
		Fcodepoints: [2]uint32{
			0: uint32(8739),
		},
	},
	1804: {
		Fname: __ccgo_ts + 19061,
		Fcodepoints: [2]uint32{
			0: uint32(8741),
		},
	},
	1805: {
		Fname: __ccgo_ts + 19077,
		Fcodepoints: [2]uint32{
			0: uint32(173),
		},
	},
	1806: {
		Fname: __ccgo_ts + 19083,
		Fcodepoints: [2]uint32{
			0: uint32(963),
		},
	},
	1807: {
		Fname: __ccgo_ts + 19091,
		Fcodepoints: [2]uint32{
			0: uint32(962),
		},
	},
	1808: {
		Fname: __ccgo_ts + 19100,
		Fcodepoints: [2]uint32{
			0: uint32(962),
		},
	},
	1809: {
		Fname: __ccgo_ts + 19109,
		Fcodepoints: [2]uint32{
			0: uint32(8764),
		},
	},
	1810: {
		Fname: __ccgo_ts + 19115,
		Fcodepoints: [2]uint32{
			0: uint32(10858),
		},
	},
	1811: {
		Fname: __ccgo_ts + 19124,
		Fcodepoints: [2]uint32{
			0: uint32(8771),
		},
	},
	1812: {
		Fname: __ccgo_ts + 19131,
		Fcodepoints: [2]uint32{
			0: uint32(8771),
		},
	},
	1813: {
		Fname: __ccgo_ts + 19139,
		Fcodepoints: [2]uint32{
			0: uint32(10910),
		},
	},
	1814: {
		Fname: __ccgo_ts + 19146,
		Fcodepoints: [2]uint32{
			0: uint32(10912),
		},
	},
	1815: {
		Fname: __ccgo_ts + 19154,
		Fcodepoints: [2]uint32{
			0: uint32(10909),
		},
	},
	1816: {
		Fname: __ccgo_ts + 19161,
		Fcodepoints: [2]uint32{
			0: uint32(10911),
		},
	},
	1817: {
		Fname: __ccgo_ts + 19169,
		Fcodepoints: [2]uint32{
			0: uint32(8774),
		},
	},
	1818: {
		Fname: __ccgo_ts + 19177,
		Fcodepoints: [2]uint32{
			0: uint32(10788),
		},
	},
	1819: {
		Fname: __ccgo_ts + 19187,
		Fcodepoints: [2]uint32{
			0: uint32(10610),
		},
	},
	1820: {
		Fname: __ccgo_ts + 19197,
		Fcodepoints: [2]uint32{
			0: uint32(8592),
		},
	},
	1821: {
		Fname: __ccgo_ts + 19205,
		Fcodepoints: [2]uint32{
			0: uint32(8726),
		},
	},
	1822: {
		Fname: __ccgo_ts + 19221,
		Fcodepoints: [2]uint32{
			0: uint32(10803),
		},
	},
	1823: {
		Fname: __ccgo_ts + 19230,
		Fcodepoints: [2]uint32{
			0: uint32(10724),
		},
	},
	1824: {
		Fname: __ccgo_ts + 19241,
		Fcodepoints: [2]uint32{
			0: uint32(8739),
		},
	},
	1825: {
		Fname: __ccgo_ts + 19248,
		Fcodepoints: [2]uint32{
			0: uint32(8995),
		},
	},
	1826: {
		Fname: __ccgo_ts + 19256,
		Fcodepoints: [2]uint32{
			0: uint32(10922),
		},
	},
	1827: {
		Fname: __ccgo_ts + 19262,
		Fcodepoints: [2]uint32{
			0: uint32(10924),
		},
	},
	1828: {
		Fname: __ccgo_ts + 19269,
		Fcodepoints: [2]uint32{
			0: uint32(10924),
			1: uint32(65024),
		},
	},
	1829: {
		Fname: __ccgo_ts + 19277,
		Fcodepoints: [2]uint32{
			0: uint32(1100),
		},
	},
	1830: {
		Fname: __ccgo_ts + 19286,
		Fcodepoints: [2]uint32{
			0: uint32(47),
		},
	},
	1831: {
		Fname: __ccgo_ts + 19292,
		Fcodepoints: [2]uint32{
			0: uint32(10692),
		},
	},
	1832: {
		Fname: __ccgo_ts + 19299,
		Fcodepoints: [2]uint32{
			0: uint32(9023),
		},
	},
	1833: {
		Fname: __ccgo_ts + 19308,
		Fcodepoints: [2]uint32{
			0: uint32(120164),
		},
	},
	1834: {
		Fname: __ccgo_ts + 19315,
		Fcodepoints: [2]uint32{
			0: uint32(9824),
		},
	},
	1835: {
		Fname: __ccgo_ts + 19324,
		Fcodepoints: [2]uint32{
			0: uint32(9824),
		},
	},
	1836: {
		Fname: __ccgo_ts + 19336,
		Fcodepoints: [2]uint32{
			0: uint32(8741),
		},
	},
	1837: {
		Fname: __ccgo_ts + 19343,
		Fcodepoints: [2]uint32{
			0: uint32(8851),
		},
	},
	1838: {
		Fname: __ccgo_ts + 19351,
		Fcodepoints: [2]uint32{
			0: uint32(8851),
			1: uint32(65024),
		},
	},
	1839: {
		Fname: __ccgo_ts + 19360,
		Fcodepoints: [2]uint32{
			0: uint32(8852),
		},
	},
	1840: {
		Fname: __ccgo_ts + 19368,
		Fcodepoints: [2]uint32{
			0: uint32(8852),
			1: uint32(65024),
		},
	},
	1841: {
		Fname: __ccgo_ts + 19377,
		Fcodepoints: [2]uint32{
			0: uint32(8847),
		},
	},
	1842: {
		Fname: __ccgo_ts + 19385,
		Fcodepoints: [2]uint32{
			0: uint32(8849),
		},
	},
	1843: {
		Fname: __ccgo_ts + 19394,
		Fcodepoints: [2]uint32{
			0: uint32(8847),
		},
	},
	1844: {
		Fname: __ccgo_ts + 19405,
		Fcodepoints: [2]uint32{
			0: uint32(8849),
		},
	},
	1845: {
		Fname: __ccgo_ts + 19418,
		Fcodepoints: [2]uint32{
			0: uint32(8848),
		},
	},
	1846: {
		Fname: __ccgo_ts + 19426,
		Fcodepoints: [2]uint32{
			0: uint32(8850),
		},
	},
	1847: {
		Fname: __ccgo_ts + 19435,
		Fcodepoints: [2]uint32{
			0: uint32(8848),
		},
	},
	1848: {
		Fname: __ccgo_ts + 19446,
		Fcodepoints: [2]uint32{
			0: uint32(8850),
		},
	},
	1849: {
		Fname: __ccgo_ts + 19459,
		Fcodepoints: [2]uint32{
			0: uint32(9633),
		},
	},
	1850: {
		Fname: __ccgo_ts + 19465,
		Fcodepoints: [2]uint32{
			0: uint32(9633),
		},
	},
	1851: {
		Fname: __ccgo_ts + 19474,
		Fcodepoints: [2]uint32{
			0: uint32(9642),
		},
	},
	1852: {
		Fname: __ccgo_ts + 19483,
		Fcodepoints: [2]uint32{
			0: uint32(9642),
		},
	},
	1853: {
		Fname: __ccgo_ts + 19490,
		Fcodepoints: [2]uint32{
			0: uint32(8594),
		},
	},
	1854: {
		Fname: __ccgo_ts + 19498,
		Fcodepoints: [2]uint32{
			0: uint32(120008),
		},
	},
	1855: {
		Fname: __ccgo_ts + 19505,
		Fcodepoints: [2]uint32{
			0: uint32(8726),
		},
	},
	1856: {
		Fname: __ccgo_ts + 19514,
		Fcodepoints: [2]uint32{
			0: uint32(8995),
		},
	},
	1857: {
		Fname: __ccgo_ts + 19523,
		Fcodepoints: [2]uint32{
			0: uint32(8902),
		},
	},
	1858: {
		Fname: __ccgo_ts + 19532,
		Fcodepoints: [2]uint32{
			0: uint32(9734),
		},
	},
	1859: {
		Fname: __ccgo_ts + 19539,
		Fcodepoints: [2]uint32{
			0: uint32(9733),
		},
	},
	1860: {
		Fname: __ccgo_ts + 19547,
		Fcodepoints: [2]uint32{
			0: uint32(1013),
		},
	},
	1861: {
		Fname: __ccgo_ts + 19565,
		Fcodepoints: [2]uint32{
			0: uint32(981),
		},
	},
	1862: {
		Fname: __ccgo_ts + 19579,
		Fcodepoints: [2]uint32{
			0: uint32(175),
		},
	},
	1863: {
		Fname: __ccgo_ts + 19587,
		Fcodepoints: [2]uint32{
			0: uint32(8834),
		},
	},
	1864: {
		Fname: __ccgo_ts + 19593,
		Fcodepoints: [2]uint32{
			0: uint32(10949),
		},
	},
	1865: {
		Fname: __ccgo_ts + 19600,
		Fcodepoints: [2]uint32{
			0: uint32(10941),
		},
	},
	1866: {
		Fname: __ccgo_ts + 19609,
		Fcodepoints: [2]uint32{
			0: uint32(8838),
		},
	},
	1867: {
		Fname: __ccgo_ts + 19616,
		Fcodepoints: [2]uint32{
			0: uint32(10947),
		},
	},
	1868: {
		Fname: __ccgo_ts + 19626,
		Fcodepoints: [2]uint32{
			0: uint32(10945),
		},
	},
	1869: {
		Fname: __ccgo_ts + 19636,
		Fcodepoints: [2]uint32{
			0: uint32(10955),
		},
	},
	1870: {
		Fname: __ccgo_ts + 19644,
		Fcodepoints: [2]uint32{
			0: uint32(8842),
		},
	},
	1871: {
		Fname: __ccgo_ts + 19652,
		Fcodepoints: [2]uint32{
			0: uint32(10943),
		},
	},
	1872: {
		Fname: __ccgo_ts + 19662,
		Fcodepoints: [2]uint32{
			0: uint32(10617),
		},
	},
	1873: {
		Fname: __ccgo_ts + 19672,
		Fcodepoints: [2]uint32{
			0: uint32(8834),
		},
	},
	1874: {
		Fname: __ccgo_ts + 19681,
		Fcodepoints: [2]uint32{
			0: uint32(8838),
		},
	},
	1875: {
		Fname: __ccgo_ts + 19692,
		Fcodepoints: [2]uint32{
			0: uint32(10949),
		},
	},
	1876: {
		Fname: __ccgo_ts + 19704,
		Fcodepoints: [2]uint32{
			0: uint32(8842),
		},
	},
	1877: {
		Fname: __ccgo_ts + 19716,
		Fcodepoints: [2]uint32{
			0: uint32(10955),
		},
	},
	1878: {
		Fname: __ccgo_ts + 19729,
		Fcodepoints: [2]uint32{
			0: uint32(10951),
		},
	},
	1879: {
		Fname: __ccgo_ts + 19738,
		Fcodepoints: [2]uint32{
			0: uint32(10965),
		},
	},
	1880: {
		Fname: __ccgo_ts + 19747,
		Fcodepoints: [2]uint32{
			0: uint32(10963),
		},
	},
	1881: {
		Fname: __ccgo_ts + 19756,
		Fcodepoints: [2]uint32{
			0: uint32(8827),
		},
	},
	1882: {
		Fname: __ccgo_ts + 19763,
		Fcodepoints: [2]uint32{
			0: uint32(10936),
		},
	},
	1883: {
		Fname: __ccgo_ts + 19776,
		Fcodepoints: [2]uint32{
			0: uint32(8829),
		},
	},
	1884: {
		Fname: __ccgo_ts + 19790,
		Fcodepoints: [2]uint32{
			0: uint32(10928),
		},
	},
	1885: {
		Fname: __ccgo_ts + 19799,
		Fcodepoints: [2]uint32{
			0: uint32(10938),
		},
	},
	1886: {
		Fname: __ccgo_ts + 19813,
		Fcodepoints: [2]uint32{
			0: uint32(10934),
		},
	},
	1887: {
		Fname: __ccgo_ts + 19824,
		Fcodepoints: [2]uint32{
			0: uint32(8937),
		},
	},
	1888: {
		Fname: __ccgo_ts + 19835,
		Fcodepoints: [2]uint32{
			0: uint32(8831),
		},
	},
	1889: {
		Fname: __ccgo_ts + 19845,
		Fcodepoints: [2]uint32{
			0: uint32(8721),
		},
	},
	1890: {
		Fname: __ccgo_ts + 19851,
		Fcodepoints: [2]uint32{
			0: uint32(9834),
		},
	},
	1891: {
		Fname: __ccgo_ts + 19858,
		Fcodepoints: [2]uint32{
			0: uint32(185),
		},
	},
	1892: {
		Fname: __ccgo_ts + 19865,
		Fcodepoints: [2]uint32{
			0: uint32(178),
		},
	},
	1893: {
		Fname: __ccgo_ts + 19872,
		Fcodepoints: [2]uint32{
			0: uint32(179),
		},
	},
	1894: {
		Fname: __ccgo_ts + 19879,
		Fcodepoints: [2]uint32{
			0: uint32(8835),
		},
	},
	1895: {
		Fname: __ccgo_ts + 19885,
		Fcodepoints: [2]uint32{
			0: uint32(10950),
		},
	},
	1896: {
		Fname: __ccgo_ts + 19892,
		Fcodepoints: [2]uint32{
			0: uint32(10942),
		},
	},
	1897: {
		Fname: __ccgo_ts + 19901,
		Fcodepoints: [2]uint32{
			0: uint32(10968),
		},
	},
	1898: {
		Fname: __ccgo_ts + 19911,
		Fcodepoints: [2]uint32{
			0: uint32(8839),
		},
	},
	1899: {
		Fname: __ccgo_ts + 19918,
		Fcodepoints: [2]uint32{
			0: uint32(10948),
		},
	},
	1900: {
		Fname: __ccgo_ts + 19928,
		Fcodepoints: [2]uint32{
			0: uint32(10185),
		},
	},
	1901: {
		Fname: __ccgo_ts + 19938,
		Fcodepoints: [2]uint32{
			0: uint32(10967),
		},
	},
	1902: {
		Fname: __ccgo_ts + 19948,
		Fcodepoints: [2]uint32{
			0: uint32(10619),
		},
	},
	1903: {
		Fname: __ccgo_ts + 19958,
		Fcodepoints: [2]uint32{
			0: uint32(10946),
		},
	},
	1904: {
		Fname: __ccgo_ts + 19968,
		Fcodepoints: [2]uint32{
			0: uint32(10956),
		},
	},
	1905: {
		Fname: __ccgo_ts + 19976,
		Fcodepoints: [2]uint32{
			0: uint32(8843),
		},
	},
	1906: {
		Fname: __ccgo_ts + 19984,
		Fcodepoints: [2]uint32{
			0: uint32(10944),
		},
	},
	1907: {
		Fname: __ccgo_ts + 19994,
		Fcodepoints: [2]uint32{
			0: uint32(8835),
		},
	},
	1908: {
		Fname: __ccgo_ts + 20003,
		Fcodepoints: [2]uint32{
			0: uint32(8839),
		},
	},
	1909: {
		Fname: __ccgo_ts + 20014,
		Fcodepoints: [2]uint32{
			0: uint32(10950),
		},
	},
	1910: {
		Fname: __ccgo_ts + 20026,
		Fcodepoints: [2]uint32{
			0: uint32(8843),
		},
	},
	1911: {
		Fname: __ccgo_ts + 20038,
		Fcodepoints: [2]uint32{
			0: uint32(10956),
		},
	},
	1912: {
		Fname: __ccgo_ts + 20051,
		Fcodepoints: [2]uint32{
			0: uint32(10952),
		},
	},
	1913: {
		Fname: __ccgo_ts + 20060,
		Fcodepoints: [2]uint32{
			0: uint32(10964),
		},
	},
	1914: {
		Fname: __ccgo_ts + 20069,
		Fcodepoints: [2]uint32{
			0: uint32(10966),
		},
	},
	1915: {
		Fname: __ccgo_ts + 20078,
		Fcodepoints: [2]uint32{
			0: uint32(8665),
		},
	},
	1916: {
		Fname: __ccgo_ts + 20086,
		Fcodepoints: [2]uint32{
			0: uint32(10534),
		},
	},
	1917: {
		Fname: __ccgo_ts + 20095,
		Fcodepoints: [2]uint32{
			0: uint32(8601),
		},
	},
	1918: {
		Fname: __ccgo_ts + 20103,
		Fcodepoints: [2]uint32{
			0: uint32(8601),
		},
	},
	1919: {
		Fname: __ccgo_ts + 20113,
		Fcodepoints: [2]uint32{
			0: uint32(10538),
		},
	},
	1920: {
		Fname: __ccgo_ts + 20122,
		Fcodepoints: [2]uint32{
			0: uint32(223),
		},
	},
	1921: {
		Fname: __ccgo_ts + 20130,
		Fcodepoints: [2]uint32{
			0: uint32(8982),
		},
	},
	1922: {
		Fname: __ccgo_ts + 20139,
		Fcodepoints: [2]uint32{
			0: uint32(964),
		},
	},
	1923: {
		Fname: __ccgo_ts + 20145,
		Fcodepoints: [2]uint32{
			0: uint32(9140),
		},
	},
	1924: {
		Fname: __ccgo_ts + 20152,
		Fcodepoints: [2]uint32{
			0: uint32(357),
		},
	},
	1925: {
		Fname: __ccgo_ts + 20161,
		Fcodepoints: [2]uint32{
			0: uint32(355),
		},
	},
	1926: {
		Fname: __ccgo_ts + 20170,
		Fcodepoints: [2]uint32{
			0: uint32(1090),
		},
	},
	1927: {
		Fname: __ccgo_ts + 20176,
		Fcodepoints: [2]uint32{
			0: uint32(8411),
		},
	},
	1928: {
		Fname: __ccgo_ts + 20183,
		Fcodepoints: [2]uint32{
			0: uint32(8981),
		},
	},
	1929: {
		Fname: __ccgo_ts + 20192,
		Fcodepoints: [2]uint32{
			0: uint32(120113),
		},
	},
	1930: {
		Fname: __ccgo_ts + 20198,
		Fcodepoints: [2]uint32{
			0: uint32(8756),
		},
	},
	1931: {
		Fname: __ccgo_ts + 20207,
		Fcodepoints: [2]uint32{
			0: uint32(8756),
		},
	},
	1932: {
		Fname: __ccgo_ts + 20219,
		Fcodepoints: [2]uint32{
			0: uint32(952),
		},
	},
	1933: {
		Fname: __ccgo_ts + 20227,
		Fcodepoints: [2]uint32{
			0: uint32(977),
		},
	},
	1934: {
		Fname: __ccgo_ts + 20238,
		Fcodepoints: [2]uint32{
			0: uint32(977),
		},
	},
	1935: {
		Fname: __ccgo_ts + 20247,
		Fcodepoints: [2]uint32{
			0: uint32(8776),
		},
	},
	1936: {
		Fname: __ccgo_ts + 20261,
		Fcodepoints: [2]uint32{
			0: uint32(8764),
		},
	},
	1937: {
		Fname: __ccgo_ts + 20272,
		Fcodepoints: [2]uint32{
			0: uint32(8201),
		},
	},
	1938: {
		Fname: __ccgo_ts + 20281,
		Fcodepoints: [2]uint32{
			0: uint32(8776),
		},
	},
	1939: {
		Fname: __ccgo_ts + 20289,
		Fcodepoints: [2]uint32{
			0: uint32(8764),
		},
	},
	1940: {
		Fname: __ccgo_ts + 20298,
		Fcodepoints: [2]uint32{
			0: uint32(254),
		},
	},
	1941: {
		Fname: __ccgo_ts + 20306,
		Fcodepoints: [2]uint32{
			0: uint32(732),
		},
	},
	1942: {
		Fname: __ccgo_ts + 20314,
		Fcodepoints: [2]uint32{
			0: uint32(215),
		},
	},
	1943: {
		Fname: __ccgo_ts + 20322,
		Fcodepoints: [2]uint32{
			0: uint32(8864),
		},
	},
	1944: {
		Fname: __ccgo_ts + 20331,
		Fcodepoints: [2]uint32{
			0: uint32(10801),
		},
	},
	1945: {
		Fname: __ccgo_ts + 20342,
		Fcodepoints: [2]uint32{
			0: uint32(10800),
		},
	},
	1946: {
		Fname: __ccgo_ts + 20351,
		Fcodepoints: [2]uint32{
			0: uint32(8749),
		},
	},
	1947: {
		Fname: __ccgo_ts + 20358,
		Fcodepoints: [2]uint32{
			0: uint32(10536),
		},
	},
	1948: {
		Fname: __ccgo_ts + 20365,
		Fcodepoints: [2]uint32{
			0: uint32(8868),
		},
	},
	1949: {
		Fname: __ccgo_ts + 20371,
		Fcodepoints: [2]uint32{
			0: uint32(9014),
		},
	},
	1950: {
		Fname: __ccgo_ts + 20380,
		Fcodepoints: [2]uint32{
			0: uint32(10993),
		},
	},
	1951: {
		Fname: __ccgo_ts + 20389,
		Fcodepoints: [2]uint32{
			0: uint32(120165),
		},
	},
	1952: {
		Fname: __ccgo_ts + 20396,
		Fcodepoints: [2]uint32{
			0: uint32(10970),
		},
	},
	1953: {
		Fname: __ccgo_ts + 20406,
		Fcodepoints: [2]uint32{
			0: uint32(10537),
		},
	},
	1954: {
		Fname: __ccgo_ts + 20413,
		Fcodepoints: [2]uint32{
			0: uint32(8244),
		},
	},
	1955: {
		Fname: __ccgo_ts + 20422,
		Fcodepoints: [2]uint32{
			0: uint32(8482),
		},
	},
	1956: {
		Fname: __ccgo_ts + 20430,
		Fcodepoints: [2]uint32{
			0: uint32(9653),
		},
	},
	1957: {
		Fname: __ccgo_ts + 20441,
		Fcodepoints: [2]uint32{
			0: uint32(9663),
		},
	},
	1958: {
		Fname: __ccgo_ts + 20456,
		Fcodepoints: [2]uint32{
			0: uint32(9667),
		},
	},
	1959: {
		Fname: __ccgo_ts + 20471,
		Fcodepoints: [2]uint32{
			0: uint32(8884),
		},
	},
	1960: {
		Fname: __ccgo_ts + 20488,
		Fcodepoints: [2]uint32{
			0: uint32(8796),
		},
	},
	1961: {
		Fname: __ccgo_ts + 20500,
		Fcodepoints: [2]uint32{
			0: uint32(9657),
		},
	},
	1962: {
		Fname: __ccgo_ts + 20516,
		Fcodepoints: [2]uint32{
			0: uint32(8885),
		},
	},
	1963: {
		Fname: __ccgo_ts + 20534,
		Fcodepoints: [2]uint32{
			0: uint32(9708),
		},
	},
	1964: {
		Fname: __ccgo_ts + 20543,
		Fcodepoints: [2]uint32{
			0: uint32(8796),
		},
	},
	1965: {
		Fname: __ccgo_ts + 20550,
		Fcodepoints: [2]uint32{
			0: uint32(10810),
		},
	},
	1966: {
		Fname: __ccgo_ts + 20561,
		Fcodepoints: [2]uint32{
			0: uint32(10809),
		},
	},
	1967: {
		Fname: __ccgo_ts + 20571,
		Fcodepoints: [2]uint32{
			0: uint32(10701),
		},
	},
	1968: {
		Fname: __ccgo_ts + 20579,
		Fcodepoints: [2]uint32{
			0: uint32(10811),
		},
	},
	1969: {
		Fname: __ccgo_ts + 20589,
		Fcodepoints: [2]uint32{
			0: uint32(9186),
		},
	},
	1970: {
		Fname: __ccgo_ts + 20600,
		Fcodepoints: [2]uint32{
			0: uint32(120009),
		},
	},
	1971: {
		Fname: __ccgo_ts + 20607,
		Fcodepoints: [2]uint32{
			0: uint32(1094),
		},
	},
	1972: {
		Fname: __ccgo_ts + 20614,
		Fcodepoints: [2]uint32{
			0: uint32(1115),
		},
	},
	1973: {
		Fname: __ccgo_ts + 20622,
		Fcodepoints: [2]uint32{
			0: uint32(359),
		},
	},
	1974: {
		Fname: __ccgo_ts + 20631,
		Fcodepoints: [2]uint32{
			0: uint32(8812),
		},
	},
	1975: {
		Fname: __ccgo_ts + 20639,
		Fcodepoints: [2]uint32{
			0: uint32(8606),
		},
	},
	1976: {
		Fname: __ccgo_ts + 20658,
		Fcodepoints: [2]uint32{
			0: uint32(8608),
		},
	},
	1977: {
		Fname: __ccgo_ts + 20678,
		Fcodepoints: [2]uint32{
			0: uint32(8657),
		},
	},
	1978: {
		Fname: __ccgo_ts + 20685,
		Fcodepoints: [2]uint32{
			0: uint32(10595),
		},
	},
	1979: {
		Fname: __ccgo_ts + 20692,
		Fcodepoints: [2]uint32{
			0: uint32(250),
		},
	},
	1980: {
		Fname: __ccgo_ts + 20701,
		Fcodepoints: [2]uint32{
			0: uint32(8593),
		},
	},
	1981: {
		Fname: __ccgo_ts + 20708,
		Fcodepoints: [2]uint32{
			0: uint32(1118),
		},
	},
	1982: {
		Fname: __ccgo_ts + 20716,
		Fcodepoints: [2]uint32{
			0: uint32(365),
		},
	},
	1983: {
		Fname: __ccgo_ts + 20725,
		Fcodepoints: [2]uint32{
			0: uint32(251),
		},
	},
	1984: {
		Fname: __ccgo_ts + 20733,
		Fcodepoints: [2]uint32{
			0: uint32(1091),
		},
	},
	1985: {
		Fname: __ccgo_ts + 20739,
		Fcodepoints: [2]uint32{
			0: uint32(8645),
		},
	},
	1986: {
		Fname: __ccgo_ts + 20747,
		Fcodepoints: [2]uint32{
			0: uint32(369),
		},
	},
	1987: {
		Fname: __ccgo_ts + 20756,
		Fcodepoints: [2]uint32{
			0: uint32(10606),
		},
	},
	1988: {
		Fname: __ccgo_ts + 20764,
		Fcodepoints: [2]uint32{
			0: uint32(10622),
		},
	},
	1989: {
		Fname: __ccgo_ts + 20773,
		Fcodepoints: [2]uint32{
			0: uint32(120114),
		},
	},
	1990: {
		Fname: __ccgo_ts + 20779,
		Fcodepoints: [2]uint32{
			0: uint32(249),
		},
	},
	1991: {
		Fname: __ccgo_ts + 20788,
		Fcodepoints: [2]uint32{
			0: uint32(8639),
		},
	},
	1992: {
		Fname: __ccgo_ts + 20796,
		Fcodepoints: [2]uint32{
			0: uint32(8638),
		},
	},
	1993: {
		Fname: __ccgo_ts + 20804,
		Fcodepoints: [2]uint32{
			0: uint32(9600),
		},
	},
	1994: {
		Fname: __ccgo_ts + 20812,
		Fcodepoints: [2]uint32{
			0: uint32(8988),
		},
	},
	1995: {
		Fname: __ccgo_ts + 20821,
		Fcodepoints: [2]uint32{
			0: uint32(8988),
		},
	},
	1996: {
		Fname: __ccgo_ts + 20832,
		Fcodepoints: [2]uint32{
			0: uint32(8975),
		},
	},
	1997: {
		Fname: __ccgo_ts + 20841,
		Fcodepoints: [2]uint32{
			0: uint32(9720),
		},
	},
	1998: {
		Fname: __ccgo_ts + 20849,
		Fcodepoints: [2]uint32{
			0: uint32(363),
		},
	},
	1999: {
		Fname: __ccgo_ts + 20857,
		Fcodepoints: [2]uint32{
			0: uint32(168),
		},
	},
	2000: {
		Fname: __ccgo_ts + 20863,
		Fcodepoints: [2]uint32{
			0: uint32(371),
		},
	},
	2001: {
		Fname: __ccgo_ts + 20871,
		Fcodepoints: [2]uint32{
			0: uint32(120166),
		},
	},
	2002: {
		Fname: __ccgo_ts + 20878,
		Fcodepoints: [2]uint32{
			0: uint32(8593),
		},
	},
	2003: {
		Fname: __ccgo_ts + 20888,
		Fcodepoints: [2]uint32{
			0: uint32(8597),
		},
	},
	2004: {
		Fname: __ccgo_ts + 20902,
		Fcodepoints: [2]uint32{
			0: uint32(8639),
		},
	},
	2005: {
		Fname: __ccgo_ts + 20918,
		Fcodepoints: [2]uint32{
			0: uint32(8638),
		},
	},
	2006: {
		Fname: __ccgo_ts + 20935,
		Fcodepoints: [2]uint32{
			0: uint32(8846),
		},
	},
	2007: {
		Fname: __ccgo_ts + 20943,
		Fcodepoints: [2]uint32{
			0: uint32(965),
		},
	},
	2008: {
		Fname: __ccgo_ts + 20950,
		Fcodepoints: [2]uint32{
			0: uint32(978),
		},
	},
	2009: {
		Fname: __ccgo_ts + 20958,
		Fcodepoints: [2]uint32{
			0: uint32(965),
		},
	},
	2010: {
		Fname: __ccgo_ts + 20968,
		Fcodepoints: [2]uint32{
			0: uint32(8648),
		},
	},
	2011: {
		Fname: __ccgo_ts + 20981,
		Fcodepoints: [2]uint32{
			0: uint32(8989),
		},
	},
	2012: {
		Fname: __ccgo_ts + 20990,
		Fcodepoints: [2]uint32{
			0: uint32(8989),
		},
	},
	2013: {
		Fname: __ccgo_ts + 21001,
		Fcodepoints: [2]uint32{
			0: uint32(8974),
		},
	},
	2014: {
		Fname: __ccgo_ts + 21010,
		Fcodepoints: [2]uint32{
			0: uint32(367),
		},
	},
	2015: {
		Fname: __ccgo_ts + 21018,
		Fcodepoints: [2]uint32{
			0: uint32(9721),
		},
	},
	2016: {
		Fname: __ccgo_ts + 21026,
		Fcodepoints: [2]uint32{
			0: uint32(120010),
		},
	},
	2017: {
		Fname: __ccgo_ts + 21033,
		Fcodepoints: [2]uint32{
			0: uint32(8944),
		},
	},
	2018: {
		Fname: __ccgo_ts + 21041,
		Fcodepoints: [2]uint32{
			0: uint32(361),
		},
	},
	2019: {
		Fname: __ccgo_ts + 21050,
		Fcodepoints: [2]uint32{
			0: uint32(9653),
		},
	},
	2020: {
		Fname: __ccgo_ts + 21057,
		Fcodepoints: [2]uint32{
			0: uint32(9652),
		},
	},
	2021: {
		Fname: __ccgo_ts + 21065,
		Fcodepoints: [2]uint32{
			0: uint32(8648),
		},
	},
	2022: {
		Fname: __ccgo_ts + 21073,
		Fcodepoints: [2]uint32{
			0: uint32(252),
		},
	},
	2023: {
		Fname: __ccgo_ts + 21080,
		Fcodepoints: [2]uint32{
			0: uint32(10663),
		},
	},
	2024: {
		Fname: __ccgo_ts + 21090,
		Fcodepoints: [2]uint32{
			0: uint32(8661),
		},
	},
	2025: {
		Fname: __ccgo_ts + 21097,
		Fcodepoints: [2]uint32{
			0: uint32(10984),
		},
	},
	2026: {
		Fname: __ccgo_ts + 21104,
		Fcodepoints: [2]uint32{
			0: uint32(10985),
		},
	},
	2027: {
		Fname: __ccgo_ts + 21112,
		Fcodepoints: [2]uint32{
			0: uint32(8872),
		},
	},
	2028: {
		Fname: __ccgo_ts + 21120,
		Fcodepoints: [2]uint32{
			0: uint32(10652),
		},
	},
	2029: {
		Fname: __ccgo_ts + 21129,
		Fcodepoints: [2]uint32{
			0: uint32(1013),
		},
	},
	2030: {
		Fname: __ccgo_ts + 21142,
		Fcodepoints: [2]uint32{
			0: uint32(1008),
		},
	},
	2031: {
		Fname: __ccgo_ts + 21153,
		Fcodepoints: [2]uint32{
			0: uint32(8709),
		},
	},
	2032: {
		Fname: __ccgo_ts + 21166,
		Fcodepoints: [2]uint32{
			0: uint32(981),
		},
	},
	2033: {
		Fname: __ccgo_ts + 21175,
		Fcodepoints: [2]uint32{
			0: uint32(982),
		},
	},
	2034: {
		Fname: __ccgo_ts + 21183,
		Fcodepoints: [2]uint32{
			0: uint32(8733),
		},
	},
	2035: {
		Fname: __ccgo_ts + 21195,
		Fcodepoints: [2]uint32{
			0: uint32(8597),
		},
	},
	2036: {
		Fname: __ccgo_ts + 21202,
		Fcodepoints: [2]uint32{
			0: uint32(1009),
		},
	},
	2037: {
		Fname: __ccgo_ts + 21211,
		Fcodepoints: [2]uint32{
			0: uint32(962),
		},
	},
	2038: {
		Fname: __ccgo_ts + 21222,
		Fcodepoints: [2]uint32{
			0: uint32(8842),
			1: uint32(65024),
		},
	},
	2039: {
		Fname: __ccgo_ts + 21237,
		Fcodepoints: [2]uint32{
			0: uint32(10955),
			1: uint32(65024),
		},
	},
	2040: {
		Fname: __ccgo_ts + 21253,
		Fcodepoints: [2]uint32{
			0: uint32(8843),
			1: uint32(65024),
		},
	},
	2041: {
		Fname: __ccgo_ts + 21268,
		Fcodepoints: [2]uint32{
			0: uint32(10956),
			1: uint32(65024),
		},
	},
	2042: {
		Fname: __ccgo_ts + 21284,
		Fcodepoints: [2]uint32{
			0: uint32(977),
		},
	},
	2043: {
		Fname: __ccgo_ts + 21295,
		Fcodepoints: [2]uint32{
			0: uint32(8882),
		},
	},
	2044: {
		Fname: __ccgo_ts + 21313,
		Fcodepoints: [2]uint32{
			0: uint32(8883),
		},
	},
	2045: {
		Fname: __ccgo_ts + 21332,
		Fcodepoints: [2]uint32{
			0: uint32(1074),
		},
	},
	2046: {
		Fname: __ccgo_ts + 21338,
		Fcodepoints: [2]uint32{
			0: uint32(8866),
		},
	},
	2047: {
		Fname: __ccgo_ts + 21346,
		Fcodepoints: [2]uint32{
			0: uint32(8744),
		},
	},
	2048: {
		Fname: __ccgo_ts + 21352,
		Fcodepoints: [2]uint32{
			0: uint32(8891),
		},
	},
	2049: {
		Fname: __ccgo_ts + 21361,
		Fcodepoints: [2]uint32{
			0: uint32(8794),
		},
	},
	2050: {
		Fname: __ccgo_ts + 21369,
		Fcodepoints: [2]uint32{
			0: uint32(8942),
		},
	},
	2051: {
		Fname: __ccgo_ts + 21378,
		Fcodepoints: [2]uint32{
			0: uint32(124),
		},
	},
	2052: {
		Fname: __ccgo_ts + 21387,
		Fcodepoints: [2]uint32{
			0: uint32(124),
		},
	},
	2053: {
		Fname: __ccgo_ts + 21394,
		Fcodepoints: [2]uint32{
			0: uint32(120115),
		},
	},
	2054: {
		Fname: __ccgo_ts + 21400,
		Fcodepoints: [2]uint32{
			0: uint32(8882),
		},
	},
	2055: {
		Fname: __ccgo_ts + 21408,
		Fcodepoints: [2]uint32{
			0: uint32(8834),
			1: uint32(8402),
		},
	},
	2056: {
		Fname: __ccgo_ts + 21416,
		Fcodepoints: [2]uint32{
			0: uint32(8835),
			1: uint32(8402),
		},
	},
	2057: {
		Fname: __ccgo_ts + 21424,
		Fcodepoints: [2]uint32{
			0: uint32(120167),
		},
	},
	2058: {
		Fname: __ccgo_ts + 21431,
		Fcodepoints: [2]uint32{
			0: uint32(8733),
		},
	},
	2059: {
		Fname: __ccgo_ts + 21439,
		Fcodepoints: [2]uint32{
			0: uint32(8883),
		},
	},
	2060: {
		Fname: __ccgo_ts + 21447,
		Fcodepoints: [2]uint32{
			0: uint32(120011),
		},
	},
	2061: {
		Fname: __ccgo_ts + 21454,
		Fcodepoints: [2]uint32{
			0: uint32(10955),
			1: uint32(65024),
		},
	},
	2062: {
		Fname: __ccgo_ts + 21463,
		Fcodepoints: [2]uint32{
			0: uint32(8842),
			1: uint32(65024),
		},
	},
	2063: {
		Fname: __ccgo_ts + 21472,
		Fcodepoints: [2]uint32{
			0: uint32(10956),
			1: uint32(65024),
		},
	},
	2064: {
		Fname: __ccgo_ts + 21481,
		Fcodepoints: [2]uint32{
			0: uint32(8843),
			1: uint32(65024),
		},
	},
	2065: {
		Fname: __ccgo_ts + 21490,
		Fcodepoints: [2]uint32{
			0: uint32(10650),
		},
	},
	2066: {
		Fname: __ccgo_ts + 21500,
		Fcodepoints: [2]uint32{
			0: uint32(373),
		},
	},
	2067: {
		Fname: __ccgo_ts + 21508,
		Fcodepoints: [2]uint32{
			0: uint32(10847),
		},
	},
	2068: {
		Fname: __ccgo_ts + 21517,
		Fcodepoints: [2]uint32{
			0: uint32(8743),
		},
	},
	2069: {
		Fname: __ccgo_ts + 21525,
		Fcodepoints: [2]uint32{
			0: uint32(8793),
		},
	},
	2070: {
		Fname: __ccgo_ts + 21534,
		Fcodepoints: [2]uint32{
			0: uint32(8472),
		},
	},
	2071: {
		Fname: __ccgo_ts + 21543,
		Fcodepoints: [2]uint32{
			0: uint32(120116),
		},
	},
	2072: {
		Fname: __ccgo_ts + 21549,
		Fcodepoints: [2]uint32{
			0: uint32(120168),
		},
	},
	2073: {
		Fname: __ccgo_ts + 21556,
		Fcodepoints: [2]uint32{
			0: uint32(8472),
		},
	},
	2074: {
		Fname: __ccgo_ts + 21561,
		Fcodepoints: [2]uint32{
			0: uint32(8768),
		},
	},
	2075: {
		Fname: __ccgo_ts + 21566,
		Fcodepoints: [2]uint32{
			0: uint32(8768),
		},
	},
	2076: {
		Fname: __ccgo_ts + 21575,
		Fcodepoints: [2]uint32{
			0: uint32(120012),
		},
	},
	2077: {
		Fname: __ccgo_ts + 21582,
		Fcodepoints: [2]uint32{
			0: uint32(8898),
		},
	},
	2078: {
		Fname: __ccgo_ts + 21589,
		Fcodepoints: [2]uint32{
			0: uint32(9711),
		},
	},
	2079: {
		Fname: __ccgo_ts + 21597,
		Fcodepoints: [2]uint32{
			0: uint32(8899),
		},
	},
	2080: {
		Fname: __ccgo_ts + 21604,
		Fcodepoints: [2]uint32{
			0: uint32(9661),
		},
	},
	2081: {
		Fname: __ccgo_ts + 21612,
		Fcodepoints: [2]uint32{
			0: uint32(120117),
		},
	},
	2082: {
		Fname: __ccgo_ts + 21618,
		Fcodepoints: [2]uint32{
			0: uint32(10234),
		},
	},
	2083: {
		Fname: __ccgo_ts + 21626,
		Fcodepoints: [2]uint32{
			0: uint32(10231),
		},
	},
	2084: {
		Fname: __ccgo_ts + 21634,
		Fcodepoints: [2]uint32{
			0: uint32(958),
		},
	},
	2085: {
		Fname: __ccgo_ts + 21639,
		Fcodepoints: [2]uint32{
			0: uint32(10232),
		},
	},
	2086: {
		Fname: __ccgo_ts + 21647,
		Fcodepoints: [2]uint32{
			0: uint32(10229),
		},
	},
	2087: {
		Fname: __ccgo_ts + 21655,
		Fcodepoints: [2]uint32{
			0: uint32(10236),
		},
	},
	2088: {
		Fname: __ccgo_ts + 21662,
		Fcodepoints: [2]uint32{
			0: uint32(8955),
		},
	},
	2089: {
		Fname: __ccgo_ts + 21669,
		Fcodepoints: [2]uint32{
			0: uint32(10752),
		},
	},
	2090: {
		Fname: __ccgo_ts + 21677,
		Fcodepoints: [2]uint32{
			0: uint32(120169),
		},
	},
	2091: {
		Fname: __ccgo_ts + 21684,
		Fcodepoints: [2]uint32{
			0: uint32(10753),
		},
	},
	2092: {
		Fname: __ccgo_ts + 21693,
		Fcodepoints: [2]uint32{
			0: uint32(10754),
		},
	},
	2093: {
		Fname: __ccgo_ts + 21702,
		Fcodepoints: [2]uint32{
			0: uint32(10233),
		},
	},
	2094: {
		Fname: __ccgo_ts + 21710,
		Fcodepoints: [2]uint32{
			0: uint32(10230),
		},
	},
	2095: {
		Fname: __ccgo_ts + 21718,
		Fcodepoints: [2]uint32{
			0: uint32(120013),
		},
	},
	2096: {
		Fname: __ccgo_ts + 21725,
		Fcodepoints: [2]uint32{
			0: uint32(10758),
		},
	},
	2097: {
		Fname: __ccgo_ts + 21734,
		Fcodepoints: [2]uint32{
			0: uint32(10756),
		},
	},
	2098: {
		Fname: __ccgo_ts + 21743,
		Fcodepoints: [2]uint32{
			0: uint32(9651),
		},
	},
	2099: {
		Fname: __ccgo_ts + 21751,
		Fcodepoints: [2]uint32{
			0: uint32(8897),
		},
	},
	2100: {
		Fname: __ccgo_ts + 21758,
		Fcodepoints: [2]uint32{
			0: uint32(8896),
		},
	},
	2101: {
		Fname: __ccgo_ts + 21767,
		Fcodepoints: [2]uint32{
			0: uint32(253),
		},
	},
	2102: {
		Fname: __ccgo_ts + 21776,
		Fcodepoints: [2]uint32{
			0: uint32(1103),
		},
	},
	2103: {
		Fname: __ccgo_ts + 21783,
		Fcodepoints: [2]uint32{
			0: uint32(375),
		},
	},
	2104: {
		Fname: __ccgo_ts + 21791,
		Fcodepoints: [2]uint32{
			0: uint32(1099),
		},
	},
	2105: {
		Fname: __ccgo_ts + 21797,
		Fcodepoints: [2]uint32{
			0: uint32(165),
		},
	},
	2106: {
		Fname: __ccgo_ts + 21803,
		Fcodepoints: [2]uint32{
			0: uint32(120118),
		},
	},
	2107: {
		Fname: __ccgo_ts + 21809,
		Fcodepoints: [2]uint32{
			0: uint32(1111),
		},
	},
	2108: {
		Fname: __ccgo_ts + 21816,
		Fcodepoints: [2]uint32{
			0: uint32(120170),
		},
	},
	2109: {
		Fname: __ccgo_ts + 21823,
		Fcodepoints: [2]uint32{
			0: uint32(120014),
		},
	},
	2110: {
		Fname: __ccgo_ts + 21830,
		Fcodepoints: [2]uint32{
			0: uint32(1102),
		},
	},
	2111: {
		Fname: __ccgo_ts + 21837,
		Fcodepoints: [2]uint32{
			0: uint32(255),
		},
	},
	2112: {
		Fname: __ccgo_ts + 21844,
		Fcodepoints: [2]uint32{
			0: uint32(378),
		},
	},
	2113: {
		Fname: __ccgo_ts + 21853,
		Fcodepoints: [2]uint32{
			0: uint32(382),
		},
	},
	2114: {
		Fname: __ccgo_ts + 21862,
		Fcodepoints: [2]uint32{
			0: uint32(1079),
		},
	},
	2115: {
		Fname: __ccgo_ts + 21868,
		Fcodepoints: [2]uint32{
			0: uint32(380),
		},
	},
	2116: {
		Fname: __ccgo_ts + 21875,
		Fcodepoints: [2]uint32{
			0: uint32(8488),
		},
	},
	2117: {
		Fname: __ccgo_ts + 21884,
		Fcodepoints: [2]uint32{
			0: uint32(950),
		},
	},
	2118: {
		Fname: __ccgo_ts + 21891,
		Fcodepoints: [2]uint32{
			0: uint32(120119),
		},
	},
	2119: {
		Fname: __ccgo_ts + 21897,
		Fcodepoints: [2]uint32{
			0: uint32(1078),
		},
	},
	2120: {
		Fname: __ccgo_ts + 21904,
		Fcodepoints: [2]uint32{
			0: uint32(8669),
		},
	},
	2121: {
		Fname: __ccgo_ts + 21914,
		Fcodepoints: [2]uint32{
			0: uint32(120171),
		},
	},
	2122: {
		Fname: __ccgo_ts + 21921,
		Fcodepoints: [2]uint32{
			0: uint32(120015),
		},
	},
	2123: {
		Fname: __ccgo_ts + 21928,
		Fcodepoints: [2]uint32{
			0: uint32(8205),
		},
	},
	2124: {
		Fname: __ccgo_ts + 21934,
		Fcodepoints: [2]uint32{
			0: uint32(8204),
		},
	},
}

type ENTITY_KEY = struct {
	Fname      uintptr
	Fname_size size_t
}

type ENTITY_KEY_tag = ENTITY_KEY

func entity_cmp(tls *libc.TLS, p_key uintptr, p_entity uintptr) (r int32) {
	var ent, key uintptr
	_, _ = ent, key
	key = p_key
	ent = p_entity
	return libc.Xstrncmp(tls, (*ENTITY_KEY)(unsafe.Pointer(key)).Fname, (*ENTITY)(unsafe.Pointer(ent)).Fname, (*ENTITY_KEY)(unsafe.Pointer(key)).Fname_size)
}

func entity_lookup(tls *libc.TLS, name uintptr, name_size size_t) (r uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _ /* key at bp+0 */ ENTITY_KEY
	*(*ENTITY_KEY)(__ccgo_ptr(bp)) = ENTITY_KEY{
		Fname:      name,
		Fname_size: name_size,
	}
	return libc.Xbsearch(tls, bp, uintptr(unsafe.Pointer(&ENTITY_MAP)), libc.Uint64FromInt64(34000)/libc.Uint64FromInt64(16), uint64(16), __ccgo_fp(entity_cmp))
}

const MD_FLAG_LATEXMATHSPANS2 = 4096
const MD_FLAG_PERMISSIVEEMAILAUTOLINKS2 = 8
const MD_FLAG_PERMISSIVEURLAUTOLINKS2 = 4
const MD_FLAG_PERMISSIVEWWWAUTOLINKS2 = 1024
const MD_FLAG_STRIKETHROUGH2 = 512
const MD_FLAG_TABLES2 = 256
const MD_FLAG_TASKLISTS2 = 2048
const MD_HTML_FLAG_DEBUG1 = 0x0001
const MD_HTML_FLAG_SKIP_UTF8_BOM1 = 0x0004
const MD_HTML_FLAG_VERBATIM_ENTITIES1 = 0x0002
const MD_HTML_FLAG_XHTML1 = 0x0008

type ghmd_md_buffer = struct {
	Fdata     uintptr
	Flength   size_t
	Fcapacity size_t
	Ffailed   int32
}

func ghmd_md_collect(tls *libc.TLS, text uintptr, size MD_SIZE, userdata uintptr) {
	var buffer, data uintptr
	var capacity size_t
	var v1 uint64
	_, _, _, _ = buffer, capacity, data, v1
	buffer = userdata
	if (*ghmd_md_buffer)(unsafe.Pointer(buffer)).Ffailed != 0 {
		return
	}
	if (*ghmd_md_buffer)(unsafe.Pointer(buffer)).Flength+uint64(size)+uint64(1) > (*ghmd_md_buffer)(unsafe.Pointer(buffer)).Fcapacity {
		if (*ghmd_md_buffer)(unsafe.Pointer(buffer)).Fcapacity != 0 {
			v1 = (*ghmd_md_buffer)(unsafe.Pointer(buffer)).Fcapacity
		} else {
			v1 = uint64(1024)
		}
		capacity = v1
		for capacity < (*ghmd_md_buffer)(unsafe.Pointer(buffer)).Flength+uint64(size)+uint64(1) {
			capacity = capacity * uint64(2)
		}
		data = libc.Xrealloc(tls, (*ghmd_md_buffer)(unsafe.Pointer(buffer)).Fdata, capacity)
		if !(data != 0) {
			(*ghmd_md_buffer)(unsafe.Pointer(buffer)).Ffailed = int32(1)
			return
		}
		(*ghmd_md_buffer)(unsafe.Pointer(buffer)).Fdata = data
		(*ghmd_md_buffer)(unsafe.Pointer(buffer)).Fcapacity = capacity
	}
	libc.Xmemcpy(tls, (*ghmd_md_buffer)(unsafe.Pointer(buffer)).Fdata+uintptr((*ghmd_md_buffer)(unsafe.Pointer(buffer)).Flength), text, uint64(size))
	*(*size_t)(__ccgo_ptr(buffer + 8)) += uint64(size)
	*(*uint8)(__ccgo_ptr((*ghmd_md_buffer)(unsafe.Pointer(buffer)).Fdata + uintptr((*ghmd_md_buffer)(unsafe.Pointer(buffer)).Flength))) = uint8(0)
}

func bridge_md4c_render(tls *libc.TLS, source uintptr, length size_t, github_dialect int32) (r uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var flags uint32
	var result, v1 int32
	var _ /* buffer at bp+0 */ ghmd_md_buffer
	_, _, _ = flags, result, v1
	*(*ghmd_md_buffer)(__ccgo_ptr(bp)) = ghmd_md_buffer{}
	if github_dialect != 0 {
		v1 = libc.Int32FromInt32(MD_FLAG_PERMISSIVEEMAILAUTOLINKS2) | libc.Int32FromInt32(MD_FLAG_PERMISSIVEURLAUTOLINKS2) | libc.Int32FromInt32(MD_FLAG_PERMISSIVEWWWAUTOLINKS2) | libc.Int32FromInt32(MD_FLAG_TABLES2) | libc.Int32FromInt32(MD_FLAG_STRIKETHROUGH2) | libc.Int32FromInt32(MD_FLAG_TASKLISTS2) | libc.Int32FromInt32(MD_FLAG_LATEXMATHSPANS2)
	} else {
		v1 = 0
	}
	flags = libc.Uint32FromInt32(v1)
	result = md_html(tls, source, uint32(length), __ccgo_fp(ghmd_md_collect), bp, flags, uint32(0))
	if result != 0 || (*(*ghmd_md_buffer)(__ccgo_ptr(bp))).Ffailed != 0 {
		libc.Xfree(tls, (*(*ghmd_md_buffer)(__ccgo_ptr(bp))).Fdata)
		return libc.UintptrFromInt32(0)
	}
	if !((*(*ghmd_md_buffer)(__ccgo_ptr(bp))).Fdata != 0) {
		(*(*ghmd_md_buffer)(__ccgo_ptr(bp))).Fdata = libc.Xcalloc(tls, uint64(1), uint64(1))
	}
	return (*(*ghmd_md_buffer)(__ccgo_ptr(bp))).Fdata
}

func bridge_md4c_free(tls *libc.TLS, value uintptr) {
	libc.Xfree(tls, value)
}

func __ccgo_fp(f interface{}) uintptr {
	type iface [2]uintptr
	return (*iface)(unsafe.Pointer(&f))[1]
}

func __ccgo_ptr(p uintptr) unsafe.Pointer {
	return unsafe.Pointer(p)
}

var __ccgo_ts = (*reflect.StringHeader)(unsafe.Pointer(&__ccgo_ts1)).Data

var __ccgo_ts1 = "\x00malloc() failed.\x00\"'=<>`\x00_.:-\x00-->\x00?>\x00>\x00]]>\x00realloc() failed.\x00.!#$%&'*+/=?^_`{|}~-\x00http\x00//\x00https\x00ftp\x00www\x00@:.\x00.-_\x00/.-_\x00&.-+_=()\x00.-+_\x00.-_+\x00({[\x00*_~\x00)}].!?,;\x00[]!\x00|\x00&\x00*_~$\x00Aborted from enter_span() callback.\x00Aborted from leave_span() callback.\x00Aborted from text() callback.\x00 \x00http://\x00mailto:\x00\n\x00Aborted from enter_block() callback.\x00Aborted from leave_block() callback.\x00Suppressing table (column_count >128)\x00pre\x00script\x00style\x00textarea\x00address\x00article\x00aside\x00base\x00basefont\x00blockquote\x00body\x00caption\x00center\x00col\x00colgroup\x00dd\x00details\x00dialog\x00dir\x00div\x00dl\x00dt\x00fieldset\x00figcaption\x00figure\x00footer\x00form\x00frame\x00frameset\x00h1\x00h2\x00h3\x00h4\x00h5\x00h6\x00head\x00header\x00hr\x00html\x00iframe\x00legend\x00li\x00link\x00main\x00menu\x00menuitem\x00nav\x00noframes\x00ol\x00optgroup\x00option\x00p\x00param\x00search\x00section\x00summary\x00table\x00tbody\x00td\x00tfoot\x00th\x00thead\x00title\x00tr\x00track\x00ul\x00![CDATA[\x00-+*\x00-_*\x00-+*.)\x00xX \x00\r\n\x00Unsupported abi_version.\x00&amp;\x00&lt;\x00&gt;\x00&quot;\x00<ol>\n\x00<ol start=\"%u\">\n\x00<li class=\"task-list-item\"><input type=\"checkbox\" class=\"task-list-item-checkbox\" disabled\x00 checked\x00<li>\x00<pre><code\x00 class=\"language-\x00\"\x00<\x00 align=\"left\">\x00 align=\"center\">\x00 align=\"right\">\x00<a href=\"\x00\" title=\"\x00\">\x00<img src=\"\x00\" alt=\"\x00\" />\x00<x-wikilink data-target=\"\x00<h1>\x00<h2>\x00<h3>\x00<h4>\x00<h5>\x00<h6>\x00<blockquote>\n\x00<ul>\n\x00<hr />\n\x00<hr>\n\x00<p>\x00<table>\n\x00<thead>\n\x00<tbody>\n\x00<tr>\n\x00</h1>\n\x00</h2>\n\x00</h3>\n\x00</h4>\n\x00</h5>\n\x00</h6>\n\x00</blockquote>\n\x00</ul>\n\x00</ol>\n\x00</li>\n\x00</code></pre>\n\x00</p>\n\x00</table>\n\x00</thead>\n\x00</tbody>\n\x00</tr>\n\x00</th>\n\x00</td>\n\x00<em>\x00<strong>\x00<u>\x00<code>\x00<del>\x00<x-equation>\x00<x-equation type=\"display\">\x00</em>\x00</strong>\x00</u>\x00</a>\x00</code>\x00</del>\x00</x-equation>\x00</x-wikilink>\x00<br />\n\x00<br>\n\x00MD4C: %s\n\x00\"&<>\x00~-_.+!*(),%#@?=;:/,+$\x00&AElig;\x00&AMP;\x00&Aacute;\x00&Abreve;\x00&Acirc;\x00&Acy;\x00&Afr;\x00&Agrave;\x00&Alpha;\x00&Amacr;\x00&And;\x00&Aogon;\x00&Aopf;\x00&ApplyFunction;\x00&Aring;\x00&Ascr;\x00&Assign;\x00&Atilde;\x00&Auml;\x00&Backslash;\x00&Barv;\x00&Barwed;\x00&Bcy;\x00&Because;\x00&Bernoullis;\x00&Beta;\x00&Bfr;\x00&Bopf;\x00&Breve;\x00&Bscr;\x00&Bumpeq;\x00&CHcy;\x00&COPY;\x00&Cacute;\x00&Cap;\x00&CapitalDifferentialD;\x00&Cayleys;\x00&Ccaron;\x00&Ccedil;\x00&Ccirc;\x00&Cconint;\x00&Cdot;\x00&Cedilla;\x00&CenterDot;\x00&Cfr;\x00&Chi;\x00&CircleDot;\x00&CircleMinus;\x00&CirclePlus;\x00&CircleTimes;\x00&ClockwiseContourIntegral;\x00&CloseCurlyDoubleQuote;\x00&CloseCurlyQuote;\x00&Colon;\x00&Colone;\x00&Congruent;\x00&Conint;\x00&ContourIntegral;\x00&Copf;\x00&Coproduct;\x00&CounterClockwiseContourIntegral;\x00&Cross;\x00&Cscr;\x00&Cup;\x00&CupCap;\x00&DD;\x00&DDotrahd;\x00&DJcy;\x00&DScy;\x00&DZcy;\x00&Dagger;\x00&Darr;\x00&Dashv;\x00&Dcaron;\x00&Dcy;\x00&Del;\x00&Delta;\x00&Dfr;\x00&DiacriticalAcute;\x00&DiacriticalDot;\x00&DiacriticalDoubleAcute;\x00&DiacriticalGrave;\x00&DiacriticalTilde;\x00&Diamond;\x00&DifferentialD;\x00&Dopf;\x00&Dot;\x00&DotDot;\x00&DotEqual;\x00&DoubleContourIntegral;\x00&DoubleDot;\x00&DoubleDownArrow;\x00&DoubleLeftArrow;\x00&DoubleLeftRightArrow;\x00&DoubleLeftTee;\x00&DoubleLongLeftArrow;\x00&DoubleLongLeftRightArrow;\x00&DoubleLongRightArrow;\x00&DoubleRightArrow;\x00&DoubleRightTee;\x00&DoubleUpArrow;\x00&DoubleUpDownArrow;\x00&DoubleVerticalBar;\x00&DownArrow;\x00&DownArrowBar;\x00&DownArrowUpArrow;\x00&DownBreve;\x00&DownLeftRightVector;\x00&DownLeftTeeVector;\x00&DownLeftVector;\x00&DownLeftVectorBar;\x00&DownRightTeeVector;\x00&DownRightVector;\x00&DownRightVectorBar;\x00&DownTee;\x00&DownTeeArrow;\x00&Downarrow;\x00&Dscr;\x00&Dstrok;\x00&ENG;\x00&ETH;\x00&Eacute;\x00&Ecaron;\x00&Ecirc;\x00&Ecy;\x00&Edot;\x00&Efr;\x00&Egrave;\x00&Element;\x00&Emacr;\x00&EmptySmallSquare;\x00&EmptyVerySmallSquare;\x00&Eogon;\x00&Eopf;\x00&Epsilon;\x00&Equal;\x00&EqualTilde;\x00&Equilibrium;\x00&Escr;\x00&Esim;\x00&Eta;\x00&Euml;\x00&Exists;\x00&ExponentialE;\x00&Fcy;\x00&Ffr;\x00&FilledSmallSquare;\x00&FilledVerySmallSquare;\x00&Fopf;\x00&ForAll;\x00&Fouriertrf;\x00&Fscr;\x00&GJcy;\x00&GT;\x00&Gamma;\x00&Gammad;\x00&Gbreve;\x00&Gcedil;\x00&Gcirc;\x00&Gcy;\x00&Gdot;\x00&Gfr;\x00&Gg;\x00&Gopf;\x00&GreaterEqual;\x00&GreaterEqualLess;\x00&GreaterFullEqual;\x00&GreaterGreater;\x00&GreaterLess;\x00&GreaterSlantEqual;\x00&GreaterTilde;\x00&Gscr;\x00&Gt;\x00&HARDcy;\x00&Hacek;\x00&Hat;\x00&Hcirc;\x00&Hfr;\x00&HilbertSpace;\x00&Hopf;\x00&HorizontalLine;\x00&Hscr;\x00&Hstrok;\x00&HumpDownHump;\x00&HumpEqual;\x00&IEcy;\x00&IJlig;\x00&IOcy;\x00&Iacute;\x00&Icirc;\x00&Icy;\x00&Idot;\x00&Ifr;\x00&Igrave;\x00&Im;\x00&Imacr;\x00&ImaginaryI;\x00&Implies;\x00&Int;\x00&Integral;\x00&Intersection;\x00&InvisibleComma;\x00&InvisibleTimes;\x00&Iogon;\x00&Iopf;\x00&Iota;\x00&Iscr;\x00&Itilde;\x00&Iukcy;\x00&Iuml;\x00&Jcirc;\x00&Jcy;\x00&Jfr;\x00&Jopf;\x00&Jscr;\x00&Jsercy;\x00&Jukcy;\x00&KHcy;\x00&KJcy;\x00&Kappa;\x00&Kcedil;\x00&Kcy;\x00&Kfr;\x00&Kopf;\x00&Kscr;\x00&LJcy;\x00&LT;\x00&Lacute;\x00&Lambda;\x00&Lang;\x00&Laplacetrf;\x00&Larr;\x00&Lcaron;\x00&Lcedil;\x00&Lcy;\x00&LeftAngleBracket;\x00&LeftArrow;\x00&LeftArrowBar;\x00&LeftArrowRightArrow;\x00&LeftCeiling;\x00&LeftDoubleBracket;\x00&LeftDownTeeVector;\x00&LeftDownVector;\x00&LeftDownVectorBar;\x00&LeftFloor;\x00&LeftRightArrow;\x00&LeftRightVector;\x00&LeftTee;\x00&LeftTeeArrow;\x00&LeftTeeVector;\x00&LeftTriangle;\x00&LeftTriangleBar;\x00&LeftTriangleEqual;\x00&LeftUpDownVector;\x00&LeftUpTeeVector;\x00&LeftUpVector;\x00&LeftUpVectorBar;\x00&LeftVector;\x00&LeftVectorBar;\x00&Leftarrow;\x00&Leftrightarrow;\x00&LessEqualGreater;\x00&LessFullEqual;\x00&LessGreater;\x00&LessLess;\x00&LessSlantEqual;\x00&LessTilde;\x00&Lfr;\x00&Ll;\x00&Lleftarrow;\x00&Lmidot;\x00&LongLeftArrow;\x00&LongLeftRightArrow;\x00&LongRightArrow;\x00&Longleftarrow;\x00&Longleftrightarrow;\x00&Longrightarrow;\x00&Lopf;\x00&LowerLeftArrow;\x00&LowerRightArrow;\x00&Lscr;\x00&Lsh;\x00&Lstrok;\x00&Lt;\x00&Map;\x00&Mcy;\x00&MediumSpace;\x00&Mellintrf;\x00&Mfr;\x00&MinusPlus;\x00&Mopf;\x00&Mscr;\x00&Mu;\x00&NJcy;\x00&Nacute;\x00&Ncaron;\x00&Ncedil;\x00&Ncy;\x00&NegativeMediumSpace;\x00&NegativeThickSpace;\x00&NegativeThinSpace;\x00&NegativeVeryThinSpace;\x00&NestedGreaterGreater;\x00&NestedLessLess;\x00&NewLine;\x00&Nfr;\x00&NoBreak;\x00&NonBreakingSpace;\x00&Nopf;\x00&Not;\x00&NotCongruent;\x00&NotCupCap;\x00&NotDoubleVerticalBar;\x00&NotElement;\x00&NotEqual;\x00&NotEqualTilde;\x00&NotExists;\x00&NotGreater;\x00&NotGreaterEqual;\x00&NotGreaterFullEqual;\x00&NotGreaterGreater;\x00&NotGreaterLess;\x00&NotGreaterSlantEqual;\x00&NotGreaterTilde;\x00&NotHumpDownHump;\x00&NotHumpEqual;\x00&NotLeftTriangle;\x00&NotLeftTriangleBar;\x00&NotLeftTriangleEqual;\x00&NotLess;\x00&NotLessEqual;\x00&NotLessGreater;\x00&NotLessLess;\x00&NotLessSlantEqual;\x00&NotLessTilde;\x00&NotNestedGreaterGreater;\x00&NotNestedLessLess;\x00&NotPrecedes;\x00&NotPrecedesEqual;\x00&NotPrecedesSlantEqual;\x00&NotReverseElement;\x00&NotRightTriangle;\x00&NotRightTriangleBar;\x00&NotRightTriangleEqual;\x00&NotSquareSubset;\x00&NotSquareSubsetEqual;\x00&NotSquareSuperset;\x00&NotSquareSupersetEqual;\x00&NotSubset;\x00&NotSubsetEqual;\x00&NotSucceeds;\x00&NotSucceedsEqual;\x00&NotSucceedsSlantEqual;\x00&NotSucceedsTilde;\x00&NotSuperset;\x00&NotSupersetEqual;\x00&NotTilde;\x00&NotTildeEqual;\x00&NotTildeFullEqual;\x00&NotTildeTilde;\x00&NotVerticalBar;\x00&Nscr;\x00&Ntilde;\x00&Nu;\x00&OElig;\x00&Oacute;\x00&Ocirc;\x00&Ocy;\x00&Odblac;\x00&Ofr;\x00&Ograve;\x00&Omacr;\x00&Omega;\x00&Omicron;\x00&Oopf;\x00&OpenCurlyDoubleQuote;\x00&OpenCurlyQuote;\x00&Or;\x00&Oscr;\x00&Oslash;\x00&Otilde;\x00&Otimes;\x00&Ouml;\x00&OverBar;\x00&OverBrace;\x00&OverBracket;\x00&OverParenthesis;\x00&PartialD;\x00&Pcy;\x00&Pfr;\x00&Phi;\x00&Pi;\x00&PlusMinus;\x00&Poincareplane;\x00&Popf;\x00&Pr;\x00&Precedes;\x00&PrecedesEqual;\x00&PrecedesSlantEqual;\x00&PrecedesTilde;\x00&Prime;\x00&Product;\x00&Proportion;\x00&Proportional;\x00&Pscr;\x00&Psi;\x00&QUOT;\x00&Qfr;\x00&Qopf;\x00&Qscr;\x00&RBarr;\x00&REG;\x00&Racute;\x00&Rang;\x00&Rarr;\x00&Rarrtl;\x00&Rcaron;\x00&Rcedil;\x00&Rcy;\x00&Re;\x00&ReverseElement;\x00&ReverseEquilibrium;\x00&ReverseUpEquilibrium;\x00&Rfr;\x00&Rho;\x00&RightAngleBracket;\x00&RightArrow;\x00&RightArrowBar;\x00&RightArrowLeftArrow;\x00&RightCeiling;\x00&RightDoubleBracket;\x00&RightDownTeeVector;\x00&RightDownVector;\x00&RightDownVectorBar;\x00&RightFloor;\x00&RightTee;\x00&RightTeeArrow;\x00&RightTeeVector;\x00&RightTriangle;\x00&RightTriangleBar;\x00&RightTriangleEqual;\x00&RightUpDownVector;\x00&RightUpTeeVector;\x00&RightUpVector;\x00&RightUpVectorBar;\x00&RightVector;\x00&RightVectorBar;\x00&Rightarrow;\x00&Ropf;\x00&RoundImplies;\x00&Rrightarrow;\x00&Rscr;\x00&Rsh;\x00&RuleDelayed;\x00&SHCHcy;\x00&SHcy;\x00&SOFTcy;\x00&Sacute;\x00&Sc;\x00&Scaron;\x00&Scedil;\x00&Scirc;\x00&Scy;\x00&Sfr;\x00&ShortDownArrow;\x00&ShortLeftArrow;\x00&ShortRightArrow;\x00&ShortUpArrow;\x00&Sigma;\x00&SmallCircle;\x00&Sopf;\x00&Sqrt;\x00&Square;\x00&SquareIntersection;\x00&SquareSubset;\x00&SquareSubsetEqual;\x00&SquareSuperset;\x00&SquareSupersetEqual;\x00&SquareUnion;\x00&Sscr;\x00&Star;\x00&Sub;\x00&Subset;\x00&SubsetEqual;\x00&Succeeds;\x00&SucceedsEqual;\x00&SucceedsSlantEqual;\x00&SucceedsTilde;\x00&SuchThat;\x00&Sum;\x00&Sup;\x00&Superset;\x00&SupersetEqual;\x00&Supset;\x00&THORN;\x00&TRADE;\x00&TSHcy;\x00&TScy;\x00&Tab;\x00&Tau;\x00&Tcaron;\x00&Tcedil;\x00&Tcy;\x00&Tfr;\x00&Therefore;\x00&Theta;\x00&ThickSpace;\x00&ThinSpace;\x00&Tilde;\x00&TildeEqual;\x00&TildeFullEqual;\x00&TildeTilde;\x00&Topf;\x00&TripleDot;\x00&Tscr;\x00&Tstrok;\x00&Uacute;\x00&Uarr;\x00&Uarrocir;\x00&Ubrcy;\x00&Ubreve;\x00&Ucirc;\x00&Ucy;\x00&Udblac;\x00&Ufr;\x00&Ugrave;\x00&Umacr;\x00&UnderBar;\x00&UnderBrace;\x00&UnderBracket;\x00&UnderParenthesis;\x00&Union;\x00&UnionPlus;\x00&Uogon;\x00&Uopf;\x00&UpArrow;\x00&UpArrowBar;\x00&UpArrowDownArrow;\x00&UpDownArrow;\x00&UpEquilibrium;\x00&UpTee;\x00&UpTeeArrow;\x00&Uparrow;\x00&Updownarrow;\x00&UpperLeftArrow;\x00&UpperRightArrow;\x00&Upsi;\x00&Upsilon;\x00&Uring;\x00&Uscr;\x00&Utilde;\x00&Uuml;\x00&VDash;\x00&Vbar;\x00&Vcy;\x00&Vdash;\x00&Vdashl;\x00&Vee;\x00&Verbar;\x00&Vert;\x00&VerticalBar;\x00&VerticalLine;\x00&VerticalSeparator;\x00&VerticalTilde;\x00&VeryThinSpace;\x00&Vfr;\x00&Vopf;\x00&Vscr;\x00&Vvdash;\x00&Wcirc;\x00&Wedge;\x00&Wfr;\x00&Wopf;\x00&Wscr;\x00&Xfr;\x00&Xi;\x00&Xopf;\x00&Xscr;\x00&YAcy;\x00&YIcy;\x00&YUcy;\x00&Yacute;\x00&Ycirc;\x00&Ycy;\x00&Yfr;\x00&Yopf;\x00&Yscr;\x00&Yuml;\x00&ZHcy;\x00&Zacute;\x00&Zcaron;\x00&Zcy;\x00&Zdot;\x00&ZeroWidthSpace;\x00&Zeta;\x00&Zfr;\x00&Zopf;\x00&Zscr;\x00&aacute;\x00&abreve;\x00&ac;\x00&acE;\x00&acd;\x00&acirc;\x00&acute;\x00&acy;\x00&aelig;\x00&af;\x00&afr;\x00&agrave;\x00&alefsym;\x00&aleph;\x00&alpha;\x00&amacr;\x00&amalg;\x00&and;\x00&andand;\x00&andd;\x00&andslope;\x00&andv;\x00&ang;\x00&ange;\x00&angle;\x00&angmsd;\x00&angmsdaa;\x00&angmsdab;\x00&angmsdac;\x00&angmsdad;\x00&angmsdae;\x00&angmsdaf;\x00&angmsdag;\x00&angmsdah;\x00&angrt;\x00&angrtvb;\x00&angrtvbd;\x00&angsph;\x00&angst;\x00&angzarr;\x00&aogon;\x00&aopf;\x00&ap;\x00&apE;\x00&apacir;\x00&ape;\x00&apid;\x00&apos;\x00&approx;\x00&approxeq;\x00&aring;\x00&ascr;\x00&ast;\x00&asymp;\x00&asympeq;\x00&atilde;\x00&auml;\x00&awconint;\x00&awint;\x00&bNot;\x00&backcong;\x00&backepsilon;\x00&backprime;\x00&backsim;\x00&backsimeq;\x00&barvee;\x00&barwed;\x00&barwedge;\x00&bbrk;\x00&bbrktbrk;\x00&bcong;\x00&bcy;\x00&bdquo;\x00&becaus;\x00&because;\x00&bemptyv;\x00&bepsi;\x00&bernou;\x00&beta;\x00&beth;\x00&between;\x00&bfr;\x00&bigcap;\x00&bigcirc;\x00&bigcup;\x00&bigodot;\x00&bigoplus;\x00&bigotimes;\x00&bigsqcup;\x00&bigstar;\x00&bigtriangledown;\x00&bigtriangleup;\x00&biguplus;\x00&bigvee;\x00&bigwedge;\x00&bkarow;\x00&blacklozenge;\x00&blacksquare;\x00&blacktriangle;\x00&blacktriangledown;\x00&blacktriangleleft;\x00&blacktriangleright;\x00&blank;\x00&blk12;\x00&blk14;\x00&blk34;\x00&block;\x00&bne;\x00&bnequiv;\x00&bnot;\x00&bopf;\x00&bot;\x00&bottom;\x00&bowtie;\x00&boxDL;\x00&boxDR;\x00&boxDl;\x00&boxDr;\x00&boxH;\x00&boxHD;\x00&boxHU;\x00&boxHd;\x00&boxHu;\x00&boxUL;\x00&boxUR;\x00&boxUl;\x00&boxUr;\x00&boxV;\x00&boxVH;\x00&boxVL;\x00&boxVR;\x00&boxVh;\x00&boxVl;\x00&boxVr;\x00&boxbox;\x00&boxdL;\x00&boxdR;\x00&boxdl;\x00&boxdr;\x00&boxh;\x00&boxhD;\x00&boxhU;\x00&boxhd;\x00&boxhu;\x00&boxminus;\x00&boxplus;\x00&boxtimes;\x00&boxuL;\x00&boxuR;\x00&boxul;\x00&boxur;\x00&boxv;\x00&boxvH;\x00&boxvL;\x00&boxvR;\x00&boxvh;\x00&boxvl;\x00&boxvr;\x00&bprime;\x00&breve;\x00&brvbar;\x00&bscr;\x00&bsemi;\x00&bsim;\x00&bsime;\x00&bsol;\x00&bsolb;\x00&bsolhsub;\x00&bull;\x00&bullet;\x00&bump;\x00&bumpE;\x00&bumpe;\x00&bumpeq;\x00&cacute;\x00&cap;\x00&capand;\x00&capbrcup;\x00&capcap;\x00&capcup;\x00&capdot;\x00&caps;\x00&caret;\x00&caron;\x00&ccaps;\x00&ccaron;\x00&ccedil;\x00&ccirc;\x00&ccups;\x00&ccupssm;\x00&cdot;\x00&cedil;\x00&cemptyv;\x00&cent;\x00&centerdot;\x00&cfr;\x00&chcy;\x00&check;\x00&checkmark;\x00&chi;\x00&cir;\x00&cirE;\x00&circ;\x00&circeq;\x00&circlearrowleft;\x00&circlearrowright;\x00&circledR;\x00&circledS;\x00&circledast;\x00&circledcirc;\x00&circleddash;\x00&cire;\x00&cirfnint;\x00&cirmid;\x00&cirscir;\x00&clubs;\x00&clubsuit;\x00&colon;\x00&colone;\x00&coloneq;\x00&comma;\x00&commat;\x00&comp;\x00&compfn;\x00&complement;\x00&complexes;\x00&cong;\x00&congdot;\x00&conint;\x00&copf;\x00&coprod;\x00&copy;\x00&copysr;\x00&crarr;\x00&cross;\x00&cscr;\x00&csub;\x00&csube;\x00&csup;\x00&csupe;\x00&ctdot;\x00&cudarrl;\x00&cudarrr;\x00&cuepr;\x00&cuesc;\x00&cularr;\x00&cularrp;\x00&cup;\x00&cupbrcap;\x00&cupcap;\x00&cupcup;\x00&cupdot;\x00&cupor;\x00&cups;\x00&curarr;\x00&curarrm;\x00&curlyeqprec;\x00&curlyeqsucc;\x00&curlyvee;\x00&curlywedge;\x00&curren;\x00&curvearrowleft;\x00&curvearrowright;\x00&cuvee;\x00&cuwed;\x00&cwconint;\x00&cwint;\x00&cylcty;\x00&dArr;\x00&dHar;\x00&dagger;\x00&daleth;\x00&darr;\x00&dash;\x00&dashv;\x00&dbkarow;\x00&dblac;\x00&dcaron;\x00&dcy;\x00&dd;\x00&ddagger;\x00&ddarr;\x00&ddotseq;\x00&deg;\x00&delta;\x00&demptyv;\x00&dfisht;\x00&dfr;\x00&dharl;\x00&dharr;\x00&diam;\x00&diamond;\x00&diamondsuit;\x00&diams;\x00&die;\x00&digamma;\x00&disin;\x00&div;\x00&divide;\x00&divideontimes;\x00&divonx;\x00&djcy;\x00&dlcorn;\x00&dlcrop;\x00&dollar;\x00&dopf;\x00&dot;\x00&doteq;\x00&doteqdot;\x00&dotminus;\x00&dotplus;\x00&dotsquare;\x00&doublebarwedge;\x00&downarrow;\x00&downdownarrows;\x00&downharpoonleft;\x00&downharpoonright;\x00&drbkarow;\x00&drcorn;\x00&drcrop;\x00&dscr;\x00&dscy;\x00&dsol;\x00&dstrok;\x00&dtdot;\x00&dtri;\x00&dtrif;\x00&duarr;\x00&duhar;\x00&dwangle;\x00&dzcy;\x00&dzigrarr;\x00&eDDot;\x00&eDot;\x00&eacute;\x00&easter;\x00&ecaron;\x00&ecir;\x00&ecirc;\x00&ecolon;\x00&ecy;\x00&edot;\x00&ee;\x00&efDot;\x00&efr;\x00&eg;\x00&egrave;\x00&egs;\x00&egsdot;\x00&el;\x00&elinters;\x00&ell;\x00&els;\x00&elsdot;\x00&emacr;\x00&empty;\x00&emptyset;\x00&emptyv;\x00&emsp13;\x00&emsp14;\x00&emsp;\x00&eng;\x00&ensp;\x00&eogon;\x00&eopf;\x00&epar;\x00&eparsl;\x00&eplus;\x00&epsi;\x00&epsilon;\x00&epsiv;\x00&eqcirc;\x00&eqcolon;\x00&eqsim;\x00&eqslantgtr;\x00&eqslantless;\x00&equals;\x00&equest;\x00&equiv;\x00&equivDD;\x00&eqvparsl;\x00&erDot;\x00&erarr;\x00&escr;\x00&esdot;\x00&esim;\x00&eta;\x00&eth;\x00&euml;\x00&euro;\x00&excl;\x00&exist;\x00&expectation;\x00&exponentiale;\x00&fallingdotseq;\x00&fcy;\x00&female;\x00&ffilig;\x00&fflig;\x00&ffllig;\x00&ffr;\x00&filig;\x00&fjlig;\x00&flat;\x00&fllig;\x00&fltns;\x00&fnof;\x00&fopf;\x00&forall;\x00&fork;\x00&forkv;\x00&fpartint;\x00&frac12;\x00&frac13;\x00&frac14;\x00&frac15;\x00&frac16;\x00&frac18;\x00&frac23;\x00&frac25;\x00&frac34;\x00&frac35;\x00&frac38;\x00&frac45;\x00&frac56;\x00&frac58;\x00&frac78;\x00&frasl;\x00&frown;\x00&fscr;\x00&gE;\x00&gEl;\x00&gacute;\x00&gamma;\x00&gammad;\x00&gap;\x00&gbreve;\x00&gcirc;\x00&gcy;\x00&gdot;\x00&ge;\x00&gel;\x00&geq;\x00&geqq;\x00&geqslant;\x00&ges;\x00&gescc;\x00&gesdot;\x00&gesdoto;\x00&gesdotol;\x00&gesl;\x00&gesles;\x00&gfr;\x00&gg;\x00&ggg;\x00&gimel;\x00&gjcy;\x00&gl;\x00&glE;\x00&gla;\x00&glj;\x00&gnE;\x00&gnap;\x00&gnapprox;\x00&gne;\x00&gneq;\x00&gneqq;\x00&gnsim;\x00&gopf;\x00&grave;\x00&gscr;\x00&gsim;\x00&gsime;\x00&gsiml;\x00&gtcc;\x00&gtcir;\x00&gtdot;\x00&gtlPar;\x00&gtquest;\x00&gtrapprox;\x00&gtrarr;\x00&gtrdot;\x00&gtreqless;\x00&gtreqqless;\x00&gtrless;\x00&gtrsim;\x00&gvertneqq;\x00&gvnE;\x00&hArr;\x00&hairsp;\x00&half;\x00&hamilt;\x00&hardcy;\x00&harr;\x00&harrcir;\x00&harrw;\x00&hbar;\x00&hcirc;\x00&hearts;\x00&heartsuit;\x00&hellip;\x00&hercon;\x00&hfr;\x00&hksearow;\x00&hkswarow;\x00&hoarr;\x00&homtht;\x00&hookleftarrow;\x00&hookrightarrow;\x00&hopf;\x00&horbar;\x00&hscr;\x00&hslash;\x00&hstrok;\x00&hybull;\x00&hyphen;\x00&iacute;\x00&ic;\x00&icirc;\x00&icy;\x00&iecy;\x00&iexcl;\x00&iff;\x00&ifr;\x00&igrave;\x00&ii;\x00&iiiint;\x00&iiint;\x00&iinfin;\x00&iiota;\x00&ijlig;\x00&imacr;\x00&image;\x00&imagline;\x00&imagpart;\x00&imath;\x00&imof;\x00&imped;\x00&in;\x00&incare;\x00&infin;\x00&infintie;\x00&inodot;\x00&int;\x00&intcal;\x00&integers;\x00&intercal;\x00&intlarhk;\x00&intprod;\x00&iocy;\x00&iogon;\x00&iopf;\x00&iota;\x00&iprod;\x00&iquest;\x00&iscr;\x00&isin;\x00&isinE;\x00&isindot;\x00&isins;\x00&isinsv;\x00&isinv;\x00&it;\x00&itilde;\x00&iukcy;\x00&iuml;\x00&jcirc;\x00&jcy;\x00&jfr;\x00&jmath;\x00&jopf;\x00&jscr;\x00&jsercy;\x00&jukcy;\x00&kappa;\x00&kappav;\x00&kcedil;\x00&kcy;\x00&kfr;\x00&kgreen;\x00&khcy;\x00&kjcy;\x00&kopf;\x00&kscr;\x00&lAarr;\x00&lArr;\x00&lAtail;\x00&lBarr;\x00&lE;\x00&lEg;\x00&lHar;\x00&lacute;\x00&laemptyv;\x00&lagran;\x00&lambda;\x00&lang;\x00&langd;\x00&langle;\x00&lap;\x00&laquo;\x00&larr;\x00&larrb;\x00&larrbfs;\x00&larrfs;\x00&larrhk;\x00&larrlp;\x00&larrpl;\x00&larrsim;\x00&larrtl;\x00&lat;\x00&latail;\x00&late;\x00&lates;\x00&lbarr;\x00&lbbrk;\x00&lbrace;\x00&lbrack;\x00&lbrke;\x00&lbrksld;\x00&lbrkslu;\x00&lcaron;\x00&lcedil;\x00&lceil;\x00&lcub;\x00&lcy;\x00&ldca;\x00&ldquo;\x00&ldquor;\x00&ldrdhar;\x00&ldrushar;\x00&ldsh;\x00&le;\x00&leftarrow;\x00&leftarrowtail;\x00&leftharpoondown;\x00&leftharpoonup;\x00&leftleftarrows;\x00&leftrightarrow;\x00&leftrightarrows;\x00&leftrightharpoons;\x00&leftrightsquigarrow;\x00&leftthreetimes;\x00&leg;\x00&leq;\x00&leqq;\x00&leqslant;\x00&les;\x00&lescc;\x00&lesdot;\x00&lesdoto;\x00&lesdotor;\x00&lesg;\x00&lesges;\x00&lessapprox;\x00&lessdot;\x00&lesseqgtr;\x00&lesseqqgtr;\x00&lessgtr;\x00&lesssim;\x00&lfisht;\x00&lfloor;\x00&lfr;\x00&lg;\x00&lgE;\x00&lhard;\x00&lharu;\x00&lharul;\x00&lhblk;\x00&ljcy;\x00&ll;\x00&llarr;\x00&llcorner;\x00&llhard;\x00&lltri;\x00&lmidot;\x00&lmoust;\x00&lmoustache;\x00&lnE;\x00&lnap;\x00&lnapprox;\x00&lne;\x00&lneq;\x00&lneqq;\x00&lnsim;\x00&loang;\x00&loarr;\x00&lobrk;\x00&longleftarrow;\x00&longleftrightarrow;\x00&longmapsto;\x00&longrightarrow;\x00&looparrowleft;\x00&looparrowright;\x00&lopar;\x00&lopf;\x00&loplus;\x00&lotimes;\x00&lowast;\x00&lowbar;\x00&loz;\x00&lozenge;\x00&lozf;\x00&lpar;\x00&lparlt;\x00&lrarr;\x00&lrcorner;\x00&lrhar;\x00&lrhard;\x00&lrm;\x00&lrtri;\x00&lsaquo;\x00&lscr;\x00&lsh;\x00&lsim;\x00&lsime;\x00&lsimg;\x00&lsqb;\x00&lsquo;\x00&lsquor;\x00&lstrok;\x00&ltcc;\x00&ltcir;\x00&ltdot;\x00&lthree;\x00&ltimes;\x00&ltlarr;\x00&ltquest;\x00&ltrPar;\x00&ltri;\x00&ltrie;\x00&ltrif;\x00&lurdshar;\x00&luruhar;\x00&lvertneqq;\x00&lvnE;\x00&mDDot;\x00&macr;\x00&male;\x00&malt;\x00&maltese;\x00&map;\x00&mapsto;\x00&mapstodown;\x00&mapstoleft;\x00&mapstoup;\x00&marker;\x00&mcomma;\x00&mcy;\x00&mdash;\x00&measuredangle;\x00&mfr;\x00&mho;\x00&micro;\x00&mid;\x00&midast;\x00&midcir;\x00&middot;\x00&minus;\x00&minusb;\x00&minusd;\x00&minusdu;\x00&mlcp;\x00&mldr;\x00&mnplus;\x00&models;\x00&mopf;\x00&mp;\x00&mscr;\x00&mstpos;\x00&mu;\x00&multimap;\x00&mumap;\x00&nGg;\x00&nGt;\x00&nGtv;\x00&nLeftarrow;\x00&nLeftrightarrow;\x00&nLl;\x00&nLt;\x00&nLtv;\x00&nRightarrow;\x00&nVDash;\x00&nVdash;\x00&nabla;\x00&nacute;\x00&nang;\x00&nap;\x00&napE;\x00&napid;\x00&napos;\x00&napprox;\x00&natur;\x00&natural;\x00&naturals;\x00&nbsp;\x00&nbump;\x00&nbumpe;\x00&ncap;\x00&ncaron;\x00&ncedil;\x00&ncong;\x00&ncongdot;\x00&ncup;\x00&ncy;\x00&ndash;\x00&ne;\x00&neArr;\x00&nearhk;\x00&nearr;\x00&nearrow;\x00&nedot;\x00&nequiv;\x00&nesear;\x00&nesim;\x00&nexist;\x00&nexists;\x00&nfr;\x00&ngE;\x00&nge;\x00&ngeq;\x00&ngeqq;\x00&ngeqslant;\x00&nges;\x00&ngsim;\x00&ngt;\x00&ngtr;\x00&nhArr;\x00&nharr;\x00&nhpar;\x00&ni;\x00&nis;\x00&nisd;\x00&niv;\x00&njcy;\x00&nlArr;\x00&nlE;\x00&nlarr;\x00&nldr;\x00&nle;\x00&nleftarrow;\x00&nleftrightarrow;\x00&nleq;\x00&nleqq;\x00&nleqslant;\x00&nles;\x00&nless;\x00&nlsim;\x00&nlt;\x00&nltri;\x00&nltrie;\x00&nmid;\x00&nopf;\x00&not;\x00&notin;\x00&notinE;\x00&notindot;\x00&notinva;\x00&notinvb;\x00&notinvc;\x00&notni;\x00&notniva;\x00&notnivb;\x00&notnivc;\x00&npar;\x00&nparallel;\x00&nparsl;\x00&npart;\x00&npolint;\x00&npr;\x00&nprcue;\x00&npre;\x00&nprec;\x00&npreceq;\x00&nrArr;\x00&nrarr;\x00&nrarrc;\x00&nrarrw;\x00&nrightarrow;\x00&nrtri;\x00&nrtrie;\x00&nsc;\x00&nsccue;\x00&nsce;\x00&nscr;\x00&nshortmid;\x00&nshortparallel;\x00&nsim;\x00&nsime;\x00&nsimeq;\x00&nsmid;\x00&nspar;\x00&nsqsube;\x00&nsqsupe;\x00&nsub;\x00&nsubE;\x00&nsube;\x00&nsubset;\x00&nsubseteq;\x00&nsubseteqq;\x00&nsucc;\x00&nsucceq;\x00&nsup;\x00&nsupE;\x00&nsupe;\x00&nsupset;\x00&nsupseteq;\x00&nsupseteqq;\x00&ntgl;\x00&ntilde;\x00&ntlg;\x00&ntriangleleft;\x00&ntrianglelefteq;\x00&ntriangleright;\x00&ntrianglerighteq;\x00&nu;\x00&num;\x00&numero;\x00&numsp;\x00&nvDash;\x00&nvHarr;\x00&nvap;\x00&nvdash;\x00&nvge;\x00&nvgt;\x00&nvinfin;\x00&nvlArr;\x00&nvle;\x00&nvlt;\x00&nvltrie;\x00&nvrArr;\x00&nvrtrie;\x00&nvsim;\x00&nwArr;\x00&nwarhk;\x00&nwarr;\x00&nwarrow;\x00&nwnear;\x00&oS;\x00&oacute;\x00&oast;\x00&ocir;\x00&ocirc;\x00&ocy;\x00&odash;\x00&odblac;\x00&odiv;\x00&odot;\x00&odsold;\x00&oelig;\x00&ofcir;\x00&ofr;\x00&ogon;\x00&ograve;\x00&ogt;\x00&ohbar;\x00&ohm;\x00&oint;\x00&olarr;\x00&olcir;\x00&olcross;\x00&oline;\x00&olt;\x00&omacr;\x00&omega;\x00&omicron;\x00&omid;\x00&ominus;\x00&oopf;\x00&opar;\x00&operp;\x00&oplus;\x00&or;\x00&orarr;\x00&ord;\x00&order;\x00&orderof;\x00&ordf;\x00&ordm;\x00&origof;\x00&oror;\x00&orslope;\x00&orv;\x00&oscr;\x00&oslash;\x00&osol;\x00&otilde;\x00&otimes;\x00&otimesas;\x00&ouml;\x00&ovbar;\x00&par;\x00&para;\x00&parallel;\x00&parsim;\x00&parsl;\x00&part;\x00&pcy;\x00&percnt;\x00&period;\x00&permil;\x00&perp;\x00&pertenk;\x00&pfr;\x00&phi;\x00&phiv;\x00&phmmat;\x00&phone;\x00&pi;\x00&pitchfork;\x00&piv;\x00&planck;\x00&planckh;\x00&plankv;\x00&plus;\x00&plusacir;\x00&plusb;\x00&pluscir;\x00&plusdo;\x00&plusdu;\x00&pluse;\x00&plusmn;\x00&plussim;\x00&plustwo;\x00&pm;\x00&pointint;\x00&popf;\x00&pound;\x00&pr;\x00&prE;\x00&prap;\x00&prcue;\x00&pre;\x00&prec;\x00&precapprox;\x00&preccurlyeq;\x00&preceq;\x00&precnapprox;\x00&precneqq;\x00&precnsim;\x00&precsim;\x00&prime;\x00&primes;\x00&prnE;\x00&prnap;\x00&prnsim;\x00&prod;\x00&profalar;\x00&profline;\x00&profsurf;\x00&prop;\x00&propto;\x00&prsim;\x00&prurel;\x00&pscr;\x00&psi;\x00&puncsp;\x00&qfr;\x00&qint;\x00&qopf;\x00&qprime;\x00&qscr;\x00&quaternions;\x00&quatint;\x00&quest;\x00&questeq;\x00&rAarr;\x00&rArr;\x00&rAtail;\x00&rBarr;\x00&rHar;\x00&race;\x00&racute;\x00&radic;\x00&raemptyv;\x00&rang;\x00&rangd;\x00&range;\x00&rangle;\x00&raquo;\x00&rarr;\x00&rarrap;\x00&rarrb;\x00&rarrbfs;\x00&rarrc;\x00&rarrfs;\x00&rarrhk;\x00&rarrlp;\x00&rarrpl;\x00&rarrsim;\x00&rarrtl;\x00&rarrw;\x00&ratail;\x00&ratio;\x00&rationals;\x00&rbarr;\x00&rbbrk;\x00&rbrace;\x00&rbrack;\x00&rbrke;\x00&rbrksld;\x00&rbrkslu;\x00&rcaron;\x00&rcedil;\x00&rceil;\x00&rcub;\x00&rcy;\x00&rdca;\x00&rdldhar;\x00&rdquo;\x00&rdquor;\x00&rdsh;\x00&real;\x00&realine;\x00&realpart;\x00&reals;\x00&rect;\x00&reg;\x00&rfisht;\x00&rfloor;\x00&rfr;\x00&rhard;\x00&rharu;\x00&rharul;\x00&rho;\x00&rhov;\x00&rightarrow;\x00&rightarrowtail;\x00&rightharpoondown;\x00&rightharpoonup;\x00&rightleftarrows;\x00&rightleftharpoons;\x00&rightrightarrows;\x00&rightsquigarrow;\x00&rightthreetimes;\x00&ring;\x00&risingdotseq;\x00&rlarr;\x00&rlhar;\x00&rlm;\x00&rmoust;\x00&rmoustache;\x00&rnmid;\x00&roang;\x00&roarr;\x00&robrk;\x00&ropar;\x00&ropf;\x00&roplus;\x00&rotimes;\x00&rpar;\x00&rpargt;\x00&rppolint;\x00&rrarr;\x00&rsaquo;\x00&rscr;\x00&rsh;\x00&rsqb;\x00&rsquo;\x00&rsquor;\x00&rthree;\x00&rtimes;\x00&rtri;\x00&rtrie;\x00&rtrif;\x00&rtriltri;\x00&ruluhar;\x00&rx;\x00&sacute;\x00&sbquo;\x00&sc;\x00&scE;\x00&scap;\x00&scaron;\x00&sccue;\x00&sce;\x00&scedil;\x00&scirc;\x00&scnE;\x00&scnap;\x00&scnsim;\x00&scpolint;\x00&scsim;\x00&scy;\x00&sdot;\x00&sdotb;\x00&sdote;\x00&seArr;\x00&searhk;\x00&searr;\x00&searrow;\x00&sect;\x00&semi;\x00&seswar;\x00&setminus;\x00&setmn;\x00&sext;\x00&sfr;\x00&sfrown;\x00&sharp;\x00&shchcy;\x00&shcy;\x00&shortmid;\x00&shortparallel;\x00&shy;\x00&sigma;\x00&sigmaf;\x00&sigmav;\x00&sim;\x00&simdot;\x00&sime;\x00&simeq;\x00&simg;\x00&simgE;\x00&siml;\x00&simlE;\x00&simne;\x00&simplus;\x00&simrarr;\x00&slarr;\x00&smallsetminus;\x00&smashp;\x00&smeparsl;\x00&smid;\x00&smile;\x00&smt;\x00&smte;\x00&smtes;\x00&softcy;\x00&sol;\x00&solb;\x00&solbar;\x00&sopf;\x00&spades;\x00&spadesuit;\x00&spar;\x00&sqcap;\x00&sqcaps;\x00&sqcup;\x00&sqcups;\x00&sqsub;\x00&sqsube;\x00&sqsubset;\x00&sqsubseteq;\x00&sqsup;\x00&sqsupe;\x00&sqsupset;\x00&sqsupseteq;\x00&squ;\x00&square;\x00&squarf;\x00&squf;\x00&srarr;\x00&sscr;\x00&ssetmn;\x00&ssmile;\x00&sstarf;\x00&star;\x00&starf;\x00&straightepsilon;\x00&straightphi;\x00&strns;\x00&sub;\x00&subE;\x00&subdot;\x00&sube;\x00&subedot;\x00&submult;\x00&subnE;\x00&subne;\x00&subplus;\x00&subrarr;\x00&subset;\x00&subseteq;\x00&subseteqq;\x00&subsetneq;\x00&subsetneqq;\x00&subsim;\x00&subsub;\x00&subsup;\x00&succ;\x00&succapprox;\x00&succcurlyeq;\x00&succeq;\x00&succnapprox;\x00&succneqq;\x00&succnsim;\x00&succsim;\x00&sum;\x00&sung;\x00&sup1;\x00&sup2;\x00&sup3;\x00&sup;\x00&supE;\x00&supdot;\x00&supdsub;\x00&supe;\x00&supedot;\x00&suphsol;\x00&suphsub;\x00&suplarr;\x00&supmult;\x00&supnE;\x00&supne;\x00&supplus;\x00&supset;\x00&supseteq;\x00&supseteqq;\x00&supsetneq;\x00&supsetneqq;\x00&supsim;\x00&supsub;\x00&supsup;\x00&swArr;\x00&swarhk;\x00&swarr;\x00&swarrow;\x00&swnwar;\x00&szlig;\x00&target;\x00&tau;\x00&tbrk;\x00&tcaron;\x00&tcedil;\x00&tcy;\x00&tdot;\x00&telrec;\x00&tfr;\x00&there4;\x00&therefore;\x00&theta;\x00&thetasym;\x00&thetav;\x00&thickapprox;\x00&thicksim;\x00&thinsp;\x00&thkap;\x00&thksim;\x00&thorn;\x00&tilde;\x00&times;\x00&timesb;\x00&timesbar;\x00&timesd;\x00&tint;\x00&toea;\x00&top;\x00&topbot;\x00&topcir;\x00&topf;\x00&topfork;\x00&tosa;\x00&tprime;\x00&trade;\x00&triangle;\x00&triangledown;\x00&triangleleft;\x00&trianglelefteq;\x00&triangleq;\x00&triangleright;\x00&trianglerighteq;\x00&tridot;\x00&trie;\x00&triminus;\x00&triplus;\x00&trisb;\x00&tritime;\x00&trpezium;\x00&tscr;\x00&tscy;\x00&tshcy;\x00&tstrok;\x00&twixt;\x00&twoheadleftarrow;\x00&twoheadrightarrow;\x00&uArr;\x00&uHar;\x00&uacute;\x00&uarr;\x00&ubrcy;\x00&ubreve;\x00&ucirc;\x00&ucy;\x00&udarr;\x00&udblac;\x00&udhar;\x00&ufisht;\x00&ufr;\x00&ugrave;\x00&uharl;\x00&uharr;\x00&uhblk;\x00&ulcorn;\x00&ulcorner;\x00&ulcrop;\x00&ultri;\x00&umacr;\x00&uml;\x00&uogon;\x00&uopf;\x00&uparrow;\x00&updownarrow;\x00&upharpoonleft;\x00&upharpoonright;\x00&uplus;\x00&upsi;\x00&upsih;\x00&upsilon;\x00&upuparrows;\x00&urcorn;\x00&urcorner;\x00&urcrop;\x00&uring;\x00&urtri;\x00&uscr;\x00&utdot;\x00&utilde;\x00&utri;\x00&utrif;\x00&uuarr;\x00&uuml;\x00&uwangle;\x00&vArr;\x00&vBar;\x00&vBarv;\x00&vDash;\x00&vangrt;\x00&varepsilon;\x00&varkappa;\x00&varnothing;\x00&varphi;\x00&varpi;\x00&varpropto;\x00&varr;\x00&varrho;\x00&varsigma;\x00&varsubsetneq;\x00&varsubsetneqq;\x00&varsupsetneq;\x00&varsupsetneqq;\x00&vartheta;\x00&vartriangleleft;\x00&vartriangleright;\x00&vcy;\x00&vdash;\x00&vee;\x00&veebar;\x00&veeeq;\x00&vellip;\x00&verbar;\x00&vert;\x00&vfr;\x00&vltri;\x00&vnsub;\x00&vnsup;\x00&vopf;\x00&vprop;\x00&vrtri;\x00&vscr;\x00&vsubnE;\x00&vsubne;\x00&vsupnE;\x00&vsupne;\x00&vzigzag;\x00&wcirc;\x00&wedbar;\x00&wedge;\x00&wedgeq;\x00&weierp;\x00&wfr;\x00&wopf;\x00&wp;\x00&wr;\x00&wreath;\x00&wscr;\x00&xcap;\x00&xcirc;\x00&xcup;\x00&xdtri;\x00&xfr;\x00&xhArr;\x00&xharr;\x00&xi;\x00&xlArr;\x00&xlarr;\x00&xmap;\x00&xnis;\x00&xodot;\x00&xopf;\x00&xoplus;\x00&xotime;\x00&xrArr;\x00&xrarr;\x00&xscr;\x00&xsqcup;\x00&xuplus;\x00&xutri;\x00&xvee;\x00&xwedge;\x00&yacute;\x00&yacy;\x00&ycirc;\x00&ycy;\x00&yen;\x00&yfr;\x00&yicy;\x00&yopf;\x00&yscr;\x00&yucy;\x00&yuml;\x00&zacute;\x00&zcaron;\x00&zcy;\x00&zdot;\x00&zeetrf;\x00&zeta;\x00&zfr;\x00&zhcy;\x00&zigrarr;\x00&zopf;\x00&zscr;\x00&zwj;\x00&zwnj;\x00"
