<script setup lang="ts">
import { ref, reactive } from "vue";
import type { FormRules } from "element-plus";
import { UploadLogFormProps } from "@/api/sys/upload-log";

const props = withDefaults(defineProps<UploadLogFormProps>(), {
  formInline: () => ({
    id: 0,
    url: null,
    provider: null,
    status: null,
    syncStatus: null,
  })
});

/** 自定义表单规则校验 */
const formRules = reactive(<FormRules>{
  name: [{ required: true, message: "名称为必填项", trigger: "blur" }]
});

const ruleFormRef = ref();
const newFormInline = ref(props.formInline);

function getRef() {
  return ruleFormRef.value;
}

defineExpose({ getRef });
</script>

<template>
  <el-form
    ref="ruleFormRef"
    :model="newFormInline"
    :rules="formRules"
    label-width="82px"
  >
  <el-form-item label="id" prop="id">
      <el-input
        v-model.number="newFormInline.id"
        clearable
        placeholder="请输入id"
      />
    </el-form-item>
  <el-form-item label="资源地址" prop="url">
      <el-input
        v-model="newFormInline.url"
        clearable
        placeholder="请输入资源地址"
      />
    </el-form-item>
  <el-form-item label="提供方" prop="provider">
      <el-input
        v-model="newFormInline.provider"
        clearable
        placeholder="请输入提供方"
      />
    </el-form-item>
  <el-form-item label="1使用中 -1 删除" prop="status">
      <el-input
        v-model.number="newFormInline.status"
        clearable
        placeholder="请输入1使用中 -1 删除"
      />
    </el-form-item>
  <el-form-item label="同步状态" prop="syncStatus">
      <el-input
        v-model.number="newFormInline.syncStatus"
        clearable
        placeholder="请输入同步状态"
      />
    </el-form-item>
  
  
  
  </el-form>
</template>
