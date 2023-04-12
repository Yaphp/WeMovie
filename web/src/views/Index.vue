<template>
    <div class="bg">
        <img class="bg-photo" :src="bg" alt="">
    </div>
</template>

<script>

import bg from '../assets/images/bg.png';

export default {
    data() {
        return {
            bg: "",
        };
    },
    mounted() {
        this.bg = bg;
    }
};
</script>

<style>
.el-main {
    padding: unset !important;
}

.bg {
    width: 100%;
    height: 100%;
    overflow: hidden;
}

.bg-photo {
    width: 70%;
    height: 70%;
}
</style>