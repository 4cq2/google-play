package acquire

import (
   "154.pages.dev/encoding/protobuf"
   "bytes"
   "errors"
   "io"
   "net/http"
   "net/url"
)

// pass
//const Device_ID = "306e9f7f4192be79"

const Device_ID = "3419ddd967f9597f"

func new_checkin() (*checkin, error) {
   req := new(http.Request)
   req.Header = make(http.Header)
   req.Header["Connection"] = []string{"Keep-Alive"}
   req.Header["Content-Type"] = []string{"application/x-protobuffer"}
   req.Header["Host"] = []string{"android.googleapis.com"}
   req.Header["User-Agent"] = []string{"Dalvik/2.1.0 (Linux; U; Android 5.0.2; Android SDK built for x86 Build/LSY66K)"}
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "android.googleapis.com"
   req.URL.Path = "/checkin"
   req.URL.Scheme = "https"
   req.Body = io.NopCloser(bytes.NewReader(checkin_body.Append(nil)))
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   var check checkin
   {
      b, err := io.ReadAll(res.Body)
      if err != nil {
         return nil, err
      }
      check.m, err = protobuf.Consume(b)
      if err != nil {
         return nil, err
      }
   }
   return &check, nil
}

type checkin struct {
   m protobuf.Message
}

var checkin_body = protobuf.Message{
   protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
   protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("1-da39a3ee5e6b4b0d3255bfef95601890afd80709")},
   protobuf.Field{Number: 4, Type: 2, Value: protobuf.Prefix{
      protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("generic_x86/sdk_google_phone_x86/generic_x86:5.0.2/LSY66K/6695550:eng/test-keys")},
         protobuf.Field{Number: 2, Type: 2, Value: protobuf.Bytes("ranchu")},
         protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("generic_x86")},
         protobuf.Field{Number: 5, Type: 2, Value: protobuf.Bytes("unknown")},
         protobuf.Field{Number: 6, Type: 2, Value: protobuf.Bytes("android-google")},
         protobuf.Field{Number: 7, Type: 0, Value: protobuf.Varint(1595298807)},
         protobuf.Field{Number: 8, Type: 0, Value: protobuf.Varint(202414013)},
         protobuf.Field{Number: 9, Type: 2, Value: protobuf.Bytes("generic_x86")},
         protobuf.Field{Number: 10, Type: 0, Value: protobuf.Varint(21)},
         protobuf.Field{Number: 11, Type: 2, Value: protobuf.Bytes("Android SDK built for x86")},
         protobuf.Field{Number: 12, Type: 2, Value: protobuf.Bytes("unknown")},
         protobuf.Field{Number: 13, Type: 2, Value: protobuf.Bytes("sdk_google_phone_x86")},
         protobuf.Field{Number: 14, Type: 0, Value: protobuf.Varint(0)},
         protobuf.Field{Number: 15, Type: 2, Value: protobuf.Prefix{
            protobuf.Field{Number: 1, Type: 0, Value: protobuf.Varint(1)},
            protobuf.Field{Number: 2, Type: 2, Value: protobuf.Bytes("android-google")},
         }},
      }},
      protobuf.Field{Number: 6, Type: 2, Value: protobuf.Bytes("310260")},
      protobuf.Field{Number: 7, Type: 2, Value: protobuf.Bytes("310260")},
      protobuf.Field{Number: 8, Type: 2, Value: protobuf.Bytes("MOBILE:LTE:")},
      protobuf.Field{Number: 9, Type: 0, Value: protobuf.Varint(0)},
      protobuf.Field{Number: 14, Type: 0, Value: protobuf.Varint(2)},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 0, Value: protobuf.Varint(5)},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(1)},
         protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("unspecified")},
         protobuf.Field{Number: 4, Type: 2, Value: protobuf.Bytes("")},
         protobuf.Field{Number: 5, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 16, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("310260")},
         protobuf.Field{Number: 2, Type: 2, Value: protobuf.Bytes("Android")},
         protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("0")},
         protobuf.Field{Number: 4, Type: 0, Value: protobuf.Varint(0)},
         protobuf.Field{Number: 4, Type: 0, Value: protobuf.Varint(1)},
         protobuf.Field{Number: 4, Type: 0, Value: protobuf.Varint(2)},
         protobuf.Field{Number: 6, Type: 2, Value: protobuf.Bytes("310260000000000")},
         protobuf.Field{Number: 8, Type: 2, Value: protobuf.Bytes("\x17L")},
      }},
      protobuf.Field{Number: 19, Type: 2, Value: protobuf.Bytes("MOBILE")},
   }},
   protobuf.Field{Number: 6, Type: 2, Value: protobuf.Bytes("en-US")},
   protobuf.Field{Number: 7, Type: 0, Value: protobuf.Varint(13981217741684452058)},
   protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("358240051111110")},
   protobuf.Field{Number: 11, Type: 2, Value: protobuf.Bytes("")},
   protobuf.Field{Number: 12, Type: 2, Value: protobuf.Bytes("America/Denver")},
   protobuf.Field{Number: 14, Type: 0, Value: protobuf.Varint(3)},
   protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("71Q6Rn2DDZl1zPDVaaeEHItd+Yg=")},
   protobuf.Field{Number: 16, Type: 2, Value: protobuf.Bytes("EMULATOR32X1X14X0")},
   protobuf.Field{Number: 18, Type: 2, Value: protobuf.Prefix{
      protobuf.Field{Number: 1, Type: 0, Value: protobuf.Varint(3)},
      protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(1)},
      protobuf.Field{Number: 3, Type: 0, Value: protobuf.Varint(2)},
      protobuf.Field{Number: 4, Type: 0, Value: protobuf.Varint(2)},
      protobuf.Field{Number: 5, Type: 0, Value: protobuf.Varint(0)},
      protobuf.Field{Number: 6, Type: 0, Value: protobuf.Varint(1)},
      protobuf.Field{Number: 7, Type: 0, Value: protobuf.Varint(400)},
      protobuf.Field{Number: 8, Type: 0, Value: protobuf.Varint(196609)},
      protobuf.Field{Number: 9, Type: 2, Value: protobuf.Bytes("android.test.runner")},
      protobuf.Field{Number: 9, Type: 2, Value: protobuf.Bytes("com.android.future.usb.accessory")},
      protobuf.Field{Number: 9, Type: 2, Value: protobuf.Bytes("com.android.location.provider")},
      protobuf.Field{Number: 9, Type: 2, Value: protobuf.Bytes("com.android.media.remotedisplay")},
      protobuf.Field{Number: 9, Type: 2, Value: protobuf.Bytes("com.android.mediadrm.signer")},
      protobuf.Field{Number: 9, Type: 2, Value: protobuf.Bytes("com.google.android.gms")},
      protobuf.Field{Number: 9, Type: 2, Value: protobuf.Bytes("com.google.android.maps")},
      protobuf.Field{Number: 9, Type: 2, Value: protobuf.Bytes("com.google.android.media.effects")},
      protobuf.Field{Number: 9, Type: 2, Value: protobuf.Bytes("javax.obex")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.audio.output")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.bluetooth")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.camera")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.camera.any")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.camera.autofocus")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.faketouch")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.location")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.location.network")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.microphone")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.screen.landscape")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.screen.portrait")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.sensor.accelerometer")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.sensor.compass")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.touchscreen")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.touchscreen.multitouch")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.touchscreen.multitouch.distinct")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.touchscreen.multitouch.jazzhand")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.usb.accessory")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.software.app_widgets")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.software.backup")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.software.connectionservice")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.software.device_admin")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.software.home_screen")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.software.input_methods")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.software.live_wallpaper")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.software.managed_users")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.software.print")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.software.voice_recognizers")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.software.webview")},
      protobuf.Field{Number: 11, Type: 2, Value: protobuf.Bytes("x86")},
      protobuf.Field{Number: 12, Type: 0, Value: protobuf.Varint(1080)},
      protobuf.Field{Number: 13, Type: 0, Value: protobuf.Varint(2160)},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("af")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("am")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("ar")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("ar-EG")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("ar-IL")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("as")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("az")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("be")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("bg")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("bg-BG")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("bn")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("bs")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("ca")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("ca-ES")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("cs")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("cs-CZ")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("da")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("da-DK")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("de")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("de-AT")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("de-CH")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("de-DE")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("de-LI")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("el")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("el-GR")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("en")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("en-AU")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("en-CA")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("en-GB")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("en-IE")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("en-IN")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("en-NZ")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("en-SG")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("en-US")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("en-ZA")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("es")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("es-ES")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("es-US")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("et")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("eu")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("fa")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("fi")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("fi-FI")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("fil")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("fil-PH")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("fr")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("fr-BE")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("fr-CA")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("fr-CH")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("fr-FR")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("gl")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("gu")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("hi")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("hi-IN")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("hr")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("hr-HR")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("hu")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("hu-HU")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("hy")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("id")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("in")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("is")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("it")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("it-CH")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("it-IT")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("iw")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("ja")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("ja-JP")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("ka")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("kk")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("km")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("kn")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("ko")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("ko-KR")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("ky")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("lo")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("lt")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("lt-LT")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("lv")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("lv-LV")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("mk")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("ml")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("mn")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("mr")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("ms")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("my")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("nb")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("nb-NO")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("ne")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("nl")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("nl-BE")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("nl-NL")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("or")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("pa")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("pl")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("pl-PL")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("pt")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("pt-BR")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("pt-PT")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("ro")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("ro-RO")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("ru")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("ru-RU")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("si")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("sk")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("sk-SK")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("sl")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("sl-SI")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("sq")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("sr")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("sr-Latn")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("sr-RS")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("sv")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("sv-SE")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("sw")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("ta")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("te")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("th")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("th-TH")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("tr")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("tr-TR")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("uk")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("uk-UA")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("ur")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("uz")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("vi")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("vi-VN")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("zh-CN")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("zh-HK")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("zh-TW")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("zu")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("ANDROID_EMU_CHECKSUM_HELPER_v1")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("ANDROID_EMU_dma_v1")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("ANDROID_EMU_gles_max_version_3_1")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("ANDROID_EMU_host_side_tracing")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("ANDROID_EMU_sync_buffer_data")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_APPLE_texture_format_BGRA8888")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_EXT_color_buffer_float")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_EXT_color_buffer_half_float")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_EXT_debug_marker")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_EXT_texture_format_BGRA8888")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_KHR_texture_compression_astc_ldr")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_EGL_image")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_EGL_image_external")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_EGL_image_external_essl3")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_EGL_sync")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_blend_equation_separate")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_blend_func_separate")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_blend_subtract")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_byte_coordinates")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_compressed_ETC1_RGB8_texture")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_compressed_paletted_texture")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_depth24")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_depth32")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_depth_texture")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_draw_texture")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_element_index_uint")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_fbo_render_mipmap")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_framebuffer_object")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_packed_depth_stencil")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_point_size_array")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_point_sprite")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_rgb8_rgba8")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_single_precision")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_stencil1")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_stencil4")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_stencil8")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_stencil_wrap")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_texture_cube_map")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_texture_env_crossbar")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_texture_float")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_texture_float_linear")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_texture_half_float")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_texture_half_float_linear")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_texture_mirored_repeat")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_texture_npot")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_vertex_array_object")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_vertex_half_float")},
      protobuf.Field{Number: 18, Type: 0, Value: protobuf.Varint(432)},
      protobuf.Field{Number: 19, Type: 0, Value: protobuf.Varint(0)},
      protobuf.Field{Number: 20, Type: 0, Value: protobuf.Varint(1588760576)},
      protobuf.Field{Number: 21, Type: 0, Value: protobuf.Varint(4)},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.audio.output")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.bluetooth")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.camera")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.camera.any")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.camera.autofocus")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.faketouch")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.location")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.location.network")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.microphone")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.screen.landscape")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.screen.portrait")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.sensor.accelerometer")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.sensor.compass")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.touchscreen")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.touchscreen.multitouch")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.touchscreen.multitouch.distinct")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.touchscreen.multitouch.jazzhand")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.usb.accessory")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.app_widgets")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.backup")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.connectionservice")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.device_admin")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.home_screen")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.input_methods")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.live_wallpaper")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.managed_users")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.print")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.voice_recognizers")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.webview")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 28, Type: 0, Value: protobuf.Varint(0)},
      protobuf.Field{Number: 30, Type: 0, Value: protobuf.Varint(2)},
   }},
   protobuf.Field{Number: 20, Type: 0, Value: protobuf.Varint(0)},
   protobuf.Field{Number: 22, Type: 0, Value: protobuf.Varint(0)},
   protobuf.Field{Number: 24, Type: 2, Value: protobuf.Bytes("CgYs_2lbIQzSEFsAAHdU9h5a1xWSAFpFJRNTS5M9AC4R4ZGh9VMLBADDJ63W0SEMAFD2hUYLAaNSAQ4qN5jJF3GjAf8ZjNa77UdCBSQ_vtIfIwS5A5XXI_rK0Z7tAhh65LU_Oxc62geZBgDroQgSVt_FsDleGZ0PHu0GvTUVpf2yW82yjZeG9mOJQGuEDfiMyzNjYeHMFgi3ONTiEVjXExcZwFnPg4_pSsED242TA91vL9pAoTENPY1N89W8A1cU7PX3onXocGSOjNH1e19ajZ-AKsuAXbRIlg6pnIy-_aEM2CgpW-rVYS5NNq9EtSbc0MmVUhGUwjJjseXecQMgr9rkwdHVg7mgAIHcdz57kV4cxM7oGcopbgGrMQWoFG-g7K0ZCxNiN1CofGQbp2utF98zKn2gMblwXxnjyFoibEfFXOqjFqFdS2olFN1Nt-KRlEd5Y5c10u39gvRgAhZM6sxt2eLCrCNw86Kgpj5IuuoLG30Lby57VjjwxnGkvs0_Vp8EnOM3LSSLWGsgIumBlfuNg5_rCh9dZ5XJH5-WtRj7wnZaQSrXYiVR3VEBVOSW10U1wnmGF5ZYqXU3okiE3y82GXlRY7_2u8sS8Ay-GJ9RqE9iJiKgK5vhK0IVenBuMf_-U65ve9nWEGch5pqBieFugOI781jNh8MhJv1Tx_C-vZO6elq_3z4RvPcZ9FC3Kv7lBLqMVQzC1wdwmpVD8JAWUXcpu9laUZ4N7hmfw8KSHABB4xWCzgNZRWcE15cN-9aRxNnsHSB6Qja06KwF2KBJETjx6bqIaW3WUDkwqe7m8reFemQvISgDklZm04gE2wpN2UdJQY9JavFW4ohqsyFJ9jtZP310Wyynuv3izSEUss-8l-ZgUsc-8_rbfrvFhTfBLcEhgF6T-pGi8_5JQWaMMhkDPFGae6TfT9wLHsXRyAiC_jT0ciKSGr_CA5wRCYewyTOmC7wCdsvQpaUVavDllZFq6NRnz1G032igUZf2bNO3MbA9Bxi3ADS6yCv65-RHKaPYrbCqotn7mXWTftH0AucTPPclOaBHQeot9tLe05gWZtFoBNVw7Hb8haZSWGx2SlD--tyoDp0GGuR1QuZN8GdzGh96Vpy_mjwPrEo1gWQmKEAuIEqppkCnsuv41wMXr4YMStY7jiplAu-LABGjcatIADaJKaWODy8tiUICmLg4YdAPoaiywrrnzfboAThrOEuYEuKV2Zem4baFlAGYEvr6krLUg9auGRKPNPmvM6V3ySYHP82p6nxd2BWsV7v0l3EYnxHTwhuuIJdxnBubsitNX18pEUuhF7L2IUJLmVXWTi3DEEim4H00f38C3CXEq7J8tew9EQy384G9hW1_Rlhhi5PLduhftQnp4kMpYi9vPfJfvTmLaHr5aDTNRwl-hzgM983fnTJqxbJRnpFK8EgpsCzghhUm4w5qqWHuJaxsyqBeiJGBeeilxcpOYSzZrqmPrN9AqRphegCcgFFQQVrwK7RhinuRvjECUIjufBuYon4ModTgr2VHTuXo8afGZUCL2P-GQ3h9_uloGoEhBi6uIo_NmHlyHBWfDzDparO311J6EtQJWAWK3ePsxJo705pnDt8Y3R8SwjyexahtGiFIPBrXM_9oLmpTkEB2qWzwCWt3tTxa6FxG6dAHru1JNAvQgxinp44_yamI6zcdLzpBijUXkXY6MRtPa3Qp5WyHvEvxzUlZI-ja69OCrxiPfA0r6KLhWea8A7ZGxVXtj-CSserukr5mlL43yP_eosQ2NRzIwrkT-3P43l6pX_yALTGXt69wsctVxum97iPBY1Yzz6N4pSNF3v4L2e4OFMGOVsxlc1fcFX_OadYSk80nA6PJYHRlCZqqO20XIRmMKXXvnOKz-xfRTWMxFeOKv9ngWyRdJOHiaKt54vhHbygXxYtreZunnbYkZ6k1Xx_MkGdl6-p6d3u-lS-7kIjAmw3NYmZuZfXpNZEqet--Rze-d6Z_ra5zhcL169PUC2748jLHhtZWuBTK8MjwWvU9V1TgJ_-ApJABqSynM1gi4U7pvEtj7seivn0rJLUIGdy_BW5sB8-d6Kjq1uvNdCMM6IgsmXXSeJkuq225Khmui5p2ZrxB3H9cgiA-elk7jw5J9qxJhI_I0taJsIHDXskzqJBnkS1XwWDuqepbY6KSwyZiFOgzVlp3o0RuhPQtrY8Gb-Vvkk7D5sAA15JTvwzC6hPM0ZSmg2g9Uc50jTtMByrxdgaWC0h3msbyUQd87aPlXQCyEzuWCcVynGL7QjneolMnNvp5JhMEaYGsffaoGooHlRZAWtNV8PQmT3r9qug1i7KJ0C4I374I1kKRwDf-vbA70kb9b65X7QvRTdlMUOK8aJS4n4o3hF8nKnCT_RzJeD6GiQxQQw7z2Ss4OSCQC4pEZ1lpLPSqSo5od0gLO3oT2orIVT24LPH08g87DzB8OWL3jaqaSWKSpOVQCmCfvBbrIDM7BLtEbl_raZaTaDrbn6GqtTYV8rXCKAUgfypLn20dewpJWn326-wCDE2lsF3-NNm2V2xA7tVdS8lSPx_sYkPvdUL068GaVHQ8VXYkI5U8lPnISPL1xqUtulpXnvm3bRF-OKIoC9iKpTNQbYr2t0wYCyUoZjq2v6hUHPO2o90bKfe9GqlUZ87eLK6CDJs7e014wKPetlvUKfQ1T91SrJFFGoN7A0h6FoE7zbBVJd7GxVPEtxUNZRfDFoqMxEwZvwmqNl5VLKLtascmVFVPzu3m4dvC__RwxCVJ8qvj21-7BXgrK0gTtsyf3vOhvwcYA81u3BDq0opFUgZ6rMnbpD8xZVwbweOxKxHKqSKGZNRI9oKEYCQsEKxBjuY7MavRYrFKgW3dgIz1Q6VqhPXhmYWkg3JlpTbgunnwPQCmBP6cDbmkJ9uAN1B1LUFEyGgcxKa5AWnocBIkyRZTzeY1tAr0XG_KDYJ99Jmo-CJCQVnuNTY-XYFtb9GPT26J434GUYuel40I9plz255q9Bk1Sj8PyoK1j2P6eG8oXyDNHETNZf0DeGns2wSRqVHLdYiE41W21pUzQv0-QQ6Y3TQltSKknGHBCdcFi5-1bEdUMyl79V04FvbBaU2jjB3QPWM7TBIBtKSZ2hmsgq9UK-CxCPyYwCZI6gjmqgAGzwc2E8KCvf5-4WzVPIShcDQgIiMsSSboaiv7YenTks16dkgO6cSB5P4Tj5rO6tp7h-3EbBkF6kG27E6JvXhHsPuEsztybxDci2Dwm67o2dH3ZWHYovgV6SENDEKdaRYa57KFBAbS6CLxtD6u90s-KHe0wUsZ2TMty20GrKjZ3qJY0j1BJFhMAw-HalRRlhXdxH96AGLB1OVXt7NwyeXcJwNVqWAY8FefAXvYPAbS1hvi2Lv6crzAHacpWdjH4qx4p71hyPqYOMedZghj78e-g-hSi2HvV2-4W9u8ZScSBn0CtvMTcCsGtC0qagEqxtRc7r-hKKTDrdAkmY1C6Dgz3WE_W33fON9WzIU255zwJ-9CQwkexGgVOoI0jj-NnsljeAxgSS0gzqEHk1FAplg8w4X5vn5dA9eyQav2tnJ_Fhq-T-2HWiutSKNy0_H0HS_h5dRQHGYZnhW_KWsKVgWlpF5GeYqi_AhsaYOR4ZpvMWwnyjIcZnPtuM5BrbCa52x8HR6GHwPnMaj9lcJC2yEKs-cgGRH1SGhYQCOM4Y-0Em41NCj-PY0Wvo2h09_9KwKM81Nj5XdBNl7B4BhjcMAqyq2X-wI2BKJpmPUpiqOadlAPhBYiK6nefARBYuZIbZUUNPleP_0AYvegZOTND01qO-lghD0uFq6lANiR4OEYfIIYetJmDqoGSJ1GLeiOoc5FORvWIphJXuFumFnic5hXGsPVa-jnS55vg-PjSUhj2Tmb_oM9FPO11l5r7NqqUcLDtLwn-5D7EiBFAa4jjrZ7gHT-sICDCQQxu1pf7bJwA6l12ej3hScinOQWQh5lcUVVrJIQDqmQ3C_AvIWsqnB8ElEYPX4Ece2LyFWTah6InrB2D4UDAguW6ufWmkNS0rZudrklQIoUENBO0rjDYsdmwys7wWOLO-u9kEuSRYqz_HBxlGXlQP9q-OvimkhO_eQ0ijKSiJmUx0nN0uHf1nGFijTffjkK-7kfiN3pwT7CgGXhFyTCgfaVGTK7i5TLA88OiIEbtnjYVucGp0L6UAO1tqj2WURKZdifmnRjLufLtdo-aMyLWBcSiaxZhZqU6tdyg6mHs3jSQyFa2skQzABT9-LBTFag8j1_QYwfyU70AbuhxXBe2PnS3arJblWf5L6aJmNFcU-Zgqjg6KETqIPVmX5fXxI6crrESgwzZOgHow_LNw27WU3DdcoEa4kGsySYS73M69ClYsf80ISlydA-E8phsT6U_f4ZuZTwwSrI8SPSAO82FHtxdBIYnMCvP-3loONHKL5FIniTJgMdVb-wyNuQVy0gNcvXD-OXBiSTqWfa7_qd40HUdAJfUBV8Qn7aCRS8_tc0eNMPoTWtJedSEWp7io80sNxKJWRAW6nOh-pejqhDpqfGSNmxJixPCq1Btt69HHNfsImMoTJtikBO4SAEe-H6wzjnbbA7rLJA2ltf0_z0vGRY3KENVAOiGTCRuxj1CSc1JF6YqmmpEx_ff4k7veJ7r-GINUf0nrhG5LBc5Pu2N4CkdDOhg4pI92R1AqzVwgpXJ75yd4X3d3mRV7_EFNZUgRHv3uLSYCnYVKoAGS0kltfcOKJGWbllPnQAloiwFVn4nHrFj2BFDhsOa5VIDibTMp11gg43Ey-x72iuy_Zj9hW3siBJiB2-ZusWxK3sUf--FCoUScCixHPKHdGi_gtx7wjU1G1O1vSCLjThDWzzYuWFYudk0s4CgKNxLsA1IZwxsuKdnx7q-CP04fGuaQuBhgL2Lz_uIQNO6OyNcIg3sIaGJQXrRFlXzX3bSQKH74HWod9p6FskhHom7kWhhwbOD-TgTQM76DfGAuPTbP3R-rp1fYt8YeqskwN14Wu_h8tx2Lr718YR37K0Q9e14klc-ENh7um1HzOkY515MHiCfNQ2SVr7GVCXaKeJyGu6TrjVyDmbncXZDnjwBpXijTLju63tBDJ6OnDJ8vO6iS-52mwWqb5bziFUCOtsBvmMpTYgQvzaGNt-47Oz9wbQmB8giqtjqrapqqANLOWAmPfb0yOUM4utZSIGCT9v1XuxpicGc0jAOYJiAqn18zgTK8Wz_lqepbbSvgEYrnR2w_pwL62EsTNjreJXSet_-r-HmEZ9OHRbzFfS1uHOYbYrm5JTTIlbcvVrzvg6oo7USCtT9G6fmESlOxv8xFTCeJC8rbqwD98Zlrcs8XMP19Z-1_NPcU9E0W40YCND0jqwhuKFU5CxwtbhFbUDd8jUSaCiR7mlfUO2mMp1nMdx7KwkpUJcBKmAT6a8RYE7kPYyTEbx37NMI522bveGAS8wOdksv9_r9nGSBhkE7GfIXUfOtIEAjse_tVb2wpVshENng2O_pbnxSqq-FG68W9mxtIAkK5GG4DHac22zleIy-_UnnAm_VWEyE6C3-gfIxjptn_FeZyItI3PEUglEVJBM_UJhgzd0xn40cvH1XajrsDjjcJBb3IimafjQLatQTKkPlXN0mRb_oyNJ3zJDdOsKwzBlsJRnWBRcP6r0WI7vrRKyQmRlySL0Dahm50ucmfOEk9LyXifbaMxcjbGkrcRY3yQm37hc59PAHAf01oBU7tzyMNHeTRPu60WuOW3kjNr5uPeZ9XawG6_WBq4GdC00FfHASZ-7zY64tGn2ZG82T7f_OBdgrij8Fay32KU3WYrsrP_-V80yu_AbHEYcz1Uygo-6YCJC6FD_wCA81ZP3DYLv9MsGP7Jc3N_sRmWCPg-elmwgFqd4MfPi0oTpNr4nHFe46M8qSMzWqTsbQxx0cPGYBUDpaMXCKdqOpGqAJm1tqiKD-wUTnCnMYnStZQGC-1jblmeTtr5-ulczQPnZG9_VBsw8NeQuFe30czggQePW7IsWWssWQzDywXNBn6tBhN6su_4j9QDA8wNrLiVMiprsdRde7dby_ac8I0CGEDUANAzdxNCbwDtHY0AHnLHwRkCdB7wM_I2l8afKKPH_k54d7LUfmwY_pK4_XZed806lHpWlqY04rJX8s4a7aEayXDdffTPc7_gS0bJ1WM2TbA-ZZTV4IPaRLK_i1ADKLSWDpx3fTJZESF1UzvARhpW8Obhq2nQ79s2_DsBK5geUaH12IRQyXXpWkMZ6traQoOdsOwbr36b3dwMuiu6kTit76nOcKSjQ8EFB7ZfHdqY5DQrvg9BySaNVD5UIWexr2qMU_KfDzKDlHSvaq5DwewURkCxmhRzyzIB27EXCDhiqL7uHpabhiWZyn2k4YHPGvafKBddes5dw6uw0o0bo-trBViwGoOpLxjtivrsZ8zZclRVEQVNUMlth6P8T874hRzNB6WQ5rXXuTNaOcqoOFGso5xBuR-BIgESwMzs7J3mbzJV8xf1vASMl7OTerJJY8Zs89MjAmGOMLugJBE-flLklER4gkGpODGnrG0XSjXmicfjW7yla8U-vHjA8IljS2nPxvbFQimdnGN2PnqUtsL6_Q6Oq_L1FRA2LJbiK4j5I0NE1yMWTrWYdDjxfGR9UBr_t8So-uKl1WjpoaQ0HTYrvXBp3fi0rdxPyGVIaevuzSq04-FiAPMRb4MbXTJ839jBhsH_mFXpvD9ZKweeyS7p9qmd6kZ9yEcJ4Jje0ZBZs-LVtat8FqQPbwsKoNxEkDsc3AjJOWrGSYSFvKWnP4m_idcdjPcsgUSqoUc7F_kX0Z0gXAsV2TQ4PanMuLWQ9cxnRcH5DbGxebNYQqv_hUHaPbTxCVcR3qIM_2KhYv41Ek1RIt7GQp6PopbbD9-5n4ASgc3YqyM4ik0aLzh9yEQc0vBHv4BOm30CbX_T-Ikiu7wajuvN-tpYOyCkBbAAxlmfG0zK3jdOkmkyUtgEUNGiFepsYVBNsNsP1WaCqDROHtrqkLAG3VHOlPurlstYU5VeHzmcUhy4YHVXktTTCH-j8XXu7i_8bsjgau8rlD6DEFB3CZghHdE1344gA_CfDeYFwsJd1qkfgY_sVwb8cH8RjZfmQPpDficNakviXrdbZl3CiY4gJpa0bBQgNW6NafPr6AVgGIP_Bp7WomkgPw27MnD4qK7K48HJoZ7cxkN2koWuGFo77-mUji9z2imfH7kZ5zKxLevkJHoRFoKkzbXwZ6vzaCpyWJIgqHaia9DnoGDZUjH6uGZ3Wf-rRzMvKm6EDVP-3N_wd5ENGYP3I17gTTAqZwbzOsNR_dwKIpO_q8N18WluY5-bPbgo82YxoCKo9Su1qLFXAyH-cWNvEStV1mhZEkiq9WMgROUaU60aHEXV6bwmylkKBNfYoQX6LIw4sAj9bxOg2zrnaH0Z89QS9KoeBrs1aAL95XqFcOlNPgPwG4YZJQ50rjtRXLcNf9-9p9_hySoBwvgnr_9roSpDiNKduDhKyYqVHfNn0cqZp0ocFyWG0L9bZpxIV79KbDZ_Gv6hYvKIuxc7KsNHgXYq-CKMUjf_q89Al_XMmaeVUgN-UOc7DSZafnZzwAmIFXdNg6JOeLHjVBpkexQMVot9LCGollx1ZBG6FcJEs4D4f_ryJPrGW6cd5nnixjdr6j12wwwWGiZdwM6boNV56s6KooTm33_1XannHsIIWBviY-VwCQhcoVx5RVQ5keplRj3XVY1zVg9_rLHIDpUZfQnYjpRPX_yoPRbnsw-DxYDIBzhtGQbvWULj7r2saRFguJBFJJgUDfTutfeFURXEqReh5CORUUbnhb4eP4-bvzrH9A_qv38NbNqXntxyR5SJmFf5ZFZks8oIYkjj5hfZHexVuiTq-W_t6HIXP30ukrEriN1EFkYuamclmhFxVOYqCbDBRdo1uw-OGilhaiNWaGwcVcdLVtQ8GZN1faP5HcOKOZzgeXgvgxm2MOzPGJzJbPqf8VqofG9L58_lJOTKss8M6SeZLBoHjHzS90wDXF2fzXaWJJlVvrgqljX4NghaZL-WxTpSsgKbK80zoTFmBF9u_1KEGbPnbSMEkNHY4uhzHXsxncq9CVoe5Bv_DGM6I2S_x6YBeve4UV2h1KwDUoV3ynp1yGTqjY7H0r3Jz4eyny2XubjY8GW9wkFhDf4FiLzl9uRkUnf0OjPKpfAeBQeoD_8U6gtbVO6rWkGKl-8Frs5YjdIJyMZwNAeqa0bX3tXlVj2yKJvBv3cwkppwwrXbWPC7mgJyVIIvkJGo20VMfgjICpCGJ3X0OE_dKQNxhaGLLaha9ejtKCRVl0ril0_0Ptn7d6i9JdUmnRFfNQUcBecDnymECRhKFMBOTipBsoXFu973jDsIgPfqqkN7oYFbeFGgHtHHMa3sVU8pyrk4pOHXuIzHwsTVBx3Lcu1eTYQhV7zBmSJg7_QMCRlNgFD7iqIgUhjL-xcLOeWJFMjISq65dUcAjzlzy3RiGmzmulmVpf4sqp0pDMhyMzB2lccA4oOYs8FYKR9GvZdLEI7OIfPoVQj_AUxJbidmCQ3izuz1fDnWHdoQG0A3c5hQMK66tvYqExRcYaMVRHi0Z4qgiP5AvE29E9KOemzw9hlDK8aOR5HVH5QdVG86m5ynCr8BGfaBYq-s81sata8vGhyR8vr5e6cV37aUfChnquYdYyqX6OSb5J5NeR42-PWbXhbUzXPAebhCK3vnsPjkZvFqv3Ra3ruuHGiZmepwZiRU-nSLbtLhFeFwIIYDzD_f2hY2g2HawAxolwr_wqkQMwHLwvFat2Gj5PGHQc081RtduBBabMKzEMX5H3aMvTv5QB89pY_AVmOeZUI23bZXxZnXPjAt0soJxSWbLzMfCCAfZLNvTG-kOXnXyYP7EBs0UHtZTF51rm-nAXwAXOgf2fUoC-1sGpCCjHL-fhXmbv74mRSQa_D_wDNgzVeFfVjOLHE7esTkt1ZejojPGTz1K4qLHKyOz5vxGPV5i79PjUqp9idYwquXESK_S-6hPhzBc-q4P-ypHADYqf2ncdkjnXPLaQd8fII3glTLYmw79d0sYLhQg72VR5hKnJCWLAzZfN2RDiXLa9D4b0GFgWlgjUmajJ0mQe0UN8mEUm_-PM7CnJ8B3BGdeIlIPPxRFNQgSk9HMONfhFkf2nAHus9OQ_PaaAxUgsDvAQR754I73lH_orcVNoEJyaFe39mrz4bdEtEByVtKpGwear_IwpTa24is4WuBWYLo6ZVkf_rPbhh-n6jpWgA63RCma8QXTp9KGmPIegKm4bFu8-6wCJDtbgCQmmfydEinXKeipoecagmuy2jJv77n4CrjM4biUrddEcs4wsIrkQhH7vESgRPteq-jFTHLNt6vWAv64lNnx5EVGPCILKX1xKVTU28g3Rn5Q-a-z3e58Jfkf2hWyd9C64NTUDe5WiBBG98qfq0GiVQiASBU-cVO0pqop6RkeLFb_i3jp2BaomWdRutWmZ7-iHBQ04FIY0sN9Mg-IJoeBKN95SIqKI6ROZ2mqmv7uV-I6iB2qArnZykVWdJjcJkfOUMphKaY0MWpLKVZzrd8L8zqPbbc1OLhZj11jokS4fsVcyKorqYHgIyaMD5cQD57_CD5evMrfH_mekcgtWDgoPG2RvE_jPm_wdqZILjFzjAfsYvRXf1SVX4I96m_kg85bJjHCI_jCTN_eUDtOPl4cwkYLJIb9PmUuBBaoaW24zNIl3-YqTM9nZ9z77Gk3-ugpszsDA-HETAy-gqkdGDZD9gYu2VPt7BhNw6bErNh-955230jM3kC4GgXDnsc44W-EX7Fq6w4tCk5mLNC1U1TtOt-Laa_tSGpf_JuXqRY4g7lJq8AgL5p8o6H7Iw4HkHZl2Q2iL2zeqhXb58YUUWcN784cEx6SkDox54jrrNvS3jzGnICjJC5XDrCMx3S-uN6iWwhk5IuMxaU-kZnqCOiDPGNaGaaShkUMlhxxS7hvOoWvf3tlu8d5FZ8KovyH4zjs0CrGlp3Kvn9T2AEsVaqqgc04Y_aLPDzfNdZF_CoW28CVEbVIgRWQCD86XO043BcdIHPk1FdUHPaMf6h8yFirNZCrKdzBXAIiVNp43HqNstSE3rCfG_70q7-cbHAebiw2SB1DqhAgEQ2THnbMceBWA0sVpQcq5fJJSDqIHJYxOwEj3-9c6MU")},
   protobuf.Field{Number: 26, Type: 0, Value: protobuf.Varint(0)},
   protobuf.Field{Number: 29, Type: 0, Value: protobuf.Varint(0)},
   protobuf.Field{Number: 30, Type: 0, Value: protobuf.Varint(0)},
   protobuf.Field{Number: 31, Type: 0, Value: protobuf.Varint(0)},
}
