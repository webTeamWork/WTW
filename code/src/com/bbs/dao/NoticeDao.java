package com.bbs.dao;

import java.util.List;

import com.bbs.model.Notice;
/**
 * 
 * @author zhangyr
 * @version 1.0
 * 2021年3月22日上午9:52:33
 */
public interface NoticeDao {

	public  List<Notice> getNotice(int pageIndex, int pageSize);

	public  void publish(Notice notice);

	public Notice getNoticeById(int noticeId);

}