<template>
    <div>
        <ul>
            <li v-for="item in items">{{ item }}</li>
        </ul>
    </div>
</template>
<script>
    export default {
        props: ['items']
    }
</script>