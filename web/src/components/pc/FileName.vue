<template>
    <el-dialog
        v-model="showDiag"
        :title="title"
        width="450px"
    >
        <div class="icon-box">
            <slot name="icon">
                <img src="/images/folder.png" class="file-icon">
            </slot>
        </div>

        <el-form label-width="70px">
            <el-input v-model="fileObject.name" clearable input-style="width:100%;"></el-input>
        </el-form>
        <template #footer>
            <span class="dialog-footer">
                <el-button @click="handleClose">取消</el-button>
                <el-button type="primary" @click="handleConfirm">确定</el-button>
            </span>
        </template>
    </el-dialog>
</template>

<script>

export default {
    name: "FileName",
    props: {
        show: Boolean,
        title: String,
        file: {
            type: Object,
            default: () => {
                return {
                    name: '新建文件夹'
                }
            }
        }
    },
    components: {},
    setup() {
        return {}
    },
    data() {
        return {
            showDiag: this.show,
            fileObject: {...this.file}
        }
    },
    methods: {
        handleConfirm() {
            this.$emit("confirm", this.fileObject)
        },
        handleClose() {
            this.$emit("close")
        },
    }
}
</script>

<style scoped>
.icon-box{
    width:100%;
    text-align:center;
}

.file-icon {
    width: 100px;
    height: 100px;
    margin: 20px auto;
    display: block;
}
</style>