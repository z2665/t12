/**
 * Created by Alkali on 15/10/19.
 */
$(function(){
    var infoHead = $("#info-head"),
        infoHeadBox = $("#info-head-box"),
        info = $("#info"),
        personalInfo = $("#personal-info"),
        updatePw = $("#update-pw"),
        pw = $("#pw"),
        intoVim = $("#into-vim"),
        ensureInfo = $("#ensure-info"),
        cancel = $("#cancel"),
        username = $("#username"),
        name = $("#name"),
        stuNum = $("#stu-num"),
        email = $("#email"),
        school = $("#school"),
        schoolInfo = $("school-info");

    infoHead.mouseover(function(){
        infoHeadBox.show();
    });
    infoHead.mouseout(function(){
        infoHeadBox.hide();
    });
    infoHeadBox.mousemove(function(){
        infoHeadBox.show();
    });

    //个人信息和修改密码切换
    info.on("click", function(){
        personalInfo.show();
        updatePw.hide();
    });
    pw.on("click", function(){
        if(vimimg){
            var r = confirm("您还没有保存，是否保存？");
            if(r){
                vimimg = false;
                quitVim();
            }else{
                return;
            }
        }
        personalInfo.hide();
        updatePw.show();
    });

    var vimimg = false;
    //进入编辑模式
    intoVim.on("click", function(){
        vim();
    });
    //退出编辑模式
    cancel.on("click", function(){
        quitVim();
    })

    function vim(){
        vimimg = true;
        intoVim.hide();
        ensureInfo.show();
        cancel.show();
        username.css("background-color", "gray");
        name.css("background-color", "gray");
        stuNum.css("background-color", "gray");
        email.removeAttr("disabled");
        school.removeAttr("disabled");
        schoolInfo.removeAttr("disabled");
    }
    function quitVim(){
        vimimg = false;
        ensureInfo.hide();
        cancel.hide();
        intoVim.show();
        email.attr("disabled", "disabled");
        school.attr("disabled", "disabled");
        schoolInfo.attr("disabled", "disabled");
        username.css("background-color", "white");
        name.css("background-color", "white");
        stuNum.css("background-color", "white");
    }
})