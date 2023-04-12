<template>
    <div class="navMenu">
        <template v-for="navMenu in navMenus">
            <!-- 最后一级菜单 -->
            <el-menu-item v-if="!navMenu.childs&&navMenu.entity" :key="navMenu.entity.id" :data="navMenu" :index="navMenu.entity.name" @click="redirect(navMenu.entity.name,navMenu.entity.alias)">
                <i :class="navMenu.entity.icon"></i>
                <span slot="title">{{ navMenu.entity.alias }}</span>
            </el-menu-item>

<!--            &lt;!&ndash; 此菜单下还有子菜单 &ndash;&gt;-->
<!--            <el-submenu v-if="navMenu.childs&&navMenu.entity"-->
<!--                        :key="navMenu.entity.id" :data="navMenu" :index="navMenu.entity.name">-->
<!--                <template slot="title">-->
<!--                    <i :class="navMenu.entity.icon"></i>-->
<!--                    <span> {{ navMenu.entity.alias }}</span>-->
<!--                </template>-->
<!--                &lt;!&ndash; 递归 &ndash;&gt;-->
<!--                <NavMenu :navMenus="navMenu.childs"></NavMenu>-->
<!--            </el-submenu>-->
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
.el-menu-item.is-active {
    color: #409EFF !important;
}
</style>
