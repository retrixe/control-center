import 'package:control_center/about.dart';
import 'package:control_center/lenovo.dart';
import 'package:control_center/nouveau.dart';
import 'package:dbus/dbus.dart';
import 'package:flutter/material.dart';
import 'package:yaru_icons/yaru_icons.dart';

abstract class SettingsPage extends StatelessWidget {
  const SettingsPage({Key? key, required this.title, required this.client})
      : super(key: key);

  final String title;
  final DBusClient client;

  Widget buildPage(BuildContext context);

  @override
  Widget build(BuildContext context) {
    const horizontalDivider = Divider(
      height: 1,
      indent: 0,
      endIndent: 0,
    );

    return Scaffold(
      appBar: AppBar(
        title: const Text('Linux Control Center'),
      ),
      body: Row(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          SizedBox(
            width: 160,
            child: ListView(
              scrollDirection: Axis.vertical,
              shrinkWrap: true,
              children: [
                ListTile(
                  selected: title == 'Lenovo',
                  leading: const Icon(YaruIcons.computer),
                  title: const Text('Lenovo'),
                  onTap: () => Navigator.pushReplacement(
                      context, pageBuilder(LenovoSettingsPage(client: client))),
                ),
                horizontalDivider,
                ListTile(
                  selected: title == 'nouveau',
                  leading: const Icon(YaruIcons.chip),
                  title: const Text('nouveau'),
                  onTap: () => Navigator.pushReplacement(context,
                      pageBuilder(NouveauSettingsPage(client: client))),
                ),
                horizontalDivider,
                ListTile(
                  selected: title == 'About',
                  leading: const Icon(YaruIcons.information),
                  title: const Text('About'),
                  onTap: () => Navigator.pushReplacement(
                      context, pageBuilder(AboutPage(client: client))),
                ),
                horizontalDivider,
              ],
            ),
          ),
          const VerticalDivider(
            width: 1,
            indent: 0,
            endIndent: 0,
          ),
          Expanded(child: buildPage(context)),
        ],
      ),
    );
  }

  Future<void> showDBusError(BuildContext context, List<Widget> text) async {
    // The navigator is never poppable unless there's a dialog.
    if (Navigator.of(context).canPop()) return;
    return showDialog<void>(
      context: context,
      barrierDismissible: false,
      builder: (BuildContext context) {
        return AlertDialog(
          title: const Text('Error'),
          content: SingleChildScrollView(
            child: ListBody(
              children: text,
            ),
          ),
          actions: <Widget>[
            TextButton(
              child: const Text('Dismiss'),
              onPressed: () => Navigator.of(context).pop(),
            ),
          ],
        );
      },
    );
  }

  pageBuilder(Widget widget) {
    return PageRouteBuilder(
      pageBuilder: (context, a1, a2) => widget,
      transitionsBuilder: (context, a1, a2, child) =>
          FadeTransition(opacity: a1, child: child),
      transitionDuration: const Duration(milliseconds: 250),
    );
  }
}

class SettingCategory extends StatelessWidget {
  const SettingCategory({Key? key, required this.child}) : super(key: key);

  final Widget child;

  @override
  Widget build(BuildContext context) {
    return Material(
      borderRadius: BorderRadius.circular(4),
      child: Container(
        child: child,
        padding: const EdgeInsets.all(16.0),
        decoration: BoxDecoration(
          borderRadius: BorderRadius.circular(4),
          border: Border.all(color: Theme.of(context).dividerColor),
        ),
      ),
    );
  }
}
