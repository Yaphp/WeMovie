<template>
    <div class="bg">
        <img class="bg-photo" :src="bg" alt="">
        <div class="bg-slogan">
            <h1 class="bg-title">{{ this.$store.state.app.name }}</h1>
        </div>
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
    width: 100%;
    height: 100%;
}

.bg-slogan {
    position: absolute;
    z-index: 999;
    top: 45%;
    left: 50%;
}

.bg-title{
    color: #579cf5;
    font-size: 60px;
}
</style>