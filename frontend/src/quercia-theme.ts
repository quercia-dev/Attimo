
import type { CustomThemeConfig } from '@skeletonlabs/tw-plugin';

export const QuercusTheme: CustomThemeConfig = {
    name: 'quercus',
    properties: {
		// =~= Theme Properties =~=
		"--theme-font-family-base": `ui-monospace, SFMono-Regular, Menlo, Monaco, Consolas, 'Liberation Mono', 'Courier New', monospace`,
		"--theme-font-family-heading": `system-ui`,
		"--theme-font-color-base": "0 0 0",
		"--theme-font-color-dark": "var(--color-surface-50)",
		"--theme-rounded-base": "12px",
		"--theme-rounded-container": "6px",
		"--theme-border-base": "2px",
		// =~= Theme On-X Colors =~=
		"--on-primary": "255 255 255",
		"--on-secondary": "255 255 255",
		"--on-tertiary": "0 0 0",
		"--on-success": "0 0 0",
		"--on-warning": "0 0 0",
		"--on-error": "var(--color-surface-100)",
		"--on-surface": "0 0 0",
		// =~= Theme Colors  =~=
		// primary | #1ea7e1 
		"--color-primary-50": "221 242 251", // #ddf2fb
		"--color-primary-100": "210 237 249", // #d2edf9
		"--color-primary-200": "199 233 248", // #c7e9f8
		"--color-primary-300": "165 220 243", // #a5dcf3
		"--color-primary-400": "98 193 234", // #62c1ea
		"--color-primary-500": "30 167 225", // #1ea7e1
		"--color-primary-600": "27 150 203", // #1b96cb
		"--color-primary-700": "23 125 169", // #177da9
		"--color-primary-800": "18 100 135", // #126487
		"--color-primary-900": "15 82 110", // #0f526e
		// secondary | #045f00
		"--color-secondary-50": "217 231 217", // #d9e7d9
		"--color-secondary-100": "205 223 204", // #cddfcc
		"--color-secondary-200": "192 215 191", // #c0d7bf
		"--color-secondary-300": "155 191 153", // #9bbf99
		"--color-secondary-400": "79 143 77", // #4f8f4d
		"--color-secondary-500": "4 95 0", // #045f00
		"--color-secondary-600": "4 86 0", // #045600
		"--color-secondary-700": "3 71 0", // #034700
		"--color-secondary-800": "2 57 0", // #023900
		"--color-secondary-900": "2 47 0", // #022f00
		// tertiary | #a9a9a9 
		"--color-tertiary-50": "242 242 242", // #f2f2f2
		"--color-tertiary-100": "238 238 238", // #eeeeee
		"--color-tertiary-200": "234 234 234", // #eaeaea
		"--color-tertiary-300": "221 221 221", // #dddddd
		"--color-tertiary-400": "195 195 195", // #c3c3c3
		"--color-tertiary-500": "169 169 169", // #a9a9a9
		"--color-tertiary-600": "152 152 152", // #989898
		"--color-tertiary-700": "127 127 127", // #7f7f7f
		"--color-tertiary-800": "101 101 101", // #656565
		"--color-tertiary-900": "83 83 83", // #535353
		// success | #87e403 
		"--color-success-50": "237 251 217", // #edfbd9
		"--color-success-100": "231 250 205", // #e7facd
		"--color-success-200": "225 248 192", // #e1f8c0
		"--color-success-300": "207 244 154", // #cff49a
		"--color-success-400": "171 236 79", // #abec4f
		"--color-success-500": "135 228 3", // #87e403
		"--color-success-600": "122 205 3", // #7acd03
		"--color-success-700": "101 171 2", // #65ab02
		"--color-success-800": "81 137 2", // #518902
		"--color-success-900": "66 112 1", // #427001
		// warning | #dec000 
		"--color-warning-50": "250 246 217", // #faf6d9
		"--color-warning-100": "248 242 204", // #f8f2cc
		"--color-warning-200": "247 239 191", // #f7efbf
		"--color-warning-300": "242 230 153", // #f2e699
		"--color-warning-400": "232 211 77", // #e8d34d
		"--color-warning-500": "222 192 0", // #dec000
		"--color-warning-600": "200 173 0", // #c8ad00
		"--color-warning-700": "167 144 0", // #a79000
		"--color-warning-800": "133 115 0", // #857300
		"--color-warning-900": "109 94 0", // #6d5e00
		// error | #ab1f24 
		"--color-error-50": "242 221 222", // #f2ddde
		"--color-error-100": "238 210 211", // #eed2d3
		"--color-error-200": "234 199 200", // #eac7c8
		"--color-error-300": "221 165 167", // #dda5a7
		"--color-error-400": "196 98 102", // #c46266
		"--color-error-500": "171 31 36", // #ab1f24
		"--color-error-600": "154 28 32", // #9a1c20
		"--color-error-700": "128 23 27", // #80171b
		"--color-error-800": "103 19 22", // #671316
		"--color-error-900": "84 15 18", // #540f12
		// surface | #128a5e 
		"--color-surface-50": "219 237 231", // #dbede7
		"--color-surface-100": "208 232 223", // #d0e8df
		"--color-surface-200": "196 226 215", // #c4e2d7
		"--color-surface-300": "160 208 191", // #a0d0bf
		"--color-surface-400": "89 173 142", // #59ad8e
		"--color-surface-500": "18 138 94", // #128a5e
		"--color-surface-600": "16 124 85", // #107c55
		"--color-surface-700": "14 104 71", // #0e6847
		"--color-surface-800": "11 83 56", // #0b5338
		"--color-surface-900": "9 68 46", // #09442e
		
	}
}