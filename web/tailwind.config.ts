import type { Config} from 'tailwindcss'

export default <Partial<Config>>{
    theme: {
        extend: {
            colors: {
                'symbol-blue':  '#001871',
                'silver-gray':  '#9ea2a2',
                'scarlet-red':  '#ba0c2f',
                'air-force-yellow':  '#ffcd00',
            }
        },
        fontFamily: {
            'heading': ['Rajdhani-Bold'],
            'body': ['Ubuntu-Regular', 'NoticiaText-Regular'],
            'body-italic': ['Ubuntu-Italic', 'NoticiaText-Italic'],
            'body-bold': ['Ubuntu-Bold', 'NoticiaText-Bold'],
            'body-bold-italic': ['Ubuntu-BoldItalic', 'NoticiaText-BoldItalic'],
        },
    }
}