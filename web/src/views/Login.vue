<template>
    <div v-loading="loading" element-loading-text="登录中..." element-loading-spinner="el-icon-loading"
         element-loading-background="rgba(0, 0, 0, 0.6)" class="login-container">

        <el-form ref="loginForm" :model="loginForm" :rules="loginRules" class="login-form" auto-complete="on" label-position="left">
            <div class="title-container">
                <h3 class="title">{{this.$store.state.app.name}}</h3>
            </div>

            <el-form-item prop="username">
                <span class="svg-container">
                    <svg-icon icon-class="user"/>
                </span>
                <el-input v-model="loginForm.username" placeholder="请输入用户名" size="large"/>
            </el-form-item>
            <el-form-item prop="password">
                <span class="svg-container">
                    <svg-icon icon-class="password"/>
                </span>
                <el-input :key="passwordType" ref="password" v-model="loginForm.password" :type="passwordType" size="large" show-password placeholder="请输入密码" name="password" tabindex="2" auto-complete="on" />
            </el-form-item>

            <div>
                <el-button type="primary" style="width:100%;margin-bottom:20px;" @click.native.prevent="handleLogin" @keyup.enter.native="handleLogin">
                    登录
                </el-button>
            </div>
        </el-form>
    </div>
</template>

<script>
import {userLogin} from "@/api/login";

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
            loginForm: {
                username: 'admin',
                password: '123'
            },
            // 登录规则
            loginRules: {
                username: [{required: true, trigger: 'blur', validator: validateUsername}],
                password: [{required: true, trigger: 'blur', validator: validatePassword}]
            },
            loading: false,
            redirect: undefined
        }
    },
    methods: {
        // 登录业务
        handleLogin() {
            this.$refs.loginForm.validate((valid) => {
                // 如果符合验证规则
                if (valid) {
                    this.loading = true
                    // 发送登录请求
                    userLogin(this.loginForm).then((res) => {
                        console.log(res)
                        this.loading = false
                        if (res.code === 200) {
                            //保存token
                            this.$store.commit('SET_TOKEN', res.data.token)
                            this.$store.commit('SET_USERINFO', res.data)
                            this.$message({
                                message: '登录成功',
                                type: 'success',
                                duration: 1500,
                                onClose: () => {
                                    this.$router.push({path: this.redirect || '/index'})
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
}

.login-form {
    width: 500px;
    margin: 0 auto;
    padding: 20px;
    border-radius: 5px;
    background-color: #fff;
}
</style>