<script setup lang="ts">
import { ref, reactive } from "vue";
import type { FormRules } from "element-plus";
import { SysEmailFormProps } from "@/api/sys/sys-email";

const props = withDefaults(defineProps<SysEmailFormProps>(), {
  formInline: () => ({
    id: 0,
    email: null,
    code: null,
    type: null,
    status: null,
    useStatus: null,
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
  <el-form-item label="主键编码" prop="id">
      <el-input
        v-model.number="newFormInline.id"
        clearable
        placeholder="请输入主键编码"
      />
    </el-form-item>
  <el-form-item label="邮箱地址" prop="email">
      <el-input
        v-model="newFormInline.email"
        clearable
        placeholder="请输入邮箱地址"
      />
    </el-form-item>
  <el-form-item label="验证码" prop="code">
      <el-input
        v-model="newFormInline.code"
        clearable
        placeholder="请输入验证码"
      />
    </el-form-item>
  <el-form-item label="类型" prop="type">
      <el-input
        v-model="newFormInline.type"
        clearable
        placeholder="请输入类型"
      />
    </el-form-item>
  <el-form-item label="状态" prop="status">
      <el-input
        v-model.number="newFormInline.status"
        clearable
        placeholder="请输入状态"
      />
    </el-form-item>
  <el-form-item label="使用状态" prop="useStatus">
      <el-input
        v-model.number="newFormInline.useStatus"
        clearable
        placeholder="请输入使用状态"
      />
    </el-form-item>
  
  
  
  </el-form>
</template>
