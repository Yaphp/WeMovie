<template>
    <div class="table-container">
        <!--Layout布局-->
        <el-row>
            <el-col :span="24">
                <el-row :gutter="20">
                    <el-col :span="6">
                        <!--搜索区域-->
                        <el-input
                            placeholder="请输入查询内容"
                            v-model="queryInfo.keyword"
                            clearable
                            @keyup="getUserList"
                            @clear="getUserList">
                        </el-input>
                    </el-col>
                    <el-col :span="2.5">
                        <el-button type="primary" @click="addDialogVisible = true">
                            添加用户
                        </el-button>
                    </el-col>
                    <el-col :span="2.5">
                        <el-button type="danger" @click="batchDeleteUser">
                            批量删除
                        </el-button>
                    </el-col>
                </el-row>
            </el-col>
            <el-col :span="24">
                <!--表格-->
                <el-table
                    :data="userList"
                    height="500"
                    ref="filterTable"
                    border
                    stripe
                    @selection-change="handleSelectionChange"
                >
                    <el-table-column type="selection" width="55"></el-table-column>
                    <el-table-column prop="id" label="id"></el-table-column>
                    <el-table-column prop="username" label="用户名"></el-table-column>
                    <el-table-column prop="phone" label="手机号"></el-table-column>
                    <el-table-column prop="email" label="邮箱"></el-table-column>
                    <el-table-column prop="role" label="角色">
                        <template #default="scope">
                            <el-tag>{{ scope.row.role === 0 ? '管理员' : '普通用户' }}</el-tag>
                        </template>
                    </el-table-column>
                    <el-table-column prop="created_at" label="创建时间"></el-table-column>
                    <el-table-column prop="lasttime_at" label="最近登录时间"></el-table-column>
                    <el-table-column label="操作" width="200" v-if="this.$store.state.user.role === 0">
                        <!-- 作用域插槽 -->
                        <template #default="scope">
                            <!--修改按钮-->
                            <el-button
                                type="primary"
                                size="small"
                                :icon="Edit"
                                @click="showEditDialog(scope.row)"
                            >修改
                            </el-button>
                            <!--删除按钮-->
                            <el-button
                                type="danger"
                                size="small"
                                :icon="Delete"
                                @click="removeUserById(scope.row.id)"
                            >删除
                            </el-button>
                        </template>
                    </el-table-column>
                </el-table>
            </el-col>
        </el-row>

        <el-row>
            <!--分页区域-->
            <el-pagination
                @size-change="handleSizeChange"
                @current-change="handleCurrentChange"
                :current-page="queryInfo.pageNo"
                :page-sizes="[10, 20, 50, 100]"
                :page-size="queryInfo.pageSize"
                layout="total, sizes, prev, pager, next, jumper"
                :total="total"
            >
            </el-pagination>
        </el-row>

        <!--添加用户的对话框-->
        <el-dialog
            title="添加用户"
            v-model="addDialogVisible"
            width="30%"
            @close="addDialogClosed"
        >
            <!--内容主体区域-->
            <el-form :model="addForm" label-width="70px" ref="addFormRef">
                <el-form-item label="用户名" prop="username">
                    <el-input v-model="addForm.username"></el-input>
                </el-form-item>
                <el-form-item label="密码" prop="password">
                    <el-input type="password" v-model="addForm.password"></el-input>
                </el-form-item>
                <el-form-item label="角色" prop="role">
                    <el-select v-model="addForm.role">
                        <el-option
                            v-for="item in options"
                            :key="item.value"
                            :label="item.label"
                            :value="item.value"
                        />
                    </el-select>
                </el-form-item>
                <el-form-item label="手机号" prop="imei">
                    <el-input v-model="addForm.phone"></el-input>
                </el-form-item>
                <el-form-item label="邮箱" prop="imei">
                    <el-input v-model="addForm.email"></el-input>
                </el-form-item>
            </el-form>
            <!--底部按钮区域-->
            <span slot="footer" class="dialog-footer">
                <el-button @click="addDialogVisible = false">取 消</el-button>
                <el-button type="primary" @click="addUser">确 定</el-button>
            </span>
        </el-dialog>


        <!--修改用户的对话框-->
        <el-dialog title="修改用户" v-model="editDialogVisible" width="30%">
            <!--内容主体区域-->
            <el-form :model="editForm" label-width="70px">
                <el-form-item label="用户名" prop="username">
                    <el-input v-model="editForm.username"></el-input>
                </el-form-item>
                <el-form-item label="密码" prop="password">
                    <el-input type="password" v-model="editForm.password"></el-input>
                </el-form-item>
                <el-form-item label="角色" prop="role">
                    <el-select v-model="editForm.role">
                        <el-option
                            v-for="item in options"
                            :key="item.value"
                            :label="item.label"
                            :value="item.value"
                        />
                    </el-select>
                </el-form-item>
                <el-form-item label="手机号" prop="imei">
                    <el-input v-model="editForm.phone"></el-input>
                </el-form-item>
                <el-form-item label="邮箱" prop="imei">
                    <el-input v-model="editForm.email"></el-input>
                </el-form-item>
            </el-form>
            <!--底部按钮区域-->
            <span slot="footer" class="dialog-footer">
        <el-button @click="editDialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="editUser">确 定</el-button>
        </span>
        </el-dialog>
    </div>
</template>

<script>
import {userAdd, userDelete, userList, userUpdate} from "@/api/user";
import {Delete, Edit} from "@element-plus/icons-vue";

export default {
    setup() {
        return {
            Edit,
            Delete
        }
    },
    data() {
        return {
            userList: [], // 用户列表
            total: 0, // 用户总数
            options: [
                {key: "管理员", value: 0, label: '管理员'},
                {key: "普通用户", value: 1, label: '普通用户'}
            ],
            // 获取用户列表的参数对象
            queryInfo: {
                keyword: "", // 查询参数
                pageNo: 1, // 当前页码
                pageSize: 10, // 每页显示条数
            },
            addDialogVisible: false, // 控制添加用户对话框是否显示
            addForm: {
                username: "",
                password: "",
                role: "",
                phone: "",
                email: "",
            },
            editDialogVisible: false, // 控制修改用户信息对话框是否显示
            editForm: {
                id: "",
                username: "",
                password: "",
                role: "",
                phone: "",
                email: ""
            },
            multipleSelection: [],
            ids: [],
        };
    },
    created() {
        // 生命周期函数
        this.getUserList();
    },
    methods: {
        getUserList() {
            userList(this.queryInfo)
                .then((res) => {
                    if (res.code === 200) {
                        //用户列表
                        this.userList = res.data.list;
                        this.total = res.data.total || 0;
                    } else {
                        this.$message.error(res.msg);
                    }
                })
                .catch((err) => {
                    console.log(err);
                });
        },

        // 监听 pageSize 改变的事件
        handleSizeChange(newSize) {
            // console.log(newSize)
            this.queryInfo.pageSize = newSize;
            // 重新发起请求用户列表
            this.getUserList();
        },

        // 监听 当前页码值 改变的事件
        handleCurrentChange(newPage) {
            // console.log(newPage)
            this.queryInfo.pageNo = newPage;
            // 重新发起请求用户列表
            this.getUserList();
        },

        // 监听 添加用户对话框的关闭事件
        addDialogClosed() {
            // 表单内容重置为空
            this.$refs.addFormRef.resetFields();
        },

        // 监听 修改用户状态
        showEditDialog(userinfo) {
            this.editDialogVisible = true;
            userinfo.password = "";
            this.editForm = userinfo;
        },

        //添加用户
        addUser() {
            userAdd(this.addForm)
                .then((res) => {
                    console.log(res)
                    if (res.code === 200) {
                        this.addDialogVisible = false;
                        this.getUserList();
                        this.$message({
                            message: "添加用户成功",
                            type: "success",
                        });
                    } else {
                        this.$message.error("添加用户失败");
                    }
                })
                .catch((err) => {
                    this.$message.error("添加用户异常");
                    console.log(err);
                });
        },

        //修改用户
        editUser() {
            userUpdate(this.editForm)
                .then((res) => {
                    if (res.code === 200) {
                        this.editDialogVisible = false;
                        this.getUserList();
                        this.$message({
                            message: "修改用户成功",
                            type: "success",
                        });
                    } else {
                        this.$message.error("修改用户失败");
                    }
                })
                .catch((err) => {
                    this.$message.error("修改用户异常");
                    console.loge(err);
                });
        },

        // 根据ID删除对应的用户信息
        async removeUserById(id) {
            // 弹框 询问用户是否删除
            const confirmResult = await this.$confirm(
                "此操作将永久删除该用户, 是否继续?",
                "提示",
                {
                    confirmButtonText: "确定",
                    cancelButtonText: "取消",
                    type: "warning",
                }
            ).catch((err) => err);

            // console.log(confirmResult)
            if (confirmResult == "confirm") {
                //删除用户
                userDelete(id)
                    .then((res) => {
                        if (res.code === 200) {
                            this.getUserList();
                            this.$message({
                                message: "删除用户成功",
                                type: "success",
                            });
                        } else {
                            this.$message.error("删除用户失败");
                        }
                    })
                    .catch((err) => {
                        this.$message.error("删除用户异常");
                        console.log(err);
                    });
            }
        },

        //批量选中事件处理
        handleSelectionChange(val) {
            console.log(val)
            this.multipleSelection = val;
        },

        //批量删除用户
        async batchDeleteUser() {
            // 弹框 询问用户是否删除
            const confirmResult = await this.$confirm(
                "此操作将永久删除用户, 是否继续?",
                "提示",
                {
                    confirmButtonText: "确定",
                    cancelButtonText: "取消",
                    type: "warning",
                }
            ).catch((err) => err);

            if (confirmResult === "confirm") {
                //向被删除的ids赋值
                this.multipleSelection.forEach((item) => {
                    this.ids.push(item.id);
                });
                userDelete(this.ids.join(","))
                    .then((res) => {
                        if (res.code === 200) {
                            this.getUserList();
                            this.$message({
                                message: "批量删除用户成功",
                                type: "success",
                            });
                        } else {
                            this.$message.error("批量删除用户失败");
                        }
                    })
                    .catch((err) => {
                        this.$message.error("批量删除用户异常");
                        console.log(err);
                    });
            }
        }
    },
};
</script>

<style>
.table-container {
    margin: 40px 20px;
}

.el-row {
    margin-bottom: 20px;
}

.el-col {
    border-radius: 4px;
}

.el-card {
    box-shadow: 0 1px 1px rgba(0, 0, 0, 0.1) !important;
    height: 60pt;
}
</style>
