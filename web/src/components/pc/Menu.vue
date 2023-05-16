<template>
    <div class="navMenu">
        <template v-for="navMenu in navMenus">
            <!-- 最后一级菜单 -->
            <el-menu-item v-if="!navMenu.childs&&navMenu.entity" :key="navMenu.entity.id" :data="navMenu" :index="navMenu.entity.name" @click="redirect(navMenu.entity.name,navMenu.entity.alias)">
                <img class="menu-item-icon" :src="'/images/'+navMenu.entity.icon+'.png'">
                <span slot="title">{{ navMenu.entity.alias }}</span>
            </el-menu-item>
        </template>
    </div>
</template>

<script>
export default {
    name: 'NavMenu',
    props: ['navMenus'],
    data() {
        return {}
    },
    methods: {
        redirect(link, name) {
            this.$router.push(link)
            this.$store.state.tagName = name;
        }
    }
}
</script>

<style>
.navMenu {
    width: 90%;
    margin: auto;
}

.el-menu-item {
    border-radius: 15px;
    margin: 5px 0;
    height: 50px;
    display: flex;
    justify-content: flex-start;
    align-items: center;
}

.menu-item-icon {
    width: 22px;
    height: 22px;
    margin-right: 20px;
}

.el-menu-item.is-active {
    color: black !important;
        background-color: rgba(132, 133, 141, 0.2);
    border-radius: 15px;
}

.el-menu-item:hover {
    background-color: rgba(132, 133, 141, 0.2);
    border-radius: 15px;
}
</style>
