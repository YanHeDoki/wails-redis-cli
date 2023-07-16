<template>
<main>
  <el-dialog
      v-model="KeyDialogVisible"
      title=创建键
      width="60%"
  >
    <el-form :model="KeyForm" label-width="120px">
      <el-form-item label="键名称">
        <el-input placeholder="请输入键名称" v-model="KeyForm.key" />
      </el-form-item>
      <el-form-item label="值">
        <el-input placeholder="请输入值" v-model="KeyForm.Kvalue" />
      </el-form-item>
      <el-form-item label="键类型">
        <el-select v-model="KeyForm.type">
        <el-option value="string" label="string"></el-option>
        </el-select>
      </el-form-item>
      <el-form-item>
        <el-button @click="createKey">创建</el-button>
        <el-button @click="KeyDialogVisible=false">取消</el-button>
      </el-form-item>
    </el-form>
  </el-dialog>

  <el-form :inline="true" :model="form" class="demo-form-inline">
    <el-form-item label="">
      <el-input v-model="form.keyword" placeholder="输入键的信息" clearable />
    </el-form-item>

    <el-form-item>
      <el-button type="primary" @click="getKeyList">查询</el-button>
    </el-form-item>
  </el-form>
  <el-button @click="KeyDialogVisible=true" style="width: 100%; margin: 12px;">创建键</el-button>
  <div v-for="item in keys" @click="selectKeyKey(item)">
  <div v-if="item===selectKey">
    <div class="item key_select_item">
      <div style="padding: 5px 12px">{{item}}</div>
      <el-popconfirm title="确认删除吗?" @confirm="KeyDelete(item)">
        <template #reference>
          <el-button text type="danger" @click.stop>删除</el-button>
        </template>
      </el-popconfirm>
    </div>
  </div>
    <div v-else>
      <div class="item key_item">
        <div style="padding: 5px 12px">{{item}}</div>
        <el-popconfirm title="确认删除吗?" @confirm="KeyDelete(item)">
          <template #reference>
            <el-button text type="danger" @click.stop>删除</el-button>
          </template>
        </el-popconfirm>
      </div>
    </div>
  </div>
</main>
</template>

<script setup>
import {ConnectionDelete, CreateKeyValue, DeleteKeyValue, KeyList} from "../../wailsjs/go/main/App.js";
  import {reactive, ref, watch} from 'vue'
  import {ElNotification} from "element-plus";

  let props =defineProps(['keyDb','keyConnIdentity'])
  let keys=ref()
  let selectKey=ref()
  let emits=defineEmits(['emit-select-key'])
  let KeyDialogVisible=ref(false)
  let KeyForm=ref({})

  watch(props,()=>{
    getKeyList()
  })
  const form = reactive({
    keyword: '',
  })

  // const onSubmit = () => {
  //   getKeyList()
  // }
  //
  function getKeyList() {
    let data={
      conn_identity:props.keyConnIdentity,
      db:props.keyDb,
      keyword:form.keyword,
    }
    KeyList(data).then(res=>{
      if (res.code!==200){
        ElNotification({
          title:res.msg,
          type:"error",
        })
        return
      }
      keys.value=res.data
    })
  }

function selectKeyKey(key){
    selectKey.value=key
  emits("emit-select-key",key)
}

function KeyDelete(key) {
  DeleteKeyValue({conn_identity:props.keyConnIdentity,db:props.keyDb,key:key}).then(res=>{
    if (res.code!==200){
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
    getKeyList()
  })
}

function createKey() {
  CreateKeyValue({conn_identity:props.keyConnIdentity,db:props.keyDb,key:KeyForm.value.key,type:KeyForm.value.type,Kvalue:KeyForm.value.Kvalue}).then(res=>{
    if (res.code!==200){
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
    KeyDialogVisible=false
    getKeyList()
  })
}


</script>

<style scoped>

.key_item{
  color: #409eff;
  background-color: #ecf5ff;
  margin-bottom: 5px;
}

.key_select_item{
  color: #67c23a;
  background-color: #f0f9eb;
  margin-bottom: 5px;
}

</style>