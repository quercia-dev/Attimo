
import type { CustomThemeConfig } from '@skeletonlabs/tw-plugin';

export const EarthTheme: CustomThemeConfig = {
    name: 'earth',
    properties: {
		// =~= Theme Properties =~=
		"--theme-font-family-base": `system-ui`,
		"--theme-font-family-heading": `system-ui`,
		"--theme-font-color-base": "0 0 0",
		"--theme-font-color-dark": "255 255 255",
		"--theme-rounded-base": "9999px",
		"--theme-rounded-container": "8px",
		"--theme-border-base": "1px",
		// =~= Theme On-X Colors =~=
		"--on-primary": "var(--color-secondary-800)",
		"--on-secondary": "var(--color-primary-300)",
		"--on-tertiary": "var(--color-secondary-800)",
		"--on-success": "var(--color-secondary-900)",
		"--on-warning": "var(--color-secondary-700)",
		"--on-error": "var(--color-secondary-50)",
		"--on-surface": "0 0 0",
		// =~= Theme Colors  =~=
		// primary | #83b8EC 
		"--color-primary-50": "236 244 252", // #ecf4fc
		"--color-primary-100": "230 241 251", // #e6f1fb
		"--color-primary-200": "224 237 250", // #e0edfa
		"--color-primary-300": "205 227 247", // #cde3f7
		"--color-primary-400": "168 205 242", // #a8cdf2
		"--color-primary-500": "131 184 236", // #83b8EC
		"--color-primary-600": "118 166 212", // #76a6d4
		"--color-primary-700": "98 138 177", // #628ab1
		"--color-primary-800": "79 110 142", // #4f6e8e
		"--color-primary-900": "64 90 116", // #405a74
		// secondary | #9d5521 
		"--color-secondary-50": "240 230 222", // #f0e6de
		"--color-secondary-100": "235 221 211", // #ebddd3
		"--color-secondary-200": "231 213 200", // #e7d5c8
		"--color-secondary-300": "216 187 166", // #d8bba6
		"--color-secondary-400": "186 136 100", // #ba8864
		"--color-secondary-500": "157 85 33", // #9d5521
		"--color-secondary-600": "141 77 30", // #8d4d1e
		"--color-secondary-700": "118 64 25", // #764019
		"--color-secondary-800": "94 51 20", // #5e3314
		"--color-secondary-900": "77 42 16", // #4d2a10
		// tertiary | #7eb851 
		"--color-tertiary-50": "236 244 229", // #ecf4e5
		"--color-tertiary-100": "229 241 220", // #e5f1dc
		"--color-tertiary-200": "223 237 212", // #dfedd4
		"--color-tertiary-300": "203 227 185", // #cbe3b9
		"--color-tertiary-400": "165 205 133", // #a5cd85
		"--color-tertiary-500": "126 184 81", // #7eb851
		"--color-tertiary-600": "113 166 73", // #71a649
		"--color-tertiary-700": "95 138 61", // #5f8a3d
		"--color-tertiary-800": "76 110 49", // #4c6e31
		"--color-tertiary-900": "62 90 40", // #3e5a28
		// success | #88d100 
		"--color-success-50": "237 248 217", // #edf8d9
		"--color-success-100": "231 246 204", // #e7f6cc
		"--color-success-200": "225 244 191", // #e1f4bf
		"--color-success-300": "207 237 153", // #cfed99
		"--color-success-400": "172 223 77", // #acdf4d
		"--color-success-500": "136 209 0", // #88d100
		"--color-success-600": "122 188 0", // #7abc00
		"--color-success-700": "102 157 0", // #669d00
		"--color-success-800": "82 125 0", // #527d00
		"--color-success-900": "67 102 0", // #436600
		// warning | #e9da35 
		"--color-warning-50": "252 249 225", // #fcf9e1
		"--color-warning-100": "251 248 215", // #fbf8d7
		"--color-warning-200": "250 246 205", // #faf6cd
		"--color-warning-300": "246 240 174", // #f6f0ae
		"--color-warning-400": "240 229 114", // #f0e572
		"--color-warning-500": "233 218 53", // #e9da35
		"--color-warning-600": "210 196 48", // #d2c430
		"--color-warning-700": "175 164 40", // #afa428
		"--color-warning-800": "140 131 32", // #8c8320
		"--color-warning-900": "114 107 26", // #726b1a
		// error | #ca4444 
		"--color-error-50": "247 227 227", // #f7e3e3
		"--color-error-100": "244 218 218", // #f4dada
		"--color-error-200": "242 208 208", // #f2d0d0
		"--color-error-300": "234 180 180", // #eab4b4
		"--color-error-400": "218 124 124", // #da7c7c
		"--color-error-500": "202 68 68", // #ca4444
		"--color-error-600": "182 61 61", // #b63d3d
		"--color-error-700": "152 51 51", // #983333
		"--color-error-800": "121 41 41", // #792929
		"--color-error-900": "99 33 33", // #632121
		// surface | #b8d6f4 
		"--color-surface-50": "244 249 253", // #f4f9fd
		"--color-surface-100": "241 247 253", // #f1f7fd
		"--color-surface-200": "237 245 252", // #edf5fc
		"--color-surface-300": "227 239 251", // #e3effb
		"--color-surface-400": "205 226 247", // #cde2f7
		"--color-surface-500": "184 214 244", // #b8d6f4
		"--color-surface-600": "166 193 220", // #a6c1dc
		"--color-surface-700": "138 161 183", // #8aa1b7
		"--color-surface-800": "110 128 146", // #6e8092
		"--color-surface-900": "90 105 120", // #5a6978
		
	}
}