<template>
    <el-dialog
        v-model="showDiag"
        title="文件分享"
        width="90%"
        top="15vh"
    >
        <slot name="icon">
            <img src="/images/shareicon.png" class="share-file-icon">
        </slot>
        <div class="share-title">
            <h3>共 {{ files.length }} 个文件</h3>
        </div>
        <el-form :model="params" label-width="120px">
            <el-form-item label="选择有效期">
                <el-select v-model="params.expire">
                    <el-option label="30天内有效" value="1"/>
                    <el-option label="永久有效" value="2"/>
                </el-select>
            </el-form-item>
            <el-form-item label="分享形式">
                <el-select v-model="params.type">
                    <el-option label="公开链接" value="1"/>
                    <el-option label="私密链接" value="2"/>
                </el-select>
            </el-form-item>
        </el-form>
        <div class="link-show" v-if="link">
            <div class="link-text">
                <div class="up">
                    分享链接：{{ link }}
                </div>
                <div class="down">
                    <el-icon class="copy-icon" size="22" @click="handleCopyLink">
                        <CopyDocument/>
                    </el-icon>
                </div>
            </div>
        </div>
        <template v-if="!link" #footer>
            <span class="dialog-footer">
                <el-button @click="handleClose">取消</el-button>
                <el-button type="primary" @click="handleConfirm">确定</el-button>
            </span>
        </template>

        <template v-if="link" #footer>
            <span class="dialog-footer">
                <el-button type="primary" plain @click="handleCopyLink">复制链接口令</el-button>
            </span>
        </template>
    </el-dialog>
</template>

<script>
import {shareAdd} from "api/share";
import {CopyDocument} from "@element-plus/icons-vue";

export default {
    name: "FileShare",
    props: {
        show: Boolean,
        file: [Object, Array]
    },
    components: {
        CopyDocument
    },
    setup() {
        return {}
    },
    data() {
        return {
            showDiag: this.show,
            files: this.file instanceof Array ? this.file : [this.file],
            params: {
                expire: "1",
                type: "1"
            },
            link: ""
        }
    },
    methods: {
        handleConfirm() {
            this.$emit("confirm", this.files, this.params);

            // ids
            let ids = this.files.map(file => file.id);
            shareAdd({
                "id": ids.join(","),
                "expire": this.params.expire,
                "type": this.params.type,
                "uid": this.$store.state.user.id
            }).then(res => {
                if (res.code === 0) {
                    this.link = res.data;
                    this.$message.success("分享成功");
                } else {
                    this.$message.error(res.msg);
                }

            }).catch(error => {
                this.$message.error(error);
            });
        },
        handleCopyLink() {
            this.$copyText(this.link).then(() => {
                this.$message.success("复制成功");
            }).catch(() => {
                this.$message.error("复制失败");
            })
        },
        handleClose() {
            this.$emit("close")
        },
    }
}
</script>

<style scoped>
.share-file-icon {
    width: 100px;
    height: 80px;
    margin: 20px auto;
    display: block;
}

.link-show {
    width: 100%;
    height: 90px;
    background-color: rgba(225, 230, 232, 0.42);
}

.link-text {
    width: 100%;
    height: auto;
    text-align: left;
    padding: 15px;
}

.down {
    width: 100%;
    text-align: right;
}

.share-title{
    margin-bottom: 10px;
    text-align: center;
}
</style>