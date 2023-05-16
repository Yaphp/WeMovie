<template>
    <div v-loading="loading" element-loading-text="注册中..." element-loading-spinner="el-icon-loading"
         element-loading-background="rgba(0, 0, 0, 0.6)" class="login-container">

        <el-form ref="registerForm" :model="registerForm" :rules="loginRules" class="login-form" auto-complete="on"
                 label-position="left">
            <div class="title-container">
                <img class="logo-img" src="/images/logo.png">
            </div>

            <div class="title-container">
                <h3 class="title">{{ this.$store.state.app.name }}</h3>
            </div>

            <el-form-item prop="username">
                <span class="svg-container">
                    <svg-icon icon-class="user"/>
                </span>
                <el-input v-model="registerForm.username" placeholder="请输入用户名" size="large"/>
            </el-form-item>
            <el-form-item prop="password">
                <span class="svg-container">
                    <svg-icon icon-class="password"/>
                </span>
                <el-input type="password" ref="password" v-model="registerForm.password" size="large" show-password
                          placeholder="请输入密码" name="password" tabindex="2" auto-complete="on"/>
            </el-form-item>
            <div>
                <el-button type="primary" style="width:100%;margin-bottom:20px;" @click.native.prevent="handleRegister"
                           @keyup.enter.native="handleRegister">
                    注册
                </el-button>

                <div class="register">
                    <router-link to="/">已有账号,立即登录</router-link>
                </div>
            </div>
        </el-form>
    </div>
</template>

<script>
import {userRegister, userLogin} from "api/login";

export default {
    name: 'Login',
    data() {
        const validateUsername = (rule, value, callback) => {
            if (!value) {
                callback(new Error('用户名不能为空！'))
            } else {
                callback()
            }
        }
        const validatePassword = (rule, value, callback) => {
            if (value.length < 3) {
                callback(new Error('密码最少为6位字符！'))
            } else {
                callback()
            }
        }
        return {
            // 登录表单
            registerForm: {
                username: '',
                password: ''
            },
            // 登录规则
            loginRules: {
                username: [{required: true, trigger: 'blur', validator: validateUsername}],
                password: [{required: true, trigger: 'blur', validator: validatePassword}]
            },
            loading: false,
            redirect: this.$route.query.redirect || null
        }
    },
    methods: {
        // 登录业务
        handleRegister() {
            this.$refs.registerForm.validate((valid) => {
                // 如果符合验证规则
                if (valid) {
                    this.loading = true
                    // 发送登录请求
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
                        console.log(result)
                        this.loading = false
                    })
                } else {
                    console.log('error submit!!')
                    return false
                }
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

.login-form {
    width: 500px;
    margin: 0 auto;
    padding: 20px;
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