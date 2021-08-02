export const asyncValidator = {
    "default": '字段 %s 验证错误',
    required: '%s 是必需的',
    "enum": '%s 必须是 %s 之一',
    whitespace: '%s 不能为空',
    date: {
        format: '%s 日期 %s 对格式 %s 无效',
        parse: '%s 日期无法解析，%s 无效',
        invalid: '%s 日期 %s 无效'
    },
    types: {
        string: '%s 不是 %s',
        method: '%s 不是 %s (function)',
        array: '%s 不是 %s %s',
        object: '%s 不是 %s',
        number: '%s 不是 %s',
        date: '%s 不是 %s',
        "boolean": '%s 不是 %s',
        integer: '%s 不是 %s',
        "float": '%s 不是 %s',
        regexp: '%s 不是有效的 %s',
        email: '%s 不是有效的 %s',
        url: '%s 不是有效的 %s',
        hex: '%s 不是有效的 %s'
    },
    string: {
        len: '%s 必须正好是 %s 个字符',
        min: '%s 必须至少有 %s 个字符',
        max: '%s 不能超过 %s 个字符',
        range: '%s 必须在 %s 和 %s 个字符之间'
    },
    number: {
        len: '%s 必须等于 %s',
        min: '%s 不能小于 %s',
        max: '%s 不能大于 %s',
        range: '%s 必须在 %s 和 %s 之间'
    },
    array: {
        len: '%s 的长度必须为 %s',
        min: '%s 的长度不能小于 %s',
        max: '%s 的长度不能大于 %s',
        range: '%s 的长度必须在 %s 和 %s 之间'
    },
    pattern: {
        mismatch: '%s 值 %s 与模式 %s 不匹配'
    },
    // clone: function clone() {
    //     var cloned = JSON.parse(JSON.stringify(this));
    //     cloned.clone = this.clone;
    //     return cloned;
    // }
}