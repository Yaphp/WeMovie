<template>
    <router-view/>

    <van-tabbar route safe-area-inset-bottom @change="onChangeTab">
        <van-tabbar-item replace to="/files" icon="orders-o">文件</van-tabbar-item>
        <van-tabbar-item replace to="/album" icon="photo-o">相册</van-tabbar-item>
        <van-tabbar-item replace to="/user" icon="user-o">我的</van-tabbar-item>
    </van-tabbar>
</template>

<script>

export default {
    name: "Home",
    components: {},
    data() {
        return {
            activeIndex: 'files',
            activeText: '文件',
            activeTitle: '',
            pathMapping: {
                'files': '文件',
                'album': '相册',
                'user': '我的'
            },

        };
    },
    mounted() {
        // console.log(window.location.href)
        let start = window.location.href.lastIndexOf('/');
        let path = window.location.href.slice(start + 1);
        this.activeIndex = path;
        this.activeText = this.pathMapping[path];
        this.$router.push(path)
        document.title = this.$store.state.app.name;
    },
    methods: {
        onChangeTab(index) {
            if (index === 2) {
                this.activeTitle = '我的';
                this.activeText = '';
                this.activeIndex = 'user';
            } else {
                this.activeTitle = '';
                this.activeIndex = Object.keys(this.pathMapping)[index];
                this.activeText = Object.values(this.pathMapping)[index];
            }

        },
        handleSearchJump() {
            this.$router.push('/search')
        }
    }
};
</script>

<style>

</style>