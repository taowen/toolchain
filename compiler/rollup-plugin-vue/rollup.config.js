import vue from 'rollup-plugin-vue'
import commonjs from 'rollup-plugin-commonjs';

export default {
    input: 'src/hello.vue',
    output: {
        format: 'esm',
        file: 'dist/hello.js'
    },
    plugins: [
        commonjs(),
        vue()
    ]
}