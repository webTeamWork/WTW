package com.bbs.utils;

import java.util.UUID;

/**
 * @author zhangyr
 * @version 1.0
 * 2021年3月16日下午9:22:35
 */
public final class Utils {
	public static String createUUID(){
		return UUID.randomUUID().toString().replace("-", "").toUpperCase();
	}

}
