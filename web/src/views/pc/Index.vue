<template>
    <div class="bg">
        <!--        <img class="bg-photo" :src="bg" alt="">-->
        <div class="bg-slogan">
            <h1 class="bg-title">{{ this.$store.state.app.name }}</h1>
        </div>
    </div>

</template>

<script>

import {ElMessage, ElMessageBox} from 'element-plus'

export default {
    data() {
        return {
            bg: "",
        };
    },
    mounted() {
        this.bg = bg;
        //用户文件列表
        if (this.$store.state.user.role === "2") {
            // this.getUserData();
        }
    },
    methods: {
        getUserData() {
            userData({
                "id": this.$store.state.user.id
            }).then((res) => {
                // console.log()
                if (!Object.keys(res.data).length) {
                    return;
                } else {
                    ElMessageBox.alert(res.data.content, '通知', {
                        confirmButtonText: '确认',
                        type: 'warning',
                        callback: () => {
                            userRead({
                                "id": res.data.id
                            }).then((res) => {
                                console.log(res);
                            }).catch((err) => {
                                console.log(err);
                            });
                        },
                    })
                }

            }).catch((err) => {
                console.log(err);
            });
        }
    }
};
</script>

<style>
.el-main {
    padding: unset !important;
}

.bg {
    width: 100%;
    height: 100%;
    overflow: hidden;
    background-color: white;
}

.bg-photo {
    width: 100%;
    height: 100%;
}

.bg-slogan {
    position: absolute;
    z-index: 999;
    top: 45%;
    left: 40%;
}

.bg-title {
    color: #579cf5;
    font-size: 60px;
}
</style>