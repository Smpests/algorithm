package com.qingsong.security;

import javax.crypto.spec.SecretKeySpec;
import javax.crypto.Mac;


class TustPayDemo {
    public static String genSign(String data, String key)  {

        try {
            //还原秘钥
            SecretKeySpec signingKey = new SecretKeySpec(key.getBytes(), "HmacSHA1");
            //实例化MAC
            Mac mac = Mac.getInstance("HmacSHA1");
            //初始化MAC
            mac.init(signingKey);
            //获得消息概要
            byte[] rawHmac = mac.doFinal(data.getBytes());
            //BASE64编码
            return java.util.Base64.getEncoder().encodeToString(rawHmac);
        } catch (Exception e){
            e.printStackTrace();
        }
        return null ;
    }
    
    public static void main(String[] args) {
        String secretKey = "U0X214AWH77PLB5SO9IG6FBM2MENZ0YMP83DNCQ9RW8GK4T0FCCVS4GWDMMQWERT";
        String unSignedString = "amount=1&email=dharmvirsingh.agmail.com@&merchantId=98b57889799740d18a9875226b8d1d26&name=srabankumar sabar&notifyUrl=https://api.yourdomain.com/test&orderNo=1284205570012600&phone=916371231700";
        System.out.print(genSign(unSignedString, secretKey));
    }
}