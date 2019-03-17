<template>
    <div>
        <ul>
            <li v-for="item in items">{{ item }}</li>
        </ul>
    </div>
</template>
<script>
    module.exports = {
        props: ['items']
    }
</script>