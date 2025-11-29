import {
    defineConfig,
    presetWind4,
} from 'unocss'

export default defineConfig({
    presets: [
        presetWind4(),
    ],
    shortcuts: [
        {'btn-primary-big': 'button-primary bg-primary text-text-onAccent py-4 rounded px-16'},
        {'btn-primary': 'button-primary bg-primary text-text-onAccent py-2 rounded px-6'},
    ],
    theme: {
        font: {
            sans: 'Nunito Sans',
            header: 'Seymour One',
        },
        colors: {
            primary: '#1A73B3',
            primaryHover: '#3C58B4',
            background: '#F5F7FA',
            surface: '#FFFFFF',
            border: '#E0E4EC',
            text: {
                primary: '#111827',
                onAccent: '#FFFFFF',
                secondary: '#6B7280',
                disabled: '#9CA3AF',
            },
            error: '#9F5D5D',
            warning: '#8E6917',
            success: '#3E7B55',
        },
        radius: {
            none: '0px',
            sm: '4px',
            DEFAULT: '15px',
            md: '15px',
            lg: '15px',
            xl: '15px',
            '2xl': '15px',
            full: '9999px',
        }
    },

})