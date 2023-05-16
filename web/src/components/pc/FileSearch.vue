<template>
    <el-dialog title="文件搜索" :modelValue="show" width="40%" draggable>
        <div class="mt-4">
            <el-input
                v-model="keyword"
                placeholder="输入文字以搜索"
                class="input-with-select"
                @keyup.enter="handleSearch"
            >
                <template #prepend>
                    <el-select v-model="type" clearable placeholder="文件类型" style="width: 115px">
                        <el-option label="文件夹" value="folder"/>
                        <el-option label="图片" value="image"/>
                        <el-option label="视频" value="video"/>
                        <el-option label="音频" value="audio"/>
                    </el-select>
                </template>
                <template #append>
                    <el-button :icon="Search" @click="handleSearch" />
                </template>
            </el-input>
        </div>
    </el-dialog>
</template>

<script>
import {Search} from "@element-plus/icons-vue";

export default {
    name: "FileSearch",
    props: {
        show: Boolean
    },
    components: {},
    setup() {
        return {
            Search: Search
        }
    },
    data() {
        return {
            keyword: "",
            type: ""
        }
    },
    methods: {
        handleSearch() {
            this.$emit("search", {"keyword": this.keyword, "type": this.type})
        }
    },
}
</script>

<style scoped>

</style>