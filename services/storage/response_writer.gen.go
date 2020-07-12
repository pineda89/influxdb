// Generated by tmpl
// https://github.com/benbjohnson/tmpl
//
// DO NOT EDIT!
// Source: response_writer.gen.go.tmpl

package storage

import (
	"github.com/pineda89/influxdb/tsdb"
)

func (w *responseWriter) getFloatPointsFrame() *ReadResponse_Frame_FloatPoints {
	var res *ReadResponse_Frame_FloatPoints
	if len(w.buffer.Float) > 0 {
		i := len(w.buffer.Float) - 1
		res = w.buffer.Float[i]
		w.buffer.Float[i] = nil
		w.buffer.Float = w.buffer.Float[:i]
	} else {
		res = &ReadResponse_Frame_FloatPoints{&ReadResponse_FloatPointsFrame{Timestamps: make([]int64, 0, batchSize), Values: make([]float64, 0, batchSize)}}
	}

	return res
}

func (w *responseWriter) putFloatPointsFrame(f *ReadResponse_Frame_FloatPoints) {
	f.FloatPoints.Timestamps = f.FloatPoints.Timestamps[:0]
	f.FloatPoints.Values = f.FloatPoints.Values[:0]
	w.buffer.Float = append(w.buffer.Float, f)
}

func (w *responseWriter) streamFloatSeries(cur tsdb.FloatBatchCursor) {
	w.sf.DataType = DataTypeFloat
	ss := len(w.res.Frames) - 1
	ts, _ := cur.Next()
	if len(ts) == 0 {
		w.sz -= w.sf.Size()
		w.putSeriesFrame(w.res.Frames[ss].Data.(*ReadResponse_Frame_Series))
		w.res.Frames = w.res.Frames[:ss]
	} else if w.sz > writeSize {
		w.flushFrames()
	}
}

func (w *responseWriter) streamFloatPoints(cur tsdb.FloatBatchCursor) {
	w.sf.DataType = DataTypeFloat
	ss := len(w.res.Frames) - 1

	p := w.getFloatPointsFrame()
	frame := p.FloatPoints
	w.res.Frames = append(w.res.Frames, ReadResponse_Frame{p})

	var (
		seriesValueCount = 0
		b                = 0
	)

	for {
		ts, vs := cur.Next()
		if len(ts) == 0 {
			break
		}

		frame.Timestamps = append(frame.Timestamps, ts...)
		frame.Values = append(frame.Values, vs...)

		b = len(frame.Timestamps)
		if b >= batchSize {
			seriesValueCount += b
			b = 0
			w.sz += frame.Size()
			if w.sz >= writeSize {
				w.flushFrames()
				if w.err != nil {
					break
				}
			}

			p = w.getFloatPointsFrame()
			frame = p.FloatPoints
			w.res.Frames = append(w.res.Frames, ReadResponse_Frame{p})
		}
	}

	seriesValueCount += b
	w.vc += seriesValueCount
	if seriesValueCount == 0 {
		w.sz -= w.sf.Size()
		w.putSeriesFrame(w.res.Frames[ss].Data.(*ReadResponse_Frame_Series))
		w.res.Frames = w.res.Frames[:ss]
	} else if w.sz > writeSize {
		w.flushFrames()
	}
}

func (w *responseWriter) getIntegerPointsFrame() *ReadResponse_Frame_IntegerPoints {
	var res *ReadResponse_Frame_IntegerPoints
	if len(w.buffer.Integer) > 0 {
		i := len(w.buffer.Integer) - 1
		res = w.buffer.Integer[i]
		w.buffer.Integer[i] = nil
		w.buffer.Integer = w.buffer.Integer[:i]
	} else {
		res = &ReadResponse_Frame_IntegerPoints{&ReadResponse_IntegerPointsFrame{Timestamps: make([]int64, 0, batchSize), Values: make([]int64, 0, batchSize)}}
	}

	return res
}

func (w *responseWriter) putIntegerPointsFrame(f *ReadResponse_Frame_IntegerPoints) {
	f.IntegerPoints.Timestamps = f.IntegerPoints.Timestamps[:0]
	f.IntegerPoints.Values = f.IntegerPoints.Values[:0]
	w.buffer.Integer = append(w.buffer.Integer, f)
}

func (w *responseWriter) streamIntegerSeries(cur tsdb.IntegerBatchCursor) {
	w.sf.DataType = DataTypeInteger
	ss := len(w.res.Frames) - 1
	ts, _ := cur.Next()
	if len(ts) == 0 {
		w.sz -= w.sf.Size()
		w.putSeriesFrame(w.res.Frames[ss].Data.(*ReadResponse_Frame_Series))
		w.res.Frames = w.res.Frames[:ss]
	} else if w.sz > writeSize {
		w.flushFrames()
	}
}

func (w *responseWriter) streamIntegerPoints(cur tsdb.IntegerBatchCursor) {
	w.sf.DataType = DataTypeInteger
	ss := len(w.res.Frames) - 1

	p := w.getIntegerPointsFrame()
	frame := p.IntegerPoints
	w.res.Frames = append(w.res.Frames, ReadResponse_Frame{p})

	var (
		seriesValueCount = 0
		b                = 0
	)

	for {
		ts, vs := cur.Next()
		if len(ts) == 0 {
			break
		}

		frame.Timestamps = append(frame.Timestamps, ts...)
		frame.Values = append(frame.Values, vs...)

		b = len(frame.Timestamps)
		if b >= batchSize {
			seriesValueCount += b
			b = 0
			w.sz += frame.Size()
			if w.sz >= writeSize {
				w.flushFrames()
				if w.err != nil {
					break
				}
			}

			p = w.getIntegerPointsFrame()
			frame = p.IntegerPoints
			w.res.Frames = append(w.res.Frames, ReadResponse_Frame{p})
		}
	}

	seriesValueCount += b
	w.vc += seriesValueCount
	if seriesValueCount == 0 {
		w.sz -= w.sf.Size()
		w.putSeriesFrame(w.res.Frames[ss].Data.(*ReadResponse_Frame_Series))
		w.res.Frames = w.res.Frames[:ss]
	} else if w.sz > writeSize {
		w.flushFrames()
	}
}

func (w *responseWriter) getUnsignedPointsFrame() *ReadResponse_Frame_UnsignedPoints {
	var res *ReadResponse_Frame_UnsignedPoints
	if len(w.buffer.Unsigned) > 0 {
		i := len(w.buffer.Unsigned) - 1
		res = w.buffer.Unsigned[i]
		w.buffer.Unsigned[i] = nil
		w.buffer.Unsigned = w.buffer.Unsigned[:i]
	} else {
		res = &ReadResponse_Frame_UnsignedPoints{&ReadResponse_UnsignedPointsFrame{Timestamps: make([]int64, 0, batchSize), Values: make([]uint64, 0, batchSize)}}
	}

	return res
}

func (w *responseWriter) putUnsignedPointsFrame(f *ReadResponse_Frame_UnsignedPoints) {
	f.UnsignedPoints.Timestamps = f.UnsignedPoints.Timestamps[:0]
	f.UnsignedPoints.Values = f.UnsignedPoints.Values[:0]
	w.buffer.Unsigned = append(w.buffer.Unsigned, f)
}

func (w *responseWriter) streamUnsignedSeries(cur tsdb.UnsignedBatchCursor) {
	w.sf.DataType = DataTypeUnsigned
	ss := len(w.res.Frames) - 1
	ts, _ := cur.Next()
	if len(ts) == 0 {
		w.sz -= w.sf.Size()
		w.putSeriesFrame(w.res.Frames[ss].Data.(*ReadResponse_Frame_Series))
		w.res.Frames = w.res.Frames[:ss]
	} else if w.sz > writeSize {
		w.flushFrames()
	}
}

func (w *responseWriter) streamUnsignedPoints(cur tsdb.UnsignedBatchCursor) {
	w.sf.DataType = DataTypeUnsigned
	ss := len(w.res.Frames) - 1

	p := w.getUnsignedPointsFrame()
	frame := p.UnsignedPoints
	w.res.Frames = append(w.res.Frames, ReadResponse_Frame{p})

	var (
		seriesValueCount = 0
		b                = 0
	)

	for {
		ts, vs := cur.Next()
		if len(ts) == 0 {
			break
		}

		frame.Timestamps = append(frame.Timestamps, ts...)
		frame.Values = append(frame.Values, vs...)

		b = len(frame.Timestamps)
		if b >= batchSize {
			seriesValueCount += b
			b = 0
			w.sz += frame.Size()
			if w.sz >= writeSize {
				w.flushFrames()
				if w.err != nil {
					break
				}
			}

			p = w.getUnsignedPointsFrame()
			frame = p.UnsignedPoints
			w.res.Frames = append(w.res.Frames, ReadResponse_Frame{p})
		}
	}

	seriesValueCount += b
	w.vc += seriesValueCount
	if seriesValueCount == 0 {
		w.sz -= w.sf.Size()
		w.putSeriesFrame(w.res.Frames[ss].Data.(*ReadResponse_Frame_Series))
		w.res.Frames = w.res.Frames[:ss]
	} else if w.sz > writeSize {
		w.flushFrames()
	}
}

func (w *responseWriter) getStringPointsFrame() *ReadResponse_Frame_StringPoints {
	var res *ReadResponse_Frame_StringPoints
	if len(w.buffer.String) > 0 {
		i := len(w.buffer.String) - 1
		res = w.buffer.String[i]
		w.buffer.String[i] = nil
		w.buffer.String = w.buffer.String[:i]
	} else {
		res = &ReadResponse_Frame_StringPoints{&ReadResponse_StringPointsFrame{Timestamps: make([]int64, 0, batchSize), Values: make([]string, 0, batchSize)}}
	}

	return res
}

func (w *responseWriter) putStringPointsFrame(f *ReadResponse_Frame_StringPoints) {
	f.StringPoints.Timestamps = f.StringPoints.Timestamps[:0]
	f.StringPoints.Values = f.StringPoints.Values[:0]
	w.buffer.String = append(w.buffer.String, f)
}

func (w *responseWriter) streamStringSeries(cur tsdb.StringBatchCursor) {
	w.sf.DataType = DataTypeString
	ss := len(w.res.Frames) - 1
	ts, _ := cur.Next()
	if len(ts) == 0 {
		w.sz -= w.sf.Size()
		w.putSeriesFrame(w.res.Frames[ss].Data.(*ReadResponse_Frame_Series))
		w.res.Frames = w.res.Frames[:ss]
	} else if w.sz > writeSize {
		w.flushFrames()
	}
}

func (w *responseWriter) streamStringPoints(cur tsdb.StringBatchCursor) {
	w.sf.DataType = DataTypeString
	ss := len(w.res.Frames) - 1

	p := w.getStringPointsFrame()
	frame := p.StringPoints
	w.res.Frames = append(w.res.Frames, ReadResponse_Frame{p})

	var (
		seriesValueCount = 0
		b                = 0
	)

	for {
		ts, vs := cur.Next()
		if len(ts) == 0 {
			break
		}

		frame.Timestamps = append(frame.Timestamps, ts...)
		frame.Values = append(frame.Values, vs...)

		b = len(frame.Timestamps)
		if b >= batchSize {
			seriesValueCount += b
			b = 0
			w.sz += frame.Size()
			if w.sz >= writeSize {
				w.flushFrames()
				if w.err != nil {
					break
				}
			}

			p = w.getStringPointsFrame()
			frame = p.StringPoints
			w.res.Frames = append(w.res.Frames, ReadResponse_Frame{p})
		}
	}

	seriesValueCount += b
	w.vc += seriesValueCount
	if seriesValueCount == 0 {
		w.sz -= w.sf.Size()
		w.putSeriesFrame(w.res.Frames[ss].Data.(*ReadResponse_Frame_Series))
		w.res.Frames = w.res.Frames[:ss]
	} else if w.sz > writeSize {
		w.flushFrames()
	}
}

func (w *responseWriter) getBooleanPointsFrame() *ReadResponse_Frame_BooleanPoints {
	var res *ReadResponse_Frame_BooleanPoints
	if len(w.buffer.Boolean) > 0 {
		i := len(w.buffer.Boolean) - 1
		res = w.buffer.Boolean[i]
		w.buffer.Boolean[i] = nil
		w.buffer.Boolean = w.buffer.Boolean[:i]
	} else {
		res = &ReadResponse_Frame_BooleanPoints{&ReadResponse_BooleanPointsFrame{Timestamps: make([]int64, 0, batchSize), Values: make([]bool, 0, batchSize)}}
	}

	return res
}

func (w *responseWriter) putBooleanPointsFrame(f *ReadResponse_Frame_BooleanPoints) {
	f.BooleanPoints.Timestamps = f.BooleanPoints.Timestamps[:0]
	f.BooleanPoints.Values = f.BooleanPoints.Values[:0]
	w.buffer.Boolean = append(w.buffer.Boolean, f)
}

func (w *responseWriter) streamBooleanSeries(cur tsdb.BooleanBatchCursor) {
	w.sf.DataType = DataTypeBoolean
	ss := len(w.res.Frames) - 1
	ts, _ := cur.Next()
	if len(ts) == 0 {
		w.sz -= w.sf.Size()
		w.putSeriesFrame(w.res.Frames[ss].Data.(*ReadResponse_Frame_Series))
		w.res.Frames = w.res.Frames[:ss]
	} else if w.sz > writeSize {
		w.flushFrames()
	}
}

func (w *responseWriter) streamBooleanPoints(cur tsdb.BooleanBatchCursor) {
	w.sf.DataType = DataTypeBoolean
	ss := len(w.res.Frames) - 1

	p := w.getBooleanPointsFrame()
	frame := p.BooleanPoints
	w.res.Frames = append(w.res.Frames, ReadResponse_Frame{p})

	var (
		seriesValueCount = 0
		b                = 0
	)

	for {
		ts, vs := cur.Next()
		if len(ts) == 0 {
			break
		}

		frame.Timestamps = append(frame.Timestamps, ts...)
		frame.Values = append(frame.Values, vs...)

		b = len(frame.Timestamps)
		if b >= batchSize {
			seriesValueCount += b
			b = 0
			w.sz += frame.Size()
			if w.sz >= writeSize {
				w.flushFrames()
				if w.err != nil {
					break
				}
			}

			p = w.getBooleanPointsFrame()
			frame = p.BooleanPoints
			w.res.Frames = append(w.res.Frames, ReadResponse_Frame{p})
		}
	}

	seriesValueCount += b
	w.vc += seriesValueCount
	if seriesValueCount == 0 {
		w.sz -= w.sf.Size()
		w.putSeriesFrame(w.res.Frames[ss].Data.(*ReadResponse_Frame_Series))
		w.res.Frames = w.res.Frames[:ss]
	} else if w.sz > writeSize {
		w.flushFrames()
	}
}
