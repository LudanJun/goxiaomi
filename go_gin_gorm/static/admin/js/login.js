$(function(){ // 页面加载完成
    loginApp.init(); // 调用初始化方法
})
var loginApp={
    init:function(){
        this.getCaptcha() // 调用获取验证码方法
        this.captchaImgChage() // 调用验证码切换方法
    },
    getCaptcha: function () {// 获取验证码
        //请求接口获取验证码 回调函数返回数据
        //Math.random() 防止浏览器缓存 
        $.get("/admin/captcha?t="+Math.random(),function(response){// 发送ajax请求
            console.log(response)// 打印返回的验证码信息
            $("#captchaId").val(response.captchaId) // 将验证码id赋值给隐藏域
            $("#captchaImg").attr("src",response.captchaImage)// 将验证码图片赋值给img标签
        })
    },
    captchaImgChage:function(){
        var that=this; // 保存this的指向
        $("#captchaImg").click(function(){ // 点击验证码图片时
            that.getCaptcha() // 调用获取验证码方法
        })
    }
}