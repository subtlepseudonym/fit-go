// Generated by reader_util_test.go.
// Only edit to bootstrap new entries or change existing entries.

package fit_test

var decodeTestFiles = [...]struct {
	folder      string
	name        string
	wantErr     bool
	fingerprint uint64
	compress    bool
	dopts       testingDecodeOpts
	skipEncode  bool
}{
	{
		"me",
		"activity-small-fenix2-run.fit",
		false,
		11297915186217814700,
		true,
		tdoAllWithDiscardLogger,
		true, // Decode mismatch due to unknown fields
	},
	{
		"fitsdk",
		"Activity.fit",
		false,
		7755175530838408024,
		true,
		tdoNone,
		false,
	},
	{
		"fitsdk",
		"MonitoringFile.fit",
		false,
		17438554766263628503,
		true,
		tdoNone,
		true, // Fails because first message has different valid fields (#36)
	},
	{
		"fitsdk",
		"Settings.fit",
		false,
		13777716833885766673,
		true,
		tdoNone,
		false,
	},
	{
		"fitsdk",
		"WeightScaleMultiUser.fit",
		false,
		5002902390457186100,
		true,
		tdoNone,
		false,
	},
	{
		"fitsdk",
		"WorkoutCustomTargetValues.fit",
		false,
		1795447305759253807,
		true,
		tdoNone,
		false,
	},
	{
		"fitsdk",
		"WorkoutIndividualSteps.fit",
		false,
		13432157117047365566,
		true,
		tdoNone,
		false,
	},
	{
		"fitsdk",
		"WorkoutRepeatGreaterThanStep.fit",
		false,
		10036597401811670605,
		true,
		tdoNone,
		false,
	},
	{
		"fitsdk",
		"WorkoutRepeatSteps.fit",
		false,
		2000919626000542509,
		true,
		tdoNone,
		false,
	},
	{
		"fitsdk",
		"WeightScaleSingleUser.fit",
		false,
		170466418772016754,
		true,
		tdoNone,
		false,
	},
	{
		"python-fitparse",
		"garmin-edge-500-activitiy.fit",
		false,
		17565016518975205314,
		true,
		tdoNone,
		false,
	},
	{
		"python-fitparse",
		"sample-activity-indoor-trainer.fit",
		false,
		6384198383276191715,
		true,
		tdoNone,
		true, // Fails because first message has different valid fields (#36)
	},
	{
		"python-fitparse",
		"compressed-speed-distance.fit",
		false,
		0,
		false,
		tdoNone,
		false,
	},
	{
		"python-fitparse",
		"antfs-dump.63.fit",
		false,
		14798701571659473253,
		true,
		tdoNone,
		true, // Fails because first message has different valid fields (#36)
	},
	{
		"sram",
		"Settings.fit",
		false,
		15334527618026463003,
		true,
		tdoNone,
		true, // Fails due to encoder using profile length for arrays (#37)
	},
	{
		"sram",
		"Settings2.fit",
		false,
		3203160144706376883,
		true,
		tdoNone,
		true, // Fails due to encoder using profile length for arrays (#37)
	},
	{
		"dcrainmaker",
		"Edge810-Vector-2013-08-16-15-35-10.fit",
		false,
		16296406423933334795,
		true,
		tdoNone,
		true, // Fails because first message has different valid fields (#36)
	},
	{
		"misc",
		"2013-02-06-12-11-14.fit",
		false,
		6015305551413305338,
		true,
		tdoNone,
		false,
	},
	{
		"misc",
		"2015-10-13-08-43-15.fit",
		false,
		3166312160825042157,
		true,
		tdoNone,
		false,
	},
	{
		"corrupt",
		"activity-filecrc.fit",
		true,
		6894792788097918620,
		true,
		tdoNone,
		false,
	},
	{
		"corrupt",
		"activity-unexpected-eof.fit",
		true,
		6859919891971491875,
		true,
		tdoNone,
		false,
	},
}
