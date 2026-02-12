export const genderOptions = [
  {
    value: 1,
    label: "男"
  },
  {
    value: 2,
    label: "女"
  }
];

export const statusOptions = [
  {
    value: 3,
    label: "正常"
  },
  {
    value: 1,
    label: "默认"
  },
  {
    value: -1,
    label: "删除"
  }
];

export const statusSimpleOptions = [
  {
    value: 1,
    label: "启用"
  },
  {
    value: -1,
    label: "关闭"
  }
];

export const findPathById = (
  tree: Array,
  targetId: number,
  path: string = ""
) => {
  for (const node of tree) {
    const currentPath = path ? `${path}/${node.name}` : node.name;
    if (node.id === targetId) {
      return currentPath;
    }
    if (node.children && node.children.length > 0) {
      const result = findPathById(node.children, targetId, currentPath);
      if (result) {
        return result;
      }
    }
  }
  return "-";
};
