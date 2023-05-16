<template>
    <van-row>
        <van-col span="2" class="back-button" @click="this.$router.back()">
            <van-icon name="arrow-left" size="16px" color="#999"/>
        </van-col>
        <van-col span="22">
            <van-search v-model="keyword" placeholder="输入关键字以搜索"
                        show-action
                        @keyup.enter="handleSearch()"
                        @search="handleSearch">
                <template #action>
                    <div @click="handleSearch">搜索</div>
                </template>
            </van-search>
        </van-col>
    </van-row>

    <van-row>
        <van-col span="24">
            <van-cell-group>
                <van-cell v-for="(h,index) in history" @click="handleJumpPage(h)" :title="h">
                    <template #value>
                        <span @click.stop="handleClearHistory(index)">清除</span>
                    </template>
                </van-cell>
            </van-cell-group>
        </van-col>
    </van-row>
</template>

<script>
export default {
    name: "Search",
    components: {},
    data() {
        return {
            keyword: "",
            history: []
        }
    },
    mounted() {
        this.handleGetSearchHistory()
    },
    methods: {
        handleSearch() {
            // 保存搜索历史
            this.handleSetSearchHistory()

            // 跳转到搜索结果页
            this.handleJumpPage(this.keyword)
        },
        handleJumpPage(keyword) {
            // 使用route.name跳转
            this.$router.push({
                name: "Result",
                query: {
                    keyword: keyword
                }
            })
        },
        handleSetSearchHistory() {
            let history = localStorage.getItem("searchHistory")
            if (history) {
                let historyArr = JSON.parse(history)
                historyArr.unshift(this.keyword) // 插入到数组头部
                historyArr = [...new Set(historyArr)]; // 去重
                historyArr = historyArr.slice(0, 10); // 只保留前10个
                localStorage.setItem("searchHistory", JSON.stringify(historyArr));// 保存到localStoragek
            } else {
                localStorage.setItem("searchHistory", JSON.stringify([this.keyword]))
            }
        },
        handleGetSearchHistory() {
            // 从localStorage中获取搜索历史
            let history = localStorage.getItem("searchHistory")
            if (history) {
                this.history = JSON.parse(history)
            }
        },
        handleClearHistory(index) {
            this.history.splice(index, 1)
            localStorage.setItem("searchHistory", JSON.stringify(this.history))
        }
    }
}
</script>

<style scoped>
.back-button {
    display: flex;
    justify-content: center;
    align-items: center;
}
</style>