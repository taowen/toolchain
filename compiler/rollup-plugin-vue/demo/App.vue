<template>
    <div>
        <hello :items="['a','b','c']"></hello>
    </div>
</template>
<script>
    import hello from '@compiler/rollup-plugin-vue'
    export default {
        components: {hello}
    }
</script>