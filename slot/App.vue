<template>
    <app-header></app-header>
</template>

<script>
    import Header from './components/header.vue';

    export default {
        name: "app",
        components:{
            'app-header': Header,
        }
    }
</script>

<style scoped>

</style>