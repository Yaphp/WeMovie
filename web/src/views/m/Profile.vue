<template>
    <el-container class="el-container-fill">
        <el-main>
            <van-row class="profile">
                <van-col span="24" class="avatar">
                    <van-uploader :before-read="handleAvatarBeforeRead" :after-read="handleUploadAvatar">
                        <van-image round width="100" height="100" fit="cover" :src="user.avatar"/>
                    </van-uploader>
                </van-col>
                <van-col span="24" class="avatar-text">
                    点击头像修改
                </van-col>
            </van-row>

            <van-cell-group>
                <van-field v-model="user.username" label="昵称"/>
                <van-field v-model="user.phone" label="手机号"/>
                <van-field v-model="user.email" label="邮箱"/>
                <van-field v-model="user.lasttime_at" readonly label="上次登录时间"/>
                <van-field type="password" v-model="user.password" label="密码"/>
                <van-field type="password" v-model="user.confirm_password" label="确认密码"/>
            </van-cell-group>

            <van-divider
                :style="{ color: '#ffffff', borderColor: '#ffffff', padding: '0 16px' }"
            >
            </van-divider>

            <van-button type="primary" size="small" block round @click="editUser">保存</van-button>
        </el-main>
    </el-container>
</template>

<script>
import {userUpdate} from "api/user";
import {upload} from "api/upload";

export default {
    setup() {
        return {}
    },
    data() {
        return {
            user: {}, // 修改用户信息表单
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
            userUpdate(this.user).then((res) => {
                if (res.code === 0) {
                    // 更新token和用户信息
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
            }).catch((err) => {
                console.debug(err);
                this.$message.error("修改信息异常");
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
    },
};
</script>

<style scoped>
.el-container-fill {
    width: 100vw;
    height: 100vh;
    background-color: #f5f5f6;
}

.avatar {
    display: flex;
    justify-content: center;
    align-items: center;
}

.avatar-text {
    margin-bottom: 20px;
    text-align: center;
    font-size: 14px;
    color: #999999;
}
</style>
