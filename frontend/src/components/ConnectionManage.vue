<template>
  <el-button :type="btnType" @click="dialogVisible=true">{{title}}</el-button>
  <el-dialog
      v-model="dialogVisible"
      :title="title"
      width="60%"
  >
    <el-form :model="form" label-width="120px">
      <el-form-item label="连接地址">
        <el-input placeholder="请输入连接地址" v-model="form.addr" />
      </el-form-item>
      <el-form-item label="连接名称">
        <el-input placeholder="请输入连接名称" v-model="form.name" />
      </el-form-item>
      <el-form-item label="端口号">
        <el-input placeholder="请输入端口号" v-model="form.port" />
      </el-form-item>
      <el-form-item label="用户名">
        <el-input placeholder="请输入用户名" v-model="form.username" />
      </el-form-item>
      <el-form-item label="密码">
        <el-input placeholder="请输入密码" v-model="form.password" />
      </el-form-item>

      <el-form-item>
        <el-button v-if="data===undefined" type="primary"  @click="createConnection">创建</el-button>
        <el-button v-else type="primary"  @click="editconnection">编辑</el-button>
        <el-button @click="dialogVisible=false">取消</el-button>
      </el-form-item>
    </el-form>
  </el-dialog>
</template>

<script setup>
import { ref,reactive } from 'vue'
import {ConnectionCreate, ConnectionEdit} from "../../wailsjs/go/main/App.js";
import {ElNotification} from "element-plus";
let props=defineProps(['title','btnType','data'])
const dialogVisible = ref(false)
let form = reactive({})
const emits=defineEmits(['emit-connection-List'])
if (props.data!==undefined){
    form=props.data
}
function createConnection() {
  ConnectionCreate(form).then(res=>{
    if (res.code!=200){
      ElNotification({
        title:res.msg,
        type:"error",
      })
    return
    }
    ElNotification({
      title:res.msg,
      type:"success",
    })
    //获取新的连接地址
    emits('emit-connection-List')
  })
}

function editconnection() {
  ConnectionEdit(form).then(res=>{
    if (res.code!=200){
      ElNotification({
        title:res.msg,
        type:res.error,
      })
      return
    }
    ElNotification({
      title:res.msg,
      type:"success",
    })
    //获取新的连接地址
    emits('emit-connection-List')
  })
}

</script>

<style scoped>

</style>