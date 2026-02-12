<script setup lang="ts">
import { ref, reactive } from "vue";
import { SysRoleFormProps, getSysRole } from "@/api/sys/sys-role";

import { usePublicHooks } from "@/utils/hooks";

const props = withDefaults(defineProps<SysRoleFormProps>(), {
  formInline: () => ({
    type: "",
    id: 0,
    name: null,
    roleKey: null,
    roleSort: 0,
    remark: null,
    menuIds: null,
    status: 0,
    halfMenuIds: null,
    higherDeptOptions: []
  })
});

const { switchStyle } = usePublicHooks();

/** 自定义表单规则校验 */
const formRules = reactive({
  name: [{ required: true, message: "名称为必填项", trigger: "blur" }]
});

const ruleFormRef = ref();
const treeRef = ref();
const newFormInline = ref(props.formInline);

function getRef() {
  return ruleFormRef.value;
}
const handleCheckChange = () => {
  // console.log(treeRef.value.getCheckedKeys());
  // console.log(treeRef.value.getHalfCheckedKeys());
  newFormInline.value.menuIds = treeRef.value.getCheckedKeys();
  newFormInline.value.halfMenuIds = treeRef.value.getHalfCheckedKeys();
};

const getDetail = () => {
  if (newFormInline.value.type == "edit") {
    getSysRole({ id: newFormInline.value.id }).then(res => {
      if (res.code == 200) {
        (newFormInline.value.name = res.data?.name ?? null),
          (newFormInline.value.roleKey = res.data?.roleKey ?? null),
          (newFormInline.value.roleSort = res.data?.roleSort ?? null),
          (newFormInline.value.remark = res.data?.remark ?? null),
          (newFormInline.value.menuIds = res.data?.menuIds ?? null);
      }
    });
  }
};
getDetail();

defineExpose({ getRef });
</script>

<template>
  <el-form
    ref="ruleFormRef"
    :model="newFormInline"
    :rules="formRules"
    label-width="82px"
  >
    <el-form-item label="角色名称" prop="name">
      <el-input
        v-model="newFormInline.name"
        clearable
        placeholder="请输入角色名称"
      />
    </el-form-item>

    <el-form-item label="角色代码" prop="roleKey">
      <el-input
        v-model="newFormInline.roleKey"
        clearable
        placeholder="请输入角色代码"
      />
    </el-form-item>
    <el-form-item label="状态" prop="status">
      <el-switch
        v-model="newFormInline.status"
        inline-prompt
        :active-value="1"
        :inactive-value="-1"
        active-text="启用"
        inactive-text="停用"
        :style="switchStyle"
      />
    </el-form-item>
    <el-form-item label="排序" prop="roleSort">
      <el-input
        v-model.number="newFormInline.roleSort"
        clearable
        placeholder="请输入排序"
      />
    </el-form-item>
    <el-form-item label="备注" prop="remark">
      <el-input
        v-model="newFormInline.remark"
        clearable
        placeholder="请输入备注"
      />
    </el-form-item>

    <el-form-item label="授权菜单" prop="menuIds">
      <el-tree
        ref="treeRef"
        :data="newFormInline.higherDeptOptions"
        @check-change="handleCheckChange"
        node-key="id"
        show-checkbox
        :default-expanded-keys="newFormInline.menuIds"
        :default-checked-keys="newFormInline.menuIds"
        :props="{
          label: 'title'
        }"
      />
    </el-form-item>
  </el-form>
</template>
