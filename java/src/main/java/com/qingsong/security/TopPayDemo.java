package com.qingsong.security;


import com.google.gson.Gson;
import com.google.gson.JsonObject;

import java.io.IOException;
import java.security.NoSuchAlgorithmException;
import java.security.spec.InvalidKeySpecException;
import java.util.Map;
import java.util.TreeMap;

public class TopPayDemo {
  // TopPay代付接入文档地址：https://id-doc.toppay.asia/?java#toppay-api-reference
  // 测试账号
  private static final String MCH_ID = "S820220524155940000001";  // 此乃测试商户号,请使用自己申请的真实商户号和对应的密钥
  private static final String PLAT_PUBLIC_KEY = "MIICdwIBADANBgkqhkiG9w0BAQEFAASCAmEwggJdAgEAAoGBALXZNUsppu7/QpsLWi9rj2gu+YEdoy4rs2sUT4A9WyK5ecXexhpE9xqwQRJq+Koft+njzYEBRLXf3iFjV1UUsJ82UCcnMvcyXXm7+valdLY2deVgW+RBQPuea9ffiYkRBuYhNXP1B14EkssbuU46O4Vje2pYz1tok/VDEe3/mJuZAgMBAAECgYB137r04qXri5iwCZbG4DDP9osSj5SDzB37zemOlbbqdlGh2jqNMqD+MrurZZ2UIQVE3b4K3jicr+U7vToR01x8b3HFjKCwQ9JBGy+DR2TfpaQnBXKsBNBEqnSPNfaOvGfjU+NeyRD1LdZ+DyZAO0/f4HUwgGVvi2D9XQfNX6ciJQJBAODsJgFvxUP2QdCpa8o5dXBUNOfk7AL2/zuRX6tJT1PoBPO0gKJN5OCo9BSBug5hz3nI7Cdv4kQf4aQQYtka3JMCQQDO+XeHv1hNnUlBBpj1UvdJfg+zCRKa6QuP1O73PJyTN26aPiI4/PD6n60D+R7nU1rUnfrcHs5g4Am51v1tzW6jAkBF08pww7xgPwFrwE1AeNomx12taHqE6R2P5q/mw1mnWO7AohcVTE10Gb+85hFXCbKZwnDeRXPfZm5sutGVSSuVAkEAhQro4dKGoYV8RBdIlVGacZxd3Rj046zv6KnYXgaRIzSbnPxZpUY6ZO/lGEu3qLz1Avz19y66PJ1mwxXGXCGUfwJBAMf1eOWmBjRXQKLlGdU3yZRH/YE6i/olgby9s1buVNPgMtYD8+Rz+hwXXWyeUPcyWuQMkY1zbVB/2gy2cV1v0NQ=";  // 商户私钥
  private static final String MCH_PUBLIC_KEY = "MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCKN/o0dEiVG3g6dwFExrjYfyNLXVE5+B7qSDndYr8yP2rM2leGNS1QbNa/KorQzHTXs4DSRyJUttZLCknPxh5RXF9CsFHj6JLOvLeNfMAnpyFvIkQBaE0sA2XsEHs5rCBj0bmWj7DKE/kCuAul/x/nS1Of468eL9qRT55U+AnbHwIDAQAB";
  private static final String payUrl = "https://id-openapi.toppay.asia/gateway/prepaidOrder";
  private static final String cashUrl = "https://id-openapi.toppay.asia/gateway/cash";
  private static final String payNotify = "your notify url";
  private static final String cashNotify = "your notify url";

  public static void main(String[] args) throws Exception {
    // 代收
    pay();
  }
  private static void pay() throws Exception {
    Map<String, String> requestParams = new TreeMap<>();
    requestParams.put("merchantCode", MCH_ID);
    // requestParams.put("orderType", "0"); // 订单类型（0-法币交易）
    // requestParams.put("method", "BCA"); // 支付渠道
    requestParams.put("orderNum", "123"); // 商户订单号
    requestParams.put("payMoney", "150");  // 订单金额
    requestParams.put("name", "JackMa");// 客户名称
    requestParams.put("phone", "082122965511"); // 手机号
    requestParams.put("email", "test@gmail.com");// 客户邮箱
    requestParams.put("notifyUrl", "http://metapi.paofu6.com/api/pay/recharge/callback?tran_id=123");// 回调地址
    requestParams.put("dateTime", "20230524235959");// 时间戳 格式 yyyyMMddHHmmss
    requestParams.put("expiryPeriod", "1440"); // 过期时间 单位(分)
    requestParams.put("productDetail", "Test Pay");// 订单详情
    requestParams.put("redirectUrl", "http://st.paofu6.com/static/html/pay/pay_form_submit.html");// 返回跳转地址（仅限收银台使用）

    StringBuilder stringBuilder = new StringBuilder();
    for (String key : requestParams.keySet()) {
      stringBuilder.append(requestParams.get(key));  // 拼接参数
    }

    String keyStr = stringBuilder.toString();  // 得到待加密的字符串
    System.out.println("keyStr:" + keyStr);
    String signedStr = TopPayRequestUtil.privateEncrypt(keyStr, TopPayRequestUtil.getPrivateKey(MCH_PRIVATE_KEY));  // 私钥加密
    requestParams.put("sign", signedStr);
    // requestParams.put("platSign", "IuWZTO1itOYYgFBMkcTH6RUfa4QXeZrXFouqtoCh/JNMITWQbPqpogkzao0yMeE74/h09qGEHZLgMe307oDF4PHLDeTjUoC3dk2IL4xTqqhpM9Dy2tGS5TorKXB537fsKr7w4ZRERXS9/EX93a5AgZZpikbdoqPFvTj1GYIp5ul7MEbZnEPCr8BKGXkg/H/P+yvIBBwJwuRGtYNEf0vzA6haC0ib9maiuirRBoQchWn9VOT/h1n5VruZkBUp2hZy7CuaHpfyzW7lApoDDQF5un7X63wzPP7RdevTls2rIWLRn4NNwGr9pP+Ni7vJPKG+divfOkMGJL0zvQPl2djaIw==");

    String postJson = new Gson().toJson(requestParams);
    System.out.println("Post Json Params:" + postJson);


    String responseJson = TopPayRequestUtil.doPost(payUrl, postJson);  // 发送 post json请求
    System.out.println("Response Msg:" + responseJson);
    // verify sign
    // boolean isValid = TopPayRequestUtil.verifySign(new Gson().fromJson(postJson, JsonObject.class), MCH_PUBLIC_KEY);
    // System.out.println("isValid: " + isValid);
  }
}
