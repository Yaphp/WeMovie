<template>
    <div v-loading="loading" element-loading-text="注册中..." element-loading-spinner="el-icon-loading"
         element-loading-background="rgba(0, 0, 0, 0.6)" class="login-container">
        <van-form @submit="handleRegister" class="m-login-form">
            <div class="title-container">
                <img class="logo-img" src="/images/logo.png">
            </div>

            <div class="title-container">
                <h3 class="title">{{ this.$store.state.app.name }}</h3>
            </div>

            <van-cell-group inset>
                <van-field
                    v-model="registerForm.username"
                    name="用户名"
                    label="用户名"
                    size="large"
                    placeholder="用户名"
                    :rules="[{ required: true, message: '请填写用户名' }]"
                />
                <van-field
                    v-model="registerForm.password"
                    type="password"
                    name="密码"
                    label="密码"
                    size="large"
                    placeholder="密码"
                    :rules="[{ required: true, message: '请填写密码' }]"
                />
            </van-cell-group>
            <div style="margin: 16px;">
                <van-button round block type="primary" native-type="submit">
                    注册
                </van-button>
            </div>
            <div class="register">
                <router-link to="/">已有账号,立即登录</router-link>
            </div>
        </van-form>
    </div>

    <van-loading v-if="loading" type="spinner"/>
</template>

<script>
import {userRegister} from "api/login";

export default {
    name: 'Login',
    data() {
        return {
            // 登录表单
            registerForm: {
                username: '',
                password: ''
            },
            loading: false,
            redirect: this.$route.query.redirect || null
        }
    },
    methods: {
        // 登录业务
        handleRegister() {
            this.loading = true
            userRegister(this.registerForm).then((res) => {
                console.log(res)
                this.loading = false
                if (res.code === 0) {
                    this.$message({
                        message: '注册成功',
                        type: 'success',
                        duration: 1000,
                        onClose: () => {
                            this.$router.push({path: '/'})
                        }
                    })
                } else {
                    this.$message({
                        message: res.msg,
                        type: 'error',
                        duration: 1500
                    })
                }
            }).catch((result) => {
            })
        }
    }
}
</script>


<style>
.logo-img {
    width: 80px;
    height: 80px;
    margin: 0 auto;
}

.login-container {
    display: flex;
    justify-content: center;
    align-self: center;
    flex-direction: column;
    width: 100%;
    height: 100%;
    background: #005aff linear-gradient(to right, rgba(0, 255, 0, 0), rgb(0 255 173 / 50%));
}

.title-container {
    margin: 30px auto;
    text-align: center;
}

.m-login-form {
    width: 90%;
    margin: 0 auto;
    padding: 20px 0 30px 0;
    border-radius: 5px;
    background-color: #fff;
}

.register {
    width: 100%;
    margin: auto;
    font-size: 14px;
    color: #409EFF;
    text-align: center;
}
</style>