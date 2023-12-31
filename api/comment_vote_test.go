package main

import (
	"testing"
	"time"
)

func TestCommentVoteBasics(t *testing.T) {
	failTestOnError(t, setupTestEnv())

	cr0, _ := commenterNew("test1@example.com", "Test1", "undefined", "http://example.com/photo.jpg", "google", "", "undefined")
	cr1, _ := commenterNew("test2@example.com", "Test2", "undefined", "http://example.com/photo.jpg", "google", "", "undefined")
	cr2, _ := commenterNew("test3@example.com", "Test3", "undefined", "http://example.com/photo.jpg", "google", "", "undefined")

	c0, _ := commentNew(cr0, "example.com", "/path.html", "root", "**foo**", "approved", time.Now().UTC())

	if err := commentVote(cr0, c0, 1, "example.com/path.html"); err != errorSelfVote {
		t.Errorf("expected err=errorSelfVote got err=%v", err)
		return
	}

	if c, _, _ := commentList("temp", "example.com", "/path.html", false); c[0].Score != 0 {
		t.Errorf("expected c[0].Score = 0 got c[0].Score = %d", c[0].Score)
		return
	}

	if err := commentVote(cr1, c0, -1, "example.com/path.html"); err != nil {
		t.Errorf("unexpected error voting: %v", err)
		return
	}

	if err := commentVote(cr2, c0, -1, "example.com/path.html"); err != nil {
		t.Errorf("unexpected error voting: %v", err)
		return
	}

	if c, _, _ := commentList("temp", "example.com", "/path.html", false); c[0].Score != -2 {
		t.Errorf("expected c[0].Score = -2 got c[0].Score = %d", c[0].Score)
		return
	}

	if err := commentVote(cr1, c0, -1, "example.com/path.html"); err != nil {
		t.Errorf("unexpected error voting: %v", err)
		return
	}

	if c, _, _ := commentList("temp", "example.com", "/path.html", false); c[0].Score != -2 {
		t.Errorf("expected c[0].Score = -2 got c[0].Score = %d", c[0].Score)
		return
	}

	if err := commentVote(cr1, c0, 0, "example.com/path.html"); err != nil {
		t.Errorf("unexpected error voting: %v", err)
		return
	}

	if c, _, _ := commentList("temp", "example.com", "/path.html", false); c[0].Score != -1 {
		t.Errorf("expected c[0].Score = -1 got c[0].Score = %d", c[0].Score)
		return
	}
}
