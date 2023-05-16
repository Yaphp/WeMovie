<template>
    <div class="progress-notify" v-bind="$attrs">
        <div class="title">
            <div class="left">
                <h5>
                    正在上传：{{ getShortFileName(title) }}
                    <span v-if="speed">
                       传输速度： {{ speed }} /s
                    </span>
                </h5>
            </div>
            <div class="close" @click="closeProgressNotify">
                <el-icon>
                    <Close/>
                </el-icon>
            </div>
        </div>
        <div class="content">
            <el-progress
                :percentage="percentage"
                :stroke-width="10"
            />
        </div>
    </div>
</template>

<script>
import {Close} from "@element-plus/icons-vue";
import {getShortFileName} from "utils/file";

export default {
    name: "ProgressNotify",
    // inheritAttrs: false,
    // 定义percentage的类型
    props: {
        percentage: Number,
        title: String,
        speed: String
    },
    components: {
        Close: Close
    },
    setup() {
        return {
            getShortFileName
        }
    },
    data() {
        return {}
    },
    watch: {
        percentage(val) {
            if (val === 100) {
                this.$emit("completed")
            }
        }
    },
    methods: {
        closeProgressNotify() {
            this.$emit("close")
        }
    },
}
</script>

<style scoped>
.progress-notify {
    position: absolute;
    top: 85vh;
    right: 2%;
    width: 320px;
    height: 100px;
    background-color: #ffffff;
    color: black;
    border: 1px solid #ebeef5;
    z-index: 2500;
    border-radius: 15px;
    box-shadow: 0px 0px 12px rgba(0, 0, 0, .12);
    display: flex;
    justify-content: center;
    align-items: flex-start;
    flex-direction: column;
}

.progress-notify .el-progress--line {
    margin-bottom: 15px;
    width: 300px;
}

.title {
    width: 96%;
    height: 50px;
    display: flex;
    justify-content: space-around;
    align-items: center;
}

.title .left {
    width: 80%;
    text-align: left;
    margin-left: 10px;
    margin-top: 15px;
}

.title .close {
    cursor: pointer;
    width: 20%;
    margin-left: 10px;
    margin-top: 15px;
    display: flex;
    justify-content: flex-end;
    align-items: center;
}

.content {
    /*margin:10px;*/
    margin: auto;
}
</style>