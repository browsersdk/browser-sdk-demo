import dayjs from "dayjs";
import editForm from "../form.vue";
import { message } from "@/utils/message";
import {
  getSysTeamPage,
  createSysTeam,
  updateSysTeam,
  delSysTeam
} from "@/api/sys/sys-team";
import { useRouter } from "vue-router";
import { addDialog } from "@/components/ReDialog";
import { type SysTeamFormItemProps } from "@/api/sys/sys-team";
import { statusOptions } from "@/api/sys/common";
import { type PaginationProps } from "@pureadmin/table";
import { reactive, ref, onMounted, h, toRaw } from "vue";

export function useRole() {
  const form = reactive({
    page: 1,
    pageSize: 10,
    id: 0,
    name: null,
    status: null
  });
  const formRef = ref();
  const dataList = ref([]);
  const loading = ref(true);
  const router = useRouter();
  //const switchLoadMap = ref({});
  //const { switchStyle } = usePublicHooks();
  const pagination = reactive<PaginationProps>({
    total: 0,
    pageSize: 10,
    currentPage: 1,
    background: true
  });
  const columns: TableColumnList = [
    {
      label: "主键",
      prop: "id",
      minWidth: 120
    },
    {
      label: "团队名",
      prop: "name",
      minWidth: 120
    },
    {
      label: "Sk",
      prop: "appSk",
      minWidth: 120
    },
    {
      label: "状态",
      prop: "status",
      minWidth: 120,
      formatter: ({ status }) => {
        return statusOptions.find(item => item.value === status)?.label;
      }
    },
    {
      label: "创建时间",
      prop: "createdAt",
      minWidth: 120,
      formatter: ({ createTime }) =>
        dayjs(createTime).format("YYYY-MM-DD HH:mm:ss")
    },
    {
      label: "更新时间",
      prop: "updatedAt",
      minWidth: 120,
      formatter: ({ createTime }) =>
        dayjs(createTime).format("YYYY-MM-DD HH:mm:ss")
    },
    {
      label: "操作",
      width: 280,
      slot: "operation"
    }
  ];

  function handleDelete(row) {
    delSysTeam({ ids: [row.id] }).then(res => {
      if (res.code == 200) {
        message(`删除成功`, { type: "success" });
        onSearch();
      } else {
        message(`删除失败`, { type: "error" });
      }
    });
  }

  function handleSizeChange(val: number) {
    form.pageSize = val;
    onSearch();
  }

  function handleCurrentChange(val: number) {
    form.page = val;
    onSearch();
  }

  function handleSelectionChange(val) {
    console.log("handleSelectionChange", val);
  }

  async function onSearch() {
    loading.value = true;
    const { data } = await getSysTeamPage(toRaw(form));
    dataList.value = data.list;
    pagination.total = data.total;
    pagination.pageSize = data.pageSize;
    pagination.currentPage = data.currentPage;

    setTimeout(() => {
      loading.value = false;
    }, 500);
  }

  const resetForm = formEl => {
    if (!formEl) return;
    formEl.resetFields();
    onSearch();
  };

  function toDept(row: SysTeamFormItemProps) {
    router.push({
      path: "/sys/sys-dept", // 确保路由配置中定义了 name 属性
      query: {
        id: row.id
      }
    });
  }

  function toRole(row: SysTeamFormItemProps) {
    router.push({
      path: "/sys/sys-role", // 确保路由配置中定义了 name 属性
      query: {
        id: row.id
      }
    });
  }

  function toMember(row: SysTeamFormItemProps) {
    router.push({
      path: "/sys/sys-member", // 确保路由配置中定义了 name 属性
      query: {
        id: row.id
      }
    });
  }

  function openDialog(title = "新增", row?: SysTeamFormItemProps) {
    addDialog({
      title: `${title}团队`,
      props: {
        formInline: {
          id: row?.id ?? 0,
          name: row?.name ?? "",
          owner: row?.owner ?? 0,
          status: row?.status ?? 0
        }
      },
      width: "48%",
      draggable: true,
      fullscreenIcon: true,
      closeOnClickModal: false,
      contentRenderer: () => h(editForm, { ref: formRef }),
      beforeSure: (done, { options }) => {
        const FormRef = formRef.value.getRef();
        const curData = options.props.formInline as SysTeamFormItemProps;
        FormRef.validate(valid => {
          if (valid) {
            // 表单规则校验通过
            if (title === "新增") {
              createSysTeam(curData).then(res => {
                if (res.code == 200) {
                  message(res.msg, {
                    type: "success"
                  });
                  onSearch(); // 刷新表格数据
                } else {
                  message(res.msg, {
                    type: "error"
                  });
                }
              });
            } else {
              updateSysTeam(curData).then(res => {
                if (res.code == 200) {
                  message(res.msg, {
                    type: "success"
                  });
                  onSearch(); // 刷新表格数据
                } else {
                  message(res.msg, {
                    type: "error"
                  });
                }
              });
            }
            done(); // 关闭弹框
          }
        });
      }
    });
  }

  /** 菜单权限 */
  function handleMenu() {
    message("等菜单管理页面开发后完善");
  }

  /** 数据权限 可自行开发 */
  // function handleDatabase() {}

  onMounted(() => {
    onSearch();
  });

  return {
    form,
    loading,
    columns,
    dataList,
    pagination,
    statusOptions,
    toDept,
    toRole,
    toMember,
    onSearch,
    resetForm,
    openDialog,
    handleMenu,
    handleDelete,
    handleSizeChange,
    handleCurrentChange,
    handleSelectionChange
  };
}
