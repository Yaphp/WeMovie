<template>
    <el-container>
        <el-main>
            <el-row>
                <!--内容主体区域-->
                <el-form :model="user" label-width="70px">
                    <el-form-item label="头像" prop="username">
                        <el-upload
                            class="avatar-uploader"
                            :show-file-list="false"
                            :http-request="handleUploadAvatar"
                            :before-upload="handleAvatarBeforeRead"
                        >
                            <img v-if="user.avatar" :src="user.avatar || 'images/avatar.png'" class="avatar-uploader"/>
                            <el-icon v-else class="avatar-uploader-icon">
                                <Plus/>
                            </el-icon>
                        </el-upload>
                        <div>点击更换头像</div>
                    </el-form-item>
                    <el-form-item label="用户名" prop="username">
                        <el-input v-model="user.username"></el-input>
                    </el-form-item>
                    <el-form-item label="手机号" prop="imei">
                        <el-input v-model="user.phone"></el-input>
                    </el-form-item>
                    <el-form-item label="邮箱" prop="imei">
                        <el-input v-model="user.email"></el-input>
                    </el-form-item>
                    <el-form-item label="密码" prop="password">
                        <el-input type="password" v-model="user.password"></el-input>
                    </el-form-item>
                    <el-form-item label="确认密码" prop="password">
                        <el-input type="password" v-model="user.confirm_password"></el-input>
                    </el-form-item>
                    <el-form-item>
                        <el-button type="primary" @click="editUser()">提交</el-button>
                    </el-form-item>
                </el-form>
            </el-row>
        </el-main>
    </el-container>

</template>

<script>
import {userUpdate} from "api/user";
import {upload} from "api/upload.js";

export default {
    setup() {
        return {}
    },
    data() {
        return {
            user: {}
        };
    },
    created() {
        this.user = {...this.$store.state.user};
    },
    methods: {
        //修改用户
        editUser() {
            // 过滤memory字段
            delete this.user.memory;
            // 如果password 和 confirm_password都不为空
            if (this.user.password || this.user.confirm_password) {
                if (this.user.password !== this.user.confirm_password) {
                    this.$message.error("两次密码不一致");
                    return;
                }
            }
            if (!this.user.password) {
                // 移除密码字段
                delete this.user.password;
            }
            delete this.user.confirm_password;
            userUpdate(this.user)
                .then((res) => {
                    if (res.code === 0) {
                        this.user.password = '';
                        this.$store.commit('SET_TOKEN', res.data.token)
                        this.$store.commit('SET_USERINFO', res.data)
                        this.$store.commit('SET_MEMORY', res.count)
                        this.$message({
                            message: "修改信息成功",
                            type: "success",
                        });
                    } else {
                        this.$message.error("修改信息失败");
                    }
                })
                .catch((err) => {
                    this.$message.error("修改信息异常");
                    console.log(err);
                });
        },

        // 校验头像大小、格式
        handleAvatarBeforeRead(file) {
            const isJPG = file.type === 'image/jpeg'
            const isPNG = file.type === 'image/png'
            const isLt2M = file.size / 1024 / 1024 < 1

            if (!isJPG && !isPNG) {
                this.$message.error('必须是JPG/PNG格式!')
            }
            if (!isLt2M) {
                this.$message.error('大小不超过 1MB!')
            }
            return (isJPG || isPNG) && isLt2M
        },

        // 上传头像
        handleUploadAvatar(file) {
            const formData = new FormData()
            formData.append('file', file.file)
            formData.append('save', 'false') // 是否保存到数据库
            upload(formData).then((res) => {
                if (res.code === 0) {
                    this.user.avatar = res.data;
                } else {
                    this.$message.error("上传头像失败");
                }
            }).catch((err) => {
                console.debug(err);
                this.$message.error("上传头像异常");
            });
        }
    }
};
</script>

<style>
.el-row {
    margin-bottom: 20px;
}

.el-col {
    border-radius: 4px;
}

.avatar-uploader {
    height: 60px;
    width: 60px;
}
</style>
